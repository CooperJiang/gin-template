package mail_account

import (
	"errors"
	"io"

	"email-manage/internal/dto/request"
	"email-manage/internal/services/mail_account"
	"email-manage/pkg/common"
	apierrors "email-manage/pkg/errors"
	"email-manage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	svc *mail_account.Service
}

func NewController(svc *mail_account.Service) *Controller {
	return &Controller{svc: svc}
}

// Stats 获取邮箱统计数据
func (c *Controller) Stats(ctx *gin.Context) {
	result, err := c.svc.GetStats(ctx)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, result, "获取成功")
}

// DailyUsage 获取每日使用量
func (c *Controller) DailyUsage(ctx *gin.Context) {
	result, err := c.svc.GetDailyUsage(ctx)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, result, "获取成功")
}

// Import 导入邮箱账号
func (c *Controller) Import(ctx *gin.Context) {
	req, err := common.ValidateRequest[request.MailAccountImportRequest](ctx)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}

	result, err := c.svc.Import(ctx, *req)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, result, "导入成功")
}

// List 列表查询
func (c *Controller) List(ctx *gin.Context) {
	var req request.MailAccountListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		apierrors.HandleError(ctx, apierrors.New(apierrors.CodeInvalidParameter, "请求参数错误"))
		return
	}

	result, err := c.svc.List(ctx, req)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, result, "获取成功")
}

// GetByID 获取单个账号详情
func (c *Controller) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		apierrors.HandleError(ctx, apierrors.New(apierrors.CodeInvalidParameter, "ID 不能为空"))
		return
	}

	item, err := c.svc.GetByID(ctx, id)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, item, "获取成功")
}

// Create 创建单个账号
func (c *Controller) Create(ctx *gin.Context) {
	req, err := common.ValidateRequest[request.CreateMailAccountRequest](ctx)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}

	item, err := c.svc.Create(ctx, *req)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, item, "创建成功")
}

// Update 更新账号
func (c *Controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		apierrors.HandleError(ctx, apierrors.New(apierrors.CodeInvalidParameter, "ID 不能为空"))
		return
	}

	req, err := common.ValidateRequest[request.UpdateMailAccountRequest](ctx)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}

	item, err := c.svc.Update(ctx, id, *req)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, item, "更新成功")
}

// UpdateStatus 更新账号状态
func (c *Controller) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		apierrors.HandleError(ctx, apierrors.New(apierrors.CodeInvalidParameter, "ID 不能为空"))
		return
	}

	req, err := common.ValidateRequest[request.UpdateMailAccountStatusRequest](ctx)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}

	item, err := c.svc.UpdateStatus(ctx, id, req.Status)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, item, "状态更新成功")
}

// FetchMails 拉取邮箱预览
func (c *Controller) FetchMails(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		apierrors.HandleError(ctx, apierrors.New(apierrors.CodeInvalidParameter, "ID 不能为空"))
		return
	}

	result, err := c.svc.FetchMailbox(ctx, id)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, result, "获取成功")
}

// FetchCodeByID 管理端按账号ID获取最新验证码
func (c *Controller) FetchCodeByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		apierrors.HandleError(ctx, apierrors.New(apierrors.CodeInvalidParameter, "ID 不能为空"))
		return
	}

	var req request.FetchCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil && !errors.Is(err, io.EOF) {
		apierrors.HandleError(ctx, apierrors.New(apierrors.CodeInvalidParameter, "无效的请求参数"))
		return
	}

	resp, err := c.svc.FetchLatestCodeByID(ctx, id, req)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, resp, "获取验证码成功")
}

// Delete 删除单个邮箱账号
func (c *Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		apierrors.HandleError(ctx, apierrors.New(apierrors.CodeInvalidParameter, "ID 不能为空"))
		return
	}

	if err := c.svc.Delete(ctx, id); err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, nil, "删除成功")
}

// BatchDelete 批量删除邮箱账号
func (c *Controller) BatchDelete(ctx *gin.Context) {
	req, err := common.ValidateRequest[request.BatchDeleteMailAccountRequest](ctx)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}

	if err := c.svc.BatchDelete(ctx, req.IDs); err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, nil, "批量删除成功")
}

// AllocateExternal 外部接口：获取一个未使用的邮箱账号（并标记为已使用）
func (c *Controller) AllocateExternal(ctx *gin.Context) {
	logger.Info("[AllocateExternal] 收到分配请求 IP=%s UA=%s", ctx.ClientIP(), ctx.GetHeader("User-Agent"))
	item, err := c.svc.AllocateOne(ctx)
	if err != nil {
		logger.Error("[AllocateExternal] 分配失败 err=%v", err)
		apierrors.HandleError(ctx, err)
		return
	}
	logger.Info("[AllocateExternal] 分配成功 email=%s id=%s", item.Email, item.ID)
	apierrors.ResponseSuccess(ctx, item, "获取邮箱成功")
}

// FetchCodeExternal 外部接口：根据账号ID获取验证码，支持传 sent_after
func (c *Controller) FetchCodeExternal(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		apierrors.HandleError(ctx, apierrors.New(apierrors.CodeInvalidParameter, "ID 不能为空"))
		return
	}

	var req request.FetchCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil && !errors.Is(err, io.EOF) {
		apierrors.HandleError(ctx, apierrors.New(apierrors.CodeInvalidParameter, "无效的请求参数"))
		return
	}

	resp, err := c.svc.FetchLatestCodeByID(ctx, id, req)
	if err != nil {
		apierrors.HandleError(ctx, err)
		return
	}
	apierrors.ResponseSuccess(ctx, resp, "获取验证码成功")
}
