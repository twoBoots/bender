package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/twoBoots/bender/cmd"
)

func TestExampleCmd(t *testing.T) {
	var buf bytes.Buffer
	cmd.RootCmd.SetOut(&buf)
	cmd.RootCmd.SetArgs([]string{"hello", "Fry"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute hello failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Hello, Fry! Welcome to Bender.") {
		t.Errorf("unexpected hello output: %q", out)
	}
}

func TestExampleCmd_Default(t *testing.T) {
	var buf bytes.Buffer
	cmd.RootCmd.SetOut(&buf)
	cmd.RootCmd.SetArgs([]string{"hello"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute hello failed: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Hello, World! Welcome to Bender.") {
		t.Errorf("unexpected default hello output: %q", out)
	}
}
