package common

import (
	"email-manage/pkg/config"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateAndParseTokenPair(t *testing.T) {
	userID := NewUUID().String()

	pair, err := GenerateTokenPair(userID, "jwt_test_user", UserRoleUser, 3)
	if err != nil {
		t.Fatalf("GenerateTokenPair failed: %v", err)
	}

	accessClaims, err := ParseTokenByType(pair.AccessToken, TokenTypeAccess)
	if err != nil {
		t.Fatalf("ParseTokenByType access failed: %v", err)
	}
	if accessClaims.UserID != userID {
		t.Fatalf("access user_id = %s, want %s", accessClaims.UserID, userID)
	}
	if accessClaims.TokenVersion != 3 {
		t.Fatalf("access token_version = %d, want 3", accessClaims.TokenVersion)
	}

	refreshClaims, err := ParseTokenByType(pair.RefreshToken, TokenTypeRefresh)
	if err != nil {
		t.Fatalf("ParseTokenByType refresh failed: %v", err)
	}
	if refreshClaims.TokenType != TokenTypeRefresh {
		t.Fatalf("refresh token_type = %s, want %s", refreshClaims.TokenType, TokenTypeRefresh)
	}
}

func TestParseTokenTypeMismatch(t *testing.T) {
	pair, err := GenerateTokenPair(NewUUID().String(), "jwt_test_user", UserRoleUser, 1)
	if err != nil {
		t.Fatalf("GenerateTokenPair failed: %v", err)
	}

	if _, err := ParseTokenByType(pair.AccessToken, TokenTypeRefresh); err == nil {
		t.Fatalf("expected token type mismatch error")
	}
}

func TestParseTokenClaimsValidation(t *testing.T) {
	cfg := config.GetConfig().JWT
	issuer, audience := resolveIssuerAudience(cfg)

	badClaims := JWTClaims{
		UserID:       NewUUID().String(),
		Username:     "jwt_claims_test",
		Role:         UserRoleUser,
		TokenType:    TokenTypeAccess,
		TokenVersion: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        "bad-claims-token-id",
			Subject:   "jwt_claims_test",
			Issuer:    issuer + "-invalid",
			Audience:  jwt.ClaimStrings{audience + "-invalid"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, badClaims).SignedString([]byte(cfg.SecretKey))
	if err != nil {
		t.Fatalf("sign token failed: %v", err)
	}

	if _, err := ParseTokenByType(tokenString, TokenTypeAccess); err == nil {
		t.Fatalf("expected claims validation error")
	}
}

func TestTokenBlacklist(t *testing.T) {
	cfg := config.GetConfig().JWT
	if !cfg.BlacklistEnabled {
		t.Skip("blacklist disabled, skip blacklist test")
	}

	pair, err := GenerateTokenPair(NewUUID().String(), "jwt_blacklist_test", UserRoleUser, 1)
	if err != nil {
		t.Fatalf("GenerateTokenPair failed: %v", err)
	}

	claims, err := ParseTokenByType(pair.AccessToken, TokenTypeAccess)
	if err != nil {
		t.Fatalf("ParseTokenByType failed: %v", err)
	}

	if err := BlacklistToken(claims.ID, time.Now().Add(time.Hour)); err != nil {
		t.Fatalf("BlacklistToken failed: %v", err)
	}

	if !IsTokenBlacklisted(claims.ID) {
		t.Fatalf("token should be blacklisted")
	}

	if _, err := ParseTokenByType(pair.AccessToken, TokenTypeAccess); err == nil {
		t.Fatalf("expected blacklisted token parse error")
	}
}
