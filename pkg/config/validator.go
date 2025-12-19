package config

import (
	"fmt"
	"strings"
	"template/pkg/logger"
)

type ConfigValidator struct {
	warnings []string
	errors   []string
}

func NewConfigValidator() *ConfigValidator {
	return &ConfigValidator{
		warnings: make([]string, 0),
		errors:   make([]string, 0),
	}
}

func (v *ConfigValidator) Validate(cfg *Config) error {
	v.validateApp(&cfg.App)
	v.validateJWT(&cfg.JWT)
	v.validateDatabase(&cfg.Database)
	v.validateRedis(&cfg.Redis)
	v.validateLog(&cfg.Log)
	v.validateMail(&cfg.Mail)

	if len(v.warnings) > 0 {
		for _, warning := range v.warnings {
			logger.Warn("Config: %s", warning)
		}
	}

	if len(v.errors) > 0 {
		errMsg := "Configuration validation failed:\n"
		for _, err := range v.errors {
			errMsg += fmt.Sprintf("  - %s\n", err)
		}
		return fmt.Errorf(errMsg)
	}

	return nil
}

func (v *ConfigValidator) validateApp(app *AppConfig) {
	if app.Port < 1 || app.Port > 65535 {
		v.errors = append(v.errors, "port must be between 1-65535")
	}

	if app.Mode != "debug" && app.Mode != "release" && app.Mode != "test" {
		v.warnings = append(v.warnings, fmt.Sprintf("mode '%s' is not standard (debug/release/test)", app.Mode))
	}

	if app.Mode == "debug" {
		v.warnings = append(v.warnings, "running in debug mode, use 'release' for production")
	}
}

func (v *ConfigValidator) validateJWT(jwt *JWTConfig) {
	if len(jwt.SecretKey) < 32 {
		v.errors = append(v.errors, "JWT secret key must be at least 32 characters")
	}

	if strings.Contains(jwt.SecretKey, "change") || strings.Contains(jwt.SecretKey, "example") {
		v.errors = append(v.errors, "please change JWT secret key from default value")
	}

	if jwt.ExpiresIn < 1 {
		v.errors = append(v.errors, "JWT expires_in must be greater than 0")
	}

	if jwt.ExpiresIn > 24*7 {
		v.warnings = append(v.warnings, fmt.Sprintf("JWT expires_in %d hours is too long", jwt.ExpiresIn))
	}
}

func (v *ConfigValidator) validateDatabase(db *DatabaseConfig) {
	if db.Driver == "mysql" {
		if db.Host == "" {
			v.warnings = append(v.warnings, "MySQL host is empty, will use SQLite")
		}

		if db.Username == "" {
			v.warnings = append(v.warnings, "MySQL username is empty, will use SQLite")
		}

		if db.Username == "root" && db.Host != "localhost" && db.Host != "127.0.0.1" {
			v.warnings = append(v.warnings, "using root user for remote MySQL is not recommended")
		}

		if db.Port != 3306 && db.Port != 0 {
			v.warnings = append(v.warnings, fmt.Sprintf("non-standard MySQL port: %d", db.Port))
		}
	}
}

func (v *ConfigValidator) validateRedis(redis *RedisConfig) {
	if redis.Host == "" {
		v.warnings = append(v.warnings, "Redis host is empty, will use memory cache")
		return
	}

	if redis.Port != 6379 && redis.Port != 0 {
		v.warnings = append(v.warnings, fmt.Sprintf("non-standard Redis port: %d", redis.Port))
	}

	if redis.DB < 0 || redis.DB > 15 {
		v.errors = append(v.errors, "Redis DB must be between 0-15")
	}
}

func (v *ConfigValidator) validateLog(log *LogConfig) {
	validLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}
	if !validLevels[log.Level] {
		v.errors = append(v.errors, fmt.Sprintf("invalid log level '%s'", log.Level))
	}

	if log.MaxSize < 1 {
		v.errors = append(v.errors, "log max_size must be greater than 0")
	}

	if log.MaxSize > 1000 {
		v.warnings = append(v.warnings, fmt.Sprintf("log max_size %dMB is too large", log.MaxSize))
	}
}

func (v *ConfigValidator) validateMail(mail *MailConfig) {
	if !mail.Enabled {
		return
	}

	if mail.Host == "" {
		v.errors = append(v.errors, "mail is enabled but SMTP host is not configured")
	}

	if mail.Port != 25 && mail.Port != 465 && mail.Port != 587 {
		v.warnings = append(v.warnings, fmt.Sprintf("non-standard SMTP port: %d", mail.Port))
	}

	if mail.Username == "" || mail.Password == "" {
		v.errors = append(v.errors, "mail is enabled but username/password is missing")
	}

	if mail.From == "" {
		v.errors = append(v.errors, "mail is enabled but from address is missing")
	}

	if strings.Contains(mail.Host, "example.com") {
		v.errors = append(v.errors, "please configure a real SMTP server")
	}
}

func ValidateConfig() error {
	validator := NewConfigValidator()
	return validator.Validate(GetConfig())
}
