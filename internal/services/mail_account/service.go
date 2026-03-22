package mail_account

import (
	"context"
	stdErrors "errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"email-manage/internal/dto/request"
	"email-manage/internal/dto/response"
	"email-manage/internal/integrations/mail_provider"
	"email-manage/internal/models"
	mailrepo "email-manage/internal/repositories/mail_account"
	apierrors "email-manage/pkg/errors"
	"email-manage/pkg/logger"
	"email-manage/pkg/security"

	"gorm.io/gorm"
)

var codeRegex = regexp.MustCompile(`\b\d{6}\b`)

// Service 封装邮箱账号的核心业务逻辑
type Service struct {
	repo        *mailrepo.Repository
	provider    mail_provider.Client
	fieldCipher *security.FieldCipher
}

func NewService(
	repo *mailrepo.Repository,
	provider mail_provider.Client,
	fieldCipher *security.FieldCipher,
) *Service {
	return &Service{
		repo:        repo,
		provider:    provider,
		fieldCipher: fieldCipher,
	}
}

// GetStats 获取邮箱统计数据（总数、已用、未用）
func (s *Service) GetStats(ctx context.Context) (*response.MailAccountStatsResponse, error) {
	_ = ctx

	counts, err := s.repo.CountByStatus()
	if err != nil {
		return nil, apierrors.New(apierrors.CodeQueryFailed, "查询邮箱统计失败")
	}

	resp := &response.MailAccountStatsResponse{}
	for _, c := range counts {
		resp.Total += c.Count
		switch c.Status {
		case models.MailAccountStatusUnused:
			resp.Unused = c.Count
		case models.MailAccountStatusUsed:
			resp.Used = c.Count
		}
	}
	return resp, nil
}

// GetDailyUsage 获取最近 30 天每日使用量
func (s *Service) GetDailyUsage(ctx context.Context) (*response.DailyUsageResponse, error) {
	_ = ctx

	items, err := s.repo.DailyUsage(30)
	if err != nil {
		return nil, apierrors.New(apierrors.CodeQueryFailed, "查询每日使用量失败")
	}

	respItems := make([]response.DailyUsageItem, 0, len(items))
	for _, item := range items {
		respItems = append(respItems, response.DailyUsageItem{
			Date:  item.Date,
			Count: item.Count,
		})
	}
	return &response.DailyUsageResponse{Items: respItems}, nil
}

// AllocateOne 分配一个未使用的邮箱账号，并标记为已使用
func (s *Service) AllocateOne(ctx context.Context) (*response.MailAccountItem, error) {
	_ = ctx

	logger.Info("[svc.AllocateOne] 开始分配邮箱 unused→used")

	account, err := s.repo.FindOneUnused()
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("[svc.AllocateOne] 没有可用的邮箱账号")
			return nil, apierrors.New(apierrors.CodeNotFound, "没有可用的邮箱账号")
		}
		return nil, apierrors.New(apierrors.CodeQueryFailed, "查询可用邮箱失败")
	}

	logger.Info("[svc.AllocateOne] 找到未使用账号 email=%s id=%s 即将标记为 used", account.Email, account.GetIDString())

	account.Status = models.MailAccountStatusUsed
	now := time.Now()
	account.AllocatedAt = &now
	if err := s.repo.Save(account); err != nil {
		return nil, apierrors.New(apierrors.CodeQueryFailed, "更新邮箱状态失败")
	}

	logger.Info("[svc.AllocateOne] 分配完成 email=%s id=%s allocatedAt=%s", account.Email, account.GetIDString(), now.Format(time.RFC3339))

	password, _ := s.decrypt(account.PasswordCiphertext)
	item := response.NewMailAccountItem(*account, password)
	return &item, nil
}

