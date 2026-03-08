package cache

import "testing"

func TestShouldRequireRedis(t *testing.T) {
	testCases := []struct {
		name             string
		mode             string
		blacklistEnabled bool
		expected         bool
	}{
		{name: "debug blacklist on", mode: "debug", blacklistEnabled: true, expected: false},
		{name: "release blacklist on", mode: "release", blacklistEnabled: true, expected: true},
		{name: "prod blacklist on", mode: "prod", blacklistEnabled: true, expected: true},
		{name: "production blacklist on", mode: "production", blacklistEnabled: true, expected: true},
		{name: "release blacklist off", mode: "release", blacklistEnabled: false, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := shouldRequireRedis(tc.mode, tc.blacklistEnabled)
			if got != tc.expected {
				t.Fatalf("shouldRequireRedis(%q, %v) = %v, want %v", tc.mode, tc.blacklistEnabled, got, tc.expected)
			}
		})
	}
}
