package routes

import (
	fileModule "email-manage/internal/modules/file"

	"github.com/gin-gonic/gin"
)

// RegisterFileRoutes 注册文件访问路由
func RegisterFileRoutes(router *gin.Engine) {
	fileModule.RegisterRoutes(router)
}
