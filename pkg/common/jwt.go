package common

import (
	stdErrors "errors"
	"strings"
	"template/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	// TokenTypeAccess 访问令牌
	TokenTypeAccess = "access"
	// TokenTypeRefresh 刷新令牌
	TokenTypeRefresh = "refresh"
)

var (
	errTokenTypeMismatch = stdErrors.New("token type mismatch")
	errTokenBlacklisted  = stdErrors.New("token blacklisted")
)

// JWTClaims 自定义JWT声明结构
type JWTClaims struct {
	UserID       string `json:"user_id"`
	Role         int    `json:"role"`
	Username     string `json:"username"`
	TokenType    string `json:"token_type"`
	TokenVersion int    `json:"token_version"`
	jwt.RegisteredClaims
}

// TokenPair 令牌对
type TokenPair struct {
	AccessToken      string
	AccessExpiresAt  time.Time
	RefreshToken     string
	RefreshExpiresAt time.Time
}

// GenerateToken 生成访问令牌（兼容旧调用）
func GenerateToken(userID UUID, username string, role int) (string, error) {
	token, _, err := generateTokenWithType(userID.String(), username, role, 1, TokenTypeAccess, getAccessTTL())
	return token, err
}

// GenerateTokenWithStringID 生成访问令牌（兼容旧调用）
func GenerateTokenWithStringID(userID string, username string, role int) (string, error) {
	token, _, err := generateTokenWithType(userID, username, role, 1, TokenTypeAccess, getAccessTTL())
	return token, err
}

// GenerateTokenPair 生成 access + refresh 令牌对
func GenerateTokenPair(userID, username string, role, tokenVersion int) (*TokenPair, error) {
	accessToken, accessExpiresAt, err := generateTokenWithType(
		userID,
		username,
		role,
		tokenVersion,
		TokenTypeAccess,
		getAccessTTL(),
	)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpiresAt, err := generateTokenWithType(
		userID,
		username,
		role,
		tokenVersion,
		TokenTypeRefresh,
		getRefreshTTL(),
	)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:      accessToken,
		AccessExpiresAt:  accessExpiresAt,
		RefreshToken:     refreshToken,
		RefreshExpiresAt: refreshExpiresAt,
	}, nil
}

// ParseToken 解析 access token（兼容旧调用）
func ParseToken(tokenString string) (*JWTClaims, error) {
	return ParseTokenByType(tokenString, TokenTypeAccess)
}

// ParseTokenByType 按类型解析 token，并执行 issuer/audience/nbf 校验
func ParseTokenByType(tokenString, expectedType string) (*JWTClaims, error) {
	jwtConfig := config.GetConfig().JWT

	claims := &JWTClaims{}
	parser := jwt.NewParser(buildParserOptions(jwtConfig)...)

	token, err := parser.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	if expectedType != "" && claims.TokenType != expectedType {
		return nil, errTokenTypeMismatch
	}

	if jwtConfig.BlacklistEnabled && claims.ID != "" && IsTokenBlacklisted(claims.ID) {
		return nil, errTokenBlacklisted
	}

	return claims, nil
}

func generateTokenWithType(
	userID string,
	username string,
	role int,
	tokenVersion int,
	tokenType string,
	ttl time.Duration,
) (string, time.Time, error) {
	jwtConfig := config.GetConfig().JWT
	issuer, audience := resolveIssuerAudience(jwtConfig)
	now := time.Now()
	expiresAt := now.Add(ttl)

	claims := JWTClaims{
		UserID:       userID,
		Username:     username,
		Role:         role,
		TokenType:    tokenType,
		TokenVersion: normalizeTokenVersion(tokenVersion),
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			Subject:   userID,
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	claims.Audience = jwt.ClaimStrings{audience}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtConfig.SecretKey))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

func normalizeTokenVersion(tokenVersion int) int {
	if tokenVersion <= 0 {
		return 1
	}
	return tokenVersion
}

func getAccessTTL() time.Duration {
	return time.Duration(config.GetConfig().JWT.ExpiresIn) * time.Hour
}

func getRefreshTTL() time.Duration {
	return time.Duration(config.GetConfig().JWT.RefreshExpiresIn) * time.Hour
}

func buildParserOptions(jwtConfig config.JWTConfig) []jwt.ParserOption {
	issuer, audience := resolveIssuerAudience(jwtConfig)

	options := []jwt.ParserOption{
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithLeeway(30 * time.Second),
		jwt.WithIssuer(issuer),
		jwt.WithAudience(audience),
	}

	return options
}

func resolveIssuerAudience(jwtConfig config.JWTConfig) (string, string) {
	issuer := strings.TrimSpace(jwtConfig.Issuer)
	if issuer == "" {
		issuer = strings.TrimSpace(config.GetConfig().App.Name)
	}
	if issuer == "" {
		issuer = "template"
	}

	audience := strings.TrimSpace(jwtConfig.Audience)
	if audience == "" {
		audience = issuer + "-users"
	}

	return issuer, audience
}
