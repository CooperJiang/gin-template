package user

import (
	cryptoRand "crypto/rand"
	stdErrors "errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"template/internal/models"
	userRepo "template/internal/repositories/user"
	"template/pkg/cache"
	"template/pkg/common"
	"template/pkg/database"
	"template/pkg/email"
	"template/pkg/errors"
	"template/pkg/utils"
	"time"

	"gorm.io/gorm"
)

const (
	codeTTL                   = 5 * time.Minute
	codeSendCooldown          = 60 * time.Second
	codeRateLimitWindow       = 10 * time.Minute
	maxCodeSendsPerWindow     = 5
	genericLoginFailedMessage = "账号或密码错误"
)

// Mailer 邮件发送抽象，便于注入和测试
// Enabled 返回邮件服务是否可用
// Send 负责发送邮件内容
type Mailer interface {
	Enabled() bool
	Send(to, subject, body string) error
}

type emailMailer struct{}

// NewEmailMailer 创建默认邮件发送器
func NewEmailMailer() Mailer {
	return &emailMailer{}
}

func (m *emailMailer) Enabled() bool {
	return email.IsMailEnabled()
}

func (m *emailMailer) Send(to, subject, body string) error {
	return email.SendMail(to, subject, body)
}

// Service 定义用户服务能力边界，便于controller注入
type Service interface {
	Login(account, password string) (map[string]interface{}, *common.TokenPair, error)
	RefreshToken(refreshToken string) (*common.TokenPair, error)
	Logout(accessToken string) error
	FindUsers() ([]models.User, error)
	FindUserByID(id string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	SendRegistrationCode(email string) error
	SendResetPasswordCode(email string) error
	SendChangeEmailCode(userID, newEmail string) error
	ValidateCode(email, code, codeType string) bool
	RegisterUser(username, emailAddr, password, code string) error
	ResetPassword(emailAddr, code, newPassword string) error
	GetUserInfo(userID string) (map[string]interface{}, error)
	UpdateProfile(userID, username, emailAddr, avatar, code string) (map[string]interface{}, error)
	ChangePassword(userID, oldPassword, newPassword string) error
}

// UserService 用户服务
type UserService struct {
	repo       *userRepo.Repository
	cacheStore cache.Cache
	mailer     Mailer
}

var defaultService Service

// NewService 创建用户服务实例（显式依赖注入）
func NewService(db *gorm.DB, cacheStore cache.Cache, mailer Mailer) *UserService {
	if cacheStore == nil {
		cacheStore = cache.GetCache()
	}
	if mailer == nil {
		mailer = NewEmailMailer()
	}

	var repo *userRepo.Repository
	if db != nil {
		repo = userRepo.NewRepository(db)
	}

	return &UserService{
		repo:       repo,
		cacheStore: cacheStore,
		mailer:     mailer,
	}
}

// InitUserService 初始化默认用户服务（兼容旧调用）
func InitUserService() {
	defaultService = NewService(database.GetDB(), cache.GetCache(), NewEmailMailer())
}

// GetUserService 获取默认用户服务实例（兼容旧调用）
func GetUserService() Service {
	if defaultService == nil {
		InitUserService()
	}
	return defaultService
}

func (s *UserService) getRepo() (*userRepo.Repository, error) {
	if s.repo == nil {
		return nil, errors.New(errors.CodeDBConnectionFailed, "数据库连接失败")
	}
	return s.repo, nil
}

type userAuthState struct {
	UserID       string
	Username     string
	Role         int
	Status       int
	TokenVersion int
}

func (s *UserService) getUserAuthState(userID string) (*userAuthState, error) {
	repo, err := s.getRepo()
	if err != nil {
		return nil, err
	}

	row, err := repo.FindAuthByID(userID)
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(errors.CodeUserNotFound, "用户不存在")
		}
		return nil, errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}

	if row.Status != common.UserStatusNormal {
		return nil, errors.New(errors.CodeUserDisabled, "账号已被禁用")
	}

	if row.TokenVersion <= 0 {
		row.TokenVersion = 1
		_, _ = repo.UpdateByID(row.ID.String(), map[string]interface{}{"token_version": 1})
	}

	return &userAuthState{
		UserID:       row.ID.String(),
		Username:     row.Username,
		Role:         row.Role,
		Status:       row.Status,
		TokenVersion: row.TokenVersion,
	}, nil
}

