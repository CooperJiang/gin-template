package main

import (
	"fmt"
	"template/internal/cron"
	"template/internal/middleware"
	"template/internal/routes"
	"template/internal/services/user"
	"template/pkg/cache"
	"template/pkg/config"
	"template/pkg/database"
	"template/pkg/email"
	"template/pkg/errors"
	"template/pkg/logger"
	"template/pkg/metrics"
	"time"

	"github.com/gin-gonic/gin"
)

// 应用版本号
const appVersion = "1.0.0"

func main() {
	// 设置时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = loc

	// 初始化各种服务
	logger.Init()
	config.InitConfig()
	database.InitDB()
	defer database.Close()
	cache.InitCache()
	email.Init()

	// 初始化定时任务
	cron.InitCronManager()
	defer cron.Stop()

	// 初始化基础服务
	user.InitUserService()

	// 设置 gin 模式
	gin.SetMode(config.GetConfig().App.Mode)

	// 使用 gin.New() 替代 gin.Default() 以便完全控制中间件
	r := gin.New()

	// 添加 recovery 和 logger 中间件
	r.Use(gin.Recovery())

	// 添加错误处理和请求ID中间件
	r.Use(errors.ErrorHandler())

	// 配置
	cfg := config.GetConfig().App

	// 添加性能监控中间件
	if cfg.EnableMetrics {
		r.Use(metrics.MetricsMiddleware())
		logger.Info("性能指标监控已启用")
	}

	// 如果是开发模式或明确启用，启用pprof调试
	if cfg.EnablePprof || cfg.Mode == gin.DebugMode {
		metrics.PprofMiddleware(r)
		logger.Info("pprof性能分析工具已启用，访问 /debug/pprof 查看")
	}

	// 添加请求速率限制
	if cfg.EnableRateLimit {
		// 获取配置的速率限制参数，如果未配置则使用默认值
		maxRequests := cfg.RateLimitRequests
		if maxRequests <= 0 {
			maxRequests = 100 // 默认每分钟100个请求
		}

		windowSeconds := cfg.RateLimitWindow
		if windowSeconds <= 0 {
			windowSeconds = 60 // 默认1分钟窗口
		}

		windowDuration := time.Duration(windowSeconds) * time.Second
		rateLimiter := metrics.NewRateLimiter(maxRequests, windowDuration)
		r.Use(rateLimiter.Middleware())
		logger.Info("请求速率限制已启用: %d请求/%v", maxRequests, windowDuration)
	}

	// 添加 CORS 中间件
	r.Use(middleware.CORSMiddleware())

	r.SetTrustedProxies([]string{"127.0.0.1", "localhost"})

	// 注册路由
	routes.RegisterRoutes(r)

	// 注册指标接口
	if cfg.EnableMetrics {
		metrics.RegisterMetricsHandlers(r)
		logger.Info("性能指标接口已注册，访问 /metrics 查看")

		// 配置慢查询阈值
		if cfg.SlowQueryThreshold > 0 {
			metrics.SlowQueryThreshold = time.Duration(cfg.SlowQueryThreshold) * time.Millisecond
			logger.Info("慢查询阈值已设置为 %dms", cfg.SlowQueryThreshold)
		}

		// 整合数据库指标监控
		db := database.GetDB()
		if db != nil {
			// 创建并初始化GORM指标插件
			slowThreshold := time.Duration(cfg.SlowQueryThreshold)
			if slowThreshold <= 0 {
				slowThreshold = 200 // 默认200毫秒
			}
			dbMetricsPlugin := metrics.NewGormPlugin(slowThreshold*time.Millisecond, true)
			if err := dbMetricsPlugin.Initialize(db); err != nil {
				logger.Error("初始化数据库指标监控失败: %v", err)
			} else {
				logger.Info("数据库指标监控已启用")
			}
		}

		// 整合Redis指标监控
		if cache.RedisAvailable() {
			// 获取Redis客户端并包装
			redisClient := cache.GetRedisClient()
			metrics.NewRedisMetricsWrapper(redisClient)
			logger.Info("Redis指标监控已启用")
		}

		// 定期记录系统指标
		interval := cfg.MetricsLogInterval
		if interval <= 0 {
			interval = 5 // 默认5分钟
		}

		go func() {
			// 立即收集一次指标
			metrics.GetMetrics().CollectSystemStats()

			// 定期收集指标
			ticker := time.NewTicker(time.Duration(interval) * time.Minute)
			defer ticker.Stop()

			for range ticker.C {
				metrics.GetMetrics().CollectSystemStats()
				logger.Debug("系统指标已收集")
			}
		}()
	}

	// 启动服务器
	logger.Info("服务启动成功，监听端口: %d，版本: %s", cfg.Port, appVersion)

	if err := r.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}
