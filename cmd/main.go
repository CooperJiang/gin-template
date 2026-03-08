package main

import (
	"fmt"
	"template/internal/app"
	"template/internal/routes"
	"template/pkg/config"
	"template/pkg/logger"
)

// 应用版本号
const appVersion = "1.0.0"

func main() {
	runtime, err := app.Bootstrap()
	if err != nil {
		panic(err)
	}
	defer runtime.Shutdown()

	routes.RegisterRoutes(runtime.Engine, runtime.Dependencies)

	// 启动服务器
	logger.Info("服务启动成功，监听端口: %d，版本: %s", config.GetConfig().App.Port, appVersion)
	addr := fmt.Sprintf(":%d", config.GetConfig().App.Port)
	if err := app.RunHTTPServer(runtime.Engine, addr); err != nil {
		panic(err)
	}
}
