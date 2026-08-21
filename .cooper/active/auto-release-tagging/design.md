# Technical Design: Automated Git Release Tagging

## 1. Architecture Overview

```mermaid
flowchart TD
    PUSH["Push / Merge to main"] --> CI["Job: CI Verification"]
    CI --> TAG_JOB["Job: Auto-Tag on Main"]
    
    subgraph Auto-Tag Logic
        TAG_JOB --> GET_VER["Extract Version from cmd/root.go (e.g. 1.0.0)"]
        GET_VER --> CHECK_TAG["Check if v${VERSION} exists on remote origin"]
        CHECK_TAG -->|Missing| CREATE_TAG["Create and push tag v${VERSION} using GITHUB_TOKEN"]
        CHECK_TAG -->|Exists| PROCEED["Proceed with existing tag"]
    end
    
    CREATE_TAG --> MATRIX["Job: Build & Release Matrix (5 platforms)"]
    PROCEED --> MATRIX
    MATRIX --> PUBLISH["Job: Publish GitHub Release (v${VERSION} + latest)"]
```

## 2. Implementation Strategy

### Workflow Configuration (`.github/workflows/release.yml`)
1. **Pre-build Auto-Tag Step / Job**:
   - Executes when `github.ref == 'refs/heads/main'`.
   - Reads `Version` from `cmd/root.go`.
   - Uses `gh api` or `git ls-remote` / `git push` to verify and push `v${VERSION}` if missing.
2. **Release Publishing**:
   - `gh release create "v${VERSION}" release-assets/* --title "Bender v${VERSION}" --generate-notes`
   - Uploads to both `v${VERSION}` tag and updates `latest`.

## 3. Security & Quality Controls
- Uses standard repository `GITHUB_TOKEN` with `contents: write` permissions.
- Idempotent execution: if `v${VERSION}` already exists, no duplicate tags or error aborts occur.
