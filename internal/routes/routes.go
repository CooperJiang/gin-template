package routes

import (
	"template/internal/middleware"
	"template/pkg/health"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	// 应用CORS中间件
	r.Use(middleware.CORSMiddleware())

	// 注册前端路由
	RegisterClientRoutes(r)

	prefix := r.Group("/api")
	version := prefix.Group("/v1")

	// 健康检查
	version.GET("/health", health.SimpleHealthHandler)
	version.GET("/health/basic", health.BasicHealthHandler)
	version.GET("/health/complete", health.CompleteHealthHandler)

	{
		// 用户相关路由
		userRoutes := version.Group("/user")
		RegisterUserRoutes(userRoutes)

		// 在这里添加其他模块路由
		// 例如：
		// productRoutes := api.Group("/product")
		// RegisterProductRoutes(productRoutes)
	}

	// 静态文件服务
	r.Static("/static", "./internal/static")
}
