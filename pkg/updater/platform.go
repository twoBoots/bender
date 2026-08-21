package updater

import (
	"fmt"
	"runtime"
	"strings"
)

const DefaultBinaryName = "bender"

// GetCurrentPlatformBinaryName returns the release binary name for the current runtime.
func GetCurrentPlatformBinaryName(binaryName string) (string, error) {
	return GetBinaryNameForPlatform(binaryName, runtime.GOOS, runtime.GOARCH)
}

// GetBinaryNameForPlatform maps GOOS and GOARCH to the release asset binary name convention.
func GetBinaryNameForPlatform(binaryName, goos, goarch string) (string, error) {
	if strings.TrimSpace(binaryName) == "" {
		binaryName = DefaultBinaryName
	}

	var normalizedArch string
	switch goarch {
	case "amd64", "x86_64":
		normalizedArch = "x86_64"
	case "arm64", "aarch64":
		normalizedArch = "aarch64"
	default:
		return "", fmt.Errorf("unsupported architecture: %s", goarch)
	}

	switch goos {
	case "darwin", "linux":
		return fmt.Sprintf("%s-%s-%s", binaryName, goos, normalizedArch), nil
	case "windows":
		if normalizedArch != "x86_64" {
			return "", fmt.Errorf("unsupported architecture for windows: %s", goarch)
		}
		return fmt.Sprintf("%s-windows-%s.exe", binaryName, normalizedArch), nil
	default:
		return "", fmt.Errorf("unsupported operating system: %s", goos)
	}
}