func normalizeEmail(emailAddr string) string {
	return strings.ToLower(strings.TrimSpace(emailAddr))
}

func (s *UserService) codeCacheKey(emailAddr, codeType string) string {
	return fmt.Sprintf("auth:code:%s:%s", codeType, normalizeEmail(emailAddr))
}

func (s *UserService) codeCooldownKey(emailAddr, codeType string) string {
	return fmt.Sprintf("auth:code:cooldown:%s:%s", codeType, normalizeEmail(emailAddr))
}

func (s *UserService) codeRateKey(emailAddr, codeType string) string {
	return fmt.Sprintf("auth:code:rate:%s:%s", codeType, normalizeEmail(emailAddr))
}

func (s *UserService) consumeSendCodeQuota(emailAddr, codeType string) error {
	if s.cacheStore == nil {
		return errors.New(errors.CodeRedisError, "缓存服务不可用")
	}

	cooldownKey := s.codeCooldownKey(emailAddr, codeType)
	if s.cacheStore.Exists(cooldownKey) {
		return errors.New(errors.CodeRateLimited, "请求过于频繁，请稍后再试")
	}

	rateKey := s.codeRateKey(emailAddr, codeType)
	count := 0
	if raw, err := s.cacheStore.Get(rateKey); err == nil {
		if parsed, parseErr := strconv.Atoi(strings.TrimSpace(raw)); parseErr == nil && parsed > 0 {
			count = parsed
		}
	}

	if count >= maxCodeSendsPerWindow {
		return errors.New(errors.CodeRateLimited, "请求过于频繁，请稍后再试")
	}

	if err := s.cacheStore.Set(rateKey, strconv.Itoa(count+1), codeRateLimitWindow); err != nil {
		return errors.New(errors.CodeRedisError, "验证码频控异常，请稍后再试")
	}
	if err := s.cacheStore.Set(cooldownKey, "1", codeSendCooldown); err != nil {
		return errors.New(errors.CodeRedisError, "验证码频控异常，请稍后再试")
	}

	return nil
}

func (s *UserService) generateVerificationCode(emailAddr, codeType string) (string, error) {
	n, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", errors.New(errors.CodeInternal, "生成验证码失败")
	}

	code := fmt.Sprintf("%06d", n.Int64())
	if err := s.cacheStore.Set(s.codeCacheKey(emailAddr, codeType), code, codeTTL); err != nil {
		return "", errors.New(errors.CodeInternal, "存储验证码失败")
	}

	return code, nil
}

// Login 用户登录
func (s *UserService) Login(account, password string) (map[string]interface{}, *common.TokenPair, error) {
	repo, err := s.getRepo()
	if err != nil {
		return nil, nil, err
	}

	userRow, err := repo.FindByAccount(strings.TrimSpace(account))
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errors.New(errors.CodeWrongPassword, genericLoginFailedMessage)
		}
		return nil, nil, errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}

	if !utils.ComparePasswords(userRow.Password, password) {
		return nil, nil, errors.New(errors.CodeWrongPassword, genericLoginFailedMessage)
	}
	if userRow.Status != common.UserStatusNormal {
		return nil, nil, errors.New(errors.CodeWrongPassword, genericLoginFailedMessage)
	}

	if userRow.TokenVersion <= 0 {
		userRow.TokenVersion = 1
		_, _ = repo.UpdateByID(userRow.ID.String(), map[string]interface{}{"token_version": 1})
	}

	tokenPair, err := common.GenerateTokenPair(userRow.ID.String(), userRow.Username, userRow.Role, userRow.TokenVersion)
	if err != nil {
		return nil, nil, errors.New(errors.CodeInternal, "生成token失败")
	}

	userInfo := map[string]interface{}{
		"id":       userRow.ID.String(),
		"username": userRow.Username,
		"email":    userRow.Email,
		"avatar":   userRow.Avatar,
		"bio":      userRow.Bio,
		"role":     userRow.Role,
		"status":   userRow.Status,
	}

	return userInfo, tokenPair, nil
}

