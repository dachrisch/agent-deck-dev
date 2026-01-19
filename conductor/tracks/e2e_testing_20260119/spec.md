# Specification: E2E Testing with tui-test

## Overview
Establish a robust end-to-end (E2E) testing framework for Agent Deck using `microsoft/tui-test`. This will allow for automated verification of the TUI's behavior, starting with a basic build-and-startup check.

## Goal
To ensure that Agent Deck builds correctly and the TUI initializes to a valid state in a real terminal environment.

## Functional Requirements
- **Test Infrastructure:** Create a `tests/e2e` directory in the Dev Project.
- **Dependency Management:** Integrate `microsoft/tui-test` to interact with and verify the TUI buffer.
- **Local Build Integration:** Automatically build the `agent-deck` binary from the submodule before test execution.
- **Tmux Integration:** Execute the startup test within a real `tmux` session to match the application's production environment.
- **First Test Case:** 
    - Build the application.
    - Launch Agent Deck.
    - Use `tui-test` to verify that a specific UI element (e.g., logo or Home tab) is rendered correctly.

## Non-Functional Requirements
- **Local Execution:** Optimized for local development and verification.
- **Environment Parity:** Tests should accurately reflect the TUI experience in a standard terminal.

## Acceptance Criteria
- [ ] `tests/e2e` directory exists in the Dev Project.
- [ ] A successful run of the startup test confirms the binary builds.
- [ ] The test confirms the app starts in `tmux` and renders initial UI elements.
- [ ] Test results are clearly reported in the console.

## Out of Scope
- Integration of this specific E2E suite into the remote CI pipeline (currently focused on local execution).
- Complex multi-session interaction tests (reserved for future tracks).
