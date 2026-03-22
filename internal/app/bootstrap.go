package app

import (
	"fmt"
	"strings"
	"email-manage/internal/cron"
	"email-manage/internal/middleware"
	"email-manage/pkg/cache"
	"email-manage/pkg/config"
	"email-manage/pkg/database"
	"email-manage/pkg/email"
	"email-manage/pkg/errors"
	"email-manage/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultTimezone = "Asia/Shanghai"

// Runtime 应用运行时对象
type Runtime struct {
	Engine       *gin.Engine
	Dependencies *Dependencies
}

// Bootstrap 初始化应用运行时所需基础设施
func Bootstrap() (*Runtime, error) {
	logger.Init()
	config.InitConfig()

	if err := setupTimezone(config.GetConfig().App.Timezone); err != nil {
		return nil, fmt.Errorf("初始化时区失败: %w", err)
	}

	database.InitDB()
	cache.InitCache()
	email.Init()

	cron.InitCronManager()

	gin.SetMode(config.GetConfig().App.Mode)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(errors.ErrorHandler())
	engine.Use(middleware.CORSMiddleware())

	if err := engine.SetTrustedProxies([]string{"127.0.0.1", "localhost"}); err != nil {
		logger.Warn("设置受信任代理失败: %v", err)
	}

	return &Runtime{
		Engine:       engine,
		Dependencies: NewDependencies(),
	}, nil
}

// Shutdown 关闭应用基础设施
func (rt *Runtime) Shutdown() {
	cron.Stop()

	if err := cache.Close(); err != nil {
		logger.Warn("关闭缓存失败: %v", err)
	}

	if err := database.Close(); err != nil {
		logger.Warn("关闭数据库失败: %v", err)
	}
}

func setupTimezone(tz string) error {
	timezone := strings.TrimSpace(tz)
	if timezone == "" {
		timezone = defaultTimezone
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	time.Local = loc
	return nil
}
