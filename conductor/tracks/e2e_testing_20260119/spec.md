# Specification: E2E Testing with creack/pty

## Overview
Establish a robust end-to-end (E2E) testing framework for Agent Deck using `github.com/creack/pty`. This allows for automated verification of the TUI's behavior by interacting with the pseudo-terminal output.

## Goal
To ensure that Agent Deck builds correctly and the TUI initializes to a valid state in an isolated terminal environment.

## Functional Requirements
- **Test Infrastructure:** Create a `tests/e2e` directory in the Dev Project.
- **Dependency Management:** Integrate `github.com/creack/pty` to interact with the TUI output.
- **Local Build Integration:** Automatically build the `agent-deck` binary from the submodule before test execution.
- **Profile Isolation:** Execute tests using a temporary profile and isolated HOME directory to avoid lock file conflicts.
- **First Test Case:** 
    - Build the application.
    - Launch Agent Deck using a PTY.
    - Verify that "Agent Deck" is rendered in the TUI output.

## Acceptance Criteria
- [x] `tests/e2e` directory exists in the Dev Project.
- [x] A successful run of the startup test confirms the binary builds.
- [x] The test confirms the app starts and renders initial UI elements.
