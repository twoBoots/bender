# Capability Specification: CLI Lifecycle & Commands

## Purpose & Scope
Defines core command-line ergonomics, version reporting, Cobra routing, and top-level entrypoints for Bender-based CLIs.

## Requirements

### Requirement: Modular Command Hierarchy
The CLI SHALL provide a Cobra-based command hierarchy with built-in version inspection and help flag routing.

#### Scenario: Version Reporting
- GIVEN a compiled CLI binary built with `-ldflags`
- WHEN the user executes `bender --version` or `bender version`
- THEN the CLI MUST print the injected semantic version and git commit metadata.

#### Scenario: Default Help & Subcommand Routing
- GIVEN the CLI root command
- WHEN the user executes with no arguments or `--help`
- THEN the CLI MUST print structured help text detailing all available subcommands.
