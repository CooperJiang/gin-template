package routes

import (
	uploadController "template/internal/controllers/upload"
	uploadModule "template/internal/modules/upload"

	"github.com/gin-gonic/gin"
)

// RegisterUploadRoutes 注册上传模块路由
func RegisterUploadRoutes(router *gin.RouterGroup, controller *uploadController.Controller) {
	uploadModule.RegisterRoutes(router, controller)
}
