package middleware

import (
	"errors"
	"strings"
	"template/internal/models"
	"template/pkg/common"
	"template/pkg/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 定义上下文中用户信息的键
const (
	ContextPayloadKey = "payload"
)

var (
	errMissingAuthHeader = errors.New("missing authorization header")
	errInvalidAuthHeader = errors.New("invalid authorization header")
	errInvalidAuthToken  = errors.New("invalid authorization token")
)

// ExtractBearerToken 从请求头提取Bearer Token
func ExtractBearerToken(c *gin.Context) (string, error) {
	authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
	if authHeader == "" {
		return "", errMissingAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", errInvalidAuthHeader
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", errInvalidAuthHeader
	}

	return token, nil
}

// parseClaimsFromRequest 解析并校验请求中的 Bearer Token
func parseClaimsFromRequest(c *gin.Context) (*common.JWTClaims, error) {
	token, err := ExtractBearerToken(c)
	if err != nil {
		return nil, err
	}

	claims, err := common.ParseTokenByType(token, common.TokenTypeAccess)
	if err != nil {
		return nil, errInvalidAuthToken
	}

	if err := validateClaimsWithUserState(claims); err != nil {
		return nil, errInvalidAuthToken
	}

	return claims, nil
}

func validateClaimsWithUserState(claims *common.JWTClaims) error {
	db := database.GetDB()
	if db == nil {
		return errInvalidAuthToken
	}

	var user models.User
	if err := db.Select("id", "status", "token_version").Where("id = ?", claims.UserID).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errInvalidAuthToken
		}
		return errInvalidAuthToken
	}

	expectedTokenVersion := user.TokenVersion
	if expectedTokenVersion <= 0 {
		expectedTokenVersion = 1
	}

	if user.Status != common.UserStatusNormal {
		return errInvalidAuthToken
	}

	if claims.TokenVersion != expectedTokenVersion {
		return errInvalidAuthToken
	}

	return nil
}

// GetUserFromContext 从上下文中获取用户信息
func GetUserFromContext(c *gin.Context) (*common.JWTClaims, error) {
	value, exists := c.Get(ContextPayloadKey)
	if !exists {
		return nil, errors.New("用户未登录")
	}
	user, ok := value.(*common.JWTClaims)
	if !ok {
		return nil, errors.New("用户信息类型错误")
	}
	return user, nil
}

// RequireAuth 基础认证中间件，验证用户是否登录
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := parseClaimsFromRequest(c)
		if err != nil {
			if errors.Is(err, errMissingAuthHeader) {
				common.Unauthorized(c, "未提供有效的认证凭证")
				c.Abort()
				return
			}
			common.Unauthorized(c, "认证凭证无效或已过期")
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set(ContextPayloadKey, claims)
		// 同时设置user_id便于控制器直接获取
		c.Set("user_id", claims.UserID)
		c.Set("token_id", claims.ID)
		c.Next()
	}
}

// RequireSuperAuth 超级管理员权限中间件
func RequireSuperAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先验证基本认证
		RequireAuth()(c)
		if c.IsAborted() {
			return
		}

		// 获取用户信息
		claims, err := GetUserFromContext(c)
		if err != nil {
			common.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		// 检查是否是超级管理员
		if claims.Role != common.UserRoleSuperAdmin {
			common.Forbidden(c, "需要超级管理员权限")
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireAdmin 管理员权限中间件
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先验证基本认证
		RequireAuth()(c)
		if c.IsAborted() {
			return
		}

		// 获取用户信息
		claims, err := GetUserFromContext(c)
		if err != nil {
			common.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		// 检查是否是管理员或超级管理员
		if claims.Role != common.UserRoleAdmin && claims.Role != common.UserRoleSuperAdmin {
			common.Forbidden(c, "需要管理员权限")
			c.Abort()
			return
		}

		c.Next()
	}
}
