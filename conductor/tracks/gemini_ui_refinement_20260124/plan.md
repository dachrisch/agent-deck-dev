# Plan: Gemini UI Refinement - Model Relocation

This plan outlines the steps to relocate the Gemini model display from the tree view to the right panel.

## Phase 1: Foundation & Cleanup

- [x] Task: TDD - Revert tree view model labels
    - [x] Update `internal/ui/home_test.go` (or equivalent) to expect simple `gemini` labels in the tree view.
    - [x] Modify `renderSessionItem` in `internal/ui/home.go` to remove the model suffix from the tool label.
    - [x] Verify that Gemini sessions now display as just `gemini` in the left panel.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Foundation & Cleanup' (Protocol in workflow.md)

## Phase 2: Right Panel Enhancements

- [x] Task: TDD - Implement right panel model indicator
    - [x] Add tests to verify the presence and content of the "Model" field in the session info block.
    - [x] Update `renderSessionDetails` (or equivalent) in `internal/ui/home.go` to include the "Model" field below "Session ID".
    - [x] Implement the styling for the field to look like a dropdown (visual border/background).
- [x] Task: TDD - Implement dynamic auto-labeling
    - [x] Write tests for the `auto (detected-model)` formatting logic.
    - [x] Update the rendering logic to check for detected models when the session is in "auto" mode.
- [x] Task: Conductor - User Manual Verification 'Phase 2: Right Panel Enhancements' (Protocol in workflow.md)

## Phase 3: Layout & Verification

- [x] Task: Ensure layout stability and theme consistency
    - [x] Verify that the right panel uses `ensureExactWidth` to prevent layout bleeding.
    - [x] Perform a manual UI check to ensure Tokyo Night theme compliance.
- [x] Task: Final E2E Verification
    - [x] Run full E2E suite (`tests/e2e`) to ensure no regressions in model detection or persistence.
- [x] Task: Conductor - User Manual Verification 'Phase 3: Layout & Verification' (Protocol in workflow.md)
