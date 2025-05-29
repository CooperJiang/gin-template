package database

import (
	"fmt"
	"template/internal/models"
	"template/pkg/common"
	"template/pkg/config"
	log "template/pkg/logger"
	"template/pkg/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return db
}

// InitDB 初始化数据库连接
func InitDB() {
	cfg := config.GetConfig().Database

	// 配置 GORM
	gormConfig := &gorm.Config{
		// 使用自定义 logger，并设置日志级别
		Logger: logger.Default.LogMode(logger.Silent), // 修改这里来控制日志级别
		NowFunc: func() time.Time {
			return time.Now().In(time.Local)
		},

		// 其他配置保持不变
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	var err error

	// 检查是否配置了MySQL数据库
	if cfg.Host == "" || cfg.Username == "" || cfg.Name == "" {
		// 未配置MySQL，使用SQLite
		log.Info("未检测到MySQL配置，将使用SQLite数据库")
		db, err = gorm.Open(sqlite.Open("app.db"), gormConfig)
		if err != nil {
			log.Fatal("连接SQLite数据库失败: %v", err)
		}
	} else {
		// 尝试使用MySQL数据库
		// 构建 DSN
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Asia%%2FShanghai",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Name,
			cfg.Charset,
		)

		// 连接数据库
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)
		if err != nil {
			log.Error("连接MySQL数据库失败: %v", err)
			log.Info("自动降级到SQLite数据库")
			// 降级到SQLite
			db, err = gorm.Open(sqlite.Open("app.db"), gormConfig)
			if err != nil {
				log.Fatal("连接SQLite数据库失败: %v", err)
			}
		} else {
			log.Info("成功连接到MySQL数据库")
		}
	}

	// 自动迁移
	if err := autoMigrate(); err != nil {
		log.Fatal("数据库迁移失败: %v", err)
	}

	// 检查并创建 root 用户
	if err := createRootUserIfNotExists(); err != nil {
		log.Fatal("创建 root 用户失败: %v", err)
	}

	log.Info("数据库连接成功")
}

// autoMigrate 自动迁移数据库结构
func autoMigrate() error {
	return db.AutoMigrate(
		&models.User{},
		&models.UploadFile{},
		&models.ChunkInfo{},
		// 文件权限管理模型
		&models.FileShare{},
		&models.FilePermission{},
		&models.TemporaryAccess{},
		// 在这里添加其他模型
	)
}

// 检查并创建 root 用户
func createRootUserIfNotExists() error {
	var count int64
	if err := db.Model(&models.User{}).Where("username = ?", "root").Count(&count).Error; err != nil {
		return err
	}

	// 如果 root 用户不存在，则创建
	if count == 0 {
		password := config.GetConfig().App.DefaultRootPass
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return err
		}
		rootUser := models.User{
			Username: "root",
			Password: hashedPassword,
			Role:     common.UserRoleSuperAdmin,
		}

		if err := db.Create(&rootUser).Error; err != nil {
			return err
		}
		log.Info("已成功创建 root 管理员用户, 密码是: %s", config.GetConfig().App.DefaultRootPass)
	}

	return nil
}

// Close 关闭数据库连接
func Close() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
