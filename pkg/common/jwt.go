package common

import (
	"template/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims 自定义JWT声明结构
type JWTClaims struct {
	UserID   string `json:"user_id"` // 改为string类型以支持UUID
	Role     int    `json:"role"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌 - 支持UUID
func GenerateToken(userID UUID, username string, role int) (string, error) {
	// 获取JWT配置
	jwtConfig := config.GetConfig().JWT

	// 设置过期时间
	expirationTime := time.Now().Add(time.Duration(jwtConfig.ExpiresIn) * time.Hour)

	claims := JWTClaims{
		UserID:   userID.String(), // 将UUID转换为字符串
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.SecretKey))
}

// GenerateTokenWithStringID 生成JWT令牌 - 支持字符串ID（兼容性函数）
func GenerateTokenWithStringID(userID string, username string, role int) (string, error) {
	// 获取JWT配置
	jwtConfig := config.GetConfig().JWT

	// 设置过期时间
	expirationTime := time.Now().Add(time.Duration(jwtConfig.ExpiresIn) * time.Hour)

	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.SecretKey))
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*JWTClaims, error) {
	// 获取JWT配置
	jwtConfig := config.GetConfig().JWT

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
