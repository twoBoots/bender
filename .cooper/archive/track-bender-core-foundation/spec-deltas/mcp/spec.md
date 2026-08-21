# Spec Delta: Embedded MCP Server

## Capability: `mcp`

### Added Requirements

+ ### Requirement: JSON-RPC 2.0 Protocol Lifecycle
+ The MCP server SHALL communicate over standard input/output (`stdio`) implementing standard MCP handshake and method dispatching.
+ 
+ #### Scenario: Server Initialization Handshake
+ - GIVEN an MCP client sending an `initialize` JSON-RPC request
+ - WHEN the server processes the request
+ - THEN it MUST respond with protocol version, server implementation name/version, and server capabilities (Tools, Resources, Prompts).
+ 
+ #### Scenario: Tool Discovery & Execution
+ - GIVEN registered tools on the MCP server
+ - WHEN a client sends `tools/list` or `tools/call`
+ - THEN the server MUST return registered tools metadata or execute the corresponding handler returning structured results.
+ 
+ ### Requirement: Multi-Client MCP Configuration Installer
+ The MCP installer SHALL inspect and safely merge the server command configuration into target AI coding assistants without destroying existing server configurations.
+ 
+ #### Scenario: Merge Configuration
+ - GIVEN an existing client JSON config (e.g. Cursor, Claude, Antigravity)
+ - WHEN running client configuration installer
+ - THEN the server configuration MUST be inserted or updated under `mcpServers.<server_name>` preserving all other keys.
