package common

import (
	"net/http"
	"template/pkg/errors"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构 (保留旧版兼容性)
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}, message string) {
	errors.ResponseSuccess(c, data, message)
}

// SuccessWithMessage 成功响应（无数据）
func SuccessWithMessage(c *gin.Context, message string) {
	errors.ResponseSuccess(c, nil, message)
}

// BadRequest 请求参数错误
func BadRequest(c *gin.Context, message string) {
	err := errors.New(errors.CodeInvalidParameter, message)
	errors.HandleError(c, err)
}

// Unauthorized 未授权
func Unauthorized(c *gin.Context, message string) {
	err := errors.New(errors.CodeUnauthorized, message)
	errors.HandleError(c, err)
}

// Forbidden 禁止访问
func Forbidden(c *gin.Context, message string) {
	err := errors.New(errors.CodeForbidden, message)
	errors.HandleError(c, err)
}

// NotFound 资源不存在
func NotFound(c *gin.Context, message string) {
	err := errors.New(errors.CodeNotFound, message)
	errors.HandleError(c, err)
}

// ServerError 服务器错误
func ServerError(c *gin.Context, message string) {
	err := errors.New(errors.CodeInternal, message)
	errors.HandleError(c, err)
}

// Fail 失败响应
func Fail(c *gin.Context, code int, message string) {
	// 从旧的HTTP状态码映射到新的错误码系统
	var errorCode errors.ErrorCode
	switch code {
	case http.StatusBadRequest:
		errorCode = errors.CodeInvalidParameter
	case http.StatusUnauthorized:
		errorCode = errors.CodeUnauthorized
	case http.StatusForbidden:
		errorCode = errors.CodeForbidden
	case http.StatusNotFound:
		errorCode = errors.CodeNotFound
	case http.StatusInternalServerError:
		errorCode = errors.CodeInternal
	default:
		errorCode = errors.ErrorCode(code)
	}

	err := errors.New(errorCode, message)
	errors.HandleError(c, err)
}

// HandleError 处理错误并返回响应
func HandleError(c *gin.Context, err error) {
	errors.HandleError(c, err)
}
