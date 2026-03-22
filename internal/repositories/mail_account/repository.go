package mail_account

import (
	"email-manage/internal/models"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ListFilter struct {
	Page          int
	PageSize      int
	Keyword       string
	Status        string
	MailboxStatus string
	ProviderReady *bool
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) baseQuery() *gorm.DB {
	return r.db.Model(&models.MailAccount{})
}

func (r *Repository) List(filter ListFilter) ([]models.MailAccount, int64, error) {
	query := r.baseQuery()
	keyword := strings.TrimSpace(strings.ToLower(filter.Keyword))
	if keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where("LOWER(email) LIKE ? OR LOWER(client_id) LIKE ? OR LOWER(remark) LIKE ?", like, like, like)
	}
	if strings.TrimSpace(filter.Status) != "" {
		query = query.Where("status = ?", strings.TrimSpace(filter.Status))
	}
	if strings.TrimSpace(filter.MailboxStatus) != "" {
		query = query.Where("mailbox_status = ?", strings.TrimSpace(filter.MailboxStatus))
	}
	if filter.ProviderReady != nil {
		query = query.Where("provider_ready = ?", *filter.ProviderReady)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []models.MailAccount
	if err := query.Order("CASE WHEN allocated_at IS NULL THEN 1 ELSE 0 END, allocated_at DESC, created_at DESC").Offset((filter.Page - 1) * filter.PageSize).Limit(filter.PageSize).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *Repository) FindByID(id string) (*models.MailAccount, error) {
	var account models.MailAccount
	err := r.db.Where("id = ?", strings.TrimSpace(id)).First(&account).Error
	return &account, err
}

func (r *Repository) FindByEmail(email string) (*models.MailAccount, error) {
	var account models.MailAccount
	err := r.db.Where("email = ?", strings.ToLower(strings.TrimSpace(email))).First(&account).Error
	return &account, err
}

func (r *Repository) Create(account *models.MailAccount) error {
	return r.db.Create(account).Error
}

func (r *Repository) UpdateByID(id string, updates map[string]interface{}) (int64, error) {
	result := r.db.Model(&models.MailAccount{}).Where("id = ?", strings.TrimSpace(id)).Updates(updates)
	return result.RowsAffected, result.Error
}

func (r *Repository) Save(account *models.MailAccount) error {
	return r.db.Save(account).Error
}

// Delete 根据 ID 删除单个邮箱账号
func (r *Repository) Delete(id string) error {
	return r.db.Where("id = ?", strings.TrimSpace(id)).Delete(&models.MailAccount{}).Error
}

// BatchDelete 根据 ID 列表批量删除邮箱账号
func (r *Repository) BatchDelete(ids []string) error {
	return r.db.Where("id IN ?", ids).Delete(&models.MailAccount{}).Error
}

// FindOneUnused 查找一个未使用的邮箱账号
func (r *Repository) FindOneUnused() (*models.MailAccount, error) {
	var account models.MailAccount
	err := r.db.
		Where("status = ?", models.MailAccountStatusUnused).
		Order("created_at ASC").
		First(&account).Error
	return &account, err
}

// StatusCount 按状态统计的结果
type StatusCount struct {
	Status string `json:"status"`
	Count  int64  `json:"count"`
}

// CountByStatus 按状态分组统计邮箱数量
func (r *Repository) CountByStatus() ([]StatusCount, error) {
	var results []StatusCount
	err := r.baseQuery().
		Select("status, COUNT(*) as count").
		Group("status").
		Scan(&results).Error
	return results, err
}

// DailyUsageItem 每日使用量统计
type DailyUsageItem struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

// DailyUsage 查询最近 N 天的每日分配量（基于 allocated_at 字段）
func (r *Repository) DailyUsage(days int) ([]DailyUsageItem, error) {
	var results []DailyUsageItem

	// 兼容 SQLite 和 MySQL 的日期格式化
	dialect := r.db.Dialector.Name()
	var dateExpr string
	switch dialect {
	case "mysql":
		dateExpr = "DATE_FORMAT(allocated_at, '%Y-%m-%d')"
	default: // sqlite
		dateExpr = "strftime('%Y-%m-%d', allocated_at)"
	}

	err := r.baseQuery().
		Select(dateExpr+" as date, COUNT(*) as count").
		Where("allocated_at IS NOT NULL AND allocated_at >= ?",
			time.Now().AddDate(0, 0, -days)).
		Group("date").
		Order("date ASC").
		Scan(&results).Error

	return results, err
}