// Import 从多行文本导入邮箱账号，支持多种格式混合
func (s *Service) Import(ctx context.Context, req request.MailAccountImportRequest) (*response.MailAccountImportResponse, error) {
	_ = ctx

	mode := strings.TrimSpace(req.Mode)
	if mode == "" {
		mode = "fill_missing"
	}

	lines := splitLines(req.Content)
	result := &response.MailAccountImportResponse{
		TotalLines: len(lines),
		Errors:     make([]response.ImportErrorItem, 0),
	}
	if len(lines) == 0 {
		return result, nil
	}

	batchNo := fmt.Sprintf("BATCH-%d", time.Now().UnixNano())
	result.BatchNo = batchNo

	for idx, rawLine := range lines {
		lineNo := idx + 1
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") {
			result.SkippedCount++
			continue
		}

		account, plainPassword, plainRefreshToken, reason := s.parseLine(line)
		if reason != "" {
			result.FailedCount++
			result.Errors = append(result.Errors, response.ImportErrorItem{
				Line:    lineNo,
				Reason:  reason,
				RawLine: maskImportLine(line),
			})
			continue
		}

		account.ImportBatchNo = batchNo
		if err := s.setSecrets(&account, plainPassword, plainRefreshToken); err != nil {
			result.FailedCount++
			result.Errors = append(result.Errors, response.ImportErrorItem{
				Line:    lineNo,
				Reason:  err.Error(),
				RawLine: maskImportLine(line),
			})
			continue
		}

		existing, err := s.repo.FindByEmail(account.Email)
		if err != nil && !stdErrors.Is(err, gorm.ErrRecordNotFound) {
			result.FailedCount++
			result.Errors = append(result.Errors, response.ImportErrorItem{
				Line:    lineNo,
				Reason:  "查询已存在账号失败",
				RawLine: maskImportLine(line),
			})
			continue
		}

		if err == nil && existing != nil {
			if mode != "fill_missing" {
				result.SkippedCount++
				continue
			}

			updated := false
			if strings.TrimSpace(existing.PasswordCiphertext) == "" && strings.TrimSpace(account.PasswordCiphertext) != "" {
				existing.PasswordCiphertext = account.PasswordCiphertext
				updated = true
			}
			if strings.TrimSpace(existing.ClientID) == "" && strings.TrimSpace(account.ClientID) != "" {
				existing.ClientID = account.ClientID
				updated = true
			}
			if strings.TrimSpace(existing.RefreshTokenCiphertext) == "" && strings.TrimSpace(account.RefreshTokenCiphertext) != "" {
				existing.RefreshTokenCiphertext = account.RefreshTokenCiphertext
				updated = true
			}
			if strings.TrimSpace(existing.SourceFormat) == "" && strings.TrimSpace(account.SourceFormat) != "" {
				existing.SourceFormat = account.SourceFormat
				updated = true
			}
			existing.ImportBatchNo = batchNo

			if !updated {
				if err := s.repo.Save(existing); err != nil {
					result.FailedCount++
					result.Errors = append(result.Errors, response.ImportErrorItem{
						Line:    lineNo,
						Reason:  "更新账号失败",
						RawLine: maskImportLine(line),
					})
					continue
				}
				result.SkippedCount++
				continue
			}

			if err := s.repo.Save(existing); err != nil {
				result.FailedCount++
				result.Errors = append(result.Errors, response.ImportErrorItem{
					Line:    lineNo,
					Reason:  "更新账号失败",
					RawLine: maskImportLine(line),
				})
				continue
			}
			result.UpdatedCount++
			continue
		}

		if err := s.repo.Create(&account); err != nil {
			if isDuplicateError(err) {
				result.FailedCount++
				result.Errors = append(result.Errors, response.ImportErrorItem{
					Line:    lineNo,
					Reason:  "邮箱已存在",
					RawLine: maskImportLine(line),
				})
				continue
			}
			result.FailedCount++
			result.Errors = append(result.Errors, response.ImportErrorItem{
				Line:    lineNo,
				Reason:  "创建账号失败",
				RawLine: maskImportLine(line),
			})
			continue
		}
		result.CreatedCount++
	}

	return result, nil
}

// Delete 删除单个邮箱账号
func (s *Service) Delete(ctx context.Context, id string) error {
	_ = ctx

	if _, err := s.getAccountByID(id); err != nil {
		return err
	}

	if err := s.repo.Delete(id); err != nil {
		return apierrors.New(apierrors.CodeQueryFailed, "删除邮箱账号失败")
	}
	return nil
}

