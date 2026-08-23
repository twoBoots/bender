# Implementation Plan: Bump Go Runtime to 1.27.0

## Phase 1: Go Toolchain & Runtime Upgrade

- [x] Task: Update Go Module Runtime Directive (a86393f)
  - [x] Sub-task: Update `go.mod` to declare `go 1.27.0`
  - [x] Sub-task: Verify `go.mod` and `go.sum` consistency
- [x] Task: Update GitHub Actions Workflows (d123678)
  - [x] Sub-task: Update `ci.yml` to use `go-version: "1.27.0"`
  - [x] Sub-task: Update `release.yml` to use `go-version: "1.27.0"`
- [x] Task: Update Installer Script & Documentation (abe797f)
  - [x] Sub-task: Update `install.sh` Tier 3 fallback message to reference Go 1.27.0+
  - [x] Sub-task: Update `README.md` and `.cooper/definition/tech-stack.md`
- [x] Task: Phase 1 Verification & Checkpoint (dcb6860) [checkpoint: dcb6860]
  - [x] Sub-task: Execute test suite across all packages (`go test ./...`)
  - [x] Sub-task: Run formatting and linting verification (`gofmt -l .`, `go vet ./...`)
  - [x] Sub-task: Perform git phase sync and record checkpoint notes
