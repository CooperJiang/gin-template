package routes

import (
	"template/internal/middleware"
	"template/internal/static"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册前端静态文件路由
func RegisterClientRoutes(r *gin.Engine) {
	// 获取嵌入的静态文件系统
	distFS := static.GetDistFS()

	// 使用静态文件处理中间件
	r.Use(middleware.StaticFileHandler(distFS))
}
