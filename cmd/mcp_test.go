package cmd_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/twoBoots/bender/cmd"
)

func TestMCPCmd_Serve(t *testing.T) {
	var in bytes.Buffer
	in.WriteString("{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"ping\"}\n")
	var out bytes.Buffer

	cmd.RootCmd.SetIn(&in)
	cmd.RootCmd.SetOut(&out)
	cmd.RootCmd.SetArgs([]string{"mcp"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute mcp failed: %v", err)
	}

	res := out.String()
	if !strings.Contains(res, "\"jsonrpc\":\"2.0\"") {
		t.Errorf("expected jsonrpc response, got %q", res)
	}
}

func TestMCPInstallCmd(t *testing.T) {
	tmpDir := t.TempDir()
	homeDir := filepath.Join(tmpDir, "home")
	_ = os.MkdirAll(homeDir, 0755)

	t.Setenv("HOME", homeDir)

	var out bytes.Buffer
	cmd.RootCmd.SetOut(&out)
	cmd.RootCmd.SetArgs([]string{"mcp", "install", "--client", "cursor,antigravity", "--non-interactive"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute mcp install failed: %v", err)
	}

	res := out.String()
	if !strings.Contains(res, "Configuring Bender MCP Server") && !strings.Contains(res, "Configured") {
		t.Errorf("unexpected mcp install output: %q", res)
	}
}
