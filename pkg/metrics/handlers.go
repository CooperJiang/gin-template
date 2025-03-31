package metrics

import (
	"template/pkg/errors"
	"time"

	"github.com/gin-gonic/gin"
)

// MetricsHandler 返回所有指标数据
func MetricsHandler(c *gin.Context) {
	m := GetMetrics()

	// 确保系统指标是最新的
	m.CollectSystemStats()

	// 返回所有指标
	metrics := m.GetAllMetrics()
	errors.ResponseSuccess(c, metrics, "获取性能指标成功")
}

// MetricsResetHandler 重置指标数据
func MetricsResetHandler(c *gin.Context) {
	m := GetMetrics()
	m.ResetStats()
	errors.ResponseSuccess(c, nil, "性能指标已重置")
}

// MetricsSystemHandler 只返回系统级指标
func MetricsSystemHandler(c *gin.Context) {
	m := GetMetrics()
	m.CollectSystemStats()

	metrics := m.GetAllMetrics()
	if systemInfo, ok := metrics["system"].(map[string]interface{}); ok {
		errors.ResponseSuccess(c, systemInfo, "获取系统指标成功")
		return
	}

	errors.ResponseSuccess(c, nil, "获取系统指标失败")
}

// MetricsRequestHandler 只返回请求相关指标
func MetricsRequestHandler(c *gin.Context) {
	m := GetMetrics()

	metrics := m.GetAllMetrics()
	if requestInfo, ok := metrics["request"].(map[string]interface{}); ok {
		errors.ResponseSuccess(c, requestInfo, "获取请求指标成功")
		return
	}

	errors.ResponseSuccess(c, nil, "获取请求指标失败")
}

// MetricsDatabaseHandler 只返回数据库相关指标
func MetricsDatabaseHandler(c *gin.Context) {
	m := GetMetrics()

	metrics := m.GetAllMetrics()
	if dbInfo, ok := metrics["database"].(map[string]interface{}); ok {
		errors.ResponseSuccess(c, dbInfo, "获取数据库指标成功")
		return
	}

	errors.ResponseSuccess(c, nil, "获取数据库指标失败")
}

// MetricsCacheHandler 只返回缓存相关指标
func MetricsCacheHandler(c *gin.Context) {
	m := GetMetrics()

	metrics := m.GetAllMetrics()
	if cacheInfo, ok := metrics["cache"].(map[string]interface{}); ok {
		errors.ResponseSuccess(c, cacheInfo, "获取缓存指标成功")
		return
	}

	errors.ResponseSuccess(c, nil, "获取缓存指标失败")
}

// RegisterMetricsHandlers 注册指标相关路由
func RegisterMetricsHandlers(r *gin.Engine) {
	// 指标路由组
	metricsGroup := r.Group("/api/v1/metrics")
	{
		metricsGroup.GET("", MetricsHandler)                  // 获取所有指标
		metricsGroup.GET("/system", MetricsSystemHandler)     // 系统指标
		metricsGroup.GET("/requests", MetricsRequestHandler)  // 请求指标
		metricsGroup.GET("/database", MetricsDatabaseHandler) // 数据库指标
		metricsGroup.GET("/cache", MetricsCacheHandler)       // 缓存指标
		metricsGroup.POST("/reset", MetricsResetHandler)      // 重置指标
	}

	// 每分钟后台收集一次系统指标
	go func() {
		ticker := time.NewTicker(time.Minute)
		for range ticker.C {
			GetMetrics().CollectSystemStats()
		}
	}()
}

// 数据库和缓存指标集成工具

// IntegrateWithDatabase 与数据库集成
func IntegrateWithDatabase(db interface{}) {
	// 检查数据库连接并更新指标
	// 这里只是一个示例，具体实现需要根据实际使用的数据库ORM来调整
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for range ticker.C {
			// 假设DB对象有Stats()方法获取连接信息
			// stats := db.Stats()
			// m := GetMetrics()
			// m.SetActiveDatabaseConnections(stats.OpenConnections, stats.MaxOpenConnections)
		}
	}()
}

// IntegrateWithCache 与缓存系统集成
func IntegrateWithCache(cache interface{}) {
	// 定期收集缓存指标
	// 这里只是一个示例，具体实现需要根据实际使用的缓存系统来调整
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for range ticker.C {
			// 假设缓存对象有Stats()方法获取统计信息
			// stats := cache.Stats()
			// m := GetMetrics()
			// m.SetCacheStats(stats.ItemCount, stats.Capacity, stats.EvictionCount, stats.ExpirationCount)
		}
	}()
}
