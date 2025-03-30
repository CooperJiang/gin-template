package dto

// SendCodeDTO 发送验证码请求
type SendCodeDTO struct {
	Email string `json:"email" binding:"required,email"`
}

// GetValidationMessages 获取验证消息
func (d *SendCodeDTO) GetValidationMessages() map[string]string {
	return map[string]string{
		"Email.required": "邮箱不能为空",
		"Email.email":    "邮箱格式不正确",
	}
}

// ResetPasswordDTO 重置密码请求
type ResetPasswordDTO struct {
	Email       string `json:"email" binding:"required,email"`
	Code        string `json:"code" binding:"required,len=6"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=20"`
}

// GetValidationMessages 获取验证消息
func (d *ResetPasswordDTO) GetValidationMessages() map[string]string {
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
