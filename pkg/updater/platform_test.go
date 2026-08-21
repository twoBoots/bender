package updater_test

import (
	"runtime"
	"testing"

	"github.com/twoBoots/bender/pkg/updater"
)

func TestGetBinaryNameForPlatform(t *testing.T) {
	tests := []struct {
		binaryName string
		goos       string
		goarch     string
		expected   string
		expectErr  bool
	}{
		{"bender", "darwin", "arm64", "bender-darwin-aarch64", false},
		{"bender", "darwin", "aarch64", "bender-darwin-aarch64", false},
		{"bender", "darwin", "amd64", "bender-darwin-x86_64", false},
		{"bender", "darwin", "x86_64", "bender-darwin-x86_64", false},
		{"bender", "linux", "amd64", "bender-linux-x86_64", false},
		{"bender", "linux", "arm64", "bender-linux-aarch64", false},
		{"bender", "windows", "amd64", "bender-windows-x86_64.exe", false},
		{"bender", "windows", "x86_64", "bender-windows-x86_64.exe", false},
		{"cooper", "linux", "amd64", "cooper-linux-x86_64", false},
		{"bender", "windows", "arm64", "", true},
		{"bender", "plan9", "amd64", "", true},
		{"bender", "linux", "mips", "", true},
		{"", "darwin", "arm64", "bender-darwin-aarch64", false}, // defaults binary name to bender if empty
	}

	for _, tt := range tests {
		t.Run(tt.binaryName+"_"+tt.goos+"_"+tt.goarch, func(t *testing.T) {
			got, err := updater.GetBinaryNameForPlatform(tt.binaryName, tt.goos, tt.goarch)
			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error for (%s, %s, %s), got nil", tt.binaryName, tt.goos, tt.goarch)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for (%s, %s, %s): %v", tt.binaryName, tt.goos, tt.goarch, err)
				}
				if got != tt.expected {
					t.Errorf("GetBinaryNameForPlatform(%q, %q, %q) = %q; want %q", tt.binaryName, tt.goos, tt.goarch, got, tt.expected)
				}
			}
		})
	}
}

func TestGetCurrentPlatformBinaryName(t *testing.T) {
	name, err := updater.GetCurrentPlatformBinaryName("bender")
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" || (runtime.GOOS == "windows" && runtime.GOARCH == "amd64") {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if name == "" {
			t.Errorf("expected non-empty binary name")
		}
	}
}
