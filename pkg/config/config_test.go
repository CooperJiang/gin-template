package config

import (
	"strings"
	"testing"
)

func TestSetDefaultsAndDerivedDefaults(t *testing.T) {
	cfg := Config{}

	setDefaults(&cfg)
	applyDerivedDefaults(&cfg)

	if cfg.App.Name != defaultAppName {
		t.Fatalf("App.Name = %q, want %q", cfg.App.Name, defaultAppName)
	}
	if cfg.App.Port != 9000 {
		t.Fatalf("App.Port = %d, want %d", cfg.App.Port, 9000)
	}
	if cfg.Database.Driver != defaultDatabaseDriver {
		t.Fatalf("Database.Driver = %q, want %q", cfg.Database.Driver, defaultDatabaseDriver)
	}
	if cfg.JWT.SecretKey != defaultJWTPlaceholder {
		t.Fatalf("JWT.SecretKey = %q, want %q", cfg.JWT.SecretKey, defaultJWTPlaceholder)
	}
	if cfg.JWT.ExpiresIn != defaultJWTAccessHours {
		t.Fatalf("JWT.ExpiresIn = %d, want %d", cfg.JWT.ExpiresIn, defaultJWTAccessHours)
	}
	if cfg.JWT.RefreshExpiresIn != defaultJWTRefreshHours {
		t.Fatalf("JWT.RefreshExpiresIn = %d, want %d", cfg.JWT.RefreshExpiresIn, defaultJWTRefreshHours)
	}
	if cfg.JWT.Issuer != defaultAppName {
		t.Fatalf("JWT.Issuer = %q, want %q", cfg.JWT.Issuer, defaultAppName)
	}
	if cfg.JWT.Audience != defaultAppName+"-users" {
		t.Fatalf("JWT.Audience = %q, want %q", cfg.JWT.Audience, defaultAppName+"-users")
	}
	if len(cfg.CORS.AllowedOrigins) == 0 {
		t.Fatal("CORS.AllowedOrigins should not be empty")
	}
}

