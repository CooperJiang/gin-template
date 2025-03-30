package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

// Config 应用配置结构
type Config struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	JWT      JWTConfig      `yaml:"jwt"`
	Log      LogConfig      `yaml:"log"`
	Mail     MailConfig     `yaml:"mail"`
}

// AppConfig 应用基础配置
type AppConfig struct {
	Name            string `yaml:"name"`
	Port            int    `yaml:"port"`
	Mode            string `yaml:"mode"`
	DefaultRootPass string `yaml:"defaultRootPass"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Charset  string `yaml:"charset"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey string `yaml:"secret_key"`
	ExpiresIn int    `yaml:"expires_in"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `yaml:"level"`
	Path       string `yaml:"path"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

// MailConfig 邮件配置
type MailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
	FromName string `yaml:"from_name"`
	ReplyTo  string `yaml:"reply_to"`
	SSL      bool   `yaml:"ssl"`
	Enabled  bool   `yaml:"enabled"`
}

var (
	config Config
	once   sync.Once
)

// InitConfig 初始化配置
func InitConfig() {
	once.Do(func() {
		// 读取配置文件
		data, err := os.ReadFile("config.yaml")
		if err != nil {
			log.Fatalf("无法读取配置文件: %v", err)
		}

		// 解析配置文件
		if err := yaml.Unmarshal(data, &config); err != nil {
			log.Fatalf("无法解析配置文件: %v", err)
		}
	})
}

// GetConfig 获取配置
func GetConfig() *Config {
	return &config
}
