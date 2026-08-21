package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/twoBoots/bender/cmd"
)

func TestRootCmd_Version(t *testing.T) {
	cmd.Version = "1.0.0"
	cmd.Commit = "abcdef"
	cmd.BuildDate = "2026-08-21"

	var buf bytes.Buffer
	cmd.RootCmd.SetOut(&buf)
	cmd.RootCmd.SetArgs([]string{"--version"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "bender") || !strings.Contains(out, "1.0.0") {
		t.Errorf("expected version output to contain 'bender' and '1.0.0', got %q", out)
	}
}

func TestRootCmd_Help(t *testing.T) {
	var buf bytes.Buffer
	cmd.RootCmd.SetOut(&buf)
	cmd.RootCmd.SetArgs([]string{"--help"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Usage:") {
		t.Errorf("expected help output to contain 'Usage:', got %q", out)
	}
}