// BatchDelete 批量删除邮箱账号
func (s *Service) BatchDelete(ctx context.Context, ids []string) error {
	_ = ctx

	if err := s.repo.BatchDelete(ids); err != nil {
		return apierrors.New(apierrors.CodeQueryFailed, "批量删除邮箱账号失败")
	}
	return nil
}

// GetByEmail 根据邮箱地址查找账号
func (s *Service) GetByEmail(ctx context.Context, email string) (*models.MailAccount, error) {
	_ = ctx

	account, err := s.repo.FindByEmail(email)
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apierrors.New(apierrors.CodeNotFound, "邮箱账号不存在")
		}
		return nil, apierrors.New(apierrors.CodeQueryFailed, "查询邮箱账号失败")
	}
	return account, nil
}

// FetchMailboxByEmail 根据邮箱地址拉取邮件，返回 (accountID, messages, error)
func (s *Service) FetchMailboxByEmail(ctx context.Context, email string) (string, []mail_provider.MailMessage, error) {
	logger.Info("[svc.FetchMailboxByEmail] 开始 email=%s", email)

	account, err := s.GetByEmail(ctx, email)
	if err != nil {
		logger.Error("[svc.FetchMailboxByEmail] GetByEmail 失败 email=%s err=%v", email, err)
		return "", nil, err
	}
	logger.Info("[svc.FetchMailboxByEmail] 找到账号 email=%s accountID=%s clientID=%s",
		account.Email, account.GetIDString(), account.ClientID)

	preview, err := s.FetchMailbox(ctx, account.GetIDString())
	if err != nil {
		logger.Error("[svc.FetchMailboxByEmail] FetchMailbox 失败 accountID=%s err=%v", account.GetIDString(), err)
		return "", nil, err
	}

	logger.Info("[svc.FetchMailboxByEmail] 拉取成功 email=%s 邮件数=%d", email, len(preview.Messages))
	for i, msg := range preview.Messages {
		logger.Info("[svc.FetchMailboxByEmail] 邮件[%d] id=%s from=%s subject=%s time=%s",
			i, msg.ID, msg.FromAddress, msg.Subject, msg.ReceivedTime.Format(time.RFC3339))
	}

	return account.GetIDString(), preview.Messages, nil
}

// GetByID 获取账号详情
func (s *Service) GetByID(ctx context.Context, id string) (*response.MailAccountItem, error) {
	_ = ctx

	account, err := s.getAccountByID(id)
	if err != nil {
		return nil, err
	}

	password, _ := s.decrypt(account.PasswordCiphertext)
	item := response.NewMailAccountItem(*account, password)
	return &item, nil
}

// List 返回邮箱账号分页列表
func (s *Service) List(ctx context.Context, req request.MailAccountListRequest) (*response.MailAccountListResponse, error) {
	_ = ctx

	page := req.Page
	size := req.Size
	if page <= 0 {
		page = 1
	}
	if size <= 0 || size > 200 {
		size = 20
	}

	var providerReady *bool
	switch strings.TrimSpace(req.ProviderReady) {
	case "true":
		t := true
		providerReady = &t
	case "false":
		f := false
		providerReady = &f
	}

	items, total, err := s.repo.List(mailrepo.ListFilter{
		Page:          page,
		PageSize:      size,
		Keyword:       req.Keyword,
		Status:        req.Status,
		MailboxStatus: req.MailboxStatus,
		ProviderReady: providerReady,
	})
	if err != nil {
		return nil, apierrors.New(apierrors.CodeQueryFailed, "查询邮箱账号列表失败")
	}

	respItems := make([]response.MailAccountItem, 0, len(items))
	for _, item := range items {
		pw, _ := s.decrypt(item.PasswordCiphertext)
		respItems = append(respItems, response.NewMailAccountItem(item, pw))
	}

	totalPages := 0
	if total > 0 {
		totalPages = int((total + int64(size) - 1) / int64(size))
	}

	return &response.MailAccountListResponse{
		Items:      respItems,
		Total:      total,
		Page:       page,
		PageSize:   size,
		TotalPages: totalPages,
	}, nil
}

