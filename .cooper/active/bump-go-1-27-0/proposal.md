# Track Proposal: Bump Go Runtime to 1.27.0

## Context & Motivation
The Bender CLI archetype and core package library currently target Go 1.22. Upgrading the Go runtime directive and CI/CD pipelines to Go 1.27.0 ensures that the project leverages modern Go standard library capabilities, compiler optimizations, and security patches while maintaining consistent runtime definitions across development, local installation, and multi-architecture GitHub Actions release workflows.

## User & Developer Benefit
- Ensures developers and automated CI/CD jobs build and test Bender against Go 1.27.0.
- Keeps release workflow binaries compiled with the Go 1.27.0 toolchain.
- Updates documentation (`README.md`, `.cooper/definition/tech-stack.md`) and installer runtime guidance (`install.sh`) to accurately reflect Go 1.27.0+ prerequisites.

## Scope Boundaries
- **In Scope**:
  - Update `go.mod` module directive to `go 1.27.0`.
  - Update `.github/workflows/ci.yml` and `.github/workflows/release.yml` `setup-go` action to `1.27.0` (or `1.27`).
  - Update `install.sh` fallback messaging to recommend `Go 1.27.0+`.
  - Update `README.md` and `.cooper/definition/tech-stack.md` documentation to specify `Go 1.27.0+`.
  - Verify all package builds and unit tests pass.
- **Out of Scope**:
  - Refactoring existing packages or altering public APIs of `pkg/updater` or `pkg/mcp`.
  - Introducing new external third-party dependencies.
