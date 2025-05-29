package routes

import (
	"template/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	// 应用CORS中间件
	r.Use(middleware.CORSMiddleware())

	// 注册前端路由
	RegisterClientRoutes(r)

	// 注册文件访问路由（独立路由，不在 /api/v1 下）
	RegisterFileRoutes(r)

	prefix := r.Group("/api")
	version := prefix.Group("/v1")

	{
		// 用户相关路由
		userRoutes := version.Group("/user")
		RegisterUserRoutes(userRoutes)

		// 上传相关路由
		RegisterUploadRoutes(version)

		// 在这里添加其他模块路由
		// 例如：
		// productRoutes := api.Group("/product")
		// RegisterProductRoutes(productRoutes)
	}

	// 静态文件服务
	r.Static("/static", "./internal/static")
}
