package main

import (
	"fmt"
	"template/internal/cron"
	"template/internal/routes"
	"template/internal/services/user"
	"template/pkg/cache"
	"template/pkg/config"
	"template/pkg/database"
	"template/pkg/email"
	"template/pkg/errors"
	"template/pkg/health"
	"template/pkg/logger"
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

	// 设置健康检查系统应用版本
	health.SetVersion(appVersion)

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

	// 添加 CORS 中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	r.SetTrustedProxies([]string{"127.0.0.1", "localhost"})

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务器
	cfg := config.GetConfig().App
	logger.Info("服务启动成功，监听端口: %d，版本: %s", cfg.Port, appVersion)
	if err := r.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}
