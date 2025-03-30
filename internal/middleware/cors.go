package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware 处理跨域请求中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的域名
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的 HTTP 方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// 设置允许的 Headers
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// 设置是否允许携带认证信息
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 设置暴露的 Headers
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")

		// 立即处理 OPTIONS 请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
