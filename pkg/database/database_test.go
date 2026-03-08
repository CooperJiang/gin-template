package database

import (
	"template/pkg/config"
	"testing"
)

func TestIsReleaseMode(t *testing.T) {
	testCases := []struct {
		name     string
		mode     string
		expected bool
	}{
		{name: "release", mode: "release", expected: true},
		{name: "production", mode: "production", expected: true},
		{name: "prod", mode: "prod", expected: true},
		{name: "mixed case", mode: "ReLeAsE", expected: true},
		{name: "with space", mode: " production ", expected: true},
		{name: "debug", mode: "debug", expected: false},
		{name: "empty", mode: "", expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := isReleaseMode(tc.mode); actual != tc.expected {
				t.Fatalf("isReleaseMode(%q) = %v, want %v", tc.mode, actual, tc.expected)
			}
		})
	}
}

func TestShouldUseMySQL(t *testing.T) {
	testCases := []struct {
		name           string
		cfg            config.DatabaseConfig
		wantMySQL      bool
		wantMySQLReady bool
	}{
		{
			name: "explicit sqlite",
			cfg: config.DatabaseConfig{
				Driver:   "sqlite",
				Host:     "localhost",
				Username: "root",
				Name:     "app",
			},
			wantMySQL:      false,
			wantMySQLReady: false,
		},
		{
			name: "explicit mysql complete config",
			cfg: config.DatabaseConfig{
				Driver:   "mysql",
				Host:     "localhost",
				Username: "root",
				Name:     "app",
			},
			wantMySQL:      true,
			wantMySQLReady: true,
		},
		{
			name: "explicit mysql incomplete config",
			cfg: config.DatabaseConfig{
				Driver:   "mysql",
				Host:     "",
				Username: "root",
				Name:     "app",
			},
			wantMySQL:      true,
			wantMySQLReady: false,
		},
		{
			name: "implicit sqlite when no mysql signal",
			cfg: config.DatabaseConfig{
				Driver:   "",
				Host:     "",
				Username: "",
				Name:     "",
			},
			wantMySQL:      false,
			wantMySQLReady: false,
		},
		{
			name: "implicit mysql incomplete by signal",
			cfg: config.DatabaseConfig{
				Driver:   "",
				Host:     "localhost",
				Username: "",
				Name:     "app",
			},
			wantMySQL:      true,
			wantMySQLReady: false,
		},
		{
			name: "implicit mysql complete by signal",
			cfg: config.DatabaseConfig{
				Driver:   "",
				Host:     "localhost",
				Username: "root",
				Name:     "app",
			},
			wantMySQL:      true,
			wantMySQLReady: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotMySQL, gotMySQLReady := shouldUseMySQL(tc.cfg)
			if gotMySQL != tc.wantMySQL || gotMySQLReady != tc.wantMySQLReady {
				t.Fatalf("shouldUseMySQL(%+v) = (%v,%v), want (%v,%v)",
					tc.cfg,
					gotMySQL,
					gotMySQLReady,
					tc.wantMySQL,
					tc.wantMySQLReady,
				)
			}
		})
	}
}

func TestResolveInitOptionsByMode(t *testing.T) {
	t.Run("debug defaults", func(t *testing.T) {
		opts := resolveInitOptionsByMode("debug", nil)
		if opts.strictMode {
			t.Fatalf("strictMode = true, want false")
		}
		if !opts.autoMigrate {
			t.Fatalf("autoMigrate = false, want true")
		}
		if !opts.bootstrapRoot {
			t.Fatalf("bootstrapRoot = false, want true")
		}
	})

	t.Run("release defaults", func(t *testing.T) {
		opts := resolveInitOptionsByMode("release", nil)
		if !opts.strictMode {
			t.Fatalf("strictMode = false, want true")
		}
		if opts.autoMigrate {
			t.Fatalf("autoMigrate = true, want false")
		}
		if opts.bootstrapRoot {
			t.Fatalf("bootstrapRoot = true, want false")
		}
	})

	t.Run("override by env lookup", func(t *testing.T) {
		mockGetEnvBool := func(key string, defaultValue bool) bool {
			switch key {
			case "APP_DB_STRICT":
				return false
			case "APP_DB_AUTO_MIGRATE":
				return true
			case "APP_DB_BOOTSTRAP_ROOT":
				return true
			default:
				return defaultValue
			}
		}

		opts := resolveInitOptionsByMode("release", mockGetEnvBool)
		if opts.strictMode {
			t.Fatalf("strictMode = true, want false")
		}
		if !opts.autoMigrate {
			t.Fatalf("autoMigrate = false, want true")
		}
		if !opts.bootstrapRoot {
			t.Fatalf("bootstrapRoot = false, want true")
		}
	})
}
