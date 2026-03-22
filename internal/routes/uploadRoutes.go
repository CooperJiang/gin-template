package routes

import (
	uploadController "email-manage/internal/controllers/upload"
	uploadModule "email-manage/internal/modules/upload"

	"github.com/gin-gonic/gin"
)

// RegisterUploadRoutes 注册上传模块路由
func RegisterUploadRoutes(router *gin.RouterGroup, controller *uploadController.Controller) {
	uploadModule.RegisterRoutes(router, controller)
}