// Create 创建单个账号
func (s *Service) Create(ctx context.Context, req request.CreateMailAccountRequest) (*response.MailAccountItem, error) {
	_ = ctx

	account := models.MailAccount{
		Email:        strings.TrimSpace(req.Email),
		ClientID:     strings.TrimSpace(req.ClientID),
		SourceFormat: "manual",
		Status:       defaultStatus(req.Status),
		Remark:       strings.TrimSpace(req.Remark),
		Tags:         strings.TrimSpace(req.Tags),
		Folder:       strings.TrimSpace(req.Folder),
	}
	if err := s.setSecrets(&account, req.Password, req.RefreshToken); err != nil {
		return nil, err
	}

	if err := s.repo.Create(&account); err != nil {
		if isDuplicateError(err) {
			return nil, apierrors.New(apierrors.CodeConflict, "邮箱已存在")
		}
		return nil, apierrors.New(apierrors.CodeQueryFailed, "创建邮箱账号失败")
	}

	plainPw, _ := s.decrypt(account.PasswordCiphertext)
	item := response.NewMailAccountItem(account, plainPw)
	return &item, nil
}

// Update 更新账号的可变字段
func (s *Service) Update(ctx context.Context, id string, req request.UpdateMailAccountRequest) (*response.MailAccountItem, error) {
	_ = ctx

	account, err := s.getAccountByID(id)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(req.Status) != "" {
		account.Status = strings.TrimSpace(req.Status)
	}
	if req.Remark != "" {
		account.Remark = strings.TrimSpace(req.Remark)
	}
	if req.Tags != "" {
		account.Tags = strings.TrimSpace(req.Tags)
	}
	if req.Folder != "" {
		account.Folder = strings.TrimSpace(req.Folder)
	}
	if req.ClientID != "" {
		account.ClientID = strings.TrimSpace(req.ClientID)
	}

	if err := s.setSecrets(account, req.Password, req.RefreshToken); err != nil {
		return nil, err
	}
	if err := s.repo.Save(account); err != nil {
		return nil, apierrors.New(apierrors.CodeQueryFailed, "更新邮箱账号失败")
	}

	updatedPw, _ := s.decrypt(account.PasswordCiphertext)
	item := response.NewMailAccountItem(*account, updatedPw)
	return &item, nil
}

// UpdateStatus 更新账号状态
func (s *Service) UpdateStatus(ctx context.Context, id string, status string) (*response.MailAccountItem, error) {
	_ = ctx

	account, err := s.getAccountByID(id)
	if err != nil {
		return nil, err
	}

	account.Status = strings.TrimSpace(status)
	if err := s.repo.Save(account); err != nil {
		return nil, apierrors.New(apierrors.CodeQueryFailed, "更新账号状态失败")
	}

	statusPw, _ := s.decrypt(account.PasswordCiphertext)
	item := response.NewMailAccountItem(*account, statusPw)
	return &item, nil
}

