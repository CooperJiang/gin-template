package cache

import (
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
		logger.Error("Redis初始化失败: %v，将使用内存缓存", err)
		// Redis初始化失败，使用内存缓存
		defaultCache = InitMemCache()
		logger.Info("内存缓存初始化成功")
	}
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