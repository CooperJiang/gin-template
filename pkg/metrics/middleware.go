package metrics

import (
	"net/http"
	"time"

	_ "net/http/pprof" // 导入pprof

	"github.com/gin-gonic/gin"
)

// MetricsMiddleware 性能监控中间件
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求处理前
		startTime := time.Now()
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		// 记录请求开始
		m := GetMetrics()
		m.RecordRequest(path)

		// 处理请求
		c.Next()

		// 请求处理后
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()
		isError := statusCode >= 400

		// 记录请求完成
		m.RecordRequestComplete(path, statusCode, duration, isError)
	}
}

// RateLimiter 请求限速中间件
type RateLimiter struct {
	maxRequests int           // 时间窗口内最大请求数
	window      time.Duration // 时间窗口大小
	clients     map[string]*clientRateLimit
	cleanup     *time.Ticker
}

type clientRateLimit struct {
	requests    int       // 当前窗口内的请求数
	windowStart time.Time // 当前窗口的开始时间
}

// NewRateLimiter 创建新的限速器
func NewRateLimiter(maxRequests int, window time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		maxRequests: maxRequests,
		window:      window,
		clients:     make(map[string]*clientRateLimit),
		cleanup:     time.NewTicker(time.Minute * 5), // 每5分钟清理过期记录
	}

	go limiter.cleanupRoutine()
	return limiter
}

// cleanupRoutine 清理过期的客户端限制记录
func (rl *RateLimiter) cleanupRoutine() {
	for range rl.cleanup.C {
		rl.cleanupExpiredClients()
	}
}

// cleanupExpiredClients 清理过期的客户端记录
func (rl *RateLimiter) cleanupExpiredClients() {
	now := time.Now()
	for ip, client := range rl.clients {
		if now.Sub(client.windowStart) > rl.window {
			delete(rl.clients, ip)
		}
	}
}

// Stop 停止限速器
func (rl *RateLimiter) Stop() {
	rl.cleanup.Stop()
}

// Middleware 创建Gin中间件
func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		// 检查是否允许请求
		if !rl.allow(ip, now) {
			// 记录请求被限制
			m := GetMetrics()

			// 获取当前限制请求计数，如果不存在则初始化为1
			rateLimitedValue, exists := m.GetCustomMetric("rate_limited_requests")
			var count int64 = 1
			if exists {
				if val, ok := rateLimitedValue.(int64); ok {
					count = val + 1
				}
			}

			// 更新限制请求计数
			m.AddCustomMetric("rate_limited_requests", count)

			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    429,
				"message": "请求速率超限，请稍后再试",
			})
			return
		}

		c.Next()
	}
}

// allow 检查是否允许请求
func (rl *RateLimiter) allow(clientIP string, now time.Time) bool {
	client, exists := rl.clients[clientIP]

	// 如果客户端记录不存在或窗口已过期，创建新记录
	if !exists || now.Sub(client.windowStart) > rl.window {
		rl.clients[clientIP] = &clientRateLimit{
			requests:    1,
			windowStart: now,
		}
		return true
	}

	// 检查请求数是否超限
	if client.requests >= rl.maxRequests {
		return false
	}

	// 增加请求计数
	client.requests++
	return true
}

// SlowQueryThreshold 慢查询阈值
var SlowQueryThreshold = 200 * time.Millisecond

// RecordDBQuery 记录数据库查询的辅助函数
func RecordDBQuery(tableName string, startTime time.Time, err error) {
	duration := time.Since(startTime)
	isError := err != nil
	isSlowQuery := duration > SlowQueryThreshold

	// 记录查询
	m := GetMetrics()
	m.RecordDatabaseQuery(tableName, duration, isError, isSlowQuery)
}

// PprofMiddleware 为应用添加pprof监控的中间件
// 注意：仅在开发环境使用，生产环境应谨慎启用
func PprofMiddleware(r *gin.Engine) {
	pprofGroup := r.Group("/debug/pprof")
	{
		pprofGroup.GET("/", gin.WrapF(http.DefaultServeMux.ServeHTTP))
		pprofGroup.GET("/cmdline", gin.WrapF(http.DefaultServeMux.ServeHTTP))
		pprofGroup.GET("/profile", gin.WrapF(http.DefaultServeMux.ServeHTTP))
		pprofGroup.GET("/symbol", gin.WrapF(http.DefaultServeMux.ServeHTTP))
		pprofGroup.GET("/trace", gin.WrapF(http.DefaultServeMux.ServeHTTP))
		pprofGroup.GET("/allocs", gin.WrapF(http.DefaultServeMux.ServeHTTP))
		pprofGroup.GET("/block", gin.WrapF(http.DefaultServeMux.ServeHTTP))
		pprofGroup.GET("/goroutine", gin.WrapF(http.DefaultServeMux.ServeHTTP))
		pprofGroup.GET("/heap", gin.WrapF(http.DefaultServeMux.ServeHTTP))
		pprofGroup.GET("/mutex", gin.WrapF(http.DefaultServeMux.ServeHTTP))
		pprofGroup.GET("/threadcreate", gin.WrapF(http.DefaultServeMux.ServeHTTP))
	}
}
