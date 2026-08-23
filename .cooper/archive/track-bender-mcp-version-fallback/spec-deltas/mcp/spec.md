# Spec Delta: Bender MCP Version Fallback Cleanup

## Capability: `mcp`

### Requirement: JSON-RPC 2.0 Protocol Lifecycle

#### Scenario: Server Version Fallback
- `+` GIVEN an MCP server initialized without an explicit version or an empty version string
- `+` WHEN the server constructs its implementation metadata
- `+` THEN it MUST default the version to `"dev"` rather than a hardcoded semantic release version.

#### Scenario: Explicit Server Version Formatting
- `+` GIVEN an MCP server initialized with an explicit semantic version string (e.g. `"1.0.0"`)
- `+` WHEN the server processes initialization requests
- `+` THEN it MUST normalize the version to include the `"v"` prefix (e.g. `"v1.0.0"`), while preserving `"dev"` without a `"v"` prefix.
