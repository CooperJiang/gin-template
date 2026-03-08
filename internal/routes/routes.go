package routes

import (
	"template/internal/app"

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

	{
		// 用户相关路由
		userRoutes := version.Group("/user")
		RegisterUserRoutes(userRoutes, deps.UserController)

		// 上传相关路由
		RegisterUploadRoutes(version, deps.UploadController)

		// 在这里添加其他模块路由
		// 例如：
		// productRoutes := api.Group("/product")
		// RegisterProductRoutes(productRoutes)
	}

	// 静态文件服务 - 暂时注释掉，使用embed版本
	// r.Static("/static", "./internal/static")
}
