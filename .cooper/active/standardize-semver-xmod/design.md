# Design: Standardize SemVer Comparison Using golang.org/x/mod/semver

## Architecture & Integration

### Dependency
Add `golang.org/x/mod` module dependency to `go.mod`.

### SemVer Normalization & Comparison Strategy
1. **Normalization (`NormalizeVersion`)**:
   - Trim whitespace and leading `v`/`V`.
   - Used for display or comparison where raw clean semver string is required.

2. **Canonical Version for `x/mod/semver`**:
   - `golang.org/x/mod/semver` requires version strings to start with `v` (e.g. `v1.2.3`).
   - If version has only major.minor (e.g. `0.1` or `v0.1`), expand to full 3-part semver (e.g. `v0.1.0`) so `semver.IsValid` and `semver.Compare` function properly.
   - If version does not start with `v`, prefix `v` before passing to `semver.Compare`.

3. **Special Cases**:
   - `"dev"` builds: Treated as older than any released semantic version (`-1`), and `"dev" == "dev"` (`0`).
   - Invalid versions fallback: If versions cannot be parsed by standard semver, fallback to string or numerical segment comparison.

### Component Updates
- `pkg/updater/semver.go`: Reimplement `CompareVersions` using `golang.org/x/mod/semver`.
- `pkg/updater/semver_test.go`: Add comprehensive table-driven tests for SemVer 2.0 prerelease rules, `"dev"` builds, and prefix handling.
