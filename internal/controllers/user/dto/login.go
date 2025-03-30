package dto

// LoginDTO 登录请求
type LoginDTO struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// GetValidationMessages 获取验证错误信息
func (r *LoginDTO) GetValidationMessages() map[string]string {
	return map[string]string{
		"Account.required":  "请输入账户名或邮箱",
		"Password.required": "请输入密码",
	}
}
