package updater_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/twoBoots/bender/pkg/updater"
)

func TestClient_FetchLatestRelease(t *testing.T) {
	mockRelease := updater.Release{
		TagName:     "v1.2.0",
		Name:        "Release 1.2.0",
		Body:        "Initial changelog",
		PublishedAt: time.Now(),
		Assets: []updater.Asset{
			{
				Name:               "bender-darwin-aarch64",
				Size:               12345,
				BrowserDownloadURL: "https://example.com/bender-darwin-aarch64",
			},
			{
				Name:               "bender-linux-x86_64",
				Size:               12345,
				BrowserDownloadURL: "https://example.com/bender-linux-x86_64",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos/twoBoots/bender/releases/latest" {
			http.NotFound(w, r)
			return
		}
		if r.Header.Get("Accept") != "application/vnd.github.v3+json" {
			t.Errorf("expected Accept header for GitHub API")
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(mockRelease)
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	rel, err := client.FetchLatestRelease("twoBoots/bender")
	if err != nil {
		t.Fatalf("FetchLatestRelease failed: %v", err)
	}

	if rel.TagName != "v1.2.0" {
		t.Errorf("got TagName = %q; want %q", rel.TagName, "v1.2.0")
	}

	asset, err := rel.FindAsset("bender", "darwin", "arm64")
	if err != nil {
		t.Fatalf("FindAsset failed: %v", err)
	}
	if asset.Name != "bender-darwin-aarch64" {
		t.Errorf("got asset name %q; want %q", asset.Name, "bender-darwin-aarch64")
	}
}

func TestClient_FetchReleaseByTag(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/repos/twoBoots/bender/releases/tags/v1.0.0" {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(updater.Release{
				TagName: "v1.0.0",
				Assets: []updater.Asset{
					{Name: "bender-linux-x86_64", BrowserDownloadURL: "http://example.com/bin"},
				},
			})
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	rel, err := client.FetchReleaseByTag("twoBoots/bender", "v1.0.0")
	if err != nil {
		t.Fatalf("FetchReleaseByTag failed: %v", err)
	}
	if rel.TagName != "v1.0.0" {
		t.Errorf("got TagName %q; want %q", rel.TagName, "v1.0.0")
	}

	// Test 404
	_, err = client.FetchReleaseByTag("twoBoots/bender", "v9.9.9")
	if err == nil {
		t.Errorf("expected error for non-existent release")
	}
}

func TestRelease_FindAsset_NotFound(t *testing.T) {
	rel := updater.Release{
		Assets: []updater.Asset{
			{Name: "bender-linux-x86_64"},
		},
	}

	_, err := rel.FindAsset("bender", "darwin", "arm64")
	if err == nil {
		t.Errorf("expected error when asset not found")
	}

	_, err = rel.FindAsset("bender", "unsupported_os", "amd64")
	if err == nil {
		t.Errorf("expected error for unsupported os")
	}
}
