package metrics

import (
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// MonitorRedisStats 定期收集Redis统计信息
func MonitorRedisStats(client *redis.Client) {
	m := GetMetrics()

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			// 获取Redis INFO统计信息
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			infoAll, err := client.Info(ctx).Result()
			cancel()

			if err != nil {
				// 记录错误
				m.AddCustomMetric("redis_info_error_count", 1)
				continue
			}

			// 解析统计信息 (简化版)
			// 实际应用中可以添加更详细的解析逻辑
			if len(infoAll) > 0 {
				// 示例：提取Redis的内存使用信息
				if strings.Contains(infoAll, "used_memory:") {
					// 实际解析代码略
					m.AddCustomMetric("redis_info_parsed", true)
				}
			}

			// 记录连接池信息
			poolStats := client.PoolStats()
			m.AddCustomMetric("redis_total_conns", poolStats.TotalConns)
			m.AddCustomMetric("redis_idle_conns", poolStats.IdleConns)
			m.AddCustomMetric("redis_stale_conns", poolStats.StaleConns)

			// 设置缓存统计信息 (示例值)
			// 实际应用中这些值应该从Redis INFO中解析获得
			m.SetCacheStats(
				10000,  // 项目数量 (示例)
				100000, // 容量 (示例)
				1000,   // 驱逐数 (示例)
				500,    // 过期数 (示例)
			)
		}
	}()
}

// RedisMetricsWrapper Redis操作指标包装器
type RedisMetricsWrapper struct {
	client  *redis.Client
	metrics *Metrics
}

// NewRedisMetricsWrapper 创建新的Redis指标包装器
func NewRedisMetricsWrapper(client *redis.Client) *RedisMetricsWrapper {
	// 启动监控
	MonitorRedisStats(client)

	// 创建包装器
	wrapper := &RedisMetricsWrapper{
		client:  client,
		metrics: GetMetrics(),
	}

	// 添加命令监听
	wrapper.MonitorCommands()

	return wrapper
}

// MonitorCommands 使用拦截器监控Redis命令
func (w *RedisMetricsWrapper) MonitorCommands() {
	// 使用自定义函数记录Redis操作指标
	go func() {
		// 每30秒采样一些命令执行情况
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		// 基本命令监控
		for range ticker.C {
			w.sampleRedisCommands()
		}
	}()
}

// sampleRedisCommands 采样Redis命令
func (w *RedisMetricsWrapper) sampleRedisCommands() {
	// 执行并记录GET命令
	ctx := context.Background()
	startTime := time.Now()
	_, err := w.client.Get(ctx, "_metrics_sample_key").Result()
	duration := time.Since(startTime)
	isHit := err != redis.Nil

	// 记录GET操作
	w.metrics.RecordCacheOperation(isHit, 0, duration)
	w.metrics.AddCustomMetric("redis_cmd_get_sample_time_ms", duration.Milliseconds())

	// 执行并记录PING命令
	startTime = time.Now()
	_, err = w.client.Ping(ctx).Result()
	duration = time.Since(startTime)

	// 记录PING操作
	w.metrics.AddCustomMetric("redis_ping_time_ms", duration.Milliseconds())
	w.metrics.AddCustomMetric("redis_available", err == nil)
}

// GetClient 获取原始Redis客户端
func (w *RedisMetricsWrapper) GetClient() *redis.Client {
	return w.client
}

// RecordCacheHit 记录缓存命中
func (w *RedisMetricsWrapper) RecordCacheHit(key string, size uint64) {
	w.metrics.RecordCacheOperation(true, size, 0)

	// 记录键级别的统计（可选）
	metricName := "cache_key_" + key + "_hit_count"
	value, exists := w.metrics.GetCustomMetric(metricName)
	var count int64 = 1
	if exists {
		if val, ok := value.(int64); ok {
			count = val + 1
		}
	}
	w.metrics.AddCustomMetric(metricName, count)
}

// RecordCacheMiss 记录缓存未命中
func (w *RedisMetricsWrapper) RecordCacheMiss(key string) {
	w.metrics.RecordCacheOperation(false, 0, 0)

	// 记录键级别的统计（可选）
	metricName := "cache_key_" + key + "_miss_count"
	value, exists := w.metrics.GetCustomMetric(metricName)
	var count int64 = 1
	if exists {
		if val, ok := value.(int64); ok {
			count = val + 1
		}
	}
	w.metrics.AddCustomMetric(metricName, count)
}
