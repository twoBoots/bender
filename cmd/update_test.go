package cmd_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/twoBoots/bender/cmd"
	"github.com/twoBoots/bender/pkg/updater"
)

func TestUpdateCmd_Check(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(updater.Release{
			TagName: "v2.0.0",
		})
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	cmd.SetUpdaterClient(client)
	defer cmd.SetUpdaterClient(nil)

	var buf bytes.Buffer
	cmd.RootCmd.SetOut(&buf)
	cmd.RootCmd.SetArgs([]string{"update", "--check", "--repo", "twoBoots/bender"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute update failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Checking for updates") || !strings.Contains(out, "v2.0.0") {
		t.Errorf("unexpected update output: %q", out)
	}
}

func TestUpdateCmd_Apply(t *testing.T) {
	tmpDir := t.TempDir()
	fakeBin := filepath.Join(tmpDir, "bender")
	_ = os.WriteFile(fakeBin, []byte("old-bin"), 0755)

	binName, err := updater.GetCurrentPlatformBinaryName("bender")
	if err != nil {
		t.Skip("skipping on unsupported platform")
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/download" {
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = w.Write([]byte("new-bin"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(updater.Release{
			TagName: "v2.0.0",
			Assets: []updater.Asset{
				{Name: binName, BrowserDownloadURL: "http://" + r.Host + "/download"},
			},
		})
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	cmd.SetUpdaterClient(client)
	defer cmd.SetUpdaterClient(nil)

	var buf bytes.Buffer
	cmd.RootCmd.SetOut(&buf)
	cmd.RootCmd.SetArgs([]string{"update", "--force", "--exec-path", fakeBin, "--repo", "twoBoots/bender"})

	err = cmd.Execute()
	if err != nil {
		t.Fatalf("Execute update failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Successfully updated") {
		t.Errorf("unexpected update output: %q", out)
	}
}
