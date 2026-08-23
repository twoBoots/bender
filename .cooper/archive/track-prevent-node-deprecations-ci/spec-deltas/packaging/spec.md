# Capability Spec Delta: Prevent Node Deprecation Warnings in CI/CD Workflows

## Capability: `packaging`

### Description
Extends the installer and release packaging capability to require non-deprecated action runtimes and environment configurations across CI/CD workflows.

---

### Added Requirements

#### Requirement: CI/CD Workflow Deprecation Prevention
All GitHub Actions CI/CD workflows SHALL execute using supported, non-deprecated action versions and configure runner environment flags to prevent Node deprecation noise.

##### Scenario: Clean CI Validation Workflow Execution
- `+` GIVEN the `.github/workflows/ci.yml` workflow
- `+` WHEN triggered on pull request or push to `main`
- `+` THEN all action dependencies MUST run on supported Node runtimes
- `+` AND workflow-level environment flags MUST suppress Node runtime deprecation warnings.

##### Scenario: Clean Release Packaging Workflow Execution
- `+` GIVEN the `.github/workflows/release.yml` workflow
- `+` WHEN triggered on push to `main` or semantic release tag
- `+` THEN all action dependencies (checkout, setup-go, upload-artifact, download-artifact) MUST execute without deprecated runtime warnings.
