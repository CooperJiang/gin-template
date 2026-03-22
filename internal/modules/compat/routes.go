package compat

import (
	compatController "email-manage/internal/controllers/compat"
	"email-manage/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 在 /api 前缀下注册兼容 team-helper 插件的路由
func RegisterRoutes(api *gin.RouterGroup, ctrl *compatController.Controller) {
	// 登录不需要鉴权
	api.POST("/login", ctrl.Login)

	// 以下接口需要 Cookie 鉴权
	authed := api.Group("", middleware.RequireCookieAuth())
	{
		authed.GET("/generate", ctrl.Generate)
		authed.GET("/user/quota", ctrl.UserQuota)
		authed.POST("/mailboxes/toggle-login", ctrl.ToggleLogin)
		authed.GET("/emails", ctrl.ListEmails)
		authed.GET("/email/:id", ctrl.GetEmail)
	}
}
