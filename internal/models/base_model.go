package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 基础模型，包含所有模型的通用字段
type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`       // 主键ID
	CreatedAt time.Time      `gorm:"not null" json:"created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"not null" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`             // 软删除时间
}

// BeforeCreate 创建前钩子
func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	m.CreatedAt = now
	m.UpdatedAt = now
	return nil
}

// BeforeUpdate 更新前钩子
func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

// IsDeleted 检查记录是否已被软删除
func (m *BaseModel) IsDeleted() bool {
	return m.DeletedAt.Valid
}

// GetCreatedAt 获取创建时间的格式化字符串
func (m *BaseModel) GetCreatedAt() string {
	return m.CreatedAt.Format("2006-01-02 15:04:05")
}

// GetUpdatedAt 获取更新时间的格式化字符串
func (m *BaseModel) GetUpdatedAt() string {
	return m.UpdatedAt.Format("2006-01-02 15:04:05")
}

// GetDeletedAt 获取删除时间的格式化字符串
func (m *BaseModel) GetDeletedAt() string {
	if m.DeletedAt.Valid {
		return m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return ""
}
