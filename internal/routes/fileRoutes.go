package routes

import (
	fileController "template/internal/controllers/file"
	"template/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterFileRoutes 注册文件访问路由
func RegisterFileRoutes(router *gin.Engine) {
	// 文件访问路由组 - 使用独立的 /files 路由
	files := router.Group("/files")
	{
		// 需要权限验证的文件访问
		authFiles := files.Group("")
		authFiles.Use(middleware.FileAuthMiddleware())
		{
			authFiles.GET("/download/:fileId", fileController.DownloadFile)     // 下载文件
			authFiles.GET("/preview/:fileId", fileController.PreviewFile)      // 预览文件
		}
	}
}
