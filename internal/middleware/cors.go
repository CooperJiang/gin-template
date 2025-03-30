package middleware

import (
	"strconv"
	"strings"
	"template/pkg/config"
	"template/pkg/logger"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware 返回基于配置的CORS中间件
func CORSMiddleware() gin.HandlerFunc {
	cfg := config.GetConfig().CORS

	// 如果未启用CORS，返回空中间件
	if !cfg.Enabled {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	logger.Info("CORS中间件已启用")

	return func(c *gin.Context) {
		// 处理Origin
		origin := c.Request.Header.Get("Origin")

		// 设置允许的来源
		if len(cfg.AllowedOrigins) > 0 {
			// 检查请求的Origin是否在允许列表中
			allowed := false
			if len(origin) > 0 {
				for _, allowedOrigin := range cfg.AllowedOrigins {
					if allowedOrigin == "*" || allowedOrigin == origin {
						allowed = true
						break
					}
				}
			}

			if allowed {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			} else if len(cfg.AllowedOrigins) == 1 && cfg.AllowedOrigins[0] == "*" {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			}
		} else {
			// 默认允许所有来源
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		// 设置允许的Headers
		if len(cfg.AllowedHeaders) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(cfg.AllowedHeaders, ", "))
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		}

		// 设置允许的Methods
		if len(cfg.AllowedMethods) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(cfg.AllowedMethods, ", "))
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		}

		// 设置是否允许携带凭证
		if cfg.AllowCredentials {
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// 设置预检请求的缓存时间
		if cfg.MaxAge > 0 {
			c.Writer.Header().Set("Access-Control-Max-Age", strconv.Itoa(cfg.MaxAge))
		}

		// 对于预检请求，直接返回成功
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
