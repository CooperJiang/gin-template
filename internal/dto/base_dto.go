package dto

import "time"

// BaseRequest 基础请求DTO
type BaseRequest struct {
	RequestID string `json:"request_id,omitempty"` // 请求ID
}

// BaseResponse 基础响应DTO
type BaseResponse struct {
	ID        uint      `json:"id"`         // 主键ID
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

// ValidationMessages 验证消息接口
type ValidationMessages interface {
	GetValidationMessages() map[string]string
}

// Convertible 可转换接口
type Convertible[T any] interface {
	ToModel() T
}

// FromModel 从模型转换接口
type FromModel[T any] interface {
	FromModel(model T) interface{}
}
