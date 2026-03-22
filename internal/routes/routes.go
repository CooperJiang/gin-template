package routes

import (
	"email-manage/internal/app"
	"email-manage/internal/modules/compat"
	"email-manage/internal/routes/mail_account"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine, deps *app.Dependencies) {
	if deps == nil {
		deps = app.NewDependencies()
	}

	// 注册前端路由
	RegisterClientRoutes(r)

	// 注册文件访问路由（独立路由，不在 /api/v1 下）
	RegisterFileRoutes(r)

	prefix := r.Group("/api")
	version := prefix.Group("/v1")

	// 注册兼容 team-helper 插件的路由（/api 下，不经过 /v1）
	compat.RegisterRoutes(prefix, deps.CompatController)

	{
		// 用户相关路由
		userRoutes := version.Group("/user")
		RegisterUserRoutes(userRoutes, deps.UserController)

		// 上传相关路由
		RegisterUploadRoutes(version, deps.UploadController)

		// 邮箱账号相关路由
		mailAccountRoutes := version.Group("/mail-accounts")
		mail_account.RegisterMailAccountRoutes(mailAccountRoutes, deps.MailAccountController)

		// 在这里添加其他模块路由
		// 例如：
		// productRoutes := api.Group("/product")
		// RegisterProductRoutes(productRoutes)
	}

	// 静态文件服务 - 暂时注释掉，使用embed版本
	// r.Static("/static", "./internal/static")
}
