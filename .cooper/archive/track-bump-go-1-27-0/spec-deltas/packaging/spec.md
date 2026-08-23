# Capability Spec Delta: Installer & Release Packaging

## Scope Diff
Updates the minimum Go runtime and compiler baseline requirement from Go 1.22+ to Go 1.27.0+ across installation fallback announcements and release packaging workflows.

## Requirement Deltas

### Requirement: 3-Tier Installation Strategy
The `install.sh` script SHALL support zero-prompt, progressive installation fallback across three tiers:
- 1. **Tier 1**: Compile locally if Go 1.22+ is available and source clone is present.
+ 1. **Tier 1**: Compile locally if Go 1.27.0+ is available and source clone is present.
2. **Tier 2**: Download pre-built binary matching OS/architecture from GitHub Releases.
3. **Tier 3**: Graceful zero-binary fallback for environments without binary access.

#### Scenario: Fallback Guidance on Missing Go Compiler
- GIVEN `install.sh` running in an environment without pre-built binary availability and without Go installed
- WHEN installation reaches Tier 3 fallback
- - THEN it MUST output guidance instructing the user to install Go 1.22+ or download the binary manually.
+ THEN it MUST output guidance instructing the user to install Go 1.27.0+ or download the binary manually.
