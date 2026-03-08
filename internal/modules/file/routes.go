package file

import (
	fileController "template/internal/controllers/file"
	"template/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册文件访问路由
func RegisterRoutes(router *gin.Engine) {
	files := router.Group("/files")
	{
		authFiles := files.Group("")
		authFiles.Use(middleware.FileAuthMiddleware())
		{
			authFiles.GET("/download/:fileId", fileController.DownloadFile)
			authFiles.GET("/preview/:fileId", fileController.PreviewFile)
			authFiles.GET("/download/:fileId/:filename", fileController.DownloadFile)
			authFiles.GET("/preview/:fileId/:filename", fileController.PreviewFile)
		}
	}
}
