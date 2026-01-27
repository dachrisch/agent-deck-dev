# Plan: Fix Claude "Action Required" Status Detection

This plan outlines the steps to enhance Claude Code's status detection to correctly identify when it is waiting for user confirmation.

## Phase 1: Test-Driven Development

- [ ] Task: TDD - Add Failing Tests for "Action Required"
    - [ ] Update `TestPromptDetector` in `internal/tmux/tmux_test.go` with cases for "Action Required", "Waiting for user confirmation", and "Allow execution of".
    - [ ] Run tests and confirm they fail.
- [ ] Task: TDD - Enhance Claude Prompt Detection
    - [ ] Modify `hasClaudePrompt` in `internal/tmux/detector.go` to include the new patterns in the `permissionPrompts` list.
    - [ ] Run tests and confirm they now pass.
- [ ] Task: Conductor - User Manual Verification 'Phase 1: Test-Driven Development' (Protocol in workflow.md)

## Phase 2: Final Verification & Quality Gates

- [ ] Task: Full Project Verification
    - [ ] Run all unit tests (`go test ./...`) and E2E tests (`cd tests/e2e && go test -v .`).
    - [ ] Run linter and static analysis (`go vet ./...`).
- [ ] Task: Conductor - User Manual Verification 'Phase 2: Final Verification & Quality Gates' (Protocol in workflow.md)
