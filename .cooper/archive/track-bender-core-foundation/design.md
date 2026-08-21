# Technical Design: Bender Core Foundation

## Package Architecture

```
bender/
├── cmd/
│   ├── root.go             # Cobra root command with -v/--version and global flags
│   ├── update.go           # 'update' command delegating to pkg/updater
│   ├── mcp.go              # 'mcp' and 'mcp install' commands delegating to pkg/mcp
│   └── example.go          # Template subcommand demonstrating CLI extension
├── pkg/
│   ├── updater/            # Generic self-updater engine
│   │   ├── client.go       # GitHub Releases API HTTP client
│   │   ├── platform.go     # Runtime OS/Arch to asset name mapping
│   │   ├── semver.go       # Semantic version comparison
│   │   └── updater.go      # Binary download and atomic swap
│   └── mcp/                # Generic Model Context Protocol engine
│       ├── protocol.go     # JSON-RPC 2.0 and MCP types
│       ├── server.go       # Stdio JSON-RPC 2.0 server
│       ├── installer.go    # Client config merger (Cursor, Claude, Antigravity, etc.)
│       └── types.go        # Tools, Resources, and Prompts interfaces
├── install.sh              # Parameterized 3-tier installation script
├── .github/workflows/
│   ├── ci.yml              # Go test, lint, race, and coverage validation
│   └── release.yml         # GitHub Actions release matrix
└── main.go                 # Binary entrypoint
```

## Reusability Strategy
- `pkg/updater` and `pkg/mcp` are designed with generic options structs (`updater.Options`, `mcp.Config`), accepting configurable GitHub repository strings, binary names, and custom tool/resource/prompt registrations.
- Projects adopting Bender can either:
  1. Fork/template the repository to start a new CLI.
  2. Import `github.com/twoBoots/bender/pkg/updater` and `github.com/twoBoots/bender/pkg/mcp` directly in `go.mod`.
