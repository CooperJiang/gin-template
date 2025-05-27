package request

import "template/internal/dto"

// LoginRequest 登录请求
type LoginRequest struct {
	dto.BaseRequest
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// GetValidationMessages 获取验证错误信息
func (r *LoginRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"Account.required":  "请输入账户名或邮箱",
		"Password.required": "请输入密码",
	}
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	dto.BaseRequest
	Username string `json:"username" binding:"required,min=2,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Code     string `json:"code" binding:"required"`
}

// GetValidationMessages 获取验证错误信息
func (r *RegisterRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"Username.required": "用户名不能为空",
		"Username.min":      "用户名长度不能小于2个字符",
		"Username.max":      "用户名长度不能超过20个字符",
		"Email.required":    "邮箱不能为空",
		"Email.email":       "请输入有效的邮箱地址",
		"Password.required": "密码不能为空",
		"Password.min":      "密码长度不能小于6个字符",
		"Password.max":      "密码长度不能超过20个字符",
		"Code.required":     "验证码不能为空",
	}
}

// SendCodeRequest 发送验证码请求
type SendCodeRequest struct {
	dto.BaseRequest
	Email string `json:"email" binding:"required,email"`
}

// GetValidationMessages 获取验证消息
func (r *SendCodeRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"Email.required": "邮箱不能为空",
		"Email.email":    "邮箱格式不正确",
	}
}

// ResetPasswordRequest 重置密码请求
type ResetPasswordRequest struct {
	dto.BaseRequest
	Email       string `json:"email" binding:"required,email"`
	Code        string `json:"code" binding:"required,len=6"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=20"`
}

// GetValidationMessages 获取验证消息
func (r *ResetPasswordRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"Email.required":       "邮箱不能为空",
		"Email.email":          "邮箱格式不正确",
		"Code.required":        "验证码不能为空",
		"Code.len":             "验证码长度必须为6位",
		"NewPassword.required": "新密码不能为空",
		"NewPassword.min":      "密码长度不能小于6位",
		"NewPassword.max":      "密码长度不能大于20位",
	}
}

// UpdateProfileRequest 更新用户资料请求
type UpdateProfileRequest struct {
	dto.BaseRequest
	Username string `json:"username,omitempty" binding:"omitempty,min=2,max=20"`
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
	Avatar   string `json:"avatar,omitempty"`
}

// GetValidationMessages 获取验证消息
func (r *UpdateProfileRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"Username.min": "用户名长度不能小于2个字符",
		"Username.max": "用户名长度不能超过20个字符",
		"Email.email":  "请输入有效的邮箱地址",
	}
}

// ChangePasswordRequest 修改密码请求
type ChangePasswordRequest struct {
	dto.BaseRequest
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=20"`
}

// GetValidationMessages 获取验证消息
func (r *ChangePasswordRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"OldPassword.required": "原密码不能为空",
		"NewPassword.required": "新密码不能为空",
		"NewPassword.min":      "密码长度不能小于6位",
		"NewPassword.max":      "密码长度不能大于20位",
	}
}
