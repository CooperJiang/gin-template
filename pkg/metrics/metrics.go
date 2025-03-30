package metrics

import (
	"runtime"
	"sync"
	"time"
)

// SystemStats 系统级别的性能统计
type SystemStats struct {
	NumGoroutine   int
	Alloc          uint64
	TotalAlloc     uint64
	Sys            uint64
	NumGC          uint32
	HeapObjects    uint64
	StackInUse     uint64
	HeapInUse      uint64
	GCPauseTotal   time.Duration
	GCPausePercent float64
}

// RequestStats 请求级别的性能统计
type RequestStats struct {
	TotalRequests      uint64
	ActiveRequests     int64
	SuccessRequests    uint64
	ErrorRequests      uint64
	TotalResponseTime  time.Duration
	MaxResponseTime    time.Duration
	MinResponseTime    time.Duration
	AvgResponseTime    time.Duration
	RequestsPerSecond  float64
	StatusCodeCount    map[int]uint64
	RouteCounts        map[string]uint64
	RouteResponseTimes map[string]time.Duration
}

// DatabaseStats 数据库级别的性能统计
type DatabaseStats struct {
	TotalQueries        uint64
	ActiveConnections   int
	MaxConnections      int
	TotalConnectionTime time.Duration
	MaxQueryTime        time.Duration
	AvgQueryTime        time.Duration
	ErrorQueries        uint64
	SlowQueries         uint64
	QueriesPerSecond    float64
	TableAccesses       map[string]uint64
}

// MemoryStats 缓存系统的性能统计
type MemoryStats struct {
	CacheHits       uint64
	CacheMisses     uint64
	CacheSize       uint64
	CacheCapacity   uint64
	CacheItemCount  uint64
	HitRatio        float64
	AvgLookupTime   time.Duration
	EvictionCount   uint64
	ExpirationCount uint64
}

// Metrics 集中管理所有性能指标
type Metrics struct {
	startTime     time.Time
	mu            sync.RWMutex
	System        SystemStats
	Request       RequestStats
	Database      DatabaseStats
	Cache         MemoryStats
	customMetrics map[string]interface{}
	lastUpdate    time.Time
	updateCounter uint64
}

// 单例模式的指标管理器
var (
	instance *Metrics
	once     sync.Once
)

// GetMetrics 获取指标管理器实例
func GetMetrics() *Metrics {
	once.Do(func() {
		instance = &Metrics{
			startTime:     time.Now(),
			lastUpdate:    time.Now(),
			customMetrics: make(map[string]interface{}),
			Request: RequestStats{
				StatusCodeCount:    make(map[int]uint64),
				RouteCounts:        make(map[string]uint64),
				RouteResponseTimes: make(map[string]time.Duration),
				MinResponseTime:    time.Hour, // 初始设为很大的值
			},
			Database: DatabaseStats{
				TableAccesses: make(map[string]uint64),
			},
		}
	})
	return instance
}

// CollectSystemStats 收集系统级别的性能指标
func (m *Metrics) CollectSystemStats() {
	m.mu.Lock()
	defer m.mu.Unlock()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	m.System.NumGoroutine = runtime.NumGoroutine()
	m.System.Alloc = memStats.Alloc
	m.System.TotalAlloc = memStats.TotalAlloc
	m.System.Sys = memStats.Sys
	m.System.NumGC = memStats.NumGC
	m.System.HeapObjects = memStats.HeapObjects
	m.System.HeapInUse = memStats.HeapInuse
	m.System.StackInUse = memStats.StackInuse

	// 计算GC暂停时间百分比
	gcTime := time.Duration(memStats.PauseTotalNs)
	totalTime := time.Since(m.startTime)
	if totalTime > 0 {
		m.System.GCPausePercent = float64(gcTime) / float64(totalTime) * 100
	}
	m.System.GCPauseTotal = time.Duration(memStats.PauseTotalNs)

	m.updateCounter++
	m.lastUpdate = time.Now()
}

// RecordRequest 记录请求开始
func (m *Metrics) RecordRequest(path string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Request.TotalRequests++
	m.Request.ActiveRequests++
	m.Request.RouteCounts[path]++
}

// RecordRequestComplete 记录请求完成
func (m *Metrics) RecordRequestComplete(path string, statusCode int, duration time.Duration, isError bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 更新请求统计
	m.Request.ActiveRequests--
	m.Request.TotalResponseTime += duration
	m.Request.StatusCodeCount[statusCode]++
	m.Request.RouteResponseTimes[path] += duration

	if isError {
		m.Request.ErrorRequests++
	} else {
		m.Request.SuccessRequests++
	}

	// 更新最大最小响应时间
	if duration > m.Request.MaxResponseTime {
		m.Request.MaxResponseTime = duration
	}
	if duration < m.Request.MinResponseTime {
		m.Request.MinResponseTime = duration
	}

	// 计算平均响应时间
	totalRequests := m.Request.SuccessRequests + m.Request.ErrorRequests
	if totalRequests > 0 {
		m.Request.AvgResponseTime = m.Request.TotalResponseTime / time.Duration(totalRequests)
	}

	// 计算每秒请求数
	uptime := time.Since(m.startTime).Seconds()
	if uptime > 0 {
		m.Request.RequestsPerSecond = float64(totalRequests) / uptime
	}
}

