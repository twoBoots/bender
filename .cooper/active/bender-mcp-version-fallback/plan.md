# Implementation Plan: Bender MCP Version Fallback Cleanup

## Phase 1: MCP Server Version Fallback & Normalization

- [x] Task 1: Add Unit Tests for Version Fallback & Normalization (Red) (8a3c4b7)
  - [x] Sub-task: Write unit tests in `pkg/mcp/server_test.go` for default fallback to `"dev"` on empty string
  - [x] Sub-task: Write unit tests for whitespace-only version fallback to `"dev"`
  - [x] Sub-task: Write unit tests for explicit `"dev"` version preserving `"dev"` without `v` prefix
  - [x] Sub-task: Write unit tests for explicit semantic version (e.g. `"1.0.0"`) normalizing to `"v1.0.0"`
  - [x] Sub-task: Confirm tests fail against existing hardcoded `"1.0.0"` implementation (Red)

- [x] Task 2: Implement Fallback & Normalization in `pkg/mcp/server.go` (Green & Refactor) (c2ec3bd)
  - [x] Sub-task: Update `NewServer` constructor logic to default empty/whitespace to `"dev"` and skip `"v"` prefix for `"dev"`
  - [x] Sub-task: Run unit tests to verify all test cases pass (Green)
  - [x] Sub-task: Run linter/formatting checks and ensure test coverage >80% (Refactor)

- [~] Task 3: Phase 1 Verification & Checkpoint
  - [ ] Sub-task: Synchronize rules and living capability specs (`git fetch origin main`)
  - [ ] Sub-task: Run full test suite (`go test -v ./...`)
  - [ ] Sub-task: Create checkpoint commit and attach verification Git Note
  - [ ] Sub-task: Synchronize branch with remote (`git push origin bender-mcp-version-fallback`)
