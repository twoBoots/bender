# Track Design: Bump Go Runtime to 1.27.0

## Architecture Decisions

### 1. Go Module Runtime Directive
- Update `go.mod` to declare `go 1.27.0`.
- Maintain standard library compatibility across `pkg/updater`, `pkg/mcp`, and `cmd/`.

### 2. CI/CD Matrix Environment
- Update GitHub Actions workflow configurations (`.github/workflows/ci.yml` and `.github/workflows/release.yml`) so `actions/setup-go@v5` uses `go-version: "1.27.0"`.
- Ensure multi-architecture compilation matrix (`linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`, `windows/amd64`) operates cleanly under the updated toolchain.

### 3. Installer & Documentation Alignment
- In `install.sh`, align Tier 3 fallback advice to prompt users for `Go 1.27.0+` if binary fetching fails.
- In `README.md` and `.cooper/definition/tech-stack.md`, update references from `Go 1.22+` to `Go 1.27.0+`.

## Risk Analysis & Mitigation
- **Toolchain Availability**: If Go 1.27.0 requires toolchain management or newer compiler features, standard Go modules and GitHub Actions `setup-go` handle official releases seamlessly.
- **Backwards Compatibility**: Codebase currently uses standard library constructs without deprecated patterns.
