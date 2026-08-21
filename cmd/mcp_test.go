package cmd_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/twoBoots/bender/cmd"
	"github.com/twoBoots/bender/pkg/mcp"
)

func TestMCPCmd_Serve(t *testing.T) {
	callParams, _ := json.Marshal(mcp.CallToolParams{Name: "get_version"})
	rawCallParams := json.RawMessage(callParams)
	reqCall := mcp.Request{
		JSONRPC: "2.0",
		ID:      newRawMsg(`1`),
		Method:  "tools/call",
		Params:  rawCallParams,
	}
	reqBytes, _ := json.Marshal(reqCall)

	var in bytes.Buffer
	in.WriteString("{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"ping\"}\n")
	in.Write(reqBytes)
	in.WriteString("\n")
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
	if !strings.Contains(res, "Bender v") {
		t.Errorf("expected get_version response in mcp output, got %q", res)
	}
}

func TestMCPInstallCmd(t *testing.T) {
	tmpDir := t.TempDir()
	homeDir := filepath.Join(tmpDir, "home")
	_ = os.MkdirAll(homeDir, 0755)

	t.Setenv("HOME", homeDir)
	defer func() {
		_ = os.RemoveAll(".cursor")
		_ = os.RemoveAll(".vscode")
	}()

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

	// Test --all flag
	var outAll bytes.Buffer
	cmd.RootCmd.SetOut(&outAll)
	cmd.RootCmd.SetArgs([]string{"mcp", "install", "--all"})
	err = cmd.Execute()
	if err != nil {
		t.Fatalf("Execute mcp install --all failed: %v", err)
	}

	// Test default non-interactive
	var outDefault bytes.Buffer
	cmd.RootCmd.SetOut(&outDefault)
	cmd.RootCmd.SetArgs([]string{"mcp", "install"})
	err = cmd.Execute()
	if err != nil {
		t.Fatalf("Execute default mcp install failed: %v", err)
	}
}

func newRawMsg(s string) *json.RawMessage {
	m := json.RawMessage(s)
	return &m
}
