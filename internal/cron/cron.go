package cron

import (
	"log"

	"github.com/robfig/cron/v3"
)

var cronManager *cron.Cron

// InitCronManager 初始化定时任务管理器
func InitCronManager() {
	// 创建一个支持秒级别的cron管理器
	cronManager = cron.New(cron.WithSeconds())

	// 注册所有定时任务
	registerTasks()

	// 启动定时任务
	cronManager.Start()
	log.Println("定时任务管理器已启动")
}

// registerTasks 注册所有定时任务
func registerTasks() {
	// 注册示例任务
	registerExampleTask()

	// 在这里注册其他定时任务
	// registerOtherTask()
}

// Stop 停止所有定时任务
func Stop() {
	if cronManager != nil {
		cronManager.Stop()
		log.Println("定时任务管理器已停止")
	}
} 