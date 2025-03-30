package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// ValidationError 验证错误
type ValidationError struct {
	Field   string
	Message string
}

// ValidateRequest 通用的请求验证函数
func ValidateRequest[T any](c *gin.Context) (*T, *ValidationError) {
	var req T

	// 如果请求结构体实现了 GetValidationMessages 接口，则获取验证消息
	var messages map[string]string
	if v, ok := any(&req).(interface{ GetValidationMessages() map[string]string }); ok {
		messages = v.GetValidationMessages()
	}

	// 绑定并验证请求
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// 构建错误信息
			for _, validationErr := range validationErrors {
				field := validationErr.StructField()
				tag := validationErr.Tag()
				msgKey := field + "." + tag
				if msg, exists := messages[msgKey]; exists {
					return nil, &ValidationError{
						Field:   field,
						Message: msg,
					}
				}
			}
			return nil, &ValidationError{
				Field:   "",
				Message: "请求参数验证失败",
			}
		}
		return nil, &ValidationError{
			Field:   "",
			Message: "无效的请求参数",
		}
	}

	return &req, nil
} 