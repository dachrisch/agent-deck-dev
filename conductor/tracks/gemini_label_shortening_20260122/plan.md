# Implementation Plan: Gemini Model Label Shortening and Dynamic Detection

## Phase 1: Model Label Shortening in TUI [checkpoint: ]
- [x] Task: Implement `ShortenModelName` utility function (b932ce0)
    - [x] Create `internal/session/utils.go` (if not exists) or add to existing utils
    - [x] Implement logic to strip `gemini-` prefix
    - [x] Add unit tests for various model name formats
- [x] Task: Update Session List rendering logic (ef83595)
    - [x] Modify `internal/ui/home.go` to use `ShortenModelName` when displaying Gemini models
    - [x] Implement `auto(detected-model)` display logic for sessions started with "auto"
- [ ] Task: Verify TUI rendering
    - [ ] Run `agent-deck` and verify Gemini session labels in the list view
- [ ] Task: Conductor - User Manual Verification 'Phase 1' (Protocol in workflow.md)

## Phase 2: Enhanced Dynamic Model Detection [checkpoint: ]
- [ ] Task: Update `tmux.Session` to support model name extraction from output
    - [ ] Define regex patterns for model name detection in session output
    - [ ] Implement `ExtractModelFromOutput` in `internal/tmux/session.go`
- [ ] Task: Integrate output-based detection into `Instance.UpdateGeminiSession`
    - [ ] Update `internal/session/instance.go` to call `ExtractModelFromOutput`
    - [ ] Ensure immediate update of `i.GeminiModel` when a new model is detected
- [ ] Task: Add Unit Tests for Dynamic Detection
    - [ ] Test `ExtractModelFromOutput` with various mock terminal outputs
- [ ] Task: Conductor - User Manual Verification 'Phase 2' (Protocol in workflow.md)

## Phase 3: Verification & E2E Testing [checkpoint: ]
- [ ] Task: Update E2E tests for label shortening
    - [ ] Modify `tests/e2e/model_switch_test.go` or add new tests to verify shortened labels in TUI
- [ ] Task: Add E2E test for output-based model detection
    - [ ] Create a test case that simulates a model change in session output and verifies the UI update
- [ ] Task: Conductor - User Manual Verification 'Phase 3' (Protocol in workflow.md)
