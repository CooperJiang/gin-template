package common

import (
	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(StatusSuccess, Response{
		Code:    StatusSuccess,
		Message: message,
		Data:    data,
	})
}

// SuccessWithMessage 成功响应（无数据）
func SuccessWithMessage(c *gin.Context, message string) {
	c.JSON(StatusSuccess, Response{
		Code:    StatusSuccess,
		Message: message,
	})
}

// BadRequest 请求参数错误
func BadRequest(c *gin.Context, message string) {
	c.JSON(StatusBadRequest, Response{
		Code:    StatusBadRequest,
		Message: message,
	})
}

// Unauthorized 未授权
func Unauthorized(c *gin.Context, message string) {
	c.JSON(StatusUnauthorized, Response{
		Code:    StatusUnauthorized,
		Message: message,
	})
}

// Forbidden 禁止访问
func Forbidden(c *gin.Context, message string) {
	c.JSON(StatusForbidden, Response{
		Code:    StatusForbidden,
		Message: message,
	})
}

// NotFound 资源不存在
func NotFound(c *gin.Context, message string) {
	c.JSON(StatusNotFound, Response{
		Code:    StatusNotFound,
		Message: message,
	})
}

// ServerError 服务器错误
func ServerError(c *gin.Context, message string) {
	c.JSON(StatusServerError, Response{
		Code:    StatusServerError,
		Message: message,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}
