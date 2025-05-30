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
			// 支持两种URL格式：
			// 1. 简单格式：/files/download/uuid 和 /files/preview/uuid
			authFiles.GET("/download/:fileId", fileController.DownloadFile)
			authFiles.GET("/preview/:fileId", fileController.PreviewFile)

			// 2. 带文件名格式：/files/download/uuid/filename.ext 和 /files/preview/uuid/filename.ext
			authFiles.GET("/download/:fileId/:filename", fileController.DownloadFile)
			authFiles.GET("/preview/:fileId/:filename", fileController.PreviewFile)
		}
	}
}
