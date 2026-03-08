package routes

import (
	userController "template/internal/controllers/user"
	userModule "template/internal/modules/user"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes 注册用户相关路由
func RegisterUserRoutes(r *gin.RouterGroup, controller *userController.Controller) {
	userModule.RegisterRoutes(r, controller)
}
