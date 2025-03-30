package logger

import (
	"context"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

// Logger 自定义日志结构体
type Logger struct {
	*log.Logger
	config    *Config
	LogLevel  logger.LogLevel
	SlowQuery time.Duration
}

// Config 日志配置
type Config struct {
	SlowThreshold time.Duration // 慢查询阈值
	Colorful      bool          // 是否启用彩色输出
	LogLevel      logger.LogLevel
}

// 全局日志实例
var Log *Logger

// 全局方法
var (
	Infof  func(format string, args ...interface{})
	Warnf  func(format string, args ...interface{})
	Errorf func(format string, args ...interface{})
	Debugf func(format string, args ...interface{})
)

// 默认配置
var defaultConfig = &Config{
	SlowThreshold: 200 * time.Millisecond, // 慢查询阈值，超过这个时间的查询会被标记为慢查询
	Colorful:      true,                   // 默认开启彩色输出
	LogLevel:      logger.Info,            // 默认日志级别为 Info
}

// New 创建一个新的日志实例
func New(config *Config) *Logger {
	if config == nil {
		config = &Config{
			SlowThreshold: 200 * time.Millisecond,
			Colorful:      true,
			LogLevel:      logger.Info,
		}
	}

	l := &Logger{
		Logger:    log.New(os.Stdout, "", log.LstdFlags),
		config:    config,
		LogLevel:  config.LogLevel,
		SlowQuery: config.SlowThreshold,
	}

	return l
}

// LogMode 设置日志级别
func (l *Logger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info 打印信息日志
func (l *Logger) Info(ctx context.Context, format string, args ...interface{}) {
	if l.LogLevel >= logger.Info {
		if l.config.Colorful {
			l.Printf(Green+"[INFO] "+format+Reset, args...)
		} else {
			l.Printf("[INFO] "+format, args...)
		}
	}
}

// Warn 打印警告日志
func (l *Logger) Warn(ctx context.Context, format string, args ...interface{}) {
	if l.LogLevel >= logger.Warn {
		if l.config.Colorful {
			l.Printf(Yellow+"[WARN] "+format+Reset, args...)
		} else {
			l.Printf("[WARN] "+format, args...)
		}
	}
}

// Error 打印错误日志
func (l *Logger) Error(ctx context.Context, format string, args ...interface{}) {
	if l.LogLevel >= logger.Error {
		if l.config.Colorful {
			l.Printf(Red+"[ERROR] "+format+Reset, args...)
		} else {
			l.Printf("[ERROR] "+format, args...)
		}
	}
}

// Trace SQL 追踪
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		l.Error(ctx, "[%.3fms] [rows:%v] %s; %s", float64(elapsed.Nanoseconds())/1e6, rows, sql, err)
		return
	}

	if l.SlowQuery != 0 && elapsed > l.SlowQuery {
		l.Warn(ctx, "[%.3fms] [rows:%v] %s; %s", float64(elapsed.Nanoseconds())/1e6, rows, sql, "SLOW SQL")
		return
	}

	l.Info(ctx, "[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
}

// InitLogger 初始化日志
func InitLogger(config *Config) {
	Log = New(config)

	// 初始化全局方法
	Infof = func(format string, args ...interface{}) {
		Log.Info(context.Background(), format, args...)
	}

	Warnf = func(format string, args ...interface{}) {
		Log.Warn(context.Background(), format, args...)
	}

	Errorf = func(format string, args ...interface{}) {
		Log.Error(context.Background(), format, args...)
	}

	Debugf = func(format string, args ...interface{}) {
		if Log.LogLevel >= logger.Info {
			Log.Info(context.Background(), "[DEBUG] "+format, args...)
		}
	}
}

// GetLogger 获取日志实例
func GetLogger() *Logger {
	if Log == nil {
		InitLogger(nil)
	}
	return Log
}

// 便捷方法
func Info(format string, args ...interface{}) {
	GetLogger().Info(context.Background(), format, args...)
}

func Warn(format string, args ...interface{}) {
	GetLogger().Warn(context.Background(), format, args...)
}

func Error(format string, args ...interface{}) {
	GetLogger().Error(context.Background(), format, args...)
}

func Debug(format string, args ...interface{}) {
	if GetLogger().LogLevel >= logger.Info {
		GetLogger().Info(context.Background(), "[DEBUG] "+format, args...)
	}
}

// Init 使用默认配置初始化日志
func Init() {
	InitLogger(defaultConfig)
}

// InitWithConfig 使用自定义配置初始化日志
func InitWithConfig(config *Config) {
	InitLogger(config)
}

// Fatal 打印错误日志并退出程序
func Fatal(format string, args ...interface{}) {
	GetLogger().Error(context.Background(), format, args...)
	os.Exit(1)
}

// ErrorExit 如果有错误就打印并退出
func ErrorExit(err error, msg string) {
	if err != nil {
		Fatal("%s: %v", msg, err)
	}
}

// ErrorReturn 如果有错误就打印并返回错误
func ErrorReturn(err error, msg string) error {
	if err != nil {
		Error("%s: %v", msg, err)
		return err
	}
	return nil
}

func PrintSeparator() {
	log.Println("")
}

func DefaultLogger(msg string) {
	log.Println(msg)
}
