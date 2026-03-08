package common

import (
	"strings"
	"template/pkg/cache"
	"template/pkg/config"
	"time"
)

const blacklistKeyPrefix = "auth:blacklist:"

// BlacklistToken 将token加入黑名单，直到过期
func BlacklistToken(tokenID string, expiresAt time.Time) error {
	if !config.GetConfig().JWT.BlacklistEnabled {
		return nil
	}

	tokenID = strings.TrimSpace(tokenID)
	if tokenID == "" {
		return nil
	}

	ttl := time.Until(expiresAt)
	if ttl <= 0 {
		return nil
	}

	return cache.GetCache().Set(blacklistKey(tokenID), "1", ttl)
}

// IsTokenBlacklisted 检查token是否在黑名单中
func IsTokenBlacklisted(tokenID string) bool {
	if !config.GetConfig().JWT.BlacklistEnabled {
		return false
	}

	tokenID = strings.TrimSpace(tokenID)
	if tokenID == "" {
		return false
	}

	return cache.GetCache().Exists(blacklistKey(tokenID))
}

func blacklistKey(tokenID string) string {
	return blacklistKeyPrefix + tokenID
}
