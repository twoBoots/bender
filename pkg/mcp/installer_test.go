package mcp_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/twoBoots/bender/pkg/mcp"
)

func TestMergeMCPServerConfig(t *testing.T) {
	// 1. Merge into empty / new config
	merged, err := mcp.MergeMCPServerConfig(nil, "bender", "bender", []string{"mcp"})
	if err != nil {
		t.Fatalf("MergeMCPServerConfig failed: %v", err)
	}

	var root map[string]interface{}
	if err := json.Unmarshal(merged, &root); err != nil {
		t.Fatalf("failed to parse merged JSON: %v", err)
	}

	servers, ok := root["mcpServers"].(map[string]interface{})
	if !ok {
		t.Fatalf("mcpServers not found in root")
	}

	benderServer, ok := servers["bender"].(map[string]interface{})
	if !ok {
		t.Fatalf("bender server entry not found")
	}
	if benderServer["command"] != "bender" {
		t.Errorf("got command %v; want bender", benderServer["command"])
	}

	// 2. Merge into existing config preserving other servers
	existing := []byte(`{
		"otherConfig": true,
		"mcpServers": {
			"existingServer": {
				"command": "node",
				"args": ["server.js"]
			}
		}
	}`)

	merged2, err := mcp.MergeMCPServerConfig(existing, "bender", "bender", []string{"mcp"})
	if err != nil {
		t.Fatalf("MergeMCPServerConfig failed: %v", err)
	}

	var root2 map[string]interface{}
	_ = json.Unmarshal(merged2, &root2)
	if root2["otherConfig"] != true {
		t.Errorf("lost otherConfig key in root")
	}
	servers2 := root2["mcpServers"].(map[string]interface{})
	if _, ok := servers2["existingServer"]; !ok {
		t.Errorf("lost existingServer entry")
	}
	if _, ok := servers2["bender"]; !ok {
		t.Errorf("failed to add bender server")
	}

	// 3. Error case - invalid JSON
	_, err = mcp.MergeMCPServerConfig([]byte("invalid json"), "bender", "bender", nil)
	if err == nil {
		t.Errorf("expected error for invalid json")
	}
}

func TestInstallClients(t *testing.T) {
	tmpDir := t.TempDir()
	homeDir := filepath.Join(tmpDir, "home")
	cwd := filepath.Join(tmpDir, "project")
	_ = os.MkdirAll(homeDir, 0755)
	_ = os.MkdirAll(cwd, 0755)

	opts := mcp.InstallerOptions{
		ServerName: "bender",
		Command:    "bender",
		Args:       []string{"mcp"},
		Cwd:        cwd,
		HomeDir:    homeDir,
		ClientIDs:  []string{"cursor", "antigravity"},
	}

	results, err := mcp.InstallClients(opts)
	if err != nil {
		t.Fatalf("InstallClients failed: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	for _, res := range results {
		if res.Error != nil {
			t.Errorf("error installing client %s: %v", res.ClientID, res.Error)
		}
		if !res.Created {
			t.Errorf("expected Created=true for new install on %s", res.ClientID)
		}
	}

	// Second run should mark Updated=true
	resultsUpdated, err := mcp.InstallClients(opts)
	if err != nil {
		t.Fatalf("second InstallClients failed: %v", err)
	}
	for _, res := range resultsUpdated {
		if !res.Updated {
			t.Errorf("expected Updated=true on second install for %s", res.ClientID)
		}
	}

	// Test "all"
	allOpts := mcp.InstallerOptions{
		Cwd:       cwd,
		HomeDir:   homeDir,
		ClientIDs: []string{"all"},
	}
	allResults, err := mcp.InstallClients(allOpts)
	if err != nil {
		t.Fatalf("InstallClients 'all' failed: %v", err)
	}
	if len(allResults) == 0 {
		t.Errorf("expected results for 'all'")
	}

	// Test unknown client ID
	unknownOpts := mcp.InstallerOptions{
		Cwd:       cwd,
		HomeDir:   homeDir,
		ClientIDs: []string{"unknown-client-xyz"},
	}
	_, err = mcp.InstallClients(unknownOpts)
	if err == nil {
		t.Errorf("expected error for unknown client ID")
	}
}

func TestGetSupportedClients(t *testing.T) {
	tmpDir := t.TempDir()
	clients := mcp.GetSupportedClients(tmpDir, tmpDir)
	if len(clients) == 0 {
		t.Errorf("expected non-empty supported clients list")
	}
}
