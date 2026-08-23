# Capability Specification: Installer & Release Packaging

## Purpose & Scope
Defines requirements for the 3-tier installation script (`install.sh`) and multi-architecture GitHub Actions release matrix.

## Requirements

### Requirement: 3-Tier Installation Strategy
The `install.sh` script SHALL support zero-prompt, progressive installation fallback across three tiers:
1. **Tier 1**: Compile locally if Go 1.27.0+ is available and source clone is present.
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

#### Scenario: Fallback Guidance on Missing Go Compiler
- GIVEN `install.sh` running in an environment without pre-built binary availability and without Go installed
- WHEN installation reaches Tier 3 fallback
- THEN it MUST output guidance instructing the user to install Go 1.27.0+ or download the binary manually.

### Requirement: Automated Git SemVer Tagging on Merge to Main
The CI/CD release workflow SHALL automatically detect the application semantic version and publish a Git tag matching `v<Version>` upon push or merge to `main` if the tag does not already exist.

#### Scenario: Tagging on Merge to Main
- GIVEN a pull request merged to `main` with `cmd.Version` set to `1.0.1`
- WHEN the `Release Binary` GitHub Actions workflow triggers on `refs/heads/main`
- THEN it MUST verify if `v1.0.1` exists on `origin`
- AND IF missing, it MUST create and push Git tag `v1.0.1` to `origin`
- AND it MUST publish the release assets under `v1.0.1` as well as updating `latest`.

#### Scenario: Idempotent Tag Handling
- GIVEN a push to `main` where Git tag `v<Version>` already exists on `origin`
- WHEN the release workflow runs
- THEN it MUST proceed without error, updating assets for `v<Version>` and `latest`.

