# Capability Specification: Embedded MCP Server

## Purpose & Scope
Defines requirements for the embedded Model Context Protocol (MCP) server running JSON-RPC 2.0 over `stdio`, supporting tools, resources, prompts, and client configuration installation.

## Requirements

### Requirement: JSON-RPC 2.0 Protocol Lifecycle
The MCP server SHALL communicate over standard input/output (`stdio`) implementing standard MCP handshake and method dispatching.

#### Scenario: Server Initialization Handshake
- GIVEN an MCP client sending an `initialize` JSON-RPC request
- WHEN the server processes the request
- THEN it MUST respond with protocol version, server implementation name/version, and server capabilities (Tools, Resources, Prompts).

#### Scenario: Server Version Fallback
- GIVEN an MCP server initialized without an explicit version or an empty version string
- WHEN the server constructs its implementation metadata
- THEN it MUST default the version to `"dev"` rather than a hardcoded semantic release version.

#### Scenario: Explicit Server Version Formatting
- GIVEN an MCP server initialized with an explicit semantic version string (e.g. `"1.0.0"`)
- WHEN the server processes initialization requests
- THEN it MUST normalize the version to include the `"v"` prefix (e.g. `"v1.0.0"`), while preserving `"dev"` without a `"v"` prefix.

#### Scenario: Tool Discovery & Execution
- GIVEN registered tools on the MCP server
- WHEN a client sends `tools/list` or `tools/call`
- THEN the server MUST return registered tools metadata or execute the corresponding handler returning structured results.

### Requirement: Multi-Client MCP Configuration Installer
The MCP installer SHALL inspect and safely merge the server command configuration into target AI coding assistants without destroying existing server configurations.

#### Scenario: Merge Configuration
- GIVEN an existing client JSON config (e.g. Cursor, Claude, Antigravity)
- WHEN running client configuration installer
- THEN the server configuration MUST be inserted or updated under `mcpServers.<server_name>` preserving all other keys.
