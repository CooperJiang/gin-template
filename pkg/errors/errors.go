package errors

import (
	"fmt"
	"runtime"
	"strings"
	"template/pkg/logger"
	"time"
)

// ErrorCode 错误码类型
type ErrorCode int

// 错误类型常量
const (
	// 通用错误码（1-999）
	CodeUnknown            ErrorCode = 1   // 未知错误
	CodeInternal           ErrorCode = 2   // 内部错误
	CodeInvalidParameter   ErrorCode = 100 // 无效参数
	CodeUnauthorized       ErrorCode = 101 // 未授权
	CodeForbidden          ErrorCode = 102 // 禁止访问
	CodeNotFound           ErrorCode = 103 // 未找到
	CodeMethodNotAllowed   ErrorCode = 104 // 方法不允许
	CodeTimeout            ErrorCode = 105 // 超时
	CodeConflict           ErrorCode = 106 // 冲突
	CodeRateLimited        ErrorCode = 107 // 频率限制
	CodeValidationFailed   ErrorCode = 108 // 验证失败
	CodeServiceUnavailable ErrorCode = 109 // 服务不可用

	// 用户相关错误码（1000-1999）
	CodeUserNotFound      ErrorCode = 1000 // 用户不存在
	CodeWrongPassword     ErrorCode = 1001 // 密码错误
	CodeUserDisabled      ErrorCode = 1002 // 用户被禁用
	CodeUserExists        ErrorCode = 1003 // 用户已存在
	CodeInvalidAuthToken  ErrorCode = 1004 // 无效的认证令牌
	CodeExpiredAuthToken  ErrorCode = 1005 // 认证令牌过期
	CodeInvalidVerifyCode ErrorCode = 1006 // 无效的验证码
	CodeEmailExists       ErrorCode = 1007 // 邮箱已存在
	CodeEmailSendFailed   ErrorCode = 1008 // 邮件发送失败

	// 数据库相关错误码（2000-2999）
	CodeDBConnectionFailed ErrorCode = 2000 // 数据库连接失败
	CodeQueryFailed        ErrorCode = 2001 // 查询失败
	CodeDBNoRecord         ErrorCode = 2002 // 记录不存在
	CodeDBDuplicate        ErrorCode = 2003 // 重复记录
	CodeDBTransaction      ErrorCode = 2004 // 事务错误

	// 第三方服务错误码（3000-3999）
	CodeThirdPartyService ErrorCode = 3000 // 第三方服务错误
	CodeRedisError        ErrorCode = 3001 // Redis错误
	CodeEmailServiceError ErrorCode = 3002 // 邮件服务错误
	CodeSMSServiceError   ErrorCode = 3003 // 短信服务错误
	CodePaymentError      ErrorCode = 3004 // 支付服务错误
)

// 错误码对应的HTTP状态码
var errorCodeToHTTPStatus = map[ErrorCode]int{
	CodeUnknown:            500,
	CodeInternal:           500,
	CodeInvalidParameter:   400,
	CodeUnauthorized:       401,
	CodeForbidden:          403,
	CodeNotFound:           404,
	CodeMethodNotAllowed:   405,
	CodeTimeout:            408,
	CodeConflict:           409,
	CodeRateLimited:        429,
	CodeValidationFailed:   400,
	CodeServiceUnavailable: 503,

	CodeUserNotFound:      404,
	CodeWrongPassword:     400,
	CodeUserDisabled:      403,
	CodeUserExists:        409,
	CodeInvalidAuthToken:  401,
	CodeExpiredAuthToken:  401,
	CodeInvalidVerifyCode: 400,
	CodeEmailExists:       409,
	CodeEmailSendFailed:   500,

	CodeDBConnectionFailed: 500,
	CodeQueryFailed:        500,
	CodeDBNoRecord:         404,
	CodeDBDuplicate:        409,
	CodeDBTransaction:      500,

	CodeThirdPartyService: 500,
	CodeRedisError:        500,
	CodeEmailServiceError: 500,
	CodeSMSServiceError:   500,
	CodePaymentError:      500,
}

