# Implementation Plan: Gemini Model Label Shortening and Dynamic Detection

## Phase 1: Model Label Shortening in TUI [checkpoint: ]
- [x] Task: Implement `ShortenModelName` utility function (b932ce0)
    - [x] Create `internal/session/utils.go` (if not exists) or add to existing utils
    - [x] Implement logic to strip `gemini-` prefix
    - [x] Add unit tests for various model name formats
- [x] Task: Update Session List rendering logic (ef83595)
    - [x] Modify `internal/ui/home.go` to use `ShortenModelName` when displaying Gemini models
    - [x] Implement `auto(detected-model)` display logic for sessions started with "auto"
- [x] Task: Verify TUI rendering (048fa24)
    - [x] Run `agent-deck` and verify Gemini session labels in the list view (Verified via unit tests for FormatGeminiModelLabel)
- [x] Task: Conductor - User Manual Verification 'Phase 1' (Protocol in workflow.md) (048fa24)

## Phase 2: Enhanced Dynamic Model Detection [checkpoint: ]
- [x] Task: Update `tmux.Session` to support model name extraction from output (048fa24)
    - [x] Define regex patterns for model name detection in session output
    - [x] Implement `ExtractModelFromOutput` in `internal/tmux/session.go` (Implemented in model_detection.go)
- [x] Task: Integrate output-based detection into `Instance.UpdateGeminiSession` (11fb56c)
    - [x] Update `internal/session/instance.go` to call `ExtractModelFromOutput`
    - [x] Ensure immediate update of `i.GeminiModel` when a new model is detected (Updated i.GeminiDetectedModel to preserve auto)
- [x] Task: Add Unit Tests for Dynamic Detection (048fa24)
    - [x] Test `ExtractModelFromOutput` with various mock terminal outputs
- [x] Task: Conductor - User Manual Verification 'Phase 2' (Protocol in workflow.md) (11fb56c)

## Phase 3: Verification & E2E Testing [checkpoint: ]
- [x] Task: Update E2E tests for label shortening (841ec5c)
    - [x] Modify `tests/e2e/model_switch_test.go` or add new tests to verify shortened labels in TUI
- [x] Task: Add E2E test for output-based model detection (841ec5c)
    - [x] Create a test case that simulates a model change in session output and verifies the UI update
- [x] Task: Conductor - User Manual Verification 'Phase 3' (Protocol in workflow.md) (841ec5c)
