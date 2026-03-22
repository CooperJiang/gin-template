package user

import (
	"email-manage/internal/dto/request"
	"email-manage/internal/dto/response"
	"email-manage/internal/middleware"
	userService "email-manage/internal/services/user"
	"email-manage/pkg/common"
	"email-manage/pkg/errors"

	"github.com/gin-gonic/gin"
)

// Controller 用户控制器
type Controller struct {
	service userService.Service
}

var defaultController *Controller

// NewController 创建用户控制器
func NewController(service userService.Service) *Controller {
	if service == nil {
		service = userService.GetUserService()
	}
	return &Controller{service: service}
}

func getDefaultController() *Controller {
	if defaultController == nil {
		defaultController = NewController(nil)
	}
	return defaultController
}

// Register 用户注册
func (ctl *Controller) Register(c *gin.Context) {
	req, err := common.ValidateRequest[request.RegisterRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := ctl.service.RegisterUser(req.Username, req.Email, req.Password, req.Code); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "注册成功")
}

// Login 用户登录
func (ctl *Controller) Login(c *gin.Context) {
	req, err := common.ValidateRequest[request.LoginRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	userInfo, tokenPair, err := ctl.service.Login(req.Account, req.Password)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	username, _ := userInfo["username"].(string)
	email, _ := userInfo["email"].(string)
	avatar, _ := userInfo["avatar"].(string)
	status, _ := userInfo["status"].(int)

	resp := response.LoginResponse{
		Token:            tokenPair.AccessToken,
		ExpiresAt:        tokenPair.AccessExpiresAt,
		RefreshToken:     tokenPair.RefreshToken,
		RefreshExpiresAt: tokenPair.RefreshExpiresAt,
		User: response.UserInfo{
			Username: username,
			Email:    email,
			Avatar:   avatar,
			Status:   status,
		},
	}

	errors.ResponseSuccess(c, resp, "登录成功")
}

// RefreshToken 刷新访问令牌
func (ctl *Controller) RefreshToken(c *gin.Context) {
	req, err := common.ValidateRequest[request.RefreshTokenRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	tokenPair, err := ctl.service.RefreshToken(req.RefreshToken)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, response.RefreshTokenResponse{
		Token:            tokenPair.AccessToken,
		ExpiresAt:        tokenPair.AccessExpiresAt,
		RefreshToken:     tokenPair.RefreshToken,
		RefreshExpiresAt: tokenPair.RefreshExpiresAt,
	}, "令牌刷新成功")
}

func getUserIDFromContext(c *gin.Context) (string, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return "", false
	}

	uid, ok := userID.(string)
	if !ok || uid == "" {
		return "", false
	}

	return uid, true
}

// GetUserInfo 获取用户信息
func (ctl *Controller) GetUserInfo(c *gin.Context) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		errors.HandleError(c, errors.New(errors.CodeUnauthorized, "未授权"))
		return
	}

	userInfo, err := ctl.service.GetUserInfo(userID)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, userInfo, "获取用户信息成功")
}

// UpdateProfile 更新用户资料
func (ctl *Controller) UpdateProfile(c *gin.Context) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		errors.HandleError(c, errors.New(errors.CodeUnauthorized, "未授权"))
		return
	}

	req, err := common.ValidateRequest[request.UpdateProfileRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	userInfo, err := ctl.service.UpdateProfile(userID, req.Username, req.Email, req.Avatar, req.Code)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, userInfo, "更新资料成功")
}

// ChangePassword 修改密码
func (ctl *Controller) ChangePassword(c *gin.Context) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		errors.HandleError(c, errors.New(errors.CodeUnauthorized, "未授权"))
		return
	}

	req, err := common.ValidateRequest[request.ChangePasswordRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := ctl.service.ChangePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "密码修改成功")
}

// Logout 用户退出登录
func (ctl *Controller) Logout(c *gin.Context) {
	token, err := middleware.ExtractBearerToken(c)
	if err != nil {
		errors.HandleError(c, errors.New(errors.CodeUnauthorized, "未授权"))
		return
	}

	if err := ctl.service.Logout(token); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "退出登录成功")
}

// SendRegistrationCode 发送注册验证码
func (ctl *Controller) SendRegistrationCode(c *gin.Context) {
	req, err := common.ValidateRequest[request.SendCodeRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := ctl.service.SendRegistrationCode(req.Email); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "验证码已发送")
}

// SendResetPasswordCode 发送重置密码验证码
func (ctl *Controller) SendResetPasswordCode(c *gin.Context) {
	req, err := common.ValidateRequest[request.SendCodeRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := ctl.service.SendResetPasswordCode(req.Email); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "验证码已发送")
}

// ResetPassword 重置密码
func (ctl *Controller) ResetPassword(c *gin.Context) {
	req, err := common.ValidateRequest[request.ResetPasswordRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := ctl.service.ResetPassword(req.Email, req.Code, req.NewPassword); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "密码重置成功")
}

// SendChangeEmailCode 发送修改邮箱验证码
func (ctl *Controller) SendChangeEmailCode(c *gin.Context) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		errors.HandleError(c, errors.New(errors.CodeUnauthorized, "未授权"))
		return
	}

	req, err := common.ValidateRequest[request.SendCodeRequest](c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := ctl.service.SendChangeEmailCode(userID, req.Email); err != nil {
		errors.HandleError(c, err)
		return
	}

	errors.ResponseSuccess(c, nil, "验证码已发送")
}

// -------- 兼容旧路由调用 --------

func Register(c *gin.Context) {
	getDefaultController().Register(c)
}

func Login(c *gin.Context) {
	getDefaultController().Login(c)
}

func RefreshToken(c *gin.Context) {
	getDefaultController().RefreshToken(c)
}

func GetUserInfo(c *gin.Context) {
	getDefaultController().GetUserInfo(c)
}

func UpdateProfile(c *gin.Context) {
	getDefaultController().UpdateProfile(c)
}

func ChangePassword(c *gin.Context) {
	getDefaultController().ChangePassword(c)
}

func Logout(c *gin.Context) {
	getDefaultController().Logout(c)
}

func SendRegistrationCode(c *gin.Context) {
	getDefaultController().SendRegistrationCode(c)
}

func SendResetPasswordCode(c *gin.Context) {
	getDefaultController().SendResetPasswordCode(c)
}

func ResetPassword(c *gin.Context) {
	getDefaultController().ResetPassword(c)
}

func SendChangeEmailCode(c *gin.Context) {
	getDefaultController().SendChangeEmailCode(c)
}
