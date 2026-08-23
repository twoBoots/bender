# Capability Specification Delta: Installer & Release Packaging

## Requirements

### Requirement: Automated Git SemVer Tagging on Merge to Main

#### Scenario: Tagging on Merge to Main
- `-` GIVEN a pull request merged to `main` with `cmd.Version` set to `1.0.0`
- `+` GIVEN a pull request merged to `main` with `cmd.Version` set to `1.0.1`
- WHEN the `Release Binary` GitHub Actions workflow triggers on `refs/heads/main`
- `-` THEN it MUST verify if `v1.0.0` exists on `origin`
- `+` THEN it MUST verify if `v1.0.1` exists on `origin`
- `-` AND IF missing, it MUST create and push Git tag `v1.0.0` to `origin`
- `+` AND IF missing, it MUST create and push Git tag `v1.0.1` to `origin`
- `-` AND it MUST publish the release assets under `v1.0.0` as well as updating `latest`.
- `+` AND it MUST publish the release assets under `v1.0.1` as well as updating `latest`.
