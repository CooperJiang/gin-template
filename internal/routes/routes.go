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

	RegisterClientRoutes(r)

	// 健康检查
	r.GET("/health", health.SimpleHealthHandler)
	r.GET("/health/basic", health.BasicHealthHandler)
	r.GET("/health/complete", health.CompleteHealthHandler)

	// API路由组
	api := r.Group("/api")
	{
		// 用户相关路由
		userRoutes := api.Group("/user")
		RegisterUserRoutes(userRoutes)

		// 在这里添加其他模块路由
		// 例如：
		// productRoutes := api.Group("/product")
		// RegisterProductRoutes(productRoutes)
	}

	// 静态文件服务
	r.Static("/static", "./internal/static")
}
