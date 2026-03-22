package models

import (
	"email-manage/pkg/common"
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	// 使用状态：未使用 / 已使用
	MailAccountStatusUnused = "unused"
	MailAccountStatusUsed   = "used"

	MailboxStatusUnknown      = "unknown"
	MailboxStatusReady        = "ready"
	MailboxStatusFetchFailed  = "fetch_failed"
	MailboxStatusTokenExpired = "token_expired"
	MailboxStatusIncomplete   = "incomplete"
)

type MailAccount struct {
	BaseModel
	Email                  string     `gorm:"size:255;not null;uniqueIndex" json:"email"`
	PasswordCiphertext     string     `gorm:"type:text" json:"-"`
	ClientID               string     `gorm:"size:255;index" json:"client_id"`
	RefreshTokenCiphertext string     `gorm:"type:text" json:"-"`
	TokenType              string     `gorm:"size:50" json:"token_type"`
	Folder                 string     `gorm:"size:100;default:inbox" json:"folder"`
	SourceFormat           string     `gorm:"size:50" json:"source_format"`
	ProviderReady          bool       `gorm:"not null;default:false;index" json:"provider_ready"`
	Status                 string     `gorm:"size:50;not null;default:unused;index" json:"status"`
	MailboxStatus          string     `gorm:"size:50;not null;default:unknown;index" json:"mailbox_status"`
	LastPermissionScope    string     `gorm:"type:text" json:"last_permission_scope"`
	LastUseLocalIP         bool       `gorm:"not null;default:false" json:"last_use_local_ip"`
	LastProviderCheckAt    *time.Time `json:"last_provider_check_at"`
	LastMailFetchAt        *time.Time `json:"last_mail_fetch_at"`
	LastCodeFetchAt        *time.Time `json:"last_code_fetch_at"`
	LastCode               string     `gorm:"size:50" json:"last_code"`
	LastCodeAt             *time.Time `json:"last_code_at"`
	LastMailSubject        string     `gorm:"size:500" json:"last_mail_subject"`
	LastMailFrom           string     `gorm:"size:255" json:"last_mail_from"`
	LastMailReceivedAt     *time.Time `json:"last_mail_received_at"`
	AllocatedAt            *time.Time `json:"allocated_at"`
	Remark                 string     `gorm:"size:1000" json:"remark"`
	Tags                   string     `gorm:"type:text" json:"tags"`
	ImportBatchNo          string     `gorm:"size:100;index" json:"import_batch_no"`
}

func (MailAccount) TableName() string {
	return "mail_account"
}

func (m *MailAccount) BeforeCreate(tx *gorm.DB) error {
	if err := m.BaseModel.BeforeCreate(tx); err != nil {
		return err
	}
	m.normalizeBeforeSave()
	return nil
}

func (m *MailAccount) BeforeUpdate(tx *gorm.DB) error {
	if err := m.BaseModel.BeforeUpdate(tx); err != nil {
		return err
	}
	m.normalizeBeforeSave()
	return nil
}

func (m *MailAccount) normalizeBeforeSave() {
	m.Email = strings.ToLower(strings.TrimSpace(m.Email))
	m.ClientID = strings.TrimSpace(m.ClientID)
	m.TokenType = strings.TrimSpace(m.TokenType)
	m.Folder = strings.TrimSpace(m.Folder)
	m.SourceFormat = strings.TrimSpace(m.SourceFormat)
	m.Remark = strings.TrimSpace(m.Remark)
	m.Tags = strings.TrimSpace(m.Tags)
	m.ImportBatchNo = strings.TrimSpace(m.ImportBatchNo)

	if m.Folder == "" {
		m.Folder = "inbox"
	}
	if strings.TrimSpace(m.Status) == "" {
		m.Status = MailAccountStatusUnused
	}

	m.syncProviderState()
}

func (m *MailAccount) syncProviderState() {
	providerReady := strings.TrimSpace(m.ClientID) != "" && strings.TrimSpace(m.RefreshTokenCiphertext) != ""
	m.ProviderReady = providerReady

	if !providerReady {
		m.MailboxStatus = MailboxStatusIncomplete
		return
	}

	if strings.TrimSpace(m.MailboxStatus) == "" || m.MailboxStatus == MailboxStatusIncomplete {
		m.MailboxStatus = MailboxStatusUnknown
	}
}

func (m *MailAccount) IsActive() bool {
	return m.Status == MailAccountStatusUnused
}

func (m *MailAccount) GetIDString() string {
	return m.ID.String()
}

var _ = common.NewUUID
