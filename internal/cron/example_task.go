package cron

import (
	"log"
	"template/pkg/logger"
)

// registerExampleTask 注册示例任务
func registerExampleTask() {
	// 每分钟执行一次的任务
	_, err := cronManager.AddFunc("0 * * * * *", func() {
		logger.Info("示例定时任务执行")
	})

	if err != nil {
		log.Printf("注册示例定时任务失败: %v", err)
	}

	// 每天凌晨00:00执行的任务
	_, err = cronManager.AddFunc("0 0 0 * * *", func() {
		logger.Info("每日定时任务执行")
	})

	if err != nil {
		log.Printf("注册每日定时任务失败: %v", err)
	}
} 