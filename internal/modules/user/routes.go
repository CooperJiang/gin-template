package user

import (
	userController "template/internal/controllers/user"
	"template/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册用户模块路由
func RegisterRoutes(r *gin.RouterGroup, controller *userController.Controller) {
	if controller == nil {
		controller = userController.NewController(nil)
	}

	// 公开路由
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/refresh-token", controller.RefreshToken)
	r.POST("/send-registration-code", controller.SendRegistrationCode)
	r.POST("/send-reset-password-code", controller.SendResetPasswordCode)
	r.POST("/reset-password", controller.ResetPassword)

	// 需要认证的路由
	protected := r.Group("")
	protected.Use(middleware.RequireAuth())
	{
		protected.GET("/info", controller.GetUserInfo)
		protected.PUT("/profile", controller.UpdateProfile)
		protected.POST("/change-password", controller.ChangePassword)
		protected.POST("/logout", controller.Logout)
		protected.POST("/send-change-email-code", controller.SendChangeEmailCode)
	}

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