// FetchMailbox 拉取邮箱最新邮件预览
func (s *Service) FetchMailbox(ctx context.Context, id string) (*response.MailPreviewResponse, error) {
	logger.Info("[svc.FetchMailbox] 开始 id=%s", id)

	account, err := s.getAccountByID(id)
	if err != nil {
		logger.Error("[svc.FetchMailbox] getAccountByID 失败 id=%s err=%v", id, err)
		return nil, err
	}
	logger.Info("[svc.FetchMailbox] 账号信息 email=%s clientID=%s tokenType=%s folder=%s hasPassword=%v hasRefreshToken=%v",
		account.Email, account.ClientID, account.TokenType, account.Folder,
		strings.TrimSpace(account.PasswordCiphertext) != "",
		strings.TrimSpace(account.RefreshTokenCiphertext) != "")

	if err := ensureMailboxFetchable(account); err != nil {
		logger.Error("[svc.FetchMailbox] ensureMailboxFetchable 失败 email=%s err=%v", account.Email, err)
		return nil, err
	}

	logger.Info("[svc.FetchMailbox] 开始 syncPermission email=%s", account.Email)
	if err := s.syncPermission(account); err != nil {
		logger.Error("[svc.FetchMailbox] syncPermission 失败 email=%s err=%v", account.Email, err)
		return nil, err
	}
	logger.Info("[svc.FetchMailbox] syncPermission 成功 email=%s tokenType=%s", account.Email, account.TokenType)

	password, err := s.decrypt(account.PasswordCiphertext)
	if err != nil {
		logger.Error("[svc.FetchMailbox] 解密密码失败 email=%s err=%v", account.Email, err)
		return nil, err
	}
	refreshToken, err := s.decrypt(account.RefreshTokenCiphertext)
	if err != nil {
		logger.Error("[svc.FetchMailbox] 解密refreshToken失败 email=%s err=%v", account.Email, err)
		return nil, err
	}

	logger.Info("[svc.FetchMailbox] 调用 provider.FetchMailbox email=%s folder=%s", account.Email, account.Folder)
	result, err := s.provider.FetchMailbox(mail_provider.AccountCredentials{
		Email:        account.Email,
		Password:     password,
		ClientID:     account.ClientID,
		RefreshToken: refreshToken,
		TokenType:    account.TokenType,
		Folder:       account.Folder,
	})
	if err != nil {
		logger.Error("[svc.FetchMailbox] provider.FetchMailbox 失败 email=%s err=%v", account.Email, err)
		now := time.Now()
		account.LastMailFetchAt = &now
		account.MailboxStatus = models.MailboxStatusFetchFailed
		_ = s.repo.Save(account)
		return nil, apierrors.New(apierrors.CodeThirdPartyService, "拉取邮件失败")
	}
	logger.Info("[svc.FetchMailbox] provider 返回 %d 封邮件 email=%s", len(result.Messages), account.Email)

	now := time.Now()
	account.LastMailFetchAt = &now
	if len(result.Messages) > 0 {
		latest := result.Messages[0]
		account.LastMailSubject = strings.TrimSpace(latest.Subject)
		account.LastMailFrom = strings.TrimSpace(latest.FromAddress)
		account.LastMailReceivedAt = &latest.ReceivedTime
	}
	account.MailboxStatus = models.MailboxStatusReady

	if err := s.repo.Save(account); err != nil {
		return nil, apierrors.New(apierrors.CodeQueryFailed, "保存邮件预览结果失败")
	}

	return &response.MailPreviewResponse{Messages: result.Messages}, nil
}

// FetchLatestCodeByID 通过账号ID获取最新验证码
// req.Keyword: 可选关键词，用于过滤邮件（如 "otp" / "code"）
// req.SentAfter: 可选 Unix 时间戳（秒），表示外部触发发送验证码的时间，只考虑该时间之后的邮件
func (s *Service) FetchLatestCodeByID(ctx context.Context, id string, req request.FetchCodeRequest) (*response.FetchCodeResponse, error) {
	account, err := s.getAccountByID(id)
	if err != nil {
		return nil, err
	}

	preview, err := s.FetchMailbox(ctx, id)
	if err != nil {
		return nil, err
	}

	// 计算时间下限：
	// 1. 优先使用外部传入的 sent_after
	// 2. 其次使用分配时间 AllocatedAt（分配邮箱时记录，确保只查分配之后的验证码）
	// 3. 再次使用上次验证码时间 LastCodeAt
	// 4. 兜底：now - 30 分钟
	now := time.Now()
	cutoff := now.Add(-30 * time.Minute)
	if req.SentAfter > 0 {
		cutoff = time.Unix(req.SentAfter, 0)
	} else if account.AllocatedAt != nil && account.AllocatedAt.After(cutoff) {
		cutoff = *account.AllocatedAt
	} else if account.LastCodeAt != nil && account.LastCodeAt.After(cutoff) {
		cutoff = *account.LastCodeAt
	}

	codeResp := extractCodeFromMessages(preview.Messages, req.Keyword, cutoff)
	if codeResp == nil {
		return nil, apierrors.New(apierrors.CodeNotFound, "未在最近邮件中匹配到验证码")
	}

	now = time.Now()
	account.LastCode = codeResp.Code
	account.LastCodeAt = codeResp.ReceivedAt
	account.LastCodeFetchAt = &now
	if err := s.repo.Save(account); err != nil {
		return nil, apierrors.New(apierrors.CodeQueryFailed, "保存验证码结果失败")
	}

	return codeResp, nil
}

