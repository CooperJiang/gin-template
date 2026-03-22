package errors

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"email-manage/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 响应结构体
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	RequestID string      `json:"request_id,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

// ErrorHandler 是一个中间件，用于统一处理错误响应
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成并设置请求ID
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)
		c.Header("X-Request-ID", requestID)

		// 设置开始时间
		start := time.Now()

		// 处理可能的panic
		defer func() {
			if r := recover(); r != nil {
				// 记录日志
				stackTrace := string(debug.Stack())
				logger.Error("[PANIC] %v\nTrace: %s", r, stackTrace)

				// 创建错误响应
				err := &Error{
					Code:      CodeInternal,
					Message:   "服务器内部错误",
					Detail:    fmt.Sprintf("%v", r),
					Stack:     stackTrace,
					Time:      time.Now(),
					RequestID: requestID,
				}

				// 响应错误
				responseError(c, err)
				c.Abort()
			}
		}()

		// 调用下一个处理器
		c.Next()

		// 计算请求处理时间
		latency := time.Since(start)

		// 记录响应日志
		logger.Info("[RESP] [%s] %s %s - %d (%dms)",
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency.Milliseconds(),
		)
	}
}

// 辅助函数：响应错误
func responseError(c *gin.Context, err error) {
	var apiErr *Error
	var statusCode int

	if e, ok := err.(*Error); ok {
		apiErr = e
		statusCode = HTTPStatus(e)
	} else {
		apiErr = &Error{
			Code:    CodeInternal,
			Message: "服务器内部错误",
			Detail:  err.Error(),
			Time:    time.Now(),
		}
		statusCode = http.StatusInternalServerError
	}

	// 获取请求ID
	if requestID := getRequestID(c); requestID != "" {
		apiErr.RequestID = requestID
	}

	// 构建安全的响应，优先使用 Detail 作为更具体的消息
	msg := apiErr.Message
	if apiErr.Detail != "" {
		msg = apiErr.Detail
	}
	response := Response{
		Code:      int(apiErr.Code),
		Message:   msg,
		RequestID: apiErr.RequestID,
		Timestamp: time.Now().Unix(),
	}

	c.JSON(statusCode, response)
}

// ResponseSuccess 成功响应
func ResponseSuccess(c *gin.Context, data interface{}, message string) {
	response := Response{
		Code:      200,
		Message:   message,
		Data:      data,
		RequestID: getRequestID(c),
		Timestamp: time.Now().Unix(),
	}
	c.JSON(http.StatusOK, response)
}

// HandleError 处理错误并返回响应
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	responseError(c, err)
}

func getRequestID(c *gin.Context) string {
	if c == nil {
		return ""
	}

	value, exists := c.Get("RequestID")
	if !exists {
		return ""
	}

	requestID, ok := value.(string)
	if !ok {
		return ""
	}

	return requestID
}
