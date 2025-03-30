package email

import (
	"errors"
	"template/pkg/config"

	"gopkg.in/gomail.v2"
)

var mailer *gomail.Dialer

// Init 初始化邮件服务
func Init() {
	cfg := config.GetConfig().Mail

	// 如果禁用邮件服务，则不初始化
	if !cfg.Enabled {
		return
	}

	// 创建邮件发送器
	mailer = gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
}

// IsMailEnabled 检查邮件服务是否可用
func IsMailEnabled() bool {
	return mailer != nil && config.GetConfig().Mail.Enabled
}

// SendMail 发送邮件
func SendMail(to, subject, body string) error {
	if !IsMailEnabled() {
		return errors.New("邮件服务未启用")
	}

	cfg := config.GetConfig().Mail

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(cfg.From, cfg.FromName))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if cfg.ReplyTo != "" {
		m.SetHeader("Reply-To", cfg.ReplyTo)
	}

	return mailer.DialAndSend(m)
}