func (s *Service) syncPermission(account *models.MailAccount) error {
	refreshToken, err := s.decrypt(account.RefreshTokenCiphertext)
	if err != nil {
		return err
	}

	result, err := s.provider.DetectPermission(account.ClientID, refreshToken)
	if err != nil {
		now := time.Now()
		account.LastProviderCheckAt = &now
		account.LastPermissionScope = ""
		account.LastUseLocalIP = false
		account.MailboxStatus = models.MailboxStatusFetchFailed
		_ = s.repo.Save(account)
		return apierrors.New(apierrors.CodeThirdPartyService, "检测邮箱权限失败")
	}

	account.TokenType = strings.TrimSpace(result.TokenType)
	account.LastPermissionScope = strings.TrimSpace(result.Scope)
	account.LastUseLocalIP = result.UseLocalIP
	account.LastProviderCheckAt = &result.CheckedAt
	account.MailboxStatus = models.MailboxStatusUnknown
	if err := s.repo.Save(account); err != nil {
		return apierrors.New(apierrors.CodeQueryFailed, "保存权限检测结果失败")
	}
	return nil
}

func (s *Service) getAccountByID(id string) (*models.MailAccount, error) {
	account, err := s.repo.FindByID(id)
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apierrors.New(apierrors.CodeNotFound, "邮箱账号不存在")
		}
		return nil, apierrors.New(apierrors.CodeQueryFailed, "查询邮箱账号失败")
	}
	return account, nil
}

func (s *Service) setSecrets(account *models.MailAccount, password, refreshToken string) error {
	pw := strings.TrimSpace(password)
	rt := strings.TrimSpace(refreshToken)
	if pw == "" && rt == "" {
		return nil
	}
	// 未配置加密密钥时，直接按明文存储，保证功能正常
	if s.fieldCipher == nil {
		if pw != "" {
			account.PasswordCiphertext = pw
		}
		if rt != "" {
			account.RefreshTokenCiphertext = rt
		}
		return nil
	}

	// 配置了加密密钥时，对敏感字段加密存储
	if pw != "" {
		enc, err := s.fieldCipher.Encrypt(pw)
		if err != nil {
			return apierrors.New(apierrors.CodeInternal, "加密邮箱密码失败")
		}
		account.PasswordCiphertext = enc
	}
	if rt != "" {
		enc, err := s.fieldCipher.Encrypt(rt)
		if err != nil {
			return apierrors.New(apierrors.CodeInternal, "加密 refresh token 失败")
		}
		account.RefreshTokenCiphertext = enc
	}

	return nil
}

func (s *Service) decrypt(cipherText string) (string, error) {
	trimmed := strings.TrimSpace(cipherText)
	if trimmed == "" {
		return "", nil
	}
	// 未配置加密密钥时，直接把存储值当作明文使用（兼容老数据）
	if s.fieldCipher == nil {
		return trimmed, nil
	}

	plain, err := s.fieldCipher.Decrypt(trimmed)
	if err != nil {
		// 解密失败时也退回使用原始值，避免影响业务流程
		return trimmed, nil
	}
	return plain, nil
}

func splitLines(content string) []string {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	if content == "" {
		return []string{}
	}
	return strings.Split(content, "\n")
}

// parseLine 解析一行账号信息，支持三种格式，返回账号基础信息 + 明文密码/token
// 格式1: 邮箱----密码----ID----token
// 格式2: 邮箱\t密码\tID\ttoken
// 格式3: 邮箱----密码
func (s *Service) parseLine(line string) (models.MailAccount, string, string, string) {
	_ = s

	if parts := strings.Split(line, "\t"); len(parts) == 4 {
		email := normalizeEmail(parts[0])
		password := strings.TrimSpace(parts[1])
		clientID := strings.TrimSpace(parts[2])
		refreshToken := strings.TrimSpace(parts[3])
		if reason := validateImportEmail(email); reason != "" {
			return models.MailAccount{}, "", "", reason
		}
		return models.MailAccount{
			Email:        email,
			ClientID:     clientID,
			SourceFormat: "format2",
			Status:       models.MailAccountStatusUnused,
		}, password, refreshToken, ""
	}

	if parts := strings.Split(line, "----"); len(parts) == 4 {
		email := normalizeEmail(parts[0])
		password := strings.TrimSpace(parts[1])
		clientID := strings.TrimSpace(parts[2])
		refreshToken := strings.TrimSpace(parts[3])
		if reason := validateImportEmail(email); reason != "" {
			return models.MailAccount{}, "", "", reason
		}
		return models.MailAccount{
			Email:        email,
			ClientID:     clientID,
			SourceFormat: "format1",
			Status:       models.MailAccountStatusUnused,
		}, password, refreshToken, ""
	}

	if parts := strings.Split(line, "----"); len(parts) == 2 {
		email := normalizeEmail(parts[0])
		password := strings.TrimSpace(parts[1])
		if reason := validateImportEmail(email); reason != "" {
			return models.MailAccount{}, "", "", reason
		}
		return models.MailAccount{
			Email:        email,
			SourceFormat: "format3",
			Status:       models.MailAccountStatusUnused,
		}, password, "", ""
	}

	return models.MailAccount{}, "", "", "无法识别的账号行格式"
}

