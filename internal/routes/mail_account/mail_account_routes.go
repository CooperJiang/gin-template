package mail_account

import (
	mailAccountController "email-manage/internal/controllers/mail_account"
	mailAccountModule "email-manage/internal/modules/mail_account"

	"github.com/gin-gonic/gin"
)

// RegisterMailAccountRoutes 注册邮箱账号相关路由
func RegisterMailAccountRoutes(r *gin.RouterGroup, controller *mailAccountController.Controller) {
	mailAccountModule.RegisterRoutes(r, controller)
}

