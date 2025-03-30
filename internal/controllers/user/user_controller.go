package user

import (
	"template/internal/controllers/user/dto"
	"template/internal/services/user"
	"template/pkg/common"
	"template/pkg/errors"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	// 验证请求
	req, err := common.ValidateRequest[dto.RegisterDTO](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	// 调用服务层进行注册
	if err := user.RegisterUser(req.Username, req.Email, req.Password, req.Code); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "注册成功")
}

// Login 用户登录
func Login(c *gin.Context) {
	// 验证请求
	req, err := common.ValidateRequest[dto.LoginDTO](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	// 调用服务层进行登录验证
	userInfo, token, err := user.Login(req.Account, req.Password)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	data := gin.H{
		"token":    token,
		"userInfo": userInfo,
	}

	errors.ResponseSuccess(c, data, "登录成功")
}

// SendRegistrationCode 发送注册验证码
func SendRegistrationCode(c *gin.Context) {
	// 验证请求
	req, err := common.ValidateRequest[dto.SendCodeDTO](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	// 发送验证码
	if err := user.SendRegistrationCode(req.Email); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "验证码已发送")
}

// SendResetPasswordCode 发送重置密码验证码
func SendResetPasswordCode(c *gin.Context) {
	// 验证请求
	req, err := common.ValidateRequest[dto.SendCodeDTO](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	// 发送验证码
	if err := user.SendResetPasswordCode(req.Email); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "验证码已发送")
}

// ResetPassword 重置密码
func ResetPassword(c *gin.Context) {
	// 验证请求
	req, err := common.ValidateRequest[dto.ResetPasswordDTO](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	// 重置密码
	if err := user.ResetPassword(req.Email, req.Code, req.NewPassword); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "密码重置成功")
}
