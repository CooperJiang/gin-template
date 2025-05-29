package routes

import (
	uploadController "template/internal/controllers/upload"
	"template/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterUploadRoutes 注册上传模块路由
func RegisterUploadRoutes(router *gin.RouterGroup) {
	// 上传相关路由组
	upload := router.Group("/upload")
	{
		// 无需认证的路由
		upload.GET("/config", uploadController.GetUploadConfig) // 获取上传配置

		// 需要认证的路由
		authUpload := upload.Group("")
		authUpload.Use(middleware.RequireAuth())
		{
			// 简单上传
			authUpload.POST("/simple", uploadController.SimpleUpload)

			// 分片上传相关
			authUpload.POST("/chunk/init", uploadController.InitChunkUpload) // 初始化分片上传
			authUpload.POST("/chunk", uploadController.UploadChunk)          // 上传分片
			authUpload.POST("/chunk/merge", uploadController.MergeChunks)    // 合并分片

			// 上传进度查询
			authUpload.GET("/progress/:fileID", uploadController.GetUploadProgress)
		}
	}
}
