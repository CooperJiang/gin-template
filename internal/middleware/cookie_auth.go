package middleware

import (
	"net/http"

	"email-manage/pkg/common"

	"github.com/gin-gonic/gin"
)

const cookieName = "iding-session"

// RequireCookieAuth 从 Cookie 读取 JWT 进行鉴权，
// 用于兼容 team-helper 插件的 /api 路由。
// 鉴权失败返回裸 JSON {"error":"..."} 而非 wrapped response。
func RequireCookieAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(cookieName)
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing session cookie"})
			c.Abort()
			return
		}

		claims, err := common.ParseTokenByType(token, common.TokenTypeAccess)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired session"})
			c.Abort()
			return
		}

		if err := validateClaimsWithUserState(claims); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired session"})
			c.Abort()
			return
		}

		c.Set(ContextPayloadKey, claims)
		c.Set("user_id", claims.UserID)
		c.Set("token_id", claims.ID)
		c.Next()
	}
}