// RefreshToken 使用刷新令牌换发新令牌
func (s *UserService) RefreshToken(refreshToken string) (*common.TokenPair, error) {
	claims, err := common.ParseTokenByType(refreshToken, common.TokenTypeRefresh)
	if err != nil {
		return nil, errors.New(errors.CodeInvalidAuthToken, "刷新令牌无效或已过期")
	}

	state, err := s.getUserAuthState(claims.UserID)
	if err != nil {
		return nil, err
	}

	if claims.TokenVersion != state.TokenVersion {
		return nil, errors.New(errors.CodeInvalidAuthToken, "登录状态已失效，请重新登录")
	}

	tokenPair, err := common.GenerateTokenPair(state.UserID, state.Username, state.Role, state.TokenVersion)
	if err != nil {
		return nil, errors.New(errors.CodeInternal, "刷新令牌失败")
	}

	if err := common.BlacklistToken(claims.ID, claims.ExpiresAt.Time); err != nil {
		return nil, errors.New(errors.CodeRedisError, "令牌状态更新失败")
	}

	return tokenPair, nil
}

// Logout 退出登录（提升token_version，使当前用户已有会话全部失效）
func (s *UserService) Logout(accessToken string) error {
	claims, err := common.ParseTokenByType(accessToken, common.TokenTypeAccess)
	if err != nil {
		return errors.New(errors.CodeInvalidAuthToken, "访问令牌无效或已过期")
	}

	repo, err := s.getRepo()
	if err != nil {
		return err
	}

	rows, err := repo.IncrementTokenVersionByID(claims.UserID, true)
	if err != nil {
		return errors.New(errors.CodeInternal, "退出登录失败")
	}
	if rows == 0 {
		return errors.New(errors.CodeInvalidAuthToken, "登录状态已失效，请重新登录")
	}

	if err := common.BlacklistToken(claims.ID, claims.ExpiresAt.Time); err != nil {
		return errors.New(errors.CodeRedisError, "退出登录失败")
	}

	return nil
}

// FindUsers 获取用户列表
func (s *UserService) FindUsers() ([]models.User, error) {
	repo, err := s.getRepo()
	if err != nil {
		return nil, err
	}

	return repo.FindAll()
}

// FindUserByID 根据ID查找用户
func (s *UserService) FindUserByID(id string) (*models.User, error) {
	repo, err := s.getRepo()
	if err != nil {
		return nil, err
	}

	return repo.FindByID(id)
}

// FindUserByEmail 根据邮箱查找用户
func (s *UserService) FindUserByEmail(emailAddr string) (*models.User, error) {
	repo, err := s.getRepo()
	if err != nil {
		return nil, err
	}

	return repo.FindByEmail(normalizeEmail(emailAddr))
}

// SendRegistrationCode 发送注册验证码
func (s *UserService) SendRegistrationCode(emailAddr string) error {
	emailAddr = normalizeEmail(emailAddr)

	if err := s.consumeSendCodeQuota(emailAddr, common.CodeTypeRegister); err != nil {
		return err
	}

	_, err := s.FindUserByEmail(emailAddr)
	if err == nil {
		return errors.New(errors.CodeEmailExists, "该邮箱已被注册")
	}
	if !stdErrors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}

	code, err := s.generateVerificationCode(emailAddr, common.CodeTypeRegister)
	if err != nil {
		return err
	}

	if err := s.sendVerificationEmail(emailAddr, code, common.CodeTypeRegister); err != nil {
		return fmt.Errorf("发送验证码失败: %v", err)
	}

	return nil
}

// SendResetPasswordCode 发送重置密码验证码
func (s *UserService) SendResetPasswordCode(emailAddr string) error {
	emailAddr = normalizeEmail(emailAddr)

	if err := s.consumeSendCodeQuota(emailAddr, common.CodeTypeResetPassword); err != nil {
		return err
	}

	_, err := s.FindUserByEmail(emailAddr)
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(errors.CodeUserNotFound, "该邮箱尚未注册")
		}
		return errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}

	code, err := s.generateVerificationCode(emailAddr, common.CodeTypeResetPassword)
	if err != nil {
		return err
	}

	if err := s.sendVerificationEmail(emailAddr, code, common.CodeTypeResetPassword); err != nil {
		return fmt.Errorf("发送验证码失败: %v", err)
	}

	return nil
}

// SendChangeEmailCode 发送修改邮箱验证码
func (s *UserService) SendChangeEmailCode(userID, newEmail string) error {
	newEmail = normalizeEmail(newEmail)

	if err := s.consumeSendCodeQuota(newEmail, common.CodeTypeChangeEmail); err != nil {
		return err
	}

	existingUser, err := s.FindUserByEmail(newEmail)
	if err == nil && existingUser.ID.String() != strings.TrimSpace(userID) {
		return errors.New(errors.CodeEmailExists, "该邮箱已被其他用户使用")
	}
	if err != nil && !stdErrors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}

	code, err := s.generateVerificationCode(newEmail, common.CodeTypeChangeEmail)
	if err != nil {
		return err
	}

	if err := s.sendVerificationEmail(newEmail, code, common.CodeTypeChangeEmail); err != nil {
		return fmt.Errorf("发送验证码失败: %v", err)
	}

	return nil
}

