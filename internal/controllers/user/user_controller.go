package user

import (
	"template/internal/dto/request"
	"template/internal/dto/response"
	"template/internal/services/user"
	"template/pkg/common"
	"template/pkg/errors"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	// 验证请求
	req, err := common.ValidateRequest[request.RegisterRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := user.RegisterUser(req.Username, req.Email, req.Password, req.Code); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "注册成功")
}

// Login 用户登录
func Login(c *gin.Context) {
	// 验证请求
	req, err := common.ValidateRequest[request.LoginRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	userInfo, token, err := user.Login(req.Account, req.Password)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	resp := response.LoginResponse{
		Token: token,
		User: response.UserInfo{
			Username: userInfo["username"].(string),
			Email:    userInfo["email"].(string),
			Avatar:   userInfo["avatar"].(string),
			Status:   userInfo["status"].(int),
		},
	}

	errors.ResponseSuccess(c, resp, "登录成功")
}

// SendRegistrationCode 发送注册验证码
func SendRegistrationCode(c *gin.Context) {
	req, err := common.ValidateRequest[request.SendCodeRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := user.SendRegistrationCode(req.Email); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "验证码已发送")
}

// SendResetPasswordCode 发送重置密码验证码
func SendResetPasswordCode(c *gin.Context) {
	req, err := common.ValidateRequest[request.SendCodeRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := user.SendResetPasswordCode(req.Email); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "验证码已发送")
}

// ResetPassword 重置密码
func ResetPassword(c *gin.Context) {
	req, err := common.ValidateRequest[request.ResetPasswordRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := user.ResetPassword(req.Email, req.Code, req.NewPassword); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "密码重置成功")
}
