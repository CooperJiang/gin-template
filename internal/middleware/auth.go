package middleware

import (
	"errors"
	"strings"
	"template/pkg/common"

	"github.com/gin-gonic/gin"
)

// 定义上下文中用户信息的键
const (
	ContextPayloadKey = "payload"
)

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
		// 从请求头获取token
		token := c.GetHeader("Authorization")
		// 如果请求头中不存在，则从查询参数获取
		if token == "" {
			token = c.Query("token")
		}

		// 去掉可能存在的Bearer前缀
		token = strings.TrimPrefix(token, "Bearer ")

		if token == "" {
			common.Unauthorized(c, "未提供有效的认证凭证")
			c.Abort()
			return
		}

		// 验证JWT令牌
		claims, err := common.ParseToken(token)
		if err != nil {
			common.Unauthorized(c, "认证凭证无效或已过期")
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set(ContextPayloadKey, claims)
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
