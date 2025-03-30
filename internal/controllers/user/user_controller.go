package user

import (
	"template/internal/controllers/user/dto"
	"template/internal/services/user"
	"template/pkg/common"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	req, validationErr := common.ValidateRequest[dto.RegisterDTO](c)
	if validationErr != nil {
		common.BadRequest(c, validationErr.Message)
		return
	}

	// 调用服务层进行注册
	err := user.RegisterUser(req.Username, req.Email, req.Password, req.Code)
	if err != nil {
		common.Fail(c, common.StatusBadRequest, err.Error())
		return
	}

	common.SuccessWithMessage(c, "注册成功")
}

// Login 用户登录
func Login(c *gin.Context) {
	req, validationErr := common.ValidateRequest[dto.LoginDTO](c)
	if validationErr != nil {
		common.BadRequest(c, validationErr.Message)
		return
	}

	// 调用服务层进行登录验证
	userInfo, token, err := user.Login(req.Account, req.Password)
	if err != nil {
		common.Fail(c, common.StatusBadRequest, err.Error())
		return
	}

	data := gin.H{
		"token":    token,
		"userInfo": userInfo,
	}

	common.Success(c, data, "登录成功")
}

// SendRegistrationCode 发送注册验证码
func SendRegistrationCode(c *gin.Context) {
	req, validationErr := common.ValidateRequest[dto.SendCodeDTO](c)
	if validationErr != nil {
		common.BadRequest(c, validationErr.Message)
		return
	}

	// 发送验证码
	if err := user.SendRegistrationCode(req.Email); err != nil {
		common.Fail(c, common.StatusBadRequest, err.Error())
		return
	}

	common.SuccessWithMessage(c, "验证码已发送")
}

// SendResetPasswordCode 发送重置密码验证码
func SendResetPasswordCode(c *gin.Context) {
	req, validationErr := common.ValidateRequest[dto.SendCodeDTO](c)
	if validationErr != nil {
		common.BadRequest(c, validationErr.Message)
		return
	}

	// 发送验证码
	if err := user.SendResetPasswordCode(req.Email); err != nil {
		common.Fail(c, common.StatusBadRequest, err.Error())
		return
	}

	common.SuccessWithMessage(c, "验证码已发送")
}

// ResetPassword 重置密码
func ResetPassword(c *gin.Context) {
	req, validationErr := common.ValidateRequest[dto.ResetPasswordDTO](c)
	if validationErr != nil {
		common.BadRequest(c, validationErr.Message)
		return
	}

	// 重置密码
	if err := user.ResetPassword(req.Email, req.Code, req.NewPassword); err != nil {
		common.Fail(c, common.StatusBadRequest, err.Error())
		return
	}

	common.SuccessWithMessage(c, "密码重置成功")
}
