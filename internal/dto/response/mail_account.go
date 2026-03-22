package response

import (
	"email-manage/internal/dto"
	"email-manage/internal/integrations/mail_provider"
	"email-manage/internal/models"
	"strings"
	"time"
)

type MailAccountItem struct {
	dto.BaseResponse
	Email                  string     `json:"email"`
	Password               string     `json:"password"`
	ClientID               string     `json:"client_id"`
	TokenType              string     `json:"token_type"`
	Folder                 string     `json:"folder"`
	SourceFormat           string     `json:"source_format"`
	ProviderReady          bool       `json:"provider_ready"`
	Status                 string     `json:"status"`
	MailboxStatus          string     `json:"mailbox_status"`
	LastPermissionScope    string     `json:"last_permission_scope"`
	LastUseLocalIP         bool       `json:"last_use_local_ip"`
	LastProviderCheckAt    *time.Time `json:"last_provider_check_at"`
	LastMailFetchAt        *time.Time `json:"last_mail_fetch_at"`
	LastCodeFetchAt        *time.Time `json:"last_code_fetch_at"`
	LastCode               string     `json:"last_code"`
	LastCodeAt             *time.Time `json:"last_code_at"`
	LastMailSubject        string     `json:"last_mail_subject"`
	LastMailFrom           string     `json:"last_mail_from"`
	LastMailReceivedAt     *time.Time `json:"last_mail_received_at"`
	AllocatedAt            *time.Time `json:"allocated_at"`
	Remark                 string     `json:"remark"`
	Tags                   string     `json:"tags"`
	ImportBatchNo          string     `json:"import_batch_no"`
	PasswordMasked         string     `json:"password_masked"`
	RefreshTokenMasked     string     `json:"refresh_token_masked"`
	RefreshTokenConfigured bool       `json:"refresh_token_configured"`
}

func NewMailAccountItem(account models.MailAccount, password string) MailAccountItem {
	return MailAccountItem{
		BaseResponse: dto.BaseResponse{
			ID:        account.ID.String(),
			CreatedAt: time.Time(account.CreatedAt),
			UpdatedAt: time.Time(account.UpdatedAt),
		},
		Email:                  account.Email,
		Password:               password,
		ClientID:               account.ClientID,
		TokenType:              account.TokenType,
		Folder:                 account.Folder,
		SourceFormat:           account.SourceFormat,
		ProviderReady:          account.ProviderReady,
		Status:                 account.Status,
		MailboxStatus:          account.MailboxStatus,
		LastPermissionScope:    account.LastPermissionScope,
		LastUseLocalIP:         account.LastUseLocalIP,
		LastProviderCheckAt:    account.LastProviderCheckAt,
		LastMailFetchAt:        account.LastMailFetchAt,
		LastCodeFetchAt:        account.LastCodeFetchAt,
		LastCode:               account.LastCode,
		LastCodeAt:             account.LastCodeAt,
		LastMailSubject:        account.LastMailSubject,
		LastMailFrom:           account.LastMailFrom,
		LastMailReceivedAt:     account.LastMailReceivedAt,
		AllocatedAt:            account.AllocatedAt,
		Remark:                 account.Remark,
		Tags:                   account.Tags,
		ImportBatchNo:          account.ImportBatchNo,
		PasswordMasked:         MaskSecret(account.PasswordCiphertext),
		RefreshTokenMasked:     MaskSecret(account.RefreshTokenCiphertext),
		RefreshTokenConfigured: strings.TrimSpace(account.RefreshTokenCiphertext) != "",
	}
}

type MailAccountListResponse struct {
	Items      []MailAccountItem `json:"items"`
	Total      int64             `json:"total"`
	Page       int               `json:"page"`
	PageSize   int               `json:"page_size"`
	TotalPages int               `json:"total_pages"`
}

type ImportErrorItem struct {
	Line    int    `json:"line"`
	Reason  string `json:"reason"`
	RawLine string `json:"raw_line"`
}

type MailAccountImportResponse struct {
	TotalLines   int               `json:"total_lines"`
	CreatedCount int               `json:"created_count"`
	UpdatedCount int               `json:"updated_count"`
	SkippedCount int               `json:"skipped_count"`
	FailedCount  int               `json:"failed_count"`
	BatchNo      string            `json:"batch_no"`
	Errors       []ImportErrorItem `json:"errors"`
}

type MailPreviewResponse struct {
	Messages []mail_provider.MailMessage `json:"messages"`
}

type FetchCodeResponse struct {
	Code           string     `json:"code"`
	MatchedSubject string     `json:"matched_subject"`
	MatchedFrom    string     `json:"matched_from"`
	ReceivedAt     *time.Time `json:"received_at"`
	BodyPreview    string     `json:"body_preview"`
}

func MaskSecret(value string) string {
	if strings.TrimSpace(value) == "" {
		return ""
	}
	return "已配置"
}

// MailAccountStatsResponse 邮箱统计响应
type MailAccountStatsResponse struct {
	Total  int64 `json:"total"`
	Unused int64 `json:"unused"`
	Used   int64 `json:"used"`
}

// DailyUsageItem 每日使用量条目
type DailyUsageItem struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

// DailyUsageResponse 每日使用量响应
type DailyUsageResponse struct {
	Items []DailyUsageItem `json:"items"`
}
