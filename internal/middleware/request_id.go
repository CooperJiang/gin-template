package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	RequestIDHeader     = "X-Request-ID"
	RequestIDContextKey = "request_id"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(RequestIDHeader)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set(RequestIDContextKey, requestID)
		c.Header(RequestIDHeader, requestID)

		c.Next()
	}
}

func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get(RequestIDContextKey); exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return uuid.New().String()
}