func TestDecodeYAMLStrictRejectsUnknownField(t *testing.T) {
	data := []byte(`
app:
  name: "demo"
  unknown_field: "x"
`)

	cfg := Config{}
	err := decodeYAMLStrict(data, &cfg)
	if err == nil {
		t.Fatal("decodeYAMLStrict() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "unknown_field") {
		t.Fatalf("decodeYAMLStrict() error = %v, want unknown_field hint", err)
	}
}

func TestLoadEnvToStructWithLookupSupportsNestedAndSlice(t *testing.T) {
	cfg := Config{}
	env := map[string]string{
		"APP_APP_PORT":                    "8081",
		"APP_JWT_BLACKLIST_ENABLED":       "false",
		"APP_CORS_ALLOWED_ORIGINS":        `["https://a.example.com","https://b.example.com"]`,
		"APP_CORS_ALLOWED_METHODS":        "GET, POST, OPTIONS",
		"APP_FRONTEND_FALLBACK_ENABLED":   "true",
		"APP_FRONTEND_FALLBACK_MESSAGE":   "maintenance",
		"APP_FRONTEND_WEB_ENABLED":        "false",
		"APP_CORS_ALLOW_CREDENTIALS":      "true",
		"APP_CORS_ALLOWED_HEADERS":        "Authorization, Content-Type",
		"APP_CORS_ENABLED":                "true",
		"APP_FRONTEND_FALLBACK_NOT_EXIST": "ignored",
	}

	lookup := func(key string) (string, bool) {
		v, ok := env[key]
		return v, ok
	}

	loadEnvToStructWithLookup("APP_APP_", &cfg.App, lookup)
	loadEnvToStructWithLookup("APP_JWT_", &cfg.JWT, lookup)
	loadEnvToStructWithLookup("APP_CORS_", &cfg.CORS, lookup)
	loadEnvToStructWithLookup("APP_FRONTEND_", &cfg.Frontend, lookup)

	if cfg.App.Port != 8081 {
		t.Fatalf("App.Port = %d, want %d", cfg.App.Port, 8081)
	}
	if cfg.JWT.BlacklistEnabled {
		t.Fatalf("JWT.BlacklistEnabled = true, want false")
	}
	if got := strings.Join(cfg.CORS.AllowedOrigins, ","); got != "https://a.example.com,https://b.example.com" {
		t.Fatalf("CORS.AllowedOrigins = %v, want 2 parsed origins", cfg.CORS.AllowedOrigins)
	}
	if got := strings.Join(cfg.CORS.AllowedMethods, ","); got != "GET,POST,OPTIONS" {
		t.Fatalf("CORS.AllowedMethods = %v, want trimmed CSV values", cfg.CORS.AllowedMethods)
	}
	if !cfg.Frontend.Fallback.Enabled {
		t.Fatalf("Frontend.Fallback.Enabled = false, want true")
	}
	if cfg.Frontend.Fallback.Message != "maintenance" {
		t.Fatalf("Frontend.Fallback.Message = %q, want %q", cfg.Frontend.Fallback.Message, "maintenance")
	}
	if cfg.Frontend.Web.Enabled {
		t.Fatalf("Frontend.Web.Enabled = true, want false")
	}
}

func TestValidateConfigReleaseRules(t *testing.T) {
	valid := newValidReleaseConfigForTest()
	if err := validateConfig(&valid); err != nil {
		t.Fatalf("validateConfig(valid) error = %v, want nil", err)
	}

	weakSecret := newValidReleaseConfigForTest()
	weakSecret.JWT.SecretKey = "12345678"
	if err := validateConfig(&weakSecret); err == nil || !strings.Contains(err.Error(), "jwt.secret_key") {
		t.Fatalf("validateConfig(weakSecret) error = %v, want jwt.secret_key validation", err)
	}

	wildcardOrigin := newValidReleaseConfigForTest()
	wildcardOrigin.CORS.AllowedOrigins = []string{"*"}
	if err := validateConfig(&wildcardOrigin); err == nil || !strings.Contains(err.Error(), "cors.allowed_origins") {
		t.Fatalf("validateConfig(wildcardOrigin) error = %v, want cors.allowed_origins validation", err)
	}

	defaultRootPass := newValidReleaseConfigForTest()
	defaultRootPass.App.DefaultRootPass = defaultRootPassword
	if err := validateConfig(&defaultRootPass); err == nil || !strings.Contains(err.Error(), "app.defaultRootPass") {
		t.Fatalf("validateConfig(defaultRootPass) error = %v, want app.defaultRootPass validation", err)
	}

	noRedisInRelease := newValidReleaseConfigForTest()
	noRedisInRelease.Redis.Host = ""
	if err := validateConfig(&noRedisInRelease); err == nil || !strings.Contains(err.Error(), "redis.host") {
		t.Fatalf("validateConfig(noRedisInRelease) error = %v, want redis.host validation", err)
	}

	releaseBlacklistDisabled := newValidReleaseConfigForTest()
	releaseBlacklistDisabled.JWT.BlacklistEnabled = false
	releaseBlacklistDisabled.Redis.Host = ""
	if err := validateConfig(&releaseBlacklistDisabled); err != nil {
		t.Fatalf("validateConfig(releaseBlacklistDisabled) error = %v, want nil", err)
	}
}

func newValidReleaseConfigForTest() Config {
	return Config{
		App: AppConfig{
			Name:            "template",
			Port:            9000,
			Mode:            "release",
			DefaultRootPass: "root-pass-2026-safe",
			Timezone:        defaultTimezone,
		},
		Database: DatabaseConfig{
			Driver:   "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Name:     "template_db",
			Charset:  "utf8mb4",
		},
		JWT: JWTConfig{
			SecretKey:        "Akm9Xr1Q2Lm7Nw8P5Td3",
			ExpiresIn:        24,
			RefreshExpiresIn: 168,
			Issuer:           "template",
			Audience:         "template-users",
			BlacklistEnabled: true,
		},
		Redis: RedisConfig{
			Host: "127.0.0.1",
			Port: 6379,
			DB:   0,
		},
		CORS: CORSConfig{
			Enabled:          true,
			AllowedOrigins:   []string{"https://example.com"},
			AllowedMethods:   []string{"GET", "POST"},
			AllowedHeaders:   []string{"Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           3600,
		},
	}
}
