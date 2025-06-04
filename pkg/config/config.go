package config

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

const (
	// 环境变量前缀
	envPrefix = "APP_"
)

// Config 应用配置结构
type Config struct {
	App      AppConfig      `yaml:"app" env:"APP"`
	Database DatabaseConfig `yaml:"database" env:"DB"`
	Redis    RedisConfig    `yaml:"redis" env:"REDIS"`
	JWT      JWTConfig      `yaml:"jwt" env:"JWT"`
	Log      LogConfig      `yaml:"log" env:"LOG"`
	Mail     MailConfig     `yaml:"mail" env:"MAIL"`
	CORS     CORSConfig     `yaml:"cors" env:"CORS"`
	Frontend FrontendConfig `yaml:"frontend" env:"FRONTEND"`
}

// AppConfig 应用基础配置
type AppConfig struct {
	Name            string `yaml:"name" env:"NAME"`
	Port            int    `yaml:"port" env:"PORT"`
	Mode            string `yaml:"mode" env:"MODE"`
	DefaultRootPass string `yaml:"defaultRootPass" env:"DEFAULT_ROOT_PASS"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string `yaml:"driver" env:"DRIVER"`
	Host     string `yaml:"host" env:"HOST"`
	Port     int    `yaml:"port" env:"PORT"`
	Username string `yaml:"username" env:"USERNAME"`
	Password string `yaml:"password" env:"PASSWORD"`
	Name     string `yaml:"name" env:"NAME"`
	Charset  string `yaml:"charset" env:"CHARSET"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `yaml:"host" env:"HOST"`
	Port     int    `yaml:"port" env:"PORT"`
	Password string `yaml:"password" env:"PASSWORD"`
	DB       int    `yaml:"db" env:"DB"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey string `yaml:"secret_key" env:"SECRET_KEY"`
	ExpiresIn int    `yaml:"expires_in" env:"EXPIRES_IN"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `yaml:"level" env:"LEVEL"`
	Path       string `yaml:"path" env:"PATH"`
	Filename   string `yaml:"filename" env:"FILENAME"`
	MaxSize    int    `yaml:"max_size" env:"MAX_SIZE"`
	MaxBackups int    `yaml:"max_backups" env:"MAX_BACKUPS"`
	MaxAge     int    `yaml:"max_age" env:"MAX_AGE"`
	Compress   bool   `yaml:"compress" env:"COMPRESS"`
}

// MailConfig 邮件配置
type MailConfig struct {
	Host     string `yaml:"host" env:"HOST"`
	Port     int    `yaml:"port" env:"PORT"`
	Username string `yaml:"username" env:"USERNAME"`
	Password string `yaml:"password" env:"PASSWORD"`
	From     string `yaml:"from" env:"FROM"`
	FromName string `yaml:"from_name" env:"FROM_NAME"`
	ReplyTo  string `yaml:"reply_to" env:"REPLY_TO"`
	SSL      bool   `yaml:"ssl" env:"SSL"`
	Enabled  bool   `yaml:"enabled" env:"ENABLED"`
}

// CORSConfig 跨域(CORS)配置
type CORSConfig struct {
	Enabled          bool     `yaml:"enabled" env:"ENABLED"`
	AllowedOrigins   []string `yaml:"allowed_origins" env:"ALLOWED_ORIGINS"`
	AllowedMethods   []string `yaml:"allowed_methods" env:"ALLOWED_METHODS"`
	AllowedHeaders   []string `yaml:"allowed_headers" env:"ALLOWED_HEADERS"`
	AllowCredentials bool     `yaml:"allow_credentials" env:"ALLOW_CREDENTIALS"`
	MaxAge           int      `yaml:"max_age" env:"MAX_AGE"`
}

// FrontendConfig 前端模块配置
type FrontendConfig struct {
	Admin    FrontendModuleConfig   `yaml:"admin" env:"ADMIN"`
	Web      FrontendModuleConfig   `yaml:"web" env:"WEB"`
	Fallback FrontendFallbackConfig `yaml:"fallback" env:"FALLBACK"`
}

