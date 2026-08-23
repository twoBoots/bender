# Track Proposal: Standardize SemVer Comparison Using golang.org/x/mod/semver

## Context & Motivation
Currently, `pkg/updater/semver.go` uses a custom string-splitting and integer-parsing implementation for semantic version comparison. While it achieved zero external dependencies, it lacks complete SemVer 2.0.0 compliance (such as complex pre-release tag precedence and validation).

The Go core team provides `golang.org/x/mod/semver`, an official, lightweight, and robust package for SemVer 2.0.0 parsing and comparison. Adopting `golang.org/x/mod/semver` standardizes version comparison behavior across the `twoBoots` ecosystem while retaining special-case ergonomics like development builds (`dev`) and unprefixed versions.

## Objectives
- Integrate official `golang.org/x/mod/semver` into `pkg/updater/semver.go`.
- Retain backwards-compatibility for `"dev"` builds and missing `"v"` prefixes.
- Expand test suite in `pkg/updater/semver_test.go` to verify standard SemVer 2.0.0 compliance.
- Maintain test coverage >80% with all tests passing.
