# Design: Version Bump 1.0.1

## Architecture & Implementation Plan

### 1. Version Constant Update
- Update `cmd.Version` in `cmd/root.go` to `"1.0.1"`.
- This ensures runtime reporting via `bender --version` and Cobra commands default to `v1.0.1` unless overridden by `-ldflags` during packaging builds.

### 2. CI/CD Release Pipeline Triggering
- The `auto-tag` job in `.github/workflows/release.yml` uses `grep` to extract `Version` from `cmd/root.go`.
- With `Version = "1.0.1"`, the workflow will check if `v1.0.1` exists on origin. Since it does not exist, the workflow will automatically create and push tag `v1.0.1` and publish release assets for `v1.0.1` and `latest`.

### 3. Verification & Testing Strategy
- Unit tests in `cmd/root_test.go` will be updated to verify `cmd.Version` reflects `"1.0.1"` and conforms to SemVer format.
- Run full repository test suite across `cmd`, `pkg/mcp`, and `pkg/updater` with 100% pass rate.
