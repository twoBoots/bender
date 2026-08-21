# 🤖 `bender`: Generic Go CLI Archetype, Self-Updater & MCP Server Engine

`bender` is the foundational template and core package library for building high-performance, production-ready Go command-line applications across the **[twoBoots](https://github.com/twoBoots)** ecosystem.

It encapsulates standardized CLI ergonomics, zero-prompt multi-tier installers, atomic GitHub Release self-updaters, and embedded **Model Context Protocol (MCP)** server engines with zero runtime dependencies.

---

## 🌟 Key Features

- **Generic Self-Updater (`pkg/updater`)**:
  - Zero external dependencies using Go standard library primitives.
  - Queries GitHub Releases API with semantic version comparison (`CompareVersions`).
  - Cross-platform asset mapping (`darwin`, `linux`, `windows` across `x86_64` and `aarch64`).
  - In-place atomic binary replacement with automatic Windows file-lock rollback and macOS quarantine stripping + ad-hoc codesigning (`codesign -s - --force`).
- **Embedded MCP Server Engine (`pkg/mcp`)**:
  - Lightweight stdio JSON-RPC 2.0 protocol engine.
  - Dynamic registration for MCP **Tools**, **Resources**, and **Prompts**.
  - Multi-client MCP configuration installer with safe JSON merging for **Cursor**, **Google Antigravity / agy**, **Claude Desktop**, **Claude Code**, **Windsurf**, and **VS Code**.
- **Standardized Cobra CLI Architecture (`cmd/`)**:
  - Modular command hierarchy with build-time `-ldflags` version, commit SHA, and build date injection.
  - Built-in `bender update`, `bender mcp`, `bender mcp install`, and sample `bender hello` commands.
- **Progressive 3-Tier Installer & CI/CD Automation**:
  - Zero-prompt `install.sh` supporting Tier 1 (local Go compilation), Tier 2 (prebuilt binary release download), and Tier 3 (fallback).
  - GitHub Actions matrix releasing 5 platform binaries (`darwin-aarch64`, `darwin-x86_64`, `linux-aarch64`, `linux-x86_64`, `windows-x86_64.exe`).

---

## 📦 Installation

### Quick Install (Remote or Local)

```bash
# Remote install via curl
curl -fsSL https://raw.githubusercontent.com/twoBoots/bender/main/install.sh | bash

# Or run locally from clone
./install.sh
```

The installer will:
1. Detect your operating system and CPU architecture.
2. Compile and install locally if Go 1.22+ is available, or download the latest prebuilt release binary from GitHub Releases.
3. Automatically strip macOS quarantine attributes and apply code signing on Darwin.
4. Register the `bender` binary globally into `/usr/local/bin` or `~/.local/bin`.

---

## 🚀 CLI Commands

```bash
# Display help and command hierarchy
bender --help

# Inspect semantic version, commit hash, and build timestamp
bender --version
bender version

# Update Bender in-place to latest or specific tag
bender update
bender update --check
bender update --target-version v1.2.0 --force

# Launch stdio Model Context Protocol (MCP) server for AI coding assistants
bender mcp

# Automatically configure Bender MCP server in AI assistants
bender mcp install
bender mcp install --client cursor,antigravity --non-interactive
bender mcp install --all

# Sample extension command
bender hello Fry
```

---

## 🤖 Model Context Protocol (MCP) Integration

`bender` natively serves JSON-RPC 2.0 MCP requests over `stdio` (`bender mcp`), allowing AI assistants like Antigravity, Claude Code, Cursor, and Windsurf to interact directly with the CLI engine.

### Automatic Configuration
To configure your local AI editors automatically:
```bash
bender mcp install
```

### Manual Configuration Example
In your editor's MCP configuration (`mcp_config.json`, `.cursor/mcp.json`, or `claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "bender": {
      "command": "bender",
      "args": ["mcp"]
    }
  }
}
```

---

## 🛠️ Reusability: Using Bender as a Library

Projects can import `pkg/updater` and `pkg/mcp` directly without external dependencies:

### 1. In-Place Self-Update in your Go CLI
```go
package main

import (
	"fmt"
	"github.com/twoBoots/bender/pkg/updater"
)

func main() {
	res, err := updater.SelfUpdate(updater.Options{
		Repo:           "my-org/my-cli",
		BinaryName:     "my-cli",
		CurrentVersion: "v1.0.0",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)
}
```

### 2. Standalone MCP Server in your Go Application
```go
package main

import (
	"context"
	"os"
	"github.com/twoBoots/bender/pkg/mcp"
)

func main() {
	srv := mcp.NewServer("my-mcp", "1.0.0", ".")

	srv.RegisterTool(mcp.Tool{
		Name:        "ping",
		Description: "Returns pong",
		InputSchema: mcp.ToolInputSchema{Type: "object"},
	}, func(ctx context.Context, args map[string]interface{}) (mcp.CallToolResult, error) {
		return mcp.NewTextResult("pong", false), nil
	})

	_ = srv.Serve(context.Background(), os.Stdin, os.Stdout)
}
```

---

## 🧪 Development & Quality Control

`bender` is developed using the **Cooper Spec-Driven Development (SDD)** lifecycle and strict TDD:

```bash
# Run unit & CLI test suite with coverage
CGO_ENABLED=0 go test -v -coverprofile=coverage.out ./...

# Verify code formatting & linting
gofmt -l .
go vet ./...

# Compile local static binary
CGO_ENABLED=0 go build -o bin/bender .
```

---

## 🔗 Architecture & Ecosystem Links

- [Product Vision](.cooper/definition/product.md)
- [Technical Stack Definition](.cooper/definition/tech-stack.md)
- [Living Capability Specs](.cooper/specs/)
  - [CLI Lifecycle Spec](.cooper/specs/cli/spec.md)
  - [Self-Updater Engine Spec](.cooper/specs/updater/spec.md)
  - [Embedded MCP Server Spec](.cooper/specs/mcp/spec.md)
  - [Packaging & Release Matrix Spec](.cooper/specs/packaging/spec.md)
- [Cooper SDD Framework](https://github.com/twoBoots/cooper)
- [Troop Worktree Isolation](https://github.com/twoBoots/troop)
- [Battery Orchestration Protocol](https://github.com/twoBoots/battery)
