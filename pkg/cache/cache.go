package cache

import (
	"strings"
	"template/pkg/config"
	"template/pkg/logger"
	"time"
)

// Cache 定义缓存接口
type Cache interface {
	Set(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	Exists(key string) bool
	TTL(key string) (time.Duration, error)
	Expire(key string, expiration time.Duration) error
	Close() error
}

var defaultCache Cache

// IsRedisEnabled 检查Redis是否启用
func IsRedisEnabled() bool {
	return redisCache != nil && defaultCache == redisCache
}

// InitCache 初始化缓存
func InitCache() {
	// 尝试初始化Redis
	if err := InitRedis(); err != nil {
		if shouldRequireRedis(config.GetConfig().App.Mode, config.GetConfig().JWT.BlacklistEnabled) {
			logger.Fatal("Redis初始化失败，release模式且启用JWT黑名单时必须可用: %v", err)
			return
		}

		defaultCache = InitMemCache()
		logger.Warn("Redis初始化失败，已降级为内存缓存: %v", err)
		return
	}

	logger.Info("Redis缓存初始化成功")
}

// GetCache 获取缓存实例
func GetCache() Cache {
	if defaultCache == nil {
		InitCache()
	}
	return defaultCache
}

// Set 设置缓存
func Set(key string, value string, expiration time.Duration) error {
	return GetCache().Set(key, value, expiration)
}

// Get 获取缓存
func Get(key string) (string, error) {
	return GetCache().Get(key)
}

// Del 删除缓存
func Del(key string) error {
	return GetCache().Del(key)
}

// Exists 检查键是否存在
func Exists(key string) bool {
	return GetCache().Exists(key)
}

// TTL 获取过期时间
func TTL(key string) (time.Duration, error) {
	return GetCache().TTL(key)
}

// Expire 设置过期时间
func Expire(key string, expiration time.Duration) error {
	return GetCache().Expire(key, expiration)
}

// Close 关闭缓存连接
func Close() error {
	if defaultCache != nil {
		return defaultCache.Close()
	}
	return nil
}

func shouldRequireRedis(mode string, blacklistEnabled bool) bool {
	if !blacklistEnabled {
		return false
	}
	return isReleaseMode(mode)
}

func isReleaseMode(mode string) bool {
	normalized := strings.ToLower(strings.TrimSpace(mode))
	return normalized == "release" || normalized == "production" || normalized == "prod"
}
