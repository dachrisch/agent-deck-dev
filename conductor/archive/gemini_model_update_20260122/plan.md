# Implementation Plan: Gemini Model Detection and Hardcoded List Update

## Phase 1: Update Hardcoded Model List [checkpoint: 05aad2a]
- [x] Task: Update model list in `internal/session/gemini.go` (05aad2a)
    - [x] Add `gemini-3-pro-preview`, `gemini-3-flash-preview`, `gemini-2.5-pro`, `gemini-2.5-flash`, and `gemini-2.5-flash-lite` to the fallback list.
- [x] Task: Verify updated model list (05aad2a)
    - [x] Add a unit test to `internal/session/gemini_test.go` to confirm the fallback list contains the new models.
- [x] Task: Conductor - User Manual Verification 'Phase 1' (Protocol in workflow.md) (05aad2a)

## Phase 2: Implement Dynamic Model Detection [checkpoint: 1231035]
- [x] Task: Implement model detection from session output (1231035)
    - [x] Update `internal/tmux/model_detection.go` to include a regex that can parse model names from the session's output stream.
- [x] Task: Integrate detection into session logic (1231035)
    - [x] Modify `internal/session/instance.go` to call the new detection logic and update the `GeminiModel` field.
- [x] Task: Add unit tests for model detection (1231035)
    - [x] Create tests in `internal/tmux/model_detection_test.go` to validate the regex against various output formats.
- [x] Task: Conductor - User Manual Verification 'Phase 2' (Protocol in workflow.md) (1231035)

## Phase 3: E2E Verification [checkpoint: 1231035]
- [x] Task: Add E2E test for model detection (1231035)
    - [x] Create a new test in `tests/e2e/model_switch_test.go` that starts a session, simulates a model change in the tmux pane, and verifies the TUI label updates.
- [x] Task: Conductor - User Manual Verification 'Phase 3' (Protocol in workflow.md) (1231035)