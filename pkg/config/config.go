package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	// 环境变量前缀
	envPrefix = "APP_"

	defaultConfigFilePath  = "config.yaml"
	defaultAppName         = "email-manage"
	defaultAppMode         = "debug"
	defaultTimezone        = "Asia/Shanghai"
	defaultRootPassword    = "123456"
	defaultJWTPlaceholder  = "change-to-your-secure-secret-key"
	defaultJWTAccessHours  = 24
	defaultJWTRefreshHours = 168
	defaultDatabaseDriver  = "sqlite"
	defaultDatabaseCharset = "utf8mb4"
	defaultDatabasePort    = 3306
	defaultRedisPort       = 6379
	defaultCORSMaxAge      = 86400
)

// Config 应用配置结构
type Config struct {
	App         AppConfig         `yaml:"app" env:"APP"`
	Database    DatabaseConfig    `yaml:"database" env:"DB"`
	Redis       RedisConfig       `yaml:"redis" env:"REDIS"`
	JWT         JWTConfig         `yaml:"jwt" env:"JWT"`
	Mail        MailConfig        `yaml:"mail" env:"MAIL"`
	CORS        CORSConfig        `yaml:"cors" env:"CORS"`
	Security    SecurityConfig    `yaml:"security" env:"SECURITY"`
	MailProvider MailProviderConfig `yaml:"mail_provider" env:"MAIL_PROVIDER"`
}

