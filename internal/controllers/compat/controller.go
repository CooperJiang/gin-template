package compat

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"email-manage/internal/integrations/mail_provider"
	mailAccountService "email-manage/internal/services/mail_account"
	userService "email-manage/internal/services/user"
	"email-manage/pkg/logger"

	"github.com/gin-gonic/gin"
)

// cachedMessage 带 TTL 的邮件缓存条目
type cachedMessage struct {
	msg       mail_provider.MailMessage
	expiresAt time.Time
}

const cacheTTL = 5 * time.Minute

// Controller 兼容 team-helper 插件的接口控制器
type Controller struct {
	userSvc        userService.Service
	mailAccountSvc *mailAccountService.Service
	messageCache   sync.Map // messageID → cachedMessage
}

// NewController 创建兼容控制器
func NewController(userSvc userService.Service, mailAccountSvc *mailAccountService.Service) *Controller {
	return &Controller{
		userSvc:        userSvc,
		mailAccountSvc: mailAccountSvc,
	}
}

// Login POST /api/login
// 入参: {"username":"...", "password":"..."}
// 成功: Set-Cookie iding-session, 200 {"ok":true}
// 失败: 401 {"error":"..."}
func (ctrl *Controller) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	_, tokenPair, err := ctrl.userSvc.Login(username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// 将 access token 写入 Cookie
	c.SetCookie("iding-session", tokenPair.AccessToken, int(tokenPair.AccessExpiresAt.Sub(time.Now()).Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// Generate GET /api/generate
// 分配一个未使用的邮箱账号
// 返回: {"email":"xxx@domain.com"}
func (ctrl *Controller) Generate(c *gin.Context) {
	logger.Info("[compat.Generate] 收到分配邮箱请求")
	item, err := ctrl.mailAccountSvc.AllocateOne(c)
	if err != nil {
		logger.Error("[compat.Generate] 分配失败 err=%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no available email account"})
		return
	}
	logger.Info("[compat.Generate] 分配成功 email=%s id=%s", item.Email, item.ID)
	c.JSON(http.StatusOK, gin.H{"email": item.Email, "password": item.Password})
}

// ToggleLogin POST /api/mailboxes/toggle-login
// 空操作，直接返回成功
func (ctrl *Controller) ToggleLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// ListEmails GET /api/emails?mailbox={email}
// 拉取邮箱最新邮件并以插件期望的格式返回
func (ctrl *Controller) ListEmails(c *gin.Context) {
	mailbox := strings.TrimSpace(c.Query("mailbox"))
	logger.Info("[compat.ListEmails] 收到请求 mailbox=%s", mailbox)

	if mailbox == "" {
		logger.Warn("[compat.ListEmails] mailbox 参数为空")
		c.JSON(http.StatusBadRequest, gin.H{"error": "mailbox parameter is required"})
		return
	}

	accountID, messages, err := ctrl.mailAccountSvc.FetchMailboxByEmail(c, mailbox)
	if err != nil {
		logger.Error("[compat.ListEmails] FetchMailboxByEmail 失败 mailbox=%s err=%v", mailbox, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch mailbox"})
		return
	}

	logger.Info("[compat.ListEmails] 成功获取邮件 mailbox=%s accountID=%s 邮件数=%d", mailbox, accountID, len(messages))

	// 缓存所有消息并构造响应
	result := make([]gin.H, 0, len(messages))
	now := time.Now()
	for i, msg := range messages {
		ctrl.messageCache.Store(msg.ID, cachedMessage{
			msg:       msg,
			expiresAt: now.Add(cacheTTL),
		})
		logger.Info("[compat.ListEmails] 邮件[%d] id=%s from=%s subject=%s time=%s bodyPreview长度=%d body长度=%d",
			i, msg.ID, msg.FromAddress, msg.Subject,
			msg.ReceivedTime.Format(time.RFC3339),
			len(msg.BodyPreview), len(msg.Body))

		result = append(result, gin.H{
			"id":      msg.ID,
			"from":    msg.FromAddress,
			"subject": msg.Subject,
			"date":    msg.ReceivedTime.Format(time.RFC3339),
			"content": msg.BodyPreview,
		})
	}

	logger.Info("[compat.ListEmails] 返回 %d 封邮件给插件", len(result))
	c.JSON(http.StatusOK, result)
}

// UserQuota GET /api/user/quota
// 对方系统用于测试鉴权是否通过的接口，返回固定假数据
func (ctrl *Controller) UserQuota(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"used":    0,
		"limit":   999999,
		"isAdmin": true,
	})
}

// GetEmail GET /api/email/:id
// 从内存缓存获取邮件详情
func (ctrl *Controller) GetEmail(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	logger.Info("[compat.GetEmail] 收到请求 id=%s", id)

	if id == "" {
		logger.Warn("[compat.GetEmail] id 参数为空")
		c.JSON(http.StatusBadRequest, gin.H{"error": "email id is required"})
		return
	}

	val, ok := ctrl.messageCache.Load(id)
	if !ok {
		// 打印当前缓存中有哪些 key，帮助定位问题
		var cachedKeys []string
		ctrl.messageCache.Range(func(key, _ interface{}) bool {
			cachedKeys = append(cachedKeys, fmt.Sprintf("%v", key))
			return len(cachedKeys) < 20 // 最多打 20 个
		})
		logger.Warn("[compat.GetEmail] 缓存未命中 id=%s 当前缓存keys(%d个)=%v", id, len(cachedKeys), cachedKeys)
		c.JSON(http.StatusNotFound, gin.H{"error": "email not found"})
		return
	}

	entry := val.(cachedMessage)
	if time.Now().After(entry.expiresAt) {
		logger.Warn("[compat.GetEmail] 缓存已过期 id=%s expiresAt=%s", id, entry.expiresAt.Format(time.RFC3339))
		ctrl.messageCache.Delete(id)
		c.JSON(http.StatusNotFound, gin.H{"error": "email not found"})
		return
	}

	msg := entry.msg
	logger.Info("[compat.GetEmail] 缓存命中 id=%s from=%s subject=%s html_content长度=%d content长度=%d",
		msg.ID, msg.FromAddress, msg.Subject, len(msg.Body), len(msg.BodyPreview))

	c.JSON(http.StatusOK, gin.H{
		"id":           msg.ID,
		"from":         msg.FromAddress,
		"subject":      msg.Subject,
		"html_content": msg.Body,
		"content":      msg.BodyPreview,
		"text":         msg.BodyPreview,
	})
}