// RecordDatabaseQuery 记录数据库查询
func (m *Metrics) RecordDatabaseQuery(tableName string, duration time.Duration, isError bool, isSlowQuery bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Database.TotalQueries++
	m.Database.TotalConnectionTime += duration
	m.Database.TableAccesses[tableName]++

	if isError {
		m.Database.ErrorQueries++
	}

	if isSlowQuery {
		m.Database.SlowQueries++
	}

	if duration > m.Database.MaxQueryTime {
		m.Database.MaxQueryTime = duration
	}

	if m.Database.TotalQueries > 0 {
		m.Database.AvgQueryTime = m.Database.TotalConnectionTime / time.Duration(m.Database.TotalQueries)
	}

	uptime := time.Since(m.startTime).Seconds()
	if uptime > 0 {
		m.Database.QueriesPerSecond = float64(m.Database.TotalQueries) / uptime
	}
}

// RecordCacheOperation 记录缓存操作
func (m *Metrics) RecordCacheOperation(hit bool, size uint64, lookupTime time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if hit {
		m.Cache.CacheHits++
	} else {
		m.Cache.CacheMisses++
	}

	totalOps := m.Cache.CacheHits + m.Cache.CacheMisses
	if totalOps > 0 {
		m.Cache.HitRatio = float64(m.Cache.CacheHits) / float64(totalOps)
	}

	m.Cache.CacheSize = size
}

// SetActiveDatabaseConnections 设置活跃数据库连接数
func (m *Metrics) SetActiveDatabaseConnections(active, max int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Database.ActiveConnections = active
	m.Database.MaxConnections = max
}

// SetCacheStats 设置缓存统计信息
func (m *Metrics) SetCacheStats(itemCount uint64, capacity uint64, evictionCount uint64, expirationCount uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Cache.CacheItemCount = itemCount
	m.Cache.CacheCapacity = capacity
	m.Cache.EvictionCount = evictionCount
	m.Cache.ExpirationCount = expirationCount
}

// AddCustomMetric 添加自定义指标
func (m *Metrics) AddCustomMetric(name string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.customMetrics[name] = value
}

// GetCustomMetric 获取自定义指标
func (m *Metrics) GetCustomMetric(name string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, exists := m.customMetrics[name]
	return value, exists
}

// GetAllMetrics 获取所有指标的快照
func (m *Metrics) GetAllMetrics() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// 更新系统指标
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	return map[string]interface{}{
		"system": map[string]interface{}{
			"uptime":           time.Since(m.startTime).String(),
			"start_time":       m.startTime.Format(time.RFC3339),
			"num_goroutine":    runtime.NumGoroutine(),
			"num_cpu":          runtime.NumCPU(),
			"alloc_bytes":      memStats.Alloc,
			"total_alloc":      memStats.TotalAlloc,
			"heap_objects":     memStats.HeapObjects,
			"gc_cycles":        memStats.NumGC,
			"gc_pause_total":   time.Duration(memStats.PauseTotalNs).String(),
			"gc_pause_percent": m.System.GCPausePercent,
		},
		"request": map[string]interface{}{
			"total":             m.Request.TotalRequests,
			"active":            m.Request.ActiveRequests,
			"success":           m.Request.SuccessRequests,
			"errors":            m.Request.ErrorRequests,
			"avg_response_time": m.Request.AvgResponseTime.String(),
			"max_response_time": m.Request.MaxResponseTime.String(),
			"min_response_time": m.Request.MinResponseTime.String(),
			"requests_per_sec":  m.Request.RequestsPerSecond,
			"status_codes":      m.Request.StatusCodeCount,
			"routes":            m.Request.RouteCounts,
		},
		"database": map[string]interface{}{
			"total_queries":      m.Database.TotalQueries,
			"active_connections": m.Database.ActiveConnections,
			"max_connections":    m.Database.MaxConnections,
			"error_queries":      m.Database.ErrorQueries,
			"slow_queries":       m.Database.SlowQueries,
			"avg_query_time":     m.Database.AvgQueryTime.String(),
			"max_query_time":     m.Database.MaxQueryTime.String(),
			"queries_per_sec":    m.Database.QueriesPerSecond,
			"table_accesses":     m.Database.TableAccesses,
		},
		"cache": map[string]interface{}{
			"hits":             m.Cache.CacheHits,
			"misses":           m.Cache.CacheMisses,
			"hit_ratio":        m.Cache.HitRatio,
			"size":             m.Cache.CacheSize,
			"capacity":         m.Cache.CacheCapacity,
			"item_count":       m.Cache.CacheItemCount,
			"eviction_count":   m.Cache.EvictionCount,
			"expiration_count": m.Cache.ExpirationCount,
		},
		"custom":      m.customMetrics,
		"last_update": m.lastUpdate.Format(time.RFC3339),
	}
}

// ResetStats 重置所有统计数据（除了系统级别的）
func (m *Metrics) ResetStats() {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 保留启动时间
	startTime := m.startTime

	// 重置请求统计
	m.Request = RequestStats{
		StatusCodeCount:    make(map[int]uint64),
		RouteCounts:        make(map[string]uint64),
		RouteResponseTimes: make(map[string]time.Duration),
		MinResponseTime:    time.Hour, // 初始设为很大的值
	}

	// 重置数据库统计
	m.Database = DatabaseStats{
		TableAccesses: make(map[string]uint64),
	}

	// 重置缓存统计
	m.Cache = MemoryStats{}

	// 重置自定义指标
	m.customMetrics = make(map[string]interface{})

	// 恢复启动时间
	m.startTime = startTime
	m.lastUpdate = time.Now()
}