// AppConfig 应用基础配置
type AppConfig struct {
	Name            string `yaml:"name" env:"NAME"`
	Port            int    `yaml:"port" env:"PORT"`
	Mode            string `yaml:"mode" env:"MODE"`
	DefaultRootPass string `yaml:"defaultRootPass" env:"DEFAULT_ROOT_PASS"`
	Timezone        string `yaml:"timezone" env:"TIMEZONE"`
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
	Enabled  bool   `yaml:"enabled" env:"ENABLED"`
	Host     string `yaml:"host" env:"HOST"`
	Port     int    `yaml:"port" env:"PORT"`
	Password string `yaml:"password" env:"PASSWORD"`
	DB       int    `yaml:"db" env:"DB"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey        string `yaml:"secret_key" env:"SECRET_KEY"`
	ExpiresIn        int    `yaml:"expires_in" env:"EXPIRES_IN"`
	RefreshExpiresIn int    `yaml:"refresh_expires_in" env:"REFRESH_EXPIRES_IN"`
	Issuer           string `yaml:"issuer" env:"ISSUER"`
	Audience         string `yaml:"audience" env:"AUDIENCE"`
	BlacklistEnabled bool   `yaml:"blacklist_enabled" env:"BLACKLIST_ENABLED"`
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

// SecurityConfig 安全相关配置
type SecurityConfig struct {
	EncryptionKey string `yaml:"encryption_key" env:"ENCRYPTION_KEY"`
}

// MailProviderConfig 邮件提供方配置
type MailProviderConfig struct {
	BaseURL string `yaml:"base_url" env:"BASE_URL"`
}

var (
	config Config
	once   sync.Once
)

// InitConfig 初始化配置
func InitConfig() {
	once.Do(func() {
		setDefaults(&config)

		if err := loadConfigFromFile(&config); err != nil {
			log.Fatalf("配置加载失败: %v", err)
		}

		loadConfigFromEnv(&config)
		applyDerivedDefaults(&config)

		if err := validateConfig(&config); err != nil {
			log.Fatalf("配置校验失败: %v", err)
		}

		log.Printf("配置加载完成，应用名称: %s, 端口: %d", config.App.Name, config.App.Port)
	})
}

// loadConfigFromFile 从配置文件加载配置（严格模式）
func loadConfigFromFile(cfg *Config) error {
	data, err := os.ReadFile(defaultConfigFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("警告: 未找到配置文件 %s，将使用默认配置和环境变量", defaultConfigFilePath)
			return nil
		}
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	if err := decodeYAMLStrict(data, cfg); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	return nil
}

func decodeYAMLStrict(data []byte, cfg *Config) error {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)
	return decoder.Decode(cfg)
}

// loadConfigFromEnv 从环境变量加载配置
func loadConfigFromEnv(cfg *Config) {
	loadEnvToStruct(envPrefix+"APP_", &cfg.App)
	loadEnvToStruct(envPrefix+"DB_", &cfg.Database)
	loadEnvToStruct(envPrefix+"REDIS_", &cfg.Redis)
	loadEnvToStruct(envPrefix+"JWT_", &cfg.JWT)
	loadEnvToStruct(envPrefix+"MAIL_", &cfg.Mail)
	loadEnvToStruct(envPrefix+"CORS_", &cfg.CORS)
	loadEnvToStruct(envPrefix+"SECURITY_", &cfg.Security)
	loadEnvToStruct(envPrefix+"MAIL_PROVIDER_", &cfg.MailProvider)
}

// loadEnvToStruct 递归加载环境变量到结构体（支持嵌套结构和字符串切片）
func loadEnvToStruct(prefix string, obj interface{}) {
	loadEnvToStructWithLookup(prefix, obj, os.LookupEnv)
}

func loadEnvToStructWithLookup(prefix string, obj interface{}, lookup func(string) (string, bool)) {
	if lookup == nil {
		lookup = os.LookupEnv
	}

	value := reflect.ValueOf(obj)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return
	}

	elem := value.Elem()
	if elem.Kind() != reflect.Struct {
		return
	}

	loadEnvToValue(prefix, elem, lookup)
}

func loadEnvToValue(prefix string, value reflect.Value, lookup func(string) (string, bool)) {
	valueType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := valueType.Field(i)
		envTag := fieldType.Tag.Get("env")
		if envTag == "" || envTag == "-" {
			continue
		}

		envKey := prefix + envTag
		if field.Kind() == reflect.Struct {
			loadEnvToValue(envKey+"_", field, lookup)
			continue
		}

		raw, exists := lookup(envKey)
		if !exists {
			continue
		}

		setFieldValue(field, raw)
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
		if intValue, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64); err == nil {
			field.SetInt(intValue)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if uintValue, err := strconv.ParseUint(strings.TrimSpace(value), 10, 64); err == nil {
			field.SetUint(uintValue)
		}
	case reflect.Float32, reflect.Float64:
		if floatValue, err := strconv.ParseFloat(strings.TrimSpace(value), 64); err == nil {
			field.SetFloat(floatValue)
		}
	case reflect.Bool:
		if boolValue, err := strconv.ParseBool(strings.TrimSpace(value)); err == nil {
			field.SetBool(boolValue)
		}
	case reflect.Slice:
		if field.Type().Elem().Kind() == reflect.String {
			field.Set(reflect.ValueOf(parseStringSlice(value)))
		}
	}
}

func parseStringSlice(value string) []string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return []string{}
	}

	if strings.HasPrefix(trimmed, "[") {
		var list []string
		if err := json.Unmarshal([]byte(trimmed), &list); err == nil {
			return normalizeStringSlice(list)
		}
	}

	return normalizeStringSlice(strings.Split(trimmed, ","))
}

func normalizeStringSlice(items []string) []string {
	result := make([]string, 0, len(items))
	for _, item := range items {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// GetConfig 获取配置
func GetConfig() *Config {
	if reflect.DeepEqual(config, Config{}) {
		InitConfig()
	}
	return &config
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

func setDefaults(cfg *Config) {
	if strings.TrimSpace(cfg.App.Name) == "" {
		cfg.App.Name = defaultAppName
	}
	if cfg.App.Port == 0 {
		cfg.App.Port = 9000
	}
	if strings.TrimSpace(cfg.App.Mode) == "" {
		cfg.App.Mode = defaultAppMode
	}
	if strings.TrimSpace(cfg.App.DefaultRootPass) == "" {
		cfg.App.DefaultRootPass = defaultRootPassword
	}
	if strings.TrimSpace(cfg.App.Timezone) == "" {
		cfg.App.Timezone = defaultTimezone
	}

	if strings.TrimSpace(cfg.Database.Driver) == "" {
		cfg.Database.Driver = defaultDatabaseDriver
	}
	if cfg.Database.Port == 0 {
		cfg.Database.Port = defaultDatabasePort
	}
	if strings.TrimSpace(cfg.Database.Name) == "" {
		cfg.Database.Name = "email_manage"
	}
	if strings.TrimSpace(cfg.Database.Charset) == "" {
		cfg.Database.Charset = defaultDatabaseCharset
	}

	if cfg.Redis.Port == 0 {
		cfg.Redis.Port = defaultRedisPort
	}

	if strings.TrimSpace(cfg.JWT.SecretKey) == "" {
		cfg.JWT.SecretKey = defaultJWTPlaceholder
	}
	if cfg.JWT.ExpiresIn <= 0 {
		cfg.JWT.ExpiresIn = defaultJWTAccessHours
	}
	if cfg.JWT.RefreshExpiresIn <= 0 {
		cfg.JWT.RefreshExpiresIn = defaultJWTRefreshHours
	}
	cfg.JWT.BlacklistEnabled = true

	cfg.CORS.Enabled = true
	if len(cfg.CORS.AllowedOrigins) == 0 {
		cfg.CORS.AllowedOrigins = []string{"*"}
	}
	if len(cfg.CORS.AllowedMethods) == 0 {
		cfg.CORS.AllowedMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	}
	if len(cfg.CORS.AllowedHeaders) == 0 {
		cfg.CORS.AllowedHeaders = []string{"Origin", "Content-Type", "Authorization"}
	}
	cfg.CORS.AllowCredentials = true
	if cfg.CORS.MaxAge == 0 {
		cfg.CORS.MaxAge = defaultCORSMaxAge
	}

	// 默认不启用外部邮件提供方，避免误调用第三方
	if strings.TrimSpace(cfg.MailProvider.BaseURL) == "" {
		cfg.MailProvider.BaseURL = "https://app.wyx66.com"
	}
}

func applyDerivedDefaults(cfg *Config) {
	if strings.TrimSpace(cfg.JWT.Issuer) == "" {
		cfg.JWT.Issuer = strings.TrimSpace(cfg.App.Name)
	}
	if strings.TrimSpace(cfg.JWT.Issuer) == "" {
		cfg.JWT.Issuer = defaultAppName
	}
	if strings.TrimSpace(cfg.JWT.Audience) == "" {
		cfg.JWT.Audience = cfg.JWT.Issuer + "-users"
	}
}

func validateConfig(cfg *Config) error {
	if cfg == nil {
		return errors.New("配置为空")
	}

	cfg.App.Mode = normalizeMode(cfg.App.Mode)
	if cfg.App.Mode == "" {
		return fmt.Errorf("app.mode 无效，仅支持 debug/release/test/production/prod")
	}

	if cfg.App.Port <= 0 || cfg.App.Port > 65535 {
		return fmt.Errorf("app.port 超出范围: %d", cfg.App.Port)
	}

	if _, err := time.LoadLocation(cfg.App.Timezone); err != nil {
		return fmt.Errorf("app.timezone 无效: %w", err)
	}

	if !isSupportedDatabaseDriver(cfg.Database.Driver) {
		return fmt.Errorf("database.driver 无效: %s", cfg.Database.Driver)
	}

	if cfg.Redis.Enabled {
		if strings.TrimSpace(cfg.Redis.Host) == "" {
			return errors.New("redis.enabled=true 时 redis.host 不能为空")
		}
		if cfg.Redis.Port <= 0 || cfg.Redis.Port > 65535 {
			return fmt.Errorf("redis.port 超出范围: %d", cfg.Redis.Port)
		}
		if cfg.Redis.DB < 0 {
			return fmt.Errorf("redis.db 不能小于0: %d", cfg.Redis.DB)
		}
	}

	if cfg.JWT.ExpiresIn <= 0 {
		return errors.New("jwt.expires_in 必须大于0")
	}
	if cfg.JWT.RefreshExpiresIn <= 0 {
		return errors.New("jwt.refresh_expires_in 必须大于0")
	}
	if strings.TrimSpace(cfg.JWT.SecretKey) == "" {
		return errors.New("jwt.secret_key 不能为空")
	}

	if cfg.Mail.Enabled {
		if err := validateMailConfig(cfg.Mail); err != nil {
			return err
		}
	}

	if err := validateCORSConfig(cfg.CORS, isReleaseMode(cfg.App.Mode)); err != nil {
		return err
	}

	if isReleaseMode(cfg.App.Mode) {
		if isWeakJWTSecret(cfg.JWT.SecretKey) {
			return errors.New("release模式下 jwt.secret_key 过弱，请使用高强度密钥")
		}

		if cfg.JWT.BlacklistEnabled && !cfg.Redis.Enabled {
			return errors.New("release模式下启用 jwt.blacklist_enabled 时必须启用 redis")
		}

		if strings.TrimSpace(cfg.App.DefaultRootPass) == "" || strings.TrimSpace(cfg.App.DefaultRootPass) == defaultRootPassword {
			return errors.New("release模式下 app.defaultRootPass 不能使用空值或默认弱密码")
		}

		if strings.EqualFold(strings.TrimSpace(cfg.Database.Driver), "mysql") {
			if strings.TrimSpace(cfg.Database.Host) == "" || strings.TrimSpace(cfg.Database.Username) == "" || strings.TrimSpace(cfg.Database.Name) == "" {
				return errors.New("release模式下 database.driver=mysql 时必须完整配置 host/username/name")
			}
		}
	}

	return nil
}

func validateMailConfig(cfg MailConfig) error {
	if strings.TrimSpace(cfg.Host) == "" {
		return errors.New("mail.enabled=true 时 mail.host 不能为空")
	}
	if cfg.Port <= 0 {
		return errors.New("mail.enabled=true 时 mail.port 必须大于0")
	}
	if strings.TrimSpace(cfg.Username) == "" {
		return errors.New("mail.enabled=true 时 mail.username 不能为空")
	}
	if strings.TrimSpace(cfg.Password) == "" {
		return errors.New("mail.enabled=true 时 mail.password 不能为空")
	}
	if strings.TrimSpace(cfg.From) == "" {
		return errors.New("mail.enabled=true 时 mail.from 不能为空")
	}
	return nil
}

func validateCORSConfig(cfg CORSConfig, strict bool) error {
	if !cfg.Enabled {
		return nil
	}

	if len(cfg.AllowedMethods) == 0 {
		return errors.New("cors.enabled=true 时 cors.allowed_methods 不能为空")
	}
	if len(cfg.AllowedHeaders) == 0 {
		return errors.New("cors.enabled=true 时 cors.allowed_headers 不能为空")
	}
	if cfg.MaxAge < 0 {
		return errors.New("cors.max_age 不能小于0")
	}

	if strict {
		if len(cfg.AllowedOrigins) == 0 {
			return errors.New("release模式下 cors.allowed_origins 不能为空")
		}
		for _, origin := range cfg.AllowedOrigins {
			trimmed := strings.TrimSpace(origin)
			if trimmed == "" {
				return errors.New("release模式下 cors.allowed_origins 不能包含空值")
			}
			if trimmed == "*" {
				return errors.New("release模式下 cors.allowed_origins 不能使用通配符*")
			}
		}
	}

	return nil
}

func normalizeMode(mode string) string {
	normalized := strings.ToLower(strings.TrimSpace(mode))
	switch normalized {
	case "debug", "release", "test", "production", "prod":
		return normalized
	default:
		return ""
	}
}

func isReleaseMode(mode string) bool {
	normalized := strings.ToLower(strings.TrimSpace(mode))
	return normalized == "release" || normalized == "production" || normalized == "prod"
}

func isSupportedDatabaseDriver(driver string) bool {
	normalized := strings.ToLower(strings.TrimSpace(driver))
	return normalized == "" || normalized == "sqlite" || normalized == "mysql"
}

func isWeakJWTSecret(secret string) bool {
	trimmed := strings.TrimSpace(secret)
	if len(trimmed) < 16 {
		return true
	}

	for _, keyword := range []string{"change-to-your-secure-secret-key", "123456", "password", "secret", "jwtsecret"} {
		if strings.Contains(strings.ToLower(trimmed), keyword) {
			return true
		}
	}

	return false
}
