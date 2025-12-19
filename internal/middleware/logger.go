package middleware

import (
	"bytes"
	"io"
	"template/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		requestID := GetRequestID(c)

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		method := c.Request.Method
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			if len(bodyBytes) > 1000 {
				requestBody = string(bodyBytes[:1000]) + "..."
			} else {
				requestBody = string(bodyBytes)
			}
		}

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		logFields := map[string]interface{}{
			"request_id":   requestID,
			"method":       method,
			"path":         path,
			"status":       statusCode,
			"latency_ms":   latency.Milliseconds(),
			"client_ip":    clientIP,
			"user_agent":   userAgent,
			"request_body": requestBody,
		}

		if len(c.Errors) > 0 {
			logFields["errors"] = c.Errors.String()
		}

		if statusCode >= 500 {
			logger.ErrorWithFields("HTTP Error", logFields)
		} else if statusCode >= 400 {
			logger.WarnWithFields("HTTP Warning", logFields)
		} else {
			logger.InfoWithFields("HTTP Request", logFields)
		}
	}
}

func SkipPaths(skipPaths []string) gin.HandlerFunc {
	skipPathsMap := make(map[string]bool)
	for _, path := range skipPaths {
		skipPathsMap[path] = true
	}

	return func(c *gin.Context) {
		if skipPathsMap[c.Request.URL.Path] {
			c.Next()
			return
		}

		LoggerMiddleware()(c)
	}
}
