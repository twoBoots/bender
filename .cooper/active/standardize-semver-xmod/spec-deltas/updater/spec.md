# Spec Delta: Self-Updater Engine

## Capability: `updater`

### Added Requirements

+ ### Requirement: Standard Semantic Version Comparison
+ The updater SHALL use standard SemVer 2.0.0 comparison rules via `golang.org/x/mod/semver` while maintaining backward compatibility for development build tags (`dev`) and unprefixed/partial SemVer strings.
+ 
+ #### Scenario: SemVer 2.0.0 Pre-release Comparison
+ - GIVEN local version `v1.0.0-rc.1` and remote version `v1.0.0-rc.2`
+ - WHEN `updater.CompareVersions` is called
+ - THEN the result MUST be `-1` (remote is newer).
+ 
+ #### Scenario: Development Build Version Comparison
+ - GIVEN local version `dev` and remote version `v1.0.0`
+ - WHEN `updater.CompareVersions` is called
+ - THEN the result MUST be `-1` (remote is newer than local development build).
