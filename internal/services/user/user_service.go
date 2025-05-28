package user

import (
	"fmt"
	"log"
	"math/rand"
	"template/internal/models"
	"template/pkg/cache"
	"template/pkg/common"
	"template/pkg/database"
	"template/pkg/email"
	"template/pkg/errors"
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
func Login(account, password string) (map[string]interface{}, string, time.Time, error) {
	db := database.GetDB()
	if db == nil {
		return nil, "", time.Time{}, errors.New(errors.CodeDBConnectionFailed, "数据库连接失败")
	}

	// 使用原始SQL查询来避免UUID扫描问题
	var userRow struct {
		ID        string `db:"id"`
		Username  string `db:"username"`
		Password  string `db:"password"`
		Email     string `db:"email"`
		Avatar    string `db:"avatar"`
		Bio       string `db:"bio"`
		Status    int    `db:"status"`
		Role      int    `db:"role"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
	}

	err := db.Raw("SELECT id, username, password, email, avatar, bio, status, role, created_at, updated_at FROM user WHERE username = ? OR email = ? LIMIT 1", account, account).Scan(&userRow).Error
	if err != nil {
		return nil, "", time.Time{}, errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}

	if userRow.ID == "" {
		return nil, "", time.Time{}, errors.New(errors.CodeUserNotFound, "用户不存在")
	}

	// 验证密码
	if !utils.ComparePasswords(userRow.Password, password) {
		return nil, "", time.Time{}, errors.New(errors.CodeWrongPassword, "密码错误")
	}

	// 检查用户状态
	if userRow.Status != common.UserStatusNormal {
		return nil, "", time.Time{}, errors.New(errors.CodeUserDisabled, "账号已被禁用")
	}

	// 解析UUID
	userID, err := common.ParseUUID(userRow.ID)
	if err != nil {
		return nil, "", time.Time{}, errors.New(errors.CodeInternal, "用户ID格式错误")
	}

	// 生成 JWT token
	token, err := common.GenerateToken(userID, userRow.Username, userRow.Role)
	if err != nil {
		return nil, "", time.Time{}, errors.New(errors.CodeInternal, "生成token失败")
	}

	// 计算过期时间
	expiresAt := time.Now().Add(time.Duration(24) * time.Hour) // 默认24小时

	// 构造用户信息
	userInfo := map[string]interface{}{
		"id":       userRow.ID,
		"username": userRow.Username,
		"email":    userRow.Email,
		"avatar":   userRow.Avatar,
		"bio":      userRow.Bio,
		"role":     userRow.Role,
		"status":   userRow.Status,
	}

	return userInfo, token, expiresAt, nil
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
		return errors.New(errors.CodeEmailExists, "该邮箱已被注册")
	}

	// 生成注册验证码
	code := generateVerificationCode(email, common.CodeTypeRegister)
	if code == "" {
		return errors.New(errors.CodeInternal, "生成验证码失败")
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
		return errors.New(errors.CodeUserNotFound, "该邮箱尚未注册")
	}

	// 生成重置密码验证码
	code := generateVerificationCode(email, common.CodeTypeResetPassword)
	if code == "" {
		return errors.New(errors.CodeInternal, "生成验证码失败")
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
		return errors.New(errors.CodeInvalidVerifyCode, "验证码无效或已过期")
	}

	// 检查用户名是否已存在
	var count int64
	db.Model(&models.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return errors.New(errors.CodeUserExists, "用户名已存在")
	}

	// 检查邮箱是否已存在
	db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return errors.New(errors.CodeEmailExists, "邮箱已被注册")
	}

	// 创建用户
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return errors.New(errors.CodeInternal, "密码加密失败")
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Status:   common.UserStatusNormal,
		Role:     common.UserRoleUser,
	}

	if err := db.Create(&user).Error; err != nil {
		return errors.New(errors.CodeInternal, "创建用户失败")
	}

	return nil
}

// sendVerificationEmail 发送验证码邮件
func sendVerificationEmail(emailAddr string, code string, codeType string) error {
	// 检查邮件服务是否可用
	if !email.IsMailEnabled() {
		return errors.New(errors.CodeEmailServiceError, "邮件服务不可用，请联系管理员")
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
		return errors.New(errors.CodeInvalidVerifyCode, "验证码无效或已过期")
	}

	// 更新密码
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New(errors.CodeInternal, "密码加密失败")
	}

	result := db.Model(&models.User{}).Where("email = ?", email).Update("password", hashedPassword)
	if result.Error != nil {
		return errors.New(errors.CodeInternal, "更新密码失败")
	}

	if result.RowsAffected == 0 {
		return errors.New(errors.CodeUserNotFound, "未找到用户")
	}

	return nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(userID string) (map[string]interface{}, error) {
	db := database.GetDB()
	if db == nil {
		return nil, errors.New(errors.CodeDBConnectionFailed, "数据库连接失败")
	}

	// 使用原始SQL查询来避免UUID扫描问题
	var userRow struct {
		ID        string `db:"id"`
		Username  string `db:"username"`
		Email     string `db:"email"`
		Avatar    string `db:"avatar"`
		Bio       string `db:"bio"`
		Status    int    `db:"status"`
		Role      int    `db:"role"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
	}

	err := db.Raw("SELECT id, username, email, avatar, bio, status, role, created_at, updated_at FROM user WHERE id = ? LIMIT 1", userID).Scan(&userRow).Error
	if err != nil {
		return nil, errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}

	if userRow.ID == "" {
		return nil, errors.New(errors.CodeUserNotFound, "用户不存在")
	}

	// 构造用户信息
	userInfo := map[string]interface{}{
		"id":         userRow.ID,
		"username":   userRow.Username,
		"email":      userRow.Email,
		"avatar":     userRow.Avatar,
		"bio":        userRow.Bio,
		"role":       userRow.Role,
		"status":     userRow.Status,
		"created_at": userRow.CreatedAt,
		"updated_at": userRow.UpdatedAt,
	}

	return userInfo, nil
}

// UpdateProfile 更新用户资料
func UpdateProfile(userID, username, email, avatar string) (map[string]interface{}, error) {
	db := database.GetDB()
	if db == nil {
		return nil, errors.New(errors.CodeDBConnectionFailed, "数据库连接失败")
	}

	// 构建更新数据
	updateData := make(map[string]interface{})
	if username != "" {
		// 检查用户名是否已被其他用户使用
		var count int64
		db.Model(&models.User{}).Where("username = ? AND id != ?", username, userID).Count(&count)
		if count > 0 {
			return nil, errors.New(errors.CodeUserExists, "用户名已被使用")
		}
		updateData["username"] = username
	}
	if email != "" {
		// 检查邮箱是否已被其他用户使用
		var count int64
		db.Model(&models.User{}).Where("email = ? AND id != ?", email, userID).Count(&count)
		if count > 0 {
			return nil, errors.New(errors.CodeEmailExists, "邮箱已被使用")
		}
		updateData["email"] = email
	}
	if avatar != "" {
		updateData["avatar"] = avatar
	}

	if len(updateData) == 0 {
		return nil, errors.New(errors.CodeInvalidParameter, "没有需要更新的数据")
	}

	// 更新用户信息
	result := db.Model(&models.User{}).Where("id = ?", userID).Updates(updateData)
	if result.Error != nil {
		return nil, errors.New(errors.CodeInternal, "更新用户信息失败")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New(errors.CodeUserNotFound, "用户不存在")
	}

	// 返回更新后的用户信息
	return GetUserInfo(userID)
}

// ChangePassword 修改密码
func ChangePassword(userID, oldPassword, newPassword string) error {
	db := database.GetDB()
	if db == nil {
		return errors.New(errors.CodeDBConnectionFailed, "数据库连接失败")
	}

	// 获取用户当前密码
	var userRow struct {
		Password string `db:"password"`
	}

	err := db.Raw("SELECT password FROM user WHERE id = ? LIMIT 1", userID).Scan(&userRow).Error
	if err != nil {
		return errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}

	if userRow.Password == "" {
		return errors.New(errors.CodeUserNotFound, "用户不存在")
	}

	// 验证原密码
	if !utils.ComparePasswords(userRow.Password, oldPassword) {
		return errors.New(errors.CodeWrongPassword, "原密码错误")
	}

	// 加密新密码
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New(errors.CodeInternal, "密码加密失败")
	}

	// 更新密码
	result := db.Model(&models.User{}).Where("id = ?", userID).Update("password", hashedPassword)
	if result.Error != nil {
		return errors.New(errors.CodeInternal, "更新密码失败")
	}

	if result.RowsAffected == 0 {
		return errors.New(errors.CodeUserNotFound, "用户不存在")
	}

	return nil
}
