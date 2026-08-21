# Track Proposal: Bender Core Foundation

## Context & Motivation
Across the `twoBoots` ecosystem, CLI tools like `battery` and `cooper` (via RFC #9) require identical core infrastructure:
1. Command hierarchy with `-ldflags` version injection.
2. Self-updating from GitHub Releases with macOS codesign and Windows file locking support.
3. Embedded Model Context Protocol (MCP) server over `stdio` with automated multi-client configuration.
4. 3-tier progressive `install.sh` and 5-target GitHub Actions matrix release automation.

Currently, `battery` has a robust, tested implementation of these patterns. `bender` will serve as the generic, reusable template and package library to standardize this foundation for `cooper` and all future Go CLIs.

## Objectives
- Extract, generalize, and test `pkg/updater` and `pkg/mcp` as zero-external-dependency Go standard library packages.
- Implement standardized Cobra CLI entrypoint in `cmd/` with `bender update`, `bender mcp`, and `bender mcp install`.
- Provide reusable 3-tier `install.sh` and CI/CD GitHub workflows.
- Achieve >80% test coverage across all packages.
