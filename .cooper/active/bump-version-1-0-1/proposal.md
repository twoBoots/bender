# Proposal: Bump CLI Version to 1.0.1 for Release

## Intent & Scope
Following the merge of several critical fixes and upgrades—including the MCP concurrency deadlock resolution, MCP server version normalization and fallback to "dev", CI/CD Node deprecation prevention, and Go 1.27.0 toolchain upgrade—this track bumps the base application version in `cmd/root.go` from `1.0.0` to `1.0.1`.

## User Value
- Allows downstream users and self-updating CLI installations (`bender update`) to detect and consume the latest patch release `v1.0.1`.
- Enables automated release tagging (`v1.0.1`) and multi-architecture asset publishing via GitHub Actions release workflow upon merge to `main`.
