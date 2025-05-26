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

	// 添加 CORS 中间件
	r.Use(middleware.CORSMiddleware())

	r.SetTrustedProxies([]string{"127.0.0.1", "localhost"})

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务器
	logger.Info("服务启动成功，监听端口: %d，版本: %s", config.GetConfig().App.Port, appVersion)

	if err := r.Run(fmt.Sprintf(":%d", config.GetConfig().App.Port)); err != nil {
		panic(err)
	}
}
