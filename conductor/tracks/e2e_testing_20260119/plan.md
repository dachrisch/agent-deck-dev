# Implementation Plan: E2E Testing with tui-test

## Phase 1: Infrastructure & Scaffolding
- [ ] Task: Create E2E test directory in Dev Project
    - [ ] Create `tests/e2e` directory
    - [ ] Initialize a Go module in `tests/e2e` (since tui-test is a Go library)
- [ ] Task: Integrate `microsoft/tui-test` dependency
    - [ ] Add `github.com/microsoft/tui-test` to the new go module
- [ ] Task: Conductor - User Manual Verification 'Phase 1' (Protocol in workflow.md)

## Phase 2: Build & Environment Utilities
- [ ] Task: Create a local build utility for the submodule
    - [ ] Implement a helper to build the `agent-deck` binary from `./agent-deck/cmd/agent-deck`
- [ ] Task: Implement Tmux session management for tests
    - [ ] Create helper to launch and clean up the TUI within a controlled tmux session
- [ ] Task: Conductor - User Manual Verification 'Phase 2' (Protocol in workflow.md)

## Phase 3: First E2E Test Development
- [ ] Task: Implement the "Startup and UI Detection" test
    - [ ] Write the test using `tui-test` to verify the logo or initial TUI view
    - [ ] Ensure the test builds the binary on-the-fly as per requirements
- [ ] Task: Conductor - User Manual Verification 'Phase 3' (Protocol in workflow.md)
