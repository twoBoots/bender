# Technology Stack Definition: Bender
 
## Primary Languages & Runtime
- **Language**: Go
- **Runtime / Engine**: Go 1.22+
- **Module Path**: `github.com/twoBoots/bender`
 
## Architecture & Frameworks
- **Application Architecture**: Modular CLI / Embedded MCP Server / Generic Package Suite
- **Core Frameworks**: [spf13/cobra](https://github.com/spf13/cobra) for CLI command routing
- **Standard Library Primitives**: `net/http`, `encoding/json`, `os`, `runtime`, `bufio` (Zero external dependencies for core `pkg/updater` and `pkg/mcp`)
 
## Testing & Quality Control
- **Test Runner / Framework**: Go standard testing package (`testing`)
- **Coverage Target**: >80% code coverage across all modules
- **Linter & Formatter**: `gofmt`, `go vet`
 
## Build & CI/CD
- **Package Manager**: Go Modules (`go.mod`)
- **CI System**: GitHub Actions (Matrix build for darwin-arm64, darwin-amd64, linux-amd64, linux-arm64, windows-amd64)
