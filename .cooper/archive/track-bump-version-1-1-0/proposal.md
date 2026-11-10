# Proposal: Bump CLI Version to 1.1.0 for Release

## Intent & Scope
Following the merge of the standardized SemVer comparison engine using `golang.org/x/mod/semver` (Track `standardize-semver-xmod`, PR #7) and the Cooper v1.2.0 workflow skills synchronization (PR #8), this track bumps the base application version in `cmd/root.go` from `1.0.1` to `1.1.0`.

## User Value
- Allows downstream users and self-updating CLI installations (`bender update`) to detect and consume the latest minor release `v1.1.0` featuring robust SemVer comparison across patch/prerelease updates.
- Enables automated release tagging (`v1.1.0`) and multi-architecture binary artifact publishing via the GitHub Actions release workflow upon merge to `main`.
