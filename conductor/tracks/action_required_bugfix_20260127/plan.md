# Plan: Fix Claude "Action Required" Status Detection

This plan outlines the steps to enhance Claude Code's status detection to correctly identify when it is waiting for user confirmation.

## Phase 1: Test-Driven Development [checkpoint: 6a84afb]

- [x] Task: TDD - Add Failing Tests for "Action Required" (5cd896b)
    - [x] Update `TestPromptDetector` in `internal/tmux/tmux_test.go` with cases for "Action Required", "Waiting for user confirmation", and "Allow execution of".
    - [x] Run tests and confirm they fail.
- [x] Task: TDD - Enhance Claude Prompt Detection (5cd896b)
    - [x] Modify `hasClaudePrompt` in `internal/tmux/detector.go` to include the new patterns in the `permissionPrompts` list.
    - [x] Run tests and confirm they now pass.
- [x] Task: Fix Goroutine Leak in LogWatcher Callback (6a84afb)
    - [x] Implement worker pool (`logUpdateChan`, `logWorker`) in `internal/ui/home.go`.
    - [x] Replace unbounded goroutine spawning with bounded worker pool.
- [x] Task: Implement Quitting Splash Screen
    - [x] Add `isQuitting` field to `Home` struct.
    - [x] Implement `renderQuittingSplash` function.
    - [x] Update `View` to show quitting splash when `isQuitting` is true.
    - [x] Update quit flow (`tryQuit`, `performQuit`, `Update`) to handle delayed shutdown with visual feedback.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Test-Driven Development' (Protocol in workflow.md) (6a84afb)

## Phase 2: Final Verification & Quality Gates [checkpoint: 6a84afb]

- [x] Task: Full Project Verification (6a84afb)
    - [x] Run all unit tests (`go test ./...`) and E2E tests (`cd tests/e2e && go test -v .`).
    - [x] Run linter and static analysis (`go vet ./...`).
- [x] Task: Conductor - User Manual Verification 'Phase 2: Final Verification & Quality Gates' (Protocol in workflow.md) (6a84afb)