// 错误码对应的错误消息（面向用户友好的消息）
var errorCodeToMessage = map[ErrorCode]string{
	CodeUnknown:            "服务器遇到了未知错误",
	CodeInternal:           "服务器内部错误",
	CodeInvalidParameter:   "无效的请求参数",
	CodeUnauthorized:       "请先登录",
	CodeForbidden:          "没有操作权限",
	CodeNotFound:           "请求的资源不存在",
	CodeMethodNotAllowed:   "不支持的请求方法",
	CodeTimeout:            "请求超时",
	CodeConflict:           "资源冲突",
	CodeRateLimited:        "请求频率过高，请稍后再试",
	CodeValidationFailed:   "数据验证失败",
	CodeServiceUnavailable: "服务暂时不可用，请稍后再试",

	CodeUserNotFound:      "用户不存在",
	CodeWrongPassword:     "密码错误",
	CodeUserDisabled:      "账号已被禁用",
	CodeUserExists:        "用户已存在",
	CodeInvalidAuthToken:  "登录凭证无效，请重新登录",
	CodeExpiredAuthToken:  "登录已过期，请重新登录",
	CodeInvalidVerifyCode: "验证码无效或已过期",
	CodeEmailExists:       "邮箱已被注册",
	CodeEmailSendFailed:   "邮件发送失败，请稍后再试",

	CodeDBConnectionFailed: "数据库连接失败",
	CodeQueryFailed:        "数据查询失败",
	CodeDBNoRecord:         "记录不存在",
	CodeDBDuplicate:        "记录已存在",
	CodeDBTransaction:      "数据库事务错误",

	CodeThirdPartyService: "第三方服务异常",
	CodeRedisError:        "缓存服务异常",
	CodeEmailServiceError: "邮件服务异常",
	CodeSMSServiceError:   "短信服务异常",
	CodePaymentError:      "支付服务异常",
}

// Error 自定义错误类型
type Error struct {
	Code      ErrorCode              // 错误码
	Message   string                 // 错误信息（用户可见）
	Detail    string                 // 详细错误信息（仅日志记录，不返回给用户）
	Stack     string                 // 错误堆栈
	Time      time.Time              // 错误发生时间
	RequestID string                 // 请求ID，用于跟踪
	Metadata  map[string]interface{} // 额外元数据
}

// Error 实现error接口
func (e *Error) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("[%d] %s: %s", e.Code, e.Message, e.Detail)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// New 创建一个新的错误
func New(code ErrorCode, detail string) *Error {
	message, ok := errorCodeToMessage[code]
	if !ok {
		message = "未知错误"
	}

	err := &Error{
		Code:     code,
		Message:  message,
		Detail:   detail,
		Time:     time.Now(),
		Metadata: make(map[string]interface{}),
	}

	// 获取调用堆栈
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(2, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])

	var stackBuilder strings.Builder
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			// 跳过runtime内部调用
			fmt.Fprintf(&stackBuilder, "%s:%d - %s\n", frame.File, frame.Line, frame.Function)
		}
		if !more {
			break
		}
	}
	err.Stack = stackBuilder.String()

	// 记录错误日志
	logger.Error("[ERROR] %s\nStack: %s", err.Error(), err.Stack)

	return err
}

// WithRequestID 添加请求ID
func (e *Error) WithRequestID(requestID string) *Error {
	e.RequestID = requestID
	return e
}

// WithMetadata 添加元数据
func (e *Error) WithMetadata(key string, value interface{}) *Error {
	e.Metadata[key] = value
	return e
}

// Is 判断错误是否为指定错误码
func Is(err error, code ErrorCode) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*Error); ok {
		return e.Code == code
	}
	return false
}

// HTTPStatus 获取错误对应的HTTP状态码
func HTTPStatus(err error) int {
	if err == nil {
		return 200
	}
	if e, ok := err.(*Error); ok {
		if status, exists := errorCodeToHTTPStatus[e.Code]; exists {
			return status
		}
	}
	return 500
}

// Wrap 包装标准错误为自定义错误
func Wrap(err error, code ErrorCode) *Error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok {
		return e
	}
	return New(code, err.Error())
}

// NewValidationError 创建验证错误
func NewValidationError(field, detail string) *Error {
	err := New(CodeValidationFailed, detail)
	err.WithMetadata("field", field)
	return err
}

// GetSafeError 获取安全的错误（移除敏感信息）
func GetSafeError(err error) *Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*Error); ok {
		// 创建一个新的错误对象，但不包含详细信息和堆栈
		return &Error{
			Code:      e.Code,
			Message:   e.Message,
			Time:      e.Time,
			RequestID: e.RequestID,
		}
	}

	// 标准错误转换为通用错误
	return &Error{
		Code:    CodeInternal,
		Message: errorCodeToMessage[CodeInternal],
		Time:    time.Now(),
	}
}
