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

	userInfo, token, expiresAt, err := user.Login(req.Account, req.Password)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	// 安全地获取用户信息
	username, _ := userInfo["username"].(string)
	email, _ := userInfo["email"].(string)
	avatar, _ := userInfo["avatar"].(string)
	status, _ := userInfo["status"].(int)

	resp := response.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User: response.UserInfo{
			Username: username,
			Email:    email,
			Avatar:   avatar,
			Status:   status,
		},
	}

	errors.ResponseSuccess(c, resp, "登录成功")
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	// 从中间件获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		errors.HandleError(c, errors.New(errors.CodeUnauthorized, "未授权"))
		return
	}

	userInfo, err := user.GetUserInfo(userID.(string))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, userInfo, "获取用户信息成功")
}

// UpdateProfile 更新用户资料
func UpdateProfile(c *gin.Context) {
	// 从中间件获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		errors.HandleError(c, errors.New(errors.CodeUnauthorized, "未授权"))
		return
	}

	req, err := common.ValidateRequest[request.UpdateProfileRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	userInfo, err := user.UpdateProfile(userID.(string), req.Username, req.Email, req.Avatar, req.Code)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, userInfo, "更新资料成功")
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	// 从中间件获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		errors.HandleError(c, errors.New(errors.CodeUnauthorized, "未授权"))
		return
	}

	req, err := common.ValidateRequest[request.ChangePasswordRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := user.ChangePassword(userID.(string), req.OldPassword, req.NewPassword); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "密码修改成功")
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

// SendChangeEmailCode 发送修改邮箱验证码
func SendChangeEmailCode(c *gin.Context) {
	// 从中间件获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		errors.HandleError(c, errors.New(errors.CodeUnauthorized, "未授权"))
		return
	}

	req, err := common.ValidateRequest[request.SendCodeRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := user.SendChangeEmailCode(userID.(string), req.Email); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "验证码已发送")
}
