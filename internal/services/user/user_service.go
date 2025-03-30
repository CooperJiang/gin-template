package user

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"template/internal/models"
	"template/pkg/cache"
	"template/pkg/common"
	"template/pkg/database"
	"template/pkg/email"
	"template/pkg/utils"
	"time"
)

var userService *UserService

// UserService 用户服务
type UserService struct {
	// 存储服务需要的依赖或状态
}

// InitUserService 初始化用户服务
func InitUserService() {
	userService = &UserService{}
}

// GetUserService 获取用户服务实例
func GetUserService() *UserService {
	return userService
}

// Login 用户登录
func Login(account, password string) (map[string]interface{}, string, error) {
	db := database.GetDB()
	var user models.User
	result := db.Where("username = ? OR email = ?", account, account).First(&user)
	if result.Error != nil {
		return nil, "", errors.New("账号或密码错误")
	}

	// 验证密码
	if !utils.ComparePasswords(user.Password, password) {
		return nil, "", errors.New("账号或密码错误")
	}

	// 检查用户状态
	if !user.IsNormal() {
		return nil, "", errors.New("账号已被禁用")
	}

	// 生成 JWT token
	token, err := common.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, "", errors.New("生成token失败")
	}

	// 构造用户信息
	userInfo := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"avatar":   user.Avatar,
		"bio":      user.Bio,
		"role":     user.Role,
		"status":   user.Status,
	}

	return userInfo, token, nil
}

// FindUsers 获取用户列表
func FindUsers() ([]models.User, error) {
	db := database.GetDB()
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}

// FindUserByID 根据ID查找用户
func FindUserByID(id string) (*models.User, error) {
	db := database.GetDB()
	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindUserByEmail 根据邮箱查找用户
func FindUserByEmail(email string) (*models.User, error) {
	db := database.GetDB()
	var user models.User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// generateVerificationCode 生成验证码
func generateVerificationCode(email string, codeType string) string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成6位随机数
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// 使用email和类型作为key，存储到缓存中，5分钟过期
	key := fmt.Sprintf("%s:%s:code", email, codeType)
	err := cache.GetCache().Set(key, code, 5*time.Minute)
	if err != nil {
		log.Printf("存储验证码到缓存失败: %v", err)
		return ""
	}

	return code
}

// SendRegistrationCode 发送注册验证码
func SendRegistrationCode(email string) error {
	// 检查邮箱是否已被注册
	_, err := FindUserByEmail(email)
	if err == nil {
		return errors.New("该邮箱已被注册")
	}

	// 生成注册验证码
	code := generateVerificationCode(email, common.CodeTypeRegister)
	if code == "" {
		return errors.New("生成验证码失败")
	}

	// 发送验证码邮件
	if err := sendVerificationEmail(email, code, common.CodeTypeRegister); err != nil {
		return fmt.Errorf("发送验证码失败: %v", err)
	}

	return nil
}

// SendResetPasswordCode 发送重置密码验证码
func SendResetPasswordCode(email string) error {
	// 检查邮箱是否存在
	_, err := FindUserByEmail(email)
	if err != nil {
		return errors.New("该邮箱尚未注册")
	}

	// 生成重置密码验证码
	code := generateVerificationCode(email, common.CodeTypeResetPassword)
	if code == "" {
		return errors.New("生成验证码失败")
	}

	// 发送验证码邮件
	if err := sendVerificationEmail(email, code, common.CodeTypeResetPassword); err != nil {
		return fmt.Errorf("发送验证码失败: %v", err)
	}

	return nil
}

// ValidateCode 验证验证码
func ValidateCode(email, code, codeType string) bool {
	key := fmt.Sprintf("%s:%s:code", email, codeType)
	cachedCode, err := cache.GetCache().Get(key)
	if err != nil {
		log.Printf("获取验证码失败: %v", err)
		return false
	}

	// 验证码匹配
	if code == cachedCode {
		// 验证成功后删除验证码
		_ = cache.GetCache().Del(key)
		return true
	}

	return false
}

// RegisterUser 注册用户
func RegisterUser(username, email, password, code string) error {
	db := database.GetDB()

	// 验证验证码
	if !ValidateCode(email, code, common.CodeTypeRegister) {
		return errors.New("验证码无效或已过期")
	}

	// 检查用户名是否已存在
	var count int64
	db.Model(&models.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return errors.New("邮箱已被注册")
	}

	// 创建用户
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return errors.New("密码加密失败")
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Status:   common.UserStatusNormal,
		Role:     common.UserRoleUser,
	}

	if err := db.Create(&user).Error; err != nil {
		return errors.New("创建用户失败")
	}

	return nil
}

// sendVerificationEmail 发送验证码邮件
func sendVerificationEmail(emailAddr string, code string, codeType string) error {
	// 检查邮件服务是否可用
	if !email.IsMailEnabled() {
		return errors.New("邮件服务不可用，请联系管理员")
	}

	var subject string
	if codeType == common.CodeTypeRegister {
		subject = "注册验证码"
	} else {
		subject = "重置密码验证码"
	}

	// 发送验证码邮件
	err := email.SendMail(emailAddr, subject, fmt.Sprintf("您的验证码是: %s，5分钟内有效。", code))
	if err != nil {
		return err
	}

	return nil
}

// ResetPassword 重置密码
func ResetPassword(email, code, newPassword string) error {
	db := database.GetDB()

	// 验证验证码
	if !ValidateCode(email, code, common.CodeTypeResetPassword) {
		return errors.New("验证码无效或已过期")
	}

	// 更新密码
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("密码加密失败")
	}

	result := db.Model(&models.User{}).Where("email = ?", email).Update("password", hashedPassword)
	if result.Error != nil {
		return errors.New("更新密码失败")
	}

	if result.RowsAffected == 0 {
		return errors.New("未找到用户")
	}

	return nil
}
