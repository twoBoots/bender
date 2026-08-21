# Capability Spec Delta: Automated Git Release Tagging

## Capability: `packaging`

### Description
Extends the release packaging capability to automate semantic Git tag creation and GitHub Release publishing upon merge to the `main` branch.

---

### Added Requirements

#### Requirement: Automated Git SemVer Tagging on Merge to Main
`+` The CI/CD release workflow SHALL automatically detect the application semantic version and publish a Git tag matching `v<Version>` upon push or merge to `main` if the tag does not already exist.

##### Scenario: Tagging on Merge to Main
- `+` GIVEN a pull request merged to `main` with `cmd.Version` set to `1.0.0`
- `+` WHEN the `Release Binary` GitHub Actions workflow triggers on `refs/heads/main`
- `+` THEN it MUST verify if `v1.0.0` exists on `origin`
- `+` AND IF missing, it MUST create and push Git tag `v1.0.0` to `origin`
- `+` AND it MUST publish the release assets under `v1.0.0` as well as updating `latest`.

##### Scenario: Idempotent Tag Handling
- `+` GIVEN a push to `main` where Git tag `v<Version>` already exists on `origin`
- `+` WHEN the release workflow runs
- `+` THEN it MUST proceed without error, updating assets for `v<Version>` and `latest`.