// ValidateCode 验证验证码
func (s *UserService) ValidateCode(emailAddr, code, codeType string) bool {
	key := s.codeCacheKey(normalizeEmail(emailAddr), codeType)
	cachedCode, err := s.cacheStore.Get(key)
	if err != nil {
		return false
	}

	if strings.TrimSpace(code) == cachedCode {
		_ = s.cacheStore.Del(key)
		return true
	}
	return false
}

// RegisterUser 注册用户
func (s *UserService) RegisterUser(username, emailAddr, password, code string) error {
	repo, err := s.getRepo()
	if err != nil {
		return err
	}

	emailAddr = normalizeEmail(emailAddr)

	if !s.ValidateCode(emailAddr, code, common.CodeTypeRegister) {
		return errors.New(errors.CodeInvalidVerifyCode, "验证码无效或已过期")
	}

	usernameCount, err := repo.CountByUsername(username)
	if err != nil {
		return errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}
	if usernameCount > 0 {
		return errors.New(errors.CodeUserExists, "用户名已存在")
	}

	emailCount, err := repo.CountByEmail(emailAddr)
	if err != nil {
		return errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}
	if emailCount > 0 {
		return errors.New(errors.CodeEmailExists, "邮箱已被注册")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return errors.New(errors.CodeInternal, "密码加密失败")
	}

	user := models.User{
		Username:     strings.TrimSpace(username),
		Email:        emailAddr,
		Password:     hashedPassword,
		Status:       common.UserStatusNormal,
		Role:         common.UserRoleUser,
		TokenVersion: 1,
	}
	if err := repo.Create(&user); err != nil {
		return errors.New(errors.CodeInternal, "创建用户失败")
	}

	return nil
}

func (s *UserService) sendVerificationEmail(emailAddr, code, codeType string) error {
	if !s.mailer.Enabled() {
		return errors.New(errors.CodeEmailServiceError, "邮件服务不可用，请联系管理员")
	}

	subject := "重置密码验证码"
	if codeType == common.CodeTypeRegister {
		subject = "注册验证码"
	} else if codeType == common.CodeTypeChangeEmail {
		subject = "修改邮箱验证码"
	}

	return s.mailer.Send(emailAddr, subject, fmt.Sprintf("您的验证码是: %s，5分钟内有效。", code))
}

