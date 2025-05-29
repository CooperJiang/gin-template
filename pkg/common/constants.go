package common

import "net/http"

// 验证码类型常量
const (
	// CodeTypeRegister 注册验证码
	CodeTypeRegister = "register"
	// CodeTypeResetPassword 重置密码验证码
	CodeTypeResetPassword = "reset_password"
	// CodeTypeChangeEmail 修改邮箱验证码
	CodeTypeChangeEmail = "change_email"
)

// 系统常量
const (
	// 状态码 (保留用于向后兼容)
	StatusSuccess             = http.StatusOK
	StatusBadRequest          = http.StatusBadRequest
	StatusUnauthorized        = http.StatusUnauthorized
	StatusForbidden           = http.StatusForbidden
	StatusNotFound            = http.StatusNotFound
	StatusServerError         = http.StatusInternalServerError
	StatusInternalServerError = http.StatusInternalServerError

	// 分页默认值
	DefaultPageSize = 10
	MaxPageSize     = 1000

	// 用户状态常量
	UserStatusNormal   = 1 // 正常状态
	UserStatusDisabled = 2 // 禁用状态
	UserStatusDeleted  = 3 // 删除状态

	// 用户角色常量
	UserRoleSuperAdmin = 1 // 超级管理员
	UserRoleAdmin      = 2 // 管理员
	UserRoleUser       = 3 // 普通用户
)

// 日志等级常量
const (
	LogLevelSuccess = 1 // 成功（绿色）
	LogLevelInfo    = 2 // 普通信息（常规）
	LogLevelWarning = 3 // 警告（warning）
	LogLevelError   = 4 // 错误（失败）
)

// 日志类别常量
const (
	LogTypeChatDialog = 1 // chat对话日志
	LogTypeOther      = 2 // 其他类型日志
)
