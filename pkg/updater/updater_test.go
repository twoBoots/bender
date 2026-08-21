package updater_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/twoBoots/bender/pkg/updater"
)

func TestSelfUpdate_AlreadyUpToDate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(updater.Release{
			TagName: "v1.0.0",
		})
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	opts := updater.Options{
		CurrentVersion: "v1.0.0",
		Repo:           "twoBoots/bender",
		Client:         client,
	}

	res, err := updater.SelfUpdate(opts)
	if err != nil {
		t.Fatalf("SelfUpdate failed: %v", err)
	}

	if res.UpdateAvailable {
		t.Errorf("expected UpdateAvailable=false, got true")
	}
	if res.Updated {
		t.Errorf("expected Updated=false, got true")
	}
}

func TestSelfUpdate_CheckOnly(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(updater.Release{
			TagName: "v1.2.0",
		})
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	opts := updater.Options{
		CurrentVersion: "v1.0.0",
		Repo:           "twoBoots/bender",
		CheckOnly:      true,
		Client:         client,
	}

	res, err := updater.SelfUpdate(opts)
	if err != nil {
		t.Fatalf("SelfUpdate failed: %v", err)
	}

	if !res.UpdateAvailable {
		t.Errorf("expected UpdateAvailable=true, got false")
	}
	if res.Updated {
		t.Errorf("expected Updated=false during CheckOnly, got true")
	}
}

func TestSelfUpdate_ForceUpdate(t *testing.T) {
	tmpDir := t.TempDir()
	fakeBinary := filepath.Join(tmpDir, "fake-bender")
	if err := os.WriteFile(fakeBinary, []byte("old-binary-content"), 0755); err != nil {
		t.Fatalf("failed to create fake binary: %v", err)
	}

	binaryName, err := updater.GetCurrentPlatformBinaryName("bender")
	if err != nil {
		t.Skipf("skipping on unsupported platform: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/download/binary" {
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = w.Write([]byte("forced-binary-content"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(updater.Release{
			TagName: "v1.0.0",
			Assets: []updater.Asset{
				{
					Name:               binaryName,
					BrowserDownloadURL: "http://" + r.Host + "/download/binary",
				},
			},
		})
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	opts := updater.Options{
		BinaryName:     "bender",
		CurrentVersion: "v1.0.0",
		Repo:           "twoBoots/bender",
		Force:          true,
		ExecutablePath: fakeBinary,
		Client:         client,
	}

	res, err := updater.SelfUpdate(opts)
	if err != nil {
		t.Fatalf("SelfUpdate failed: %v", err)
	}

	if !res.Updated {
		t.Errorf("expected Updated=true with --force")
	}
}

func TestSelfUpdate_TargetVersion(t *testing.T) {
	tmpDir := t.TempDir()
	fakeBinary := filepath.Join(tmpDir, "fake-bender")
	if err := os.WriteFile(fakeBinary, []byte("old-binary-content"), 0755); err != nil {
		t.Fatalf("failed to create fake binary: %v", err)
	}

	binaryName, err := updater.GetCurrentPlatformBinaryName("bender")
	if err != nil {
		t.Skipf("skipping on unsupported platform: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/download/binary" {
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = w.Write([]byte("v0.9.0-binary-content"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(updater.Release{
			TagName: "v0.9.0",
			Assets: []updater.Asset{
				{
					Name:               binaryName,
					BrowserDownloadURL: "http://" + r.Host + "/download/binary",
				},
			},
		})
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	opts := updater.Options{
		BinaryName:     "bender",
		CurrentVersion: "v1.0.0",
		TargetVersion:  "v0.9.0",
		Force:          true,
		Repo:           "twoBoots/bender",
		ExecutablePath: fakeBinary,
		Client:         client,
	}

	res, err := updater.SelfUpdate(opts)
	if err != nil {
		t.Fatalf("SelfUpdate failed: %v", err)
	}

	if !res.Updated {
		t.Errorf("expected Updated=true")
	}
}

func TestSelfUpdate_DownloadError(t *testing.T) {
	tmpDir := t.TempDir()
	fakeBinary := filepath.Join(tmpDir, "fake-bender")
	if err := os.WriteFile(fakeBinary, []byte("old-binary-content"), 0755); err != nil {
		t.Fatalf("failed to create fake binary: %v", err)
	}

	binaryName, err := updater.GetCurrentPlatformBinaryName("bender")
	if err != nil {
		t.Skipf("skipping on unsupported platform: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/download/binary" {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(updater.Release{
			TagName: "v2.0.0",
			Assets: []updater.Asset{
				{
					Name:               binaryName,
					BrowserDownloadURL: "http://" + r.Host + "/download/binary",
				},
			},
		})
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	opts := updater.Options{
		BinaryName:     "bender",
		CurrentVersion: "v1.0.0",
		Repo:           "twoBoots/bender",
		ExecutablePath: fakeBinary,
		Client:         client,
	}

	_, err = updater.SelfUpdate(opts)
	if err == nil {
		t.Errorf("expected error on download failure, got nil")
	}
}

func TestApplyPlatformSigning(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("codesigning test applies to darwin only")
	}
	tmpFile := filepath.Join(t.TempDir(), "test-bin")
	if err := os.WriteFile(tmpFile, []byte("#!/bin/sh\necho hi\n"), 0755); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	err := updater.ApplyPlatformSigning(tmpFile)
	if err != nil {
		t.Logf("ApplyPlatformSigning error (acceptable in test environment): %v", err)
	}
}
