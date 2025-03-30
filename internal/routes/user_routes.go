package routes

import (
	userController "template/internal/controllers/user"
	"template/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes 注册用户相关路由
func RegisterUserRoutes(r *gin.RouterGroup) {
	// 公开路由
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.POST("/send-registration-code", userController.SendRegistrationCode)
	r.POST("/send-reset-password-code", userController.SendResetPasswordCode)
	r.POST("/reset-password", userController.ResetPassword)

	// 用户路由组，需要登录才能访问
	userGroup := r.Group("/user")
	userGroup.Use(middleware.RequireAuth())
	{
		// 在这里添加需要登录的用户接口
	}

	// 管理员路由组
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.RequireAdmin())
	{
		// 在这里添加管理员接口
	}

	// 超级管理员路由组
	superGroup := r.Group("/super")
	superGroup.Use(middleware.RequireSuperAuth())
	{
		// 在这里添加超级管理员接口
	}
}
