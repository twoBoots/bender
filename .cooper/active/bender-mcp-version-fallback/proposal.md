# Proposal: Bender MCP Version Fallback Cleanup (bender)

## Intent
Remove the hardcoded `"1.0.0"` fallback version in `pkg/mcp/server.go` and replace it with `"dev"` (producing `"dev"` without a `"v"` prefix). Ensure tests cover default fallback behavior, explicit development version handling, and explicit release version normalization without hardcoding release versions into the MCP engine library.

## Scope & Boundaries
- Modify `pkg/mcp/server.go` `NewServer` constructor to default empty or whitespace versions to `"dev"`.
- Ensure `"dev"` is preserved as `"dev"` without adding a `"v"` prefix.
- Ensure semantic versions (e.g., `"1.0.0"`) continue to be normalized with a `"v"` prefix (e.g., `"v1.0.0"`).
- Extend `pkg/mcp/server_test.go` to test fallback and normalization scenarios across `NewServer` and `HandleRequest("initialize")`.

## User Benefit
Decouples the core MCP server library from hardcoded release version strings and provides clear, predictable development version metadata (`"dev"`) when instantiated without explicit build metadata.