// FrontendModuleConfig 前端模块具体配置
type FrontendModuleConfig struct {
	Enabled     bool   `yaml:"enabled" env:"ENABLED"`
	RoutePrefix string `yaml:"route_prefix" env:"ROUTE_PREFIX"`
	Title       string `yaml:"title" env:"TITLE"`
	Description string `yaml:"description" env:"DESCRIPTION"`
}

// FrontendFallbackConfig 前端备用页面配置
type FrontendFallbackConfig struct {
	Enabled bool   `yaml:"enabled" env:"ENABLED"`
	Message string `yaml:"message" env:"MESSAGE"`
}

var (
	config Config
	once   sync.Once
)

// InitConfig 初始化配置
func InitConfig() {
	once.Do(func() {
		// 先从配置文件读取默认配置
		loadConfigFromFile(&config)

		// 然后从环境变量中覆盖配置
		loadConfigFromEnv(&config)

		// 打印加载配置信息
		log.Printf("配置加载完成，应用名称: %s, 端口: %d", config.App.Name, config.App.Port)
	})
}

// loadConfigFromFile 从配置文件加载配置
func loadConfigFromFile(cfg *Config) {
	// 读取配置文件
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("警告: 无法读取配置文件: %v，将使用默认配置和环境变量", err)
		return
	}

	// 解析配置文件
	if err := yaml.Unmarshal(data, cfg); err != nil {
		log.Printf("警告: 无法解析配置文件: %v，将使用默认配置和环境变量", err)
	}
}

// loadConfigFromEnv 从环境变量加载配置
func loadConfigFromEnv(cfg *Config) {
	// 处理App配置的环境变量
	loadEnvToStruct(envPrefix+"APP_", &cfg.App)

	// 处理Database配置的环境变量
	loadEnvToStruct(envPrefix+"DB_", &cfg.Database)

	// 处理Redis配置的环境变量
	loadEnvToStruct(envPrefix+"REDIS_", &cfg.Redis)

	// 处理JWT配置的环境变量
	loadEnvToStruct(envPrefix+"JWT_", &cfg.JWT)

	// 处理Log配置的环境变量
	loadEnvToStruct(envPrefix+"LOG_", &cfg.Log)

	// 处理Mail配置的环境变量
	loadEnvToStruct(envPrefix+"MAIL_", &cfg.Mail)

	// 处理CORS配置的环境变量
	loadEnvToStruct(envPrefix+"CORS_", &cfg.CORS)

	// 处理Frontend配置的环境变量
	loadEnvToStruct(envPrefix+"FRONTEND_", &cfg.Frontend)
}

// loadEnvToStruct 加载环境变量到结构体
func loadEnvToStruct(prefix string, obj interface{}) {
	// 获取所有环境变量
	for _, env := range os.Environ() {
		if !strings.HasPrefix(env, prefix) {
			continue
		}

		parts := strings.SplitN(env, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimPrefix(parts[0], prefix)
		value := parts[1]

		// 将环境变量的值设置到结构体中
		setStructField(obj, key, value)
	}
}

// setStructField 设置结构体字段的值
func setStructField(obj interface{}, key string, value string) {
	// 使用反射设置字段值
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()

	// 遍历结构体的所有字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// 根据env标签找到对应的环境变量字段
		envKey := field.Tag.Get("env")
		if envKey == key {
			// 找到匹配的字段，根据字段类型设置值
			setFieldValue(v.Field(i), value)
			return
		}
	}
}

// setFieldValue 根据字段类型设置值
func setFieldValue(field reflect.Value, value string) {
	if !field.CanSet() {
		return
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			field.SetInt(intValue)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if uintValue, err := strconv.ParseUint(value, 10, 64); err == nil {
			field.SetUint(uintValue)
		}
	case reflect.Float32, reflect.Float64:
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			field.SetFloat(floatValue)
		}
	case reflect.Bool:
		if boolValue, err := strconv.ParseBool(value); err == nil {
			field.SetBool(boolValue)
		}
	}
}

// GetConfig 获取配置
func GetConfig() *Config {
	// 确保配置已初始化
	if reflect.DeepEqual(config, Config{}) {
		InitConfig()
	}
	return &config
}

// GetEnvString 从环境变量获取字符串值
func GetEnvString(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetEnvInt 从环境变量获取整数值
func GetEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// GetEnvBool 从环境变量获取布尔值
func GetEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
