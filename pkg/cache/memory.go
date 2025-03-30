package cache

import (
	"errors"
	"sync"
	"time"
)

// MemCache 内存缓存实现
type MemCache struct {
	data  map[string]memCacheItem
	mutex sync.RWMutex
}

// memCacheItem 内存缓存项
type memCacheItem struct {
	value      string
	expiration time.Time
}

// InitMemCache 初始化内存缓存
func InitMemCache() *MemCache {
	cache := &MemCache{
		data: make(map[string]memCacheItem),
	}

	// 启动一个协程定期清理过期项
	go cache.cleanupLoop()

	return cache
}

// cleanupLoop 定期清理过期项
func (c *MemCache) cleanupLoop() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		now := time.Now()
		for key, item := range c.data {
			if !item.expiration.IsZero() && item.expiration.Before(now) {
				delete(c.data, key)
			}
		}
		c.mutex.Unlock()
	}
}

// Set 设置缓存
func (c *MemCache) Set(key string, value string, expiration time.Duration) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var expirationTime time.Time
	if expiration > 0 {
		expirationTime = time.Now().Add(expiration)
	}

	c.data[key] = memCacheItem{
		value:      value,
		expiration: expirationTime,
	}

	return nil
}

// Get 获取缓存
func (c *MemCache) Get(key string) (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, exists := c.data[key]
	if !exists {
		return "", errors.New("缓存键不存在")
	}

	// 检查是否过期
	if !item.expiration.IsZero() && item.expiration.Before(time.Now()) {
		delete(c.data, key)
		return "", errors.New("缓存键已过期")
	}

	return item.value, nil
}

// Del 删除缓存
func (c *MemCache) Del(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
	return nil
}

// Exists 检查键是否存在
func (c *MemCache) Exists(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, exists := c.data[key]
	if !exists {
		return false
	}

	// 检查是否过期
	if !item.expiration.IsZero() && item.expiration.Before(time.Now()) {
		delete(c.data, key)
		return false
	}

	return true
}

// TTL 获取过期时间
func (c *MemCache) TTL(key string) (time.Duration, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, exists := c.data[key]
	if !exists {
		return 0, errors.New("缓存键不存在")
	}

	// 检查是否过期
	if !item.expiration.IsZero() && item.expiration.Before(time.Now()) {
		delete(c.data, key)
		return 0, errors.New("缓存键已过期")
	}

	// 如果没有设置过期时间
	if item.expiration.IsZero() {
		return -1, nil
	}

	// 计算剩余时间
	return time.Until(item.expiration), nil
}

// Expire 设置过期时间
func (c *MemCache) Expire(key string, expiration time.Duration) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, exists := c.data[key]
	if !exists {
		return errors.New("缓存键不存在")
	}

	// 设置新的过期时间
	if expiration > 0 {
		item.expiration = time.Now().Add(expiration)
	} else {
		item.expiration = time.Time{}
	}

	c.data[key] = item
	return nil
}

// Close 关闭内存缓存
func (c *MemCache) Close() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data = make(map[string]memCacheItem)
	return nil
} 