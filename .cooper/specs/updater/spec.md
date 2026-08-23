# Capability Specification: Self-Updater Engine

## Purpose & Scope
Defines requirements for checking upstream GitHub releases, comparing semantic versions, downloading platform-specific binaries, and performing atomic in-place binary upgrades.

## Requirements

### Requirement: GitHub Release Version Check
The updater SHALL query the GitHub Releases API for the latest published tag and determine if an update is available.

#### Scenario: Newer Version Available
- GIVEN a running CLI with version `v1.0.0`
- WHEN `updater.SelfUpdate` is called and remote release tag is `v1.1.0`
- THEN `Result.UpdateAvailable` MUST be true and `Result.LatestVersion` MUST equal `v1.1.0`.

#### Scenario: Already Up To Date
- GIVEN a running CLI with version `v1.0.0`
- WHEN `updater.SelfUpdate` is called and remote release tag is `v1.0.0`
- THEN `Result.UpdateAvailable` MUST be false and no download occurs unless `Force` is specified.

### Requirement: Platform Binary Asset Resolution
The updater SHALL map the current OS (`darwin`, `linux`, `windows`) and architecture (`amd64`/`x86_64`, `arm64`/`aarch64`) to the corresponding release asset binary name convention.

#### Scenario: Platform Binary Resolution
- GIVEN `runtime.GOOS == "darwin"` and `runtime.GOARCH == "arm64"`
- WHEN resolving platform binary name for binary `bender`
- THEN the resolved asset name MUST be `bender-darwin-aarch64`.

### Requirement: Atomic In-Place Upgrade
The updater SHALL download the new binary to a temporary file in the executable's directory, set executable permissions (`0755`), and replace the active executable atomically.

#### Scenario: macOS Quarantine and Code Signing
- GIVEN an update running on Darwin
- WHEN the new binary is written
- THEN the updater MUST strip `com.apple.quarantine` and apply ad-hoc code signature (`codesign -s - --force`).

### Requirement: Standard Semantic Version Comparison
The updater SHALL use standard SemVer 2.0.0 comparison rules via `golang.org/x/mod/semver` while maintaining backward compatibility for development build tags (`dev`) and unprefixed/partial SemVer strings.

#### Scenario: SemVer 2.0.0 Pre-release Comparison
- GIVEN local version `v1.0.0-rc.1` and remote version `v1.0.0-rc.2`
- WHEN `updater.CompareVersions` is called
- THEN the result MUST be `-1` (remote is newer).

#### Scenario: Development Build Version Comparison
- GIVEN local version `dev` and remote version `v1.0.0`
- WHEN `updater.CompareVersions` is called
- THEN the result MUST be `-1` (remote is newer than local development build).
