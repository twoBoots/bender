# Implementation Plan: Bender Core Foundation

Track ID: `track-bender-core-foundation`
Status: `in_progress`

## Phase 1: Generic Self-Updater Engine (`pkg/updater`)

- [x] Task 1.1: Semantic Versioning Parser & Comparator (c5c3d19)
  - [x] Sub-task: Write unit tests for Semver parsing, cleaning, and comparison (Red)
  - [x] Sub-task: Implement `pkg/updater/semver.go` (Green)
  - [x] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [x] Task 1.2: Platform & Architecture Asset Resolution (24f58b4)
  - [x] Sub-task: Write unit tests for OS/Arch mapping across Darwin, Linux, Windows (Red)
  - [x] Sub-task: Implement `pkg/updater/platform.go` (Green)
  - [x] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [x] Task 1.3: GitHub Releases API Client (f800bca)
  - [x] Sub-task: Write HTTP mock tests for `FetchLatestRelease` and `FetchReleaseByTag` (Red)
  - [x] Sub-task: Implement `pkg/updater/client.go` with zero external dependencies (Green)
  - [x] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [x] Task 1.4: In-Place Atomic Binary Replacement & Codesigning (0400881)
  - [x] Sub-task: Write unit tests for atomic swap, temp file generation, and rollback (Red)
  - [x] Sub-task: Implement `pkg/updater/updater.go` with macOS quarantine stripping & codesign (Green)
  - [x] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [x] Task 1.5: Phase 1 Verification & Checkpoint [checkpoint: 42132d5]
  - [x] Sub-task: Run all package tests with `-race` and verify coverage >80%
  - [x] Sub-task: Synchronize upstream rules (`git fetch origin main`) and record checkpoint

## Phase 2: Generic MCP Server Engine (`pkg/mcp`)

- [ ] Task 2.1: JSON-RPC 2.0 Protocol Models & Serialization
  - [ ] Sub-task: Write tests for JSON-RPC 2.0 request/response/error marshaling (Red)
  - [ ] Sub-task: Implement `pkg/mcp/protocol.go` and MCP protocol types (Green)
  - [ ] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [ ] Task 2.2: Stdio MCP Server & Dynamic Dispatcher
  - [ ] Sub-task: Write tests for `initialize`, `tools/list`, `tools/call`, `resources/*`, `prompts/*` (Red)
  - [ ] Sub-task: Implement `pkg/mcp/server.go` and handler registration (Green)
  - [ ] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [ ] Task 2.3: Multi-Client MCP Configuration Installer
  - [ ] Sub-task: Write tests for merging config into Cursor, Claude Desktop, Antigravity JSON (Red)
  - [ ] Sub-task: Implement `pkg/mcp/installer.go` with path discovery and atomic JSON merging (Green)
  - [ ] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [ ] Task 2.4: Phase 2 Verification & Checkpoint
  - [ ] Sub-task: Run all package tests with `-race` and verify coverage >80%
  - [ ] Sub-task: Synchronize upstream rules (`git fetch origin main`) and record checkpoint

## Phase 3: Cobra CLI Interface & Commands (`cmd/` & `main.go`)

- [ ] Task 3.1: Root Command & Version Injection
  - [ ] Sub-task: Write CLI tests for `-v`, `--version`, `version` output (Red)
  - [ ] Sub-task: Implement `cmd/root.go` and `main.go` with `-ldflags` variables (Green)
  - [ ] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [ ] Task 3.2: Self-Update CLI Subcommand
  - [ ] Sub-task: Write CLI tests for `bender update` and `--force` flags (Red)
  - [ ] Sub-task: Implement `cmd/update.go` delegating to `pkg/updater` (Green)
  - [ ] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [ ] Task 3.3: MCP CLI Subcommands
  - [ ] Sub-task: Write CLI tests for `bender mcp` stdio server and `bender mcp install` (Red)
  - [ ] Sub-task: Implement `cmd/mcp.go` delegating to `pkg/mcp` (Green)
  - [ ] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [ ] Task 3.4: Example / Extensibility Subcommand
  - [ ] Sub-task: Write CLI test for template subcommand demonstrating consumer extension (Red)
  - [ ] Sub-task: Implement `cmd/example.go` (Green)
  - [ ] Sub-task: Refactor & verify test coverage >80% (Refactor)
- [ ] Task 3.5: Phase 3 Verification & Checkpoint
  - [ ] Sub-task: Run all CLI suite tests and verify coverage >80%
  - [ ] Sub-task: Synchronize upstream rules (`git fetch origin main`) and record checkpoint

## Phase 4: Packaging, Automation & Final Validation

- [ ] Task 4.1: Parameterized 3-Tier `install.sh`
  - [ ] Sub-task: Write shell validation tests for Tier 1 (local build), Tier 2 (release download), Tier 3 (fallback)
  - [ ] Sub-task: Implement and chmod `install.sh`
- [ ] Task 4.2: CI & Multi-Arch Release GitHub Workflows
  - [ ] Sub-task: Implement `.github/workflows/ci.yml` (test, race, lint, coverage)
  - [ ] Sub-task: Implement `.github/workflows/release.yml` (5-target matrix: linux/amd64, linux/arm64, darwin/amd64, darwin/arm64, windows/amd64)
- [ ] Task 4.3: End-to-End Test Suite & Living Spec Reconciliation
  - [ ] Sub-task: Verify entire test suite with race detector and coverage >80%
  - [ ] Sub-task: Reconcile spec-deltas into living specs (`.cooper/specs/`)
- [ ] Task 4.4: Phase 4 Verification & Final Checkpoint
  - [ ] Sub-task: Synchronize with `origin/main` and verify clean build
  - [ ] Sub-task: Push checkpoint commits to remote
