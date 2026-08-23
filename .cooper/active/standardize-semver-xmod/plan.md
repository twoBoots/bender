# Implementation Plan: Standardize SemVer Comparison Using golang.org/x/mod/semver

## Phase 1: Dependency Integration & SemVer Engine Refactor
- [ ] Task: Add golang.org/x/mod Dependency
  - [ ] Sub-task: Run `go get golang.org/x/mod/semver` and verify `go.mod`/`go.sum`
- [ ] Task: SemVer Parsing & Comparison Refactor (TDD)
  - [ ] Sub-task: Expand unit tests in `pkg/updater/semver_test.go` covering SemVer 2.0 edge cases and dev builds (Red)
  - [ ] Sub-task: Implement `pkg/updater/semver.go` using `golang.org/x/mod/semver` (Green)
  - [ ] Sub-task: Refactor, verify linting, and ensure >80% test coverage (Refactor)
- [ ] Task: Phase 1 Verification & Checkpoint
  - [ ] Sub-task: Run full test suite (`go test -race ./...`)
  - [ ] Sub-task: Commit phase checkpoint and sync notes
