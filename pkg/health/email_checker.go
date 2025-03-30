package health

import (
	"template/pkg/config"
	"template/pkg/email"
)

// EmailChecker 邮件服务健康检查器
type EmailChecker struct{}

// Name 返回检查项名称
func (c *EmailChecker) Name() string {
	return "email"
}

// Check 执行健康检查
func (c *EmailChecker) Check() (Status, map[string]interface{}) {
	// 检查邮件服务是否启用
	if !email.IsMailEnabled() {
		return StatusUp, map[string]interface{}{
			"status":  "disabled",
			"message": "邮件服务未启用，这是正常的",
		}
	}

	// 获取邮件配置
	mailConfig := config.GetConfig().Mail

	details := map[string]interface{}{
		"host":    mailConfig.Host,
		"port":    mailConfig.Port,
		"from":    mailConfig.From,
		"ssl":     mailConfig.SSL,
		"enabled": mailConfig.Enabled,
	}

	// 注意：不执行实际SMTP连接测试，避免发送不必要的邮件
	// 在完整检查中可以考虑添加SMTP连接测试，但不发送邮件

	return StatusUp, details
}

// Type 返回检查类型
func (c *EmailChecker) Type() CheckType {
	return CheckTypeComplete // 不是基础检查
}

// init 注册邮件健康检查器
func init() {
	RegisterChecker(&EmailChecker{})
}
