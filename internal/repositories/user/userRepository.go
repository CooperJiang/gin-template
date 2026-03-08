package user

import (
	"strings"
	"template/internal/models"
	"template/pkg/common"

	"gorm.io/gorm"
)

// Repository 用户仓储
type Repository struct {
	db *gorm.DB
}

// NewRepository 创建用户仓储
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// FindAll 获取所有用户
func (r *Repository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

// FindByID 根据ID获取用户
func (r *Repository) FindByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", strings.TrimSpace(id)).First(&user).Error
	return &user, err
}

// FindByEmail 根据邮箱获取用户
func (r *Repository) FindByEmail(emailAddr string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", strings.TrimSpace(emailAddr)).First(&user).Error
	return &user, err
}

// FindByAccount 根据账户(用户名/邮箱)查询用户
func (r *Repository) FindByAccount(account string) (*models.User, error) {
	var user models.User
	account = strings.TrimSpace(account)
	err := r.db.Where("username = ? OR email = ?", account, account).First(&user).Error
	return &user, err
}

// FindAuthByID 查询鉴权所需字段
func (r *Repository) FindAuthByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.
		Select("id", "username", "role", "status", "token_version").
		Where("id = ?", strings.TrimSpace(id)).
		Take(&user).Error
	return &user, err
}

// FindPasswordByID 查询密码字段
func (r *Repository) FindPasswordByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.
		Select("id", "password").
		Where("id = ?", strings.TrimSpace(id)).
		Take(&user).Error
	return &user, err
}

// FindPublicByID 查询公开用户资料字段
func (r *Repository) FindPublicByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.
		Select("id", "username", "email", "avatar", "bio", "status", "role", "created_at", "updated_at").
		Where("id = ?", strings.TrimSpace(id)).
		Take(&user).Error
	return &user, err
}

// CountByUsername 统计用户名数量
func (r *Repository) CountByUsername(username string) (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("username = ?", strings.TrimSpace(username)).Count(&count).Error
	return count, err
}

// CountByEmail 统计邮箱数量
func (r *Repository) CountByEmail(emailAddr string) (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", strings.TrimSpace(emailAddr)).Count(&count).Error
	return count, err
}

// CountByUsernameExcludeID 统计排除指定ID后的用户名数量
func (r *Repository) CountByUsernameExcludeID(username, excludeID string) (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).
		Where("username = ? AND id != ?", strings.TrimSpace(username), strings.TrimSpace(excludeID)).
		Count(&count).Error
	return count, err
}

// CountByEmailExcludeID 统计排除指定ID后的邮箱数量
func (r *Repository) CountByEmailExcludeID(emailAddr, excludeID string) (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).
		Where("email = ? AND id != ?", strings.TrimSpace(emailAddr), strings.TrimSpace(excludeID)).
		Count(&count).Error
	return count, err
}

// Create 创建用户
func (r *Repository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// UpdateByID 根据ID更新用户
func (r *Repository) UpdateByID(id string, updates map[string]interface{}) (int64, error) {
	result := r.db.Model(&models.User{}).Where("id = ?", strings.TrimSpace(id)).Updates(updates)
	return result.RowsAffected, result.Error
}

// UpdatePasswordByEmail 根据邮箱更新密码并提升token_version
func (r *Repository) UpdatePasswordByEmail(emailAddr, hashedPassword string) (int64, error) {
	result := r.db.Model(&models.User{}).Where("email = ?", strings.TrimSpace(emailAddr)).Updates(map[string]interface{}{
		"password":      hashedPassword,
		"token_version": gorm.Expr("token_version + 1"),
	})
	return result.RowsAffected, result.Error
}

// UpdatePasswordByID 根据ID更新密码并提升token_version
func (r *Repository) UpdatePasswordByID(id, hashedPassword string) (int64, error) {
	result := r.db.Model(&models.User{}).Where("id = ?", strings.TrimSpace(id)).Updates(map[string]interface{}{
		"password":      hashedPassword,
		"token_version": gorm.Expr("token_version + 1"),
	})
	return result.RowsAffected, result.Error
}

// IncrementTokenVersionByID 根据ID提升token_version
func (r *Repository) IncrementTokenVersionByID(id string, requireNormal bool) (int64, error) {
	query := r.db.Model(&models.User{}).Where("id = ?", strings.TrimSpace(id))
	if requireNormal {
		query = query.Where("status = ?", common.UserStatusNormal)
	}

	result := query.Update("token_version", gorm.Expr("token_version + 1"))
	return result.RowsAffected, result.Error
}
