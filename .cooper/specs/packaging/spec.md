# Capability Specification: Installer & Release Packaging

## Purpose & Scope
Defines requirements for the 3-tier installation script (`install.sh`) and multi-architecture GitHub Actions release matrix.

## Requirements

### Requirement: 3-Tier Installation Strategy
The `install.sh` script SHALL support zero-prompt, progressive installation fallback across three tiers:
1. **Tier 1**: Compile locally if Go is available and source clone is present.
2. **Tier 2**: Download pre-built binary matching OS/architecture from GitHub Releases.
3. **Tier 3**: Graceful zero-binary fallback for environments without binary access.

#### Scenario: Local Compilation Tier
- GIVEN `install.sh` running in a local clone with `go` in `$PATH`
- WHEN installation executes
- THEN it MUST compile the binary, apply macOS codesign/quarantine fixes if on Darwin, and link to `/usr/local/bin` or `~/.local/bin`.

#### Scenario: Remote Binary Download Tier
- GIVEN `install.sh` executed via `curl | bash` without Go installed
- WHEN installation executes
- THEN it MUST detect OS/arch and download the prebuilt binary from GitHub Releases.
