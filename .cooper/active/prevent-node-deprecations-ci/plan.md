# Implementation Plan: Prevent Node Deprecation Warnings in CI/CD Workflows

## Phase 1: Workflow Updates & Action Modernization
- [~] Task 1: Update GitHub Actions workflows (`ci.yml`, `release.yml`) with modern action versions and runner flags
  - [ ] Sub-task: Configure top-level environment variables (`FORCE_JAVASCRIPT_ACTIONS_TO_NODE20: true`, `NODE_NO_WARNINGS: "1"`) in `ci.yml`
  - [ ] Sub-task: Configure top-level environment variables in `release.yml`
  - [ ] Sub-task: Audit and pin all action steps to supported major versions
- [ ] Task 2: Phase 1 Verification & Checkpoint
  - [ ] Sub-task: Validate workflow YAML syntax and consistency
  - [ ] Sub-task: Run local test suite (`go test -v -coverprofile=coverage.out ./...`) and verify >80% coverage
  - [ ] Sub-task: Record Git Notes checkpoint summary
