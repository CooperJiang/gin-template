package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"template/pkg/cache"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimitConfig struct {
	MaxRequests int
	Window      time.Duration
	PerUser     bool
	PerIP       bool
	KeyFunc     func(*gin.Context) string
}

type inMemoryStore struct {
	mu    sync.RWMutex
	store map[string]*rateLimitEntry
}

type rateLimitEntry struct {
	count     int
	expiresAt time.Time
}

var memStore = &inMemoryStore{
	store: make(map[string]*rateLimitEntry),
}

func RateLimitMiddleware(config RateLimitConfig) gin.HandlerFunc {
	if config.MaxRequests == 0 {
		config.MaxRequests = 100
	}
	if config.Window == 0 {
		config.Window = time.Minute
	}
	if !config.PerUser && !config.PerIP && config.KeyFunc == nil {
		config.PerIP = true
	}

	return func(c *gin.Context) {
		key := generateRateLimitKey(c, config)
		allowed, remaining, resetTime := checkRateLimit(key, config)

		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", config.MaxRequests))
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
		c.Header("X-RateLimit-Reset", fmt.Sprintf("%d", resetTime.Unix()))

		if !allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code":        http.StatusTooManyRequests,
				"message":     "Too many requests",
				"error":       "rate_limit_exceeded",
				"retry_after": resetTime.Unix(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func generateRateLimitKey(c *gin.Context, config RateLimitConfig) string {
	if config.KeyFunc != nil {
		return "ratelimit:" + config.KeyFunc(c)
	}

	var key string

	if config.PerUser {
		if userID, exists := c.Get("user_id"); exists {
			key = fmt.Sprintf("ratelimit:user:%v", userID)
		} else {
			key = fmt.Sprintf("ratelimit:ip:%s", c.ClientIP())
		}
	} else if config.PerIP {
		key = fmt.Sprintf("ratelimit:ip:%s", c.ClientIP())
	}

	return key
}

func checkRateLimit(key string, config RateLimitConfig) (allowed bool, remaining int, resetTime time.Time) {
	now := time.Now()
	resetTime = now.Add(config.Window)

	cacheClient := cache.GetCache()
	if cacheClient != nil {
		return checkRateLimitWithRedis(key, config, now, resetTime)
	}

	return checkRateLimitWithMemory(key, config, now, resetTime)
}

func checkRateLimitWithRedis(key string, config RateLimitConfig, now time.Time, resetTime time.Time) (bool, int, time.Time) {
	cacheClient := cache.GetCache()

	countStr, err := cacheClient.Get(key)
	if err != nil || countStr == "" {
		_ = cacheClient.Set(key, "1", config.Window)
		return true, config.MaxRequests - 1, resetTime
	}

	var count int
	fmt.Sscanf(countStr, "%d", &count)

	if count >= config.MaxRequests {
		return false, 0, resetTime
	}

	_ = cacheClient.Set(key, fmt.Sprintf("%d", count+1), config.Window)
	return true, config.MaxRequests - count - 1, resetTime
}

func checkRateLimitWithMemory(key string, config RateLimitConfig, now time.Time, resetTime time.Time) (bool, int, time.Time) {
	memStore.mu.Lock()
	defer memStore.mu.Unlock()

	for k, entry := range memStore.store {
		if entry.expiresAt.Before(now) {
			delete(memStore.store, k)
		}
	}

	entry, exists := memStore.store[key]
	if !exists {
		memStore.store[key] = &rateLimitEntry{
			count:     1,
			expiresAt: resetTime,
		}
		return true, config.MaxRequests - 1, resetTime
	}

	if entry.expiresAt.Before(now) {
		entry.count = 1
		entry.expiresAt = resetTime
		return true, config.MaxRequests - 1, resetTime
	}

	if entry.count >= config.MaxRequests {
		return false, 0, entry.expiresAt
	}

	entry.count++
	return true, config.MaxRequests - entry.count, entry.expiresAt
}

func DefaultRateLimiter() gin.HandlerFunc {
	return RateLimitMiddleware(RateLimitConfig{
		MaxRequests: 100,
		Window:      time.Minute,
		PerIP:       true,
	})
}

func StrictRateLimiter() gin.HandlerFunc {
	return RateLimitMiddleware(RateLimitConfig{
		MaxRequests: 20,
		Window:      time.Minute,
		PerIP:       true,
	})
}

func AuthRateLimiter() gin.HandlerFunc {
	return RateLimitMiddleware(RateLimitConfig{
		MaxRequests: 5,
		Window:      5 * time.Minute,
		PerIP:       true,
	})
}

func UserRateLimiter() gin.HandlerFunc {
	return RateLimitMiddleware(RateLimitConfig{
		MaxRequests: 1000,
		Window:      time.Hour,
		PerUser:     true,
	})
}
