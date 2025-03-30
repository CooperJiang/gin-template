package metrics

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// GormMetricsPlugin GORM插件用于收集数据库性能指标
type GormMetricsPlugin struct {
	metrics         *Metrics
	slowQueryTime   time.Duration
	enableTableStat bool
}

// NewGormPlugin 创建新的GORM指标收集插件
func NewGormPlugin(slowQueryThreshold time.Duration, enableTableStat bool) *GormMetricsPlugin {
	return &GormMetricsPlugin{
		metrics:         GetMetrics(),
		slowQueryTime:   slowQueryThreshold,
		enableTableStat: enableTableStat,
	}
}

// Name 插件名称
func (p *GormMetricsPlugin) Name() string {
	return "GormMetricsPlugin"
}

// Initialize 初始化插件
func (p *GormMetricsPlugin) Initialize(db *gorm.DB) error {
	// 添加回调函数
	if err := db.Callback().Create().Before("gorm:create").Register("metrics:before_create", p.before()); err != nil {
		return err
	}
	if err := db.Callback().Create().After("gorm:create").Register("metrics:after_create", p.after("create")); err != nil {
		return err
	}

	if err := db.Callback().Query().Before("gorm:query").Register("metrics:before_query", p.before()); err != nil {
		return err
	}
	if err := db.Callback().Query().After("gorm:query").Register("metrics:after_query", p.after("query")); err != nil {
		return err
	}

	if err := db.Callback().Update().Before("gorm:update").Register("metrics:before_update", p.before()); err != nil {
		return err
	}
	if err := db.Callback().Update().After("gorm:update").Register("metrics:after_update", p.after("update")); err != nil {
		return err
	}

	if err := db.Callback().Delete().Before("gorm:delete").Register("metrics:before_delete", p.before()); err != nil {
		return err
	}
	if err := db.Callback().Delete().After("gorm:delete").Register("metrics:after_delete", p.after("delete")); err != nil {
		return err
	}

	// 定时收集数据库连接统计
	go p.collectDBStats(db)

	return nil
}

// before 执行操作前的回调
func (p *GormMetricsPlugin) before() func(db *gorm.DB) {
	return func(db *gorm.DB) {
		startTime := time.Now()
		db.Set("metrics:start_time", startTime)
	}
}

// after 执行操作后的回调
func (p *GormMetricsPlugin) after(operation string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		// 从上下文获取开始时间
		val, exists := db.Get("metrics:start_time")
		if !exists {
			return
		}
		startTime, ok := val.(time.Time)
		if !ok {
			return
		}

		// 计算执行时间
		duration := time.Since(startTime)

		// 获取表名
		var tableName string
		if db.Statement.Table != "" {
			tableName = db.Statement.Table
		} else if db.Statement.Schema != nil {
			tableName = db.Statement.Schema.Table
		} else {
			tableName = "unknown"
		}

		// 判断是否为慢查询
		isSlowQuery := duration >= p.slowQueryTime

		// 记录指标
		p.metrics.RecordDatabaseQuery(tableName, duration, db.Error != nil, isSlowQuery)

		// 记录表级别的统计（可选）
		if p.enableTableStat {
			metricName := "db_table_" + tableName + "_" + operation + "_count"
			value, exists := p.metrics.GetCustomMetric(metricName)
			var count int64 = 1
			if exists {
				if val, ok := value.(int64); ok {
					count = val + 1
				}
			}
			p.metrics.AddCustomMetric(metricName, count)
		}
	}
}

// collectDBStats 定期收集数据库连接统计
func (p *GormMetricsPlugin) collectDBStats(db *gorm.DB) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// 对于GORM v2，我们需要使用原始的sqlDB
		sqlDB, err := db.DB()
		if err != nil {
			continue
		}

		stats := sqlDB.Stats()
		p.metrics.SetActiveDatabaseConnections(stats.InUse, stats.MaxOpenConnections)

		// 添加更多连接池指标
		p.metrics.AddCustomMetric("db_open_connections", stats.OpenConnections)
		p.metrics.AddCustomMetric("db_idle_connections", stats.Idle)
		p.metrics.AddCustomMetric("db_wait_count", stats.WaitCount)
		p.metrics.AddCustomMetric("db_max_idle_closed", stats.MaxIdleClosed)
		p.metrics.AddCustomMetric("db_max_lifetime_closed", stats.MaxLifetimeClosed)
	}
}

// CreateDBMetricsContext 创建带有指标的数据库上下文
func CreateDBMetricsContext(ctx context.Context, operation string, tableName string) context.Context {
	ctx = context.WithValue(ctx, "db_metrics_start_time", time.Now())
	ctx = context.WithValue(ctx, "db_metrics_operation", operation)
	ctx = context.WithValue(ctx, "db_metrics_table", tableName)
	return ctx
}

// RecordDBMetricsFromContext 从上下文中记录数据库指标
func RecordDBMetricsFromContext(ctx context.Context, err error) {
	// 获取开始时间
	startTimeVal := ctx.Value("db_metrics_start_time")
	if startTimeVal == nil {
		return
	}
	startTime, ok := startTimeVal.(time.Time)
	if !ok {
		return
	}

	// 获取操作和表名
	operation := ctx.Value("db_metrics_operation")
	tableName := ctx.Value("db_metrics_table")

	tableNameStr := "unknown"
	if tableName != nil {
		if str, ok := tableName.(string); ok {
			tableNameStr = str
		}
	}

	// 记录指标
	RecordDBQuery(tableNameStr, startTime, err)

	// 可以添加更多自定义指标
	if operation != nil {
		if opStr, ok := operation.(string); ok {
			m := GetMetrics()
			metricName := "db_operation_" + opStr + "_count"
			value, exists := m.GetCustomMetric(metricName)
			var count int64 = 1
			if exists {
				if val, ok := value.(int64); ok {
					count = val + 1
				}
			}
			m.AddCustomMetric(metricName, count)
		}
	}
}
