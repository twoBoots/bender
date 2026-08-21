# Track Proposal: Automated Git Release Tagging on Merge to Main

- **Track ID**: `auto-release-tagging`
- **Type**: `feature`
- **Status**: Planning

---

## 1. Context & Motivation
`bender` is the foundational Go CLI archetype and core engine library (`pkg/updater`, `pkg/mcp`) consumed across the `twoBoots` multi-barrel ecosystem. Downstream Go modules (`cooper`, `battery`) rely on Git SemVer tags (e.g. `v1.0.0`) on `origin` to resolve dependencies via `go get github.com/twoBoots/bender@v1.0.0`.

Currently, release tagging requires manual tagging steps, and merges to `main` without an existing tag publish only under an untagged `latest` asset or unindexed pseudo-version.

## 2. Goals & User Benefit
1. Automatically detect `Version` defined in `cmd/root.go` on push/merge to `main`.
2. Automatically create and push the Git tag `v${VERSION}` if it does not already exist on remote `origin`.
3. Publish GitHub Release assets under the semantic version tag (`v${VERSION}`) alongside `latest`.
4. Ensure downstream Go barrels can immediately resolve tagged Bender versions without manual intervention.

## 3. Scope Boundaries
- Update `.github/workflows/release.yml` with auto-tagging logic.
- Ensure proper permissions (`contents: write`) for `GITHUB_TOKEN`.
- Document automated release workflow in `README.md`.
