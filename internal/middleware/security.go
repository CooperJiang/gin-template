package middleware

import (
	"github.com/gin-gonic/gin"
)

type SecurityHeaders struct {
	ContentSecurityPolicy   string
	XFrameOptions           string
	XContentTypeOptions     string
	XXSSProtection          string
	ReferrerPolicy          string
	PermissionsPolicy       string
	StrictTransportSecurity string
}

func SecurityHeadersMiddleware(config ...*SecurityHeaders) gin.HandlerFunc {
	defaultConfig := &SecurityHeaders{
		ContentSecurityPolicy:   "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self' data:; connect-src 'self'",
		XFrameOptions:           "SAMEORIGIN",
		XContentTypeOptions:     "nosniff",
		XXSSProtection:          "1; mode=block",
		ReferrerPolicy:          "strict-origin-when-cross-origin",
		PermissionsPolicy:       "geolocation=(), microphone=(), camera=()",
		StrictTransportSecurity: "",
	}

	if len(config) > 0 && config[0] != nil {
		defaultConfig = config[0]
	}

	return func(c *gin.Context) {
		if defaultConfig.ContentSecurityPolicy != "" {
			c.Header("Content-Security-Policy", defaultConfig.ContentSecurityPolicy)
		}
		if defaultConfig.XFrameOptions != "" {
			c.Header("X-Frame-Options", defaultConfig.XFrameOptions)
		}
		if defaultConfig.XContentTypeOptions != "" {
			c.Header("X-Content-Type-Options", defaultConfig.XContentTypeOptions)
		}
		if defaultConfig.XXSSProtection != "" {
			c.Header("X-XSS-Protection", defaultConfig.XXSSProtection)
		}
		if defaultConfig.ReferrerPolicy != "" {
			c.Header("Referrer-Policy", defaultConfig.ReferrerPolicy)
		}
		if defaultConfig.PermissionsPolicy != "" {
			c.Header("Permissions-Policy", defaultConfig.PermissionsPolicy)
		}
		if defaultConfig.StrictTransportSecurity != "" {
			c.Header("Strict-Transport-Security", defaultConfig.StrictTransportSecurity)
		}

		c.Header("X-Powered-By", "")
		c.Header("Server", "")

		c.Next()
	}
}

func SecureHeadersForProduction() gin.HandlerFunc {
	return SecurityHeadersMiddleware(&SecurityHeaders{
		ContentSecurityPolicy:   "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self' data: https:; font-src 'self'; connect-src 'self'; frame-ancestors 'none'; base-uri 'self'; form-action 'self'",
		XFrameOptions:           "DENY",
		XContentTypeOptions:     "nosniff",
		XXSSProtection:          "1; mode=block",
		ReferrerPolicy:          "no-referrer",
		PermissionsPolicy:       "geolocation=(), microphone=(), camera=(), payment=(), usb=(), magnetometer=(), gyroscope=(), accelerometer=()",
		StrictTransportSecurity: "max-age=63072000; includeSubDomains; preload",
	})
}

func SecureHeadersForDevelopment() gin.HandlerFunc {
	return SecurityHeadersMiddleware(&SecurityHeaders{
		ContentSecurityPolicy:   "default-src 'self' 'unsafe-inline' 'unsafe-eval'; img-src 'self' data: https: http:; connect-src 'self' ws: wss: http: https:",
		XFrameOptions:           "SAMEORIGIN",
		XContentTypeOptions:     "nosniff",
		XXSSProtection:          "1; mode=block",
		ReferrerPolicy:          "strict-origin-when-cross-origin",
		PermissionsPolicy:       "",
		StrictTransportSecurity: "",
	})
}
