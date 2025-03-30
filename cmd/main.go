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
	"template/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

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

	// 使用 gin.Default() 替代 gin.New()
	r := gin.Default()

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
	if err := r.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}