func defaultStatus(status string) string {
	status = strings.TrimSpace(status)
	if status == "" {
		return models.MailAccountStatusUnused
	}
	return status
}

func normalizeEmail(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func validateImportEmail(email string) string {
	if email == "" {
		return "邮箱为空"
	}
	if !strings.Contains(email, "@") || strings.HasPrefix(email, "@") || strings.HasSuffix(email, "@") {
		return "邮箱格式不正确"
	}
	return ""
}

func ensureMailboxFetchable(account *models.MailAccount) error {
	if account == nil {
		return apierrors.New(apierrors.CodeNotFound, "邮箱账号不存在")
	}
	if strings.TrimSpace(account.PasswordCiphertext) == "" {
		return apierrors.New(apierrors.CodeInvalidParameter, "当前账号缺少邮箱密码，无法拉取邮件")
	}
	return nil
}

func extractCodeFromMessages(messages []mail_provider.MailMessage, keyword string, notBefore time.Time) *response.FetchCodeResponse {
	keyword = strings.TrimSpace(strings.ToLower(keyword))

	for _, message := range messages {
		// 只考虑 notBefore 之后收到的邮件，避免返回很久以前的旧验证码
		if message.ReceivedTime.Before(notBefore) {
			continue
		}

		if keyword != "" {
			target := strings.ToLower(message.Subject + " " + message.BodyPreview + " " + message.Body)
			if !strings.Contains(target, keyword) {
				continue
			}
		}

		if code := firstSixDigitCode(message.Subject); code != "" {
			return newFetchCodeResponse(message, code)
		}
		if code := firstSixDigitCode(message.BodyPreview); code != "" {
			return newFetchCodeResponse(message, code)
		}
		if code := firstSixDigitCode(message.Body); code != "" {
			return newFetchCodeResponse(message, code)
		}
	}

	return nil
}

func firstSixDigitCode(text string) string {
	return codeRegex.FindString(text)
}

func newFetchCodeResponse(message mail_provider.MailMessage, code string) *response.FetchCodeResponse {
	receivedAt := message.ReceivedTime
	return &response.FetchCodeResponse{
		Code:           code,
		MatchedSubject: message.Subject,
		MatchedFrom:    strings.TrimSpace(message.FromAddress),
		ReceivedAt:     &receivedAt,
		BodyPreview:    strings.TrimSpace(message.BodyPreview),
	}
}

func maskImportLine(line string) string {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return ""
	}

	if strings.Contains(trimmed, "\t") {
		parts := strings.Split(trimmed, "\t")
		for i := 1; i < len(parts); i++ {
			parts[i] = "****"
		}
		return strings.Join(parts, "\t")
	}

	if strings.Contains(trimmed, "----") {
		parts := strings.Split(trimmed, "----")
		for i := 1; i < len(parts); i++ {
			parts[i] = "****"
		}
		return strings.Join(parts, "----")
	}

	return "****"
}

func isDuplicateError(err error) bool {
	if err == nil {
		return false
	}
	text := strings.ToLower(err.Error())
	return strings.Contains(text, "duplicate") || strings.Contains(text, "unique constraint") || strings.Contains(text, "duplicated key")
}
