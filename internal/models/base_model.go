package models

import (
	"template/pkg/common"
	"time"

	"gorm.io/gorm"
)

// BaseModel 基础模型，所有模型都应该嵌入此结构
type BaseModel struct {
	ID        common.UUID     `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt common.JSONTime `json:"created_at"`
	UpdatedAt common.JSONTime `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"`
}

// BeforeCreate 创建前的钩子，自动生成UUID
func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID.IsZero() {
		m.ID = common.NewUUID()
	}
	return nil
}

// GetID 获取ID字符串
func (m *BaseModel) GetID() string {
	return m.ID.String()
}

// SetID 设置ID
func (m *BaseModel) SetID(id string) error {
	uuid, err := common.ParseUUID(id)
	if err != nil {
		return err
	}
	m.ID = uuid
	return nil
}

// IsNew 判断是否为新记录
func (m *BaseModel) IsNew() bool {
	return m.ID.IsZero()
}

// GetCreatedAt 获取创建时间
func (m *BaseModel) GetCreatedAt() time.Time {
	return time.Time(m.CreatedAt)
}

// GetUpdatedAt 获取更新时间
func (m *BaseModel) GetUpdatedAt() time.Time {
	return time.Time(m.UpdatedAt)
}

// BeforeUpdate 更新前钩子
func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = common.JSONTime(time.Now())
	return nil
}

// IsDeleted 检查记录是否已被软删除
func (m *BaseModel) IsDeleted() bool {
	return m.DeletedAt.Valid
}

// GetDeletedAt 获取删除时间的格式化字符串
func (m *BaseModel) GetDeletedAt() string {
	if m.DeletedAt.Valid {
		return m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return ""
}
