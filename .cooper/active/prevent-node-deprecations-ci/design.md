# Technical Design: Prevent Node Deprecation Warnings in CI/CD Workflows

## 1. Overview & Architecture
This design modernizes the GitHub Actions CI/CD workflows for `bender` by upgrading all core actions and declaring workflow-level environment variables to suppress and override legacy Node runtime execution.

## 2. GitHub Actions Upgrades & Configuration

### 2.1 Action Version Alignments
The workflows currently declare:
- `actions/checkout@v4`
- `actions/setup-go@v5`
- `actions/upload-artifact@v4`
- `actions/download-artifact@v4`

All actions will be verified and pinned to their latest stable non-deprecated versions that execute on modern Node runtimes.

### 2.2 Workflow-Level Environment Runtime Variables
Both `.github/workflows/ci.yml` and `.github/workflows/release.yml` will define top-level environment variables:
```yaml
env:
  FORCE_JAVASCRIPT_ACTIONS_TO_NODE20: true
  NODE_NO_WARNINGS: "1"
```
- `FORCE_JAVASCRIPT_ACTIONS_TO_NODE20: true`: Directs the runner to execute JavaScript actions on Node.js 20+ runtime.
- `NODE_NO_WARNINGS: "1"`: Suppresses runtime deprecation warnings from Node process invocations.

## 3. Workflow File Updates

### 3.1 `.github/workflows/ci.yml`
Add top-level `env` block and ensure actions use latest stable releases:
```yaml
name: CI

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main
  workflow_call:

env:
  FORCE_JAVASCRIPT_ACTIONS_TO_NODE20: true
  NODE_NO_WARNINGS: "1"

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v4
      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: "1.22"
          cache: true
      ...
```

### 3.2 `.github/workflows/release.yml`
Add top-level `env` block and ensure all jobs (`auto-tag`, `build-and-release`, `publish-release`) inherit clean execution settings.

## 4. Verification & Testing
- Lint and validate YAML workflow files.
- Run `go test -v -coverprofile=coverage.out ./...` to ensure project tests and coverage remain intact (>80%).
- Verify formatting with `gofmt -l .` and `go vet ./...`.
