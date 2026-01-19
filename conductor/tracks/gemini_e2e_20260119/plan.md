# Implementation Plan: E2E Test for Gemini Model Switching

## Phase 1: Test Infrastructure & Mocking
- [x] Task: Implement Gemini API Mocking for tests (f72ec90)
    - [x] Add support for local model listing overrides in `internal/session/gemini.go` if necessary
- [x] Task: Create `tests/e2e/model_switch_test.go` scaffolding (f6efea6)
    - [x] Setup temporary home and profile for the test
- [~] Task: Conductor - User Manual Verification 'Phase 1' (Protocol in workflow.md)

## Phase 2: Test Automation Logic
- [ ] Task: Automate Gemini Session Creation & Model Dialog
    - [ ] Implement PTY input sequence to create a session and press `Ctrl+G`
- [ ] Task: Implement Model Selection & Verification
    - [ ] Send `Enter` to select a model
    - [ ] Implement `tmux` command check to verify the new process arguments
- [ ] Task: Conductor - User Manual Verification 'Phase 2' (Protocol in workflow.md)
