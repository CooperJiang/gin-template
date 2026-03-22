package mail_account

import (
	mailAccountController "email-manage/internal/controllers/mail_account"
	"email-manage/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册邮箱账号相关路由
func RegisterRoutes(r *gin.RouterGroup, controller *mailAccountController.Controller) {
	auth := r.Group("")
	auth.Use(middleware.RequireAuth())
	{
		// 固定路径路由必须在 /:id 之前注册，避免被参数路由覆盖
		auth.GET("/stats", controller.Stats)
		auth.GET("/daily-usage", controller.DailyUsage)
		auth.POST("/import", controller.Import)
		auth.GET("/", controller.List)
		auth.GET("/:id", controller.GetByID)
		auth.POST("/", controller.Create)
		auth.PUT("/:id", controller.Update)
		auth.PATCH("/:id/status", controller.UpdateStatus)
		auth.DELETE("/:id", controller.Delete)
		auth.POST("/batch-delete", controller.BatchDelete)
		auth.POST("/:id/fetch-mails", controller.FetchMails)
		auth.POST("/:id/fetch-code", controller.FetchCodeByID)
	}

	// 对外开放接口已迁移到 compat 路由（/api/generate, /api/emails, /api/email/:id）
	// 此处不再注册无鉴权接口，避免邮箱被意外消耗
}
