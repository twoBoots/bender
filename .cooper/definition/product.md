# Product Definition: Bender

## Vision
Bender is the foundational template and core archetype for building high-performance Go-based CLIs in the `twoBoots` ecosystem. It encapsulates standardized CLI ergonomics, zero-prompt multi-tier installers, GitHub Release self-updaters, and embedded Model Context Protocol (MCP) server support.

## Target Audience
- Developers building CLI tools within the `twoBoots` ecosystem (e.g. `battery`, `cooper`, and future tools).
- AI coding agents orchestrating workflows, tools, and updates via the Model Context Protocol.
- End users installing, updating, and executing standalone static binaries across macOS, Linux, and Windows.

## Core Capabilities & Initial Scope
- **Generic Self-Updater (`pkg/updater`)**: Zero-dependency engine to query GitHub Releases, compare semver, download platform-specific binaries, perform atomic in-place executable replacement with Windows rename rollback, and handle macOS quarantine/codesigning.
- **Embedded MCP Server (`pkg/mcp`)**: Lightweight stdio JSON-RPC 2.0 protocol engine for registering tools, resources, and prompts, plus automated client configuration installation (for Cursor, Google Antigravity, Claude Desktop, Claude Code, Windsurf, VS Code).
- **Standardized Cobra CLI Foundation (`cmd/`)**: Modular Cobra command hierarchy with version information injection via `-ldflags`, verbose/quiet logging, and consistent error handling.
- **Tiered Installer & Release Matrix**: 3-tier `install.sh` script (local Go build -> pre-built release binary -> fallback) and GitHub Actions release matrix for 5 targets (`darwin-aarch64`, `darwin-x86_64`, `linux-aarch64`, `linux-x86_64`, `windows-x86_64.exe`).

## Quality & Non-Functional Goals
- **Zero Runtime Dependencies**: Single static Go binary without external runtime requirements.
- **High Test Coverage**: >80% code coverage across all core packages.
- **High Performance & Low Latency**: Fast startup (<10ms for CLI execution) and low memory footprint (<15MB RAM for MCP server lifecycle).
