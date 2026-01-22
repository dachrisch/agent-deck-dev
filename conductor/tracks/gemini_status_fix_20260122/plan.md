# Implementation Plan: Fix Gemini Status Detection

## Phase 1: Research and Reproduce [checkpoint: c99d070]
- [x] Task: Create a reproduction E2E test (c99d070)
    - [x] Create `tests/e2e/gemini_status_test.go`
    - [x] Implement a test that starts a Gemini session, sends a long-running prompt (e.g., "tell me a 500 word story"), and asserts that the status becomes `Running` during generation.
    - [x] Verify that the test fails as expected (Status stays `Idle` or `Waiting`).
- [x] Task: Conductor - User Manual Verification 'Phase 1' (Protocol in workflow.md) (c99d070)

## Phase 2: Enhance Status Detection Logic [checkpoint: 86ad2eb]
- [x] Task: Expand busy indicator detection in `internal/tmux/tmux.go` (86ad2eb)
    - [x] Increase the line range for spinner detection in `hasBusyIndicator` from 5 to 15 lines.
    - [x] Add Gemini-specific busy indicators (e.g., streaming output patterns) to `hasBusyIndicator`.
- [x] Task: Refine status mapping in `internal/session/instance.go` (86ad2eb)
    - [x] Ensure `UpdateStatus` correctly handles the transition from `Waiting` to `Running` for Gemini sessions.
- [x] Task: Add unit tests for the improved detection logic (86ad2eb)
    - [x] Add test cases to `internal/tmux/tmux_test.go` using captured Gemini output strings to verify `hasBusyIndicator` returns `true`.
- [x] Task: Conductor - User Manual Verification 'Phase 2' (Protocol in workflow.md) (86ad2eb)

## Phase 3: Final Verification [checkpoint: 86ad2eb]
- [x] Task: Run full E2E suite (86ad2eb)
    - [x] Ensure `TestGeminiStatus` passes.
    - [x] Verify no regressions in Claude or OpenCode status detection.
- [x] Task: Conductor - User Manual Verification 'Phase 3' (Protocol in workflow.md) (86ad2eb)
