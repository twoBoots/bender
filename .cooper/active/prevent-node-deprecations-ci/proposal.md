# Track Proposal: Prevent Node Deprecation Warnings in CI/CD Workflows

- **Track ID**: `prevent-node-deprecations-ci`
- **Type**: `chore`
- **Status**: Planning

---

## 1. Context & Motivation
GitHub Actions runners continually deprecate older Node.js runtimes (such as Node.js 16 and Node.js 20) and transition actions to modern runtimes (Node.js 20 / Node.js 24). Running older action versions or running without runner runtime flags results in noisy annotations, deprecation warnings, and potential pipeline failure when older runtimes are terminated.

In `bender`, the CI validation workflow (`.github/workflows/ci.yml`) and multi-architecture release workflow (`.github/workflows/release.yml`) should execute cleanly without runner deprecation warnings by updating all action dependencies to their latest major non-deprecated versions and configuring workflow runtime flags.

## 2. Goals & User Benefit
1. **Zero Deprecation Noise**: Eliminate Node.js runtime deprecation warnings across all GitHub Actions workflow runs.
2. **Action Modernization**: Upgrade all action steps (`actions/checkout`, `actions/setup-go`, `actions/upload-artifact`, `actions/download-artifact`) to modern, actively maintained versions.
3. **Resilience & Future-Proofing**: Add workflow-level runner environment flags (`FORCE_JAVASCRIPT_ACTIONS_TO_NODE20: true` and `NODE_NO_WARNINGS: "1"`) ensuring clean workflow execution across runner upgrades.
4. **Living Spec Grounding**: Extend the `packaging` capability specification to codify clean, deprecation-free CI/CD execution standards.

## 3. Scope Boundaries
- Update `.github/workflows/ci.yml` action versions and runner environment configuration.
- Update `.github/workflows/release.yml` action versions and runner environment configuration.
- Verify workflow syntax and local build/test validation.
- Register and track under Cooper Spec-Driven Development.
