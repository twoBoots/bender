# Technical Design: Bender MCP Version Fallback Cleanup

## Overview
This design details the replacement of the hardcoded `"1.0.0"` fallback version in the embedded MCP server constructor (`pkg/mcp/server.go`) with `"dev"` (producing `"dev"` without a `"v"` prefix). It specifies the version normalization logic and corresponding unit test coverage.

## Architecture & Logic Changes

### 1. Version Normalization in `pkg/mcp/server.go`
In `NewServer(name, version, cwd string) *Server`:
```go
ver := strings.TrimSpace(version)
if ver == "" {
    ver = "dev"
} else if ver != "dev" && !strings.HasPrefix(ver, "v") {
    ver = "v" + ver
}
```

### 2. Behavioral Rules Matrix
| Input `version` | Normalized `Server.version` | Rationale |
|---|---|---|
| `""` (empty) | `"dev"` | Default development fallback |
| `"   "` (whitespace) | `"dev"` | Trimmed development fallback |
| `"dev"` | `"dev"` | Explicit dev mode, un-prefixed |
| `"1.0.0"` | `"v1.0.0"` | Semantic release version normalized with `v` prefix |
| `"v1.0.0"` | `"v1.0.0"` | Pre-normalized semantic release version preserved |

### 3. Component Breakdown
- `pkg/mcp/server.go`: Updates to `NewServer` version initialization logic.
- `pkg/mcp/server_test.go`:
  - Unit tests asserting `srv.Version()` and `InitializeResult.ServerInfo.Version` for empty string, whitespace string, explicit `"dev"`, `"1.0.0"`, and `"v1.0.0"`.

## Risk & Compatibility Analysis
- **Breaking Changes**: None. Callers providing release versions via build flags (e.g. `cmd/mcp.go` passing `cmd.Version`) continue to produce normalized versions (e.g. `"v1.0.0"`).
- **Test Compatibility**: Existing tests passing explicit versions (e.g. `"1.0.0"`, `"0.1.0"`) will continue to pass.
