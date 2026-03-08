package upload

import (
	uploadController "template/internal/controllers/upload"
	"template/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册上传模块路由
func RegisterRoutes(router *gin.RouterGroup, controller *uploadController.Controller) {
	if controller == nil {
		controller = uploadController.NewController(nil)
	}

	upload := router.Group("/upload")
	{
		upload.GET("/config", controller.GetUploadConfig)

		authUpload := upload.Group("")
		authUpload.Use(middleware.RequireAuth())
		{
			authUpload.POST("/simple", controller.SimpleUpload)
			authUpload.POST("/chunk/init", controller.InitChunkUpload)
			authUpload.POST("/chunk", controller.UploadChunk)
			authUpload.POST("/chunk/merge", controller.MergeChunks)
			authUpload.GET("/progress/:fileID", controller.GetUploadProgress)
		}
	}
}
