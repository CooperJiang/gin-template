package routes

import (
	"net/http"
	"template/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	// 应用CORS中间件
	r.Use(middleware.CORSMiddleware())

	// 健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "服务运行正常",
			"service": "gin-template",
		})
	})

	// 注册前端路由
	RegisterClientRoutes(r)

	prefix := r.Group("/api")
	version := prefix.Group("/v1")

	// API健康检查路由
	version.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "API服务运行正常",
			"service": "gin-template",
			"version": "v1",
		})
	})

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
