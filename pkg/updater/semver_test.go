package updater_test

import (
	"testing"

	"github.com/twoBoots/bender/pkg/updater"
)

func TestNormalizeVersion(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"v1.2.3", "1.2.3"},
		{"V1.2.3", "1.2.3"},
		{"  v2.0.0  ", "2.0.0"},
		{"1.0.0", "1.0.0"},
		{"dev", "dev"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := updater.NormalizeVersion(tt.input)
			if got != tt.expected {
				t.Errorf("NormalizeVersion(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		v1       string
		v2       string
		expected int
	}{
		{"1.0.0", "1.0.0", 0},
		{"v1.0.0", "1.0.0", 0},
		{"1.0.0", "1.0.1", -1},
		{"1.1.0", "1.0.1", 1},
		{"2.0.0", "1.9.9", 1},
		{"1.0.0-rc1", "1.0.0", -1},
		{"1.0.0", "1.0.0-rc1", 1},
		{"1.0.0-rc1", "1.0.0-rc2", -1},
		{"1.0.0-rc2", "1.0.0-rc1", 1},
		{"1.0.0-rc1", "1.0.0-rc1", 0},
		{"dev", "1.0.0", -1},
		{"1.0.0", "dev", 1},
		{"dev", "dev", 0},
		{"0.1", "0.1.0", 0},
		{"0.1.1", "0.1", 1},
		{"0.1", "0.1.1", -1},
		{"1.0.0-beta.2", "1.0.0-beta.11", -1},
		{"1.0.0-beta.11", "1.0.0-beta.2", 1},
		{"1.0.0+20130313144700", "1.0.0+20120313144700", 0},
		{"1.0.0-alpha", "1.0.0-alpha.1", -1},
		{"1.0.0-alpha.1", "1.0.0-alpha.beta", -1},
		{"1.0.0-alpha.beta", "1.0.0-beta", -1},
		{"1.0.0-beta", "1.0.0-beta.2", -1},
		{"1.0.0-beta.11", "1.0.0-rc.1", -1},
		{"1.0.0-rc.1", "1.0.0", -1},
		{"V1.0.0", "v1.0.0", 0},
		{"1", "2", -1},
		{"2", "1", 1},
		{"1-rc1", "1.0.0", -1},
		{"0.1-rc1", "0.1.0", -1},
		{"invalid-version", "1.0.0", -1},
		{"1.0.0", "invalid-version", 1},
		{"invalid-a", "invalid-b", -1},
		{"invalid-b", "invalid-a", 1},
		{"", "1.0.0", -1},
		{"1.0.0", "", 1},
		{"", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.v1+"_vs_"+tt.v2, func(t *testing.T) {
			got := updater.CompareVersions(tt.v1, tt.v2)
			if got != tt.expected {
				t.Errorf("CompareVersions(%q, %q) = %d; want %d", tt.v1, tt.v2, got, tt.expected)
			}
		})
	}
}