// ResetPassword 重置密码
func (s *UserService) ResetPassword(emailAddr, code, newPassword string) error {
	repo, err := s.getRepo()
	if err != nil {
		return err
	}

	emailAddr = normalizeEmail(emailAddr)

	if !s.ValidateCode(emailAddr, code, common.CodeTypeResetPassword) {
		return errors.New(errors.CodeInvalidVerifyCode, "验证码无效或已过期")
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New(errors.CodeInternal, "密码加密失败")
	}

	rows, err := repo.UpdatePasswordByEmail(emailAddr, hashedPassword)
	if err != nil {
		return errors.New(errors.CodeInternal, "更新密码失败")
	}
	if rows == 0 {
		return errors.New(errors.CodeUserNotFound, "未找到用户")
	}

	return nil
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(userID string) (map[string]interface{}, error) {
	repo, err := s.getRepo()
	if err != nil {
		return nil, err
	}

	userRow, err := repo.FindPublicByID(userID)
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(errors.CodeUserNotFound, "用户不存在")
		}
		return nil, errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}

	userInfo := map[string]interface{}{
		"id":         userRow.ID.String(),
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
func (s *UserService) UpdateProfile(userID, username, emailAddr, avatar, code string) (map[string]interface{}, error) {
	repo, err := s.getRepo()
	if err != nil {
		return nil, err
	}

	updateData := make(map[string]interface{})
	if strings.TrimSpace(username) != "" {
		count, countErr := repo.CountByUsernameExcludeID(username, userID)
		if countErr != nil {
			return nil, errors.New(errors.CodeQueryFailed, "数据库查询失败")
		}
		if count > 0 {
			return nil, errors.New(errors.CodeUserExists, "用户名已被使用")
		}
		updateData["username"] = strings.TrimSpace(username)
	}

	if strings.TrimSpace(emailAddr) != "" {
		normalizedEmail := normalizeEmail(emailAddr)
		if strings.TrimSpace(code) == "" {
			return nil, errors.New(errors.CodeInvalidParameter, "修改邮箱必须提供验证码")
		}
		if !s.ValidateCode(normalizedEmail, code, common.CodeTypeChangeEmail) {
			return nil, errors.New(errors.CodeInvalidVerifyCode, "验证码无效或已过期")
		}

		count, countErr := repo.CountByEmailExcludeID(normalizedEmail, userID)
		if countErr != nil {
			return nil, errors.New(errors.CodeQueryFailed, "数据库查询失败")
		}
		if count > 0 {
			return nil, errors.New(errors.CodeEmailExists, "邮箱已被使用")
		}
		updateData["email"] = normalizedEmail
	}

	if strings.TrimSpace(avatar) != "" {
		updateData["avatar"] = strings.TrimSpace(avatar)
	}
	if len(updateData) == 0 {
		return nil, errors.New(errors.CodeInvalidParameter, "没有需要更新的数据")
	}

	rows, err := repo.UpdateByID(userID, updateData)
	if err != nil {
		return nil, errors.New(errors.CodeInternal, "更新用户信息失败")
	}
	if rows == 0 {
		return nil, errors.New(errors.CodeUserNotFound, "用户不存在")
	}

	return s.GetUserInfo(userID)
}

// ChangePassword 修改密码
func (s *UserService) ChangePassword(userID, oldPassword, newPassword string) error {
	repo, err := s.getRepo()
	if err != nil {
		return err
	}

	userRow, err := repo.FindPasswordByID(userID)
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(errors.CodeUserNotFound, "用户不存在")
		}
		return errors.New(errors.CodeQueryFailed, "数据库查询失败")
	}

	if !utils.ComparePasswords(userRow.Password, oldPassword) {
		return errors.New(errors.CodeWrongPassword, "原密码错误")
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New(errors.CodeInternal, "密码加密失败")
	}

	rows, err := repo.UpdatePasswordByID(userID, hashedPassword)
	if err != nil {
		return errors.New(errors.CodeInternal, "更新密码失败")
	}
	if rows == 0 {
		return errors.New(errors.CodeUserNotFound, "用户不存在")
	}

	return nil
}

// -------------------- 兼容旧调用 --------------------

func Login(account, password string) (map[string]interface{}, *common.TokenPair, error) {
	return GetUserService().Login(account, password)
}

func RefreshToken(refreshToken string) (*common.TokenPair, error) {
	return GetUserService().RefreshToken(refreshToken)
}

func Logout(accessToken string) error {
	return GetUserService().Logout(accessToken)
}

func FindUsers() ([]models.User, error) {
	return GetUserService().FindUsers()
}

func FindUserByID(id string) (*models.User, error) {
	return GetUserService().FindUserByID(id)
}

func FindUserByEmail(emailAddr string) (*models.User, error) {
	return GetUserService().FindUserByEmail(emailAddr)
}

func SendRegistrationCode(emailAddr string) error {
	return GetUserService().SendRegistrationCode(emailAddr)
}

func SendResetPasswordCode(emailAddr string) error {
	return GetUserService().SendResetPasswordCode(emailAddr)
}

func SendChangeEmailCode(userID, newEmail string) error {
	return GetUserService().SendChangeEmailCode(userID, newEmail)
}

func ValidateCode(emailAddr, code, codeType string) bool {
	return GetUserService().ValidateCode(emailAddr, code, codeType)
}

func RegisterUser(username, emailAddr, password, code string) error {
	return GetUserService().RegisterUser(username, emailAddr, password, code)
}

func ResetPassword(emailAddr, code, newPassword string) error {
	return GetUserService().ResetPassword(emailAddr, code, newPassword)
}

func GetUserInfo(userID string) (map[string]interface{}, error) {
	return GetUserService().GetUserInfo(userID)
}

func UpdateProfile(userID, username, emailAddr, avatar, code string) (map[string]interface{}, error) {
	return GetUserService().UpdateProfile(userID, username, emailAddr, avatar, code)
}

func ChangePassword(userID, oldPassword, newPassword string) error {
	return GetUserService().ChangePassword(userID, oldPassword, newPassword)
}
