# Plan: Update Hardcoded Gemini Model List

This plan outlines the steps to update the fallback Gemini model list to include the latest version 3 and 2.5 models.

## Phase 1: Gemini Model Updates (TDD)

- [x] Task: TDD - Add Test Case for Updated Model List (1112d61)
    - [x] Update `internal/session/gemini_test.go` to verify that `GetAvailableGeminiModels()` returns the new list when the API is unavailable.
    - [x] Run the test and confirm it fails.
- [x] Task: TDD - Update Hardcoded Model List (1112d61)
    - [x] Modify `geminiModelFallback` in `internal/session/gemini.go` to include the five new models.
    - [x] Run the test and confirm it passes.
- [~] Task: Conductor - User Manual Verification 'Phase 1: Gemini Model Updates (TDD)' (Protocol in workflow.md)

## Phase 2: Final Verification & Quality Gates [checkpoint: 86d9199]

- [x] Task: Full Project Verification (86d9199)
    - [x] Run all unit tests (`go test ./...`) and E2E tests (`cd tests/e2e && go test -v .`).
    - [x] Run linter and static analysis (`go vet ./...`).
- [x] Task: Conductor - User Manual Verification 'Phase 2: Final Verification & Quality Gates' (Protocol in workflow.md) (86d9199)
