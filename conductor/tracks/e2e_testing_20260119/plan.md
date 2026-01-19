# Implementation Plan: E2E Testing with creack/pty

## Phase 1: Infrastructure & Scaffolding [checkpoint: 823ea09]
- [x] Task: Create E2E test directory in Dev Project (5b558b2)
- [x] Task: Integrate pty dependency (1f81d99)
- [x] Task: Conductor - User Manual Verification 'Phase 1'

## Phase 2: Build & Environment Utilities [checkpoint: 250174d]
- [x] Task: Create a local build utility for the submodule (6162a94)
- [x] Task: Implement Tmux session management for tests (70c7f62)
- [x] Task: Conductor - User Manual Verification 'Phase 2'

## Phase 3: First E2E Test Development
- [x] Task: Implement the "Startup and UI Detection" test (6ea4d7b)
    - [x] Write the test using `creack/pty` to verify the initial TUI view
    - [x] Ensure the test builds the binary on-the-fly
    - [x] Implement profile isolation for tests
- [~] Task: Conductor - User Manual Verification 'Phase 3' (Protocol in workflow.md)
