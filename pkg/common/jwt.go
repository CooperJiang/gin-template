package common

import (
	"template/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims 自定义JWT声明结构
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Role     int    `json:"role"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID uint, username string, role int) (string, error) {
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