package database

import (
	"fmt"
	"strings"
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

	// 明确使用modernc.org/sqlite作为SQLite驱动
	_ "modernc.org/sqlite"
)

var db *gorm.DB

type initOptions struct {
	strictMode    bool
	autoMigrate   bool
	bootstrapRoot bool
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return db
}

// InitDB 初始化数据库连接
func InitDB() {
	cfg := config.GetConfig().Database
	opts := resolveInitOptions()

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
	db, err = openPreferredDatabase(cfg, gormConfig, opts)
	if err != nil {
		log.Fatal("初始化数据库连接失败: %v", err)
	}

	// 自动迁移
	if opts.autoMigrate {
		if err := autoMigrate(); err != nil {
			log.Fatal("数据库迁移失败: %v", err)
		}
	} else {
		log.Info("已跳过自动迁移（APP_DB_AUTO_MIGRATE=false）")
	}

	// 检查并创建 root 用户
	if opts.bootstrapRoot {
		if err := createRootUserIfNotExists(); err != nil {
			log.Fatal("创建 root 用户失败: %v", err)
		}
	} else {
		log.Info("已跳过 root 用户初始化（APP_DB_BOOTSTRAP_ROOT=false）")
	}

	log.Info("数据库连接成功")
}

func resolveInitOptions() initOptions {
	return resolveInitOptionsByMode(config.GetConfig().App.Mode, config.GetEnvBool)
}

func resolveInitOptionsByMode(mode string, getEnvBool func(key string, defaultValue bool) bool) initOptions {
	if getEnvBool == nil {
		getEnvBool = func(_ string, defaultValue bool) bool { return defaultValue }
	}

	isStrictDefault := isReleaseMode(mode)

	return initOptions{
		strictMode:    getEnvBool("APP_DB_STRICT", isStrictDefault),
		autoMigrate:   getEnvBool("APP_DB_AUTO_MIGRATE", !isStrictDefault),
		bootstrapRoot: getEnvBool("APP_DB_BOOTSTRAP_ROOT", !isStrictDefault),
	}
}

func isReleaseMode(mode string) bool {
	normalized := strings.ToLower(strings.TrimSpace(mode))
	return normalized == "release" || normalized == "production" || normalized == "prod"
}

func openPreferredDatabase(cfg config.DatabaseConfig, gormConfig *gorm.Config, opts initOptions) (*gorm.DB, error) {
	wantMySQL, mysqlReady := shouldUseMySQL(cfg)
	if !wantMySQL {
		log.Info("当前配置使用SQLite数据库")
		return openSQLite(gormConfig)
	}

	if !mysqlReady {
		if opts.strictMode {
			return nil, fmt.Errorf("MySQL配置不完整，当前运行模式禁止自动降级到SQLite")
		}
		log.Warn("MySQL配置不完整，自动降级到SQLite数据库")
		return openSQLite(gormConfig)
	}

	mysqlDB, err := openMySQL(cfg, gormConfig)
	if err == nil {
		log.Info("成功连接到MySQL数据库")
		return mysqlDB, nil
	}

	if opts.strictMode {
		return nil, fmt.Errorf("连接MySQL数据库失败: %w", err)
	}

	log.Error("连接MySQL数据库失败: %v", err)
	log.Info("自动降级到SQLite数据库")
	return openSQLite(gormConfig)
}

func shouldUseMySQL(cfg config.DatabaseConfig) (wantMySQL bool, mysqlReady bool) {
	driver := strings.ToLower(strings.TrimSpace(cfg.Driver))
	mysqlSignal := strings.TrimSpace(cfg.Host) != "" || strings.TrimSpace(cfg.Username) != "" || strings.TrimSpace(cfg.Name) != ""

	switch driver {
	case "sqlite":
		return false, false
	case "mysql":
		return true, isMySQLConfigReady(cfg)
	default:
		if mysqlSignal {
			return true, isMySQLConfigReady(cfg)
		}
		return false, false
	}
}

func isMySQLConfigReady(cfg config.DatabaseConfig) bool {
	return strings.TrimSpace(cfg.Host) != "" &&
		strings.TrimSpace(cfg.Username) != "" &&
		strings.TrimSpace(cfg.Name) != ""
}

func openMySQL(cfg config.DatabaseConfig, gormConfig *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Asia%%2FShanghai",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.Charset,
	)

	return gorm.Open(mysql.Open(dsn), gormConfig)
}

func openSQLite(gormConfig *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file:app.db?cache=shared&mode=rwc"), gormConfig)
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
		log.Info("已成功创建 root 管理员用户，请立即通过管理接口修改默认密码")
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
