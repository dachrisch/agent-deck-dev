# Plan: Gemini Refactoring and Code Cleanup

This plan outlines the steps to refactor the Gemini session management logic and consolidate utility functions.

## Phase 1: Utility Consolidation [checkpoint: df6bade]

- [x] Task: TDD - Consolidate findNewestFile
    - [x] Create `internal/session/utils_test.go` with tests for `findNewestFile`.
    - [x] Create `internal/session/utils.go` and implement `findNewestFile`.
    - [x] Update `internal/session/gemini.go` to remove duplicate implementations.
    - [x] Verify all tests pass and coverage is maintained.
- [x] Task: Conductor - User Manual Verification 'Phase 1: Utility Consolidation' (Protocol in workflow.md)

## Phase 2: UpdateGeminiSession Refactoring

- [x] Task: TDD - Extract syncGeminiSessionFromTmux
    - [x] Add unit tests for `syncGeminiSessionFromTmux` in `internal/session/gemini_refactor_test.go`.
    - [x] Implement `syncGeminiSessionFromTmux` in `internal/session/instance.go`.
    - [x] Update `UpdateGeminiSession` to use this helper.
- [x] Task: TDD - Extract syncGeminiSessionFromFile
    - [x] Add unit tests for `syncGeminiSessionFromFile`.
    - [x] Implement `syncGeminiSessionFromFile` in `internal/session/instance.go`.
    - [x] Update `UpdateGeminiSession` to use this helper.
- [x] Task: TDD - Extract updateGeminiAnalytics
    - [x] Add unit tests for `updateGeminiAnalytics`.
    - [x] Implement `updateGeminiAnalytics` in `internal/session/instance.go`.
    - [x] Update `UpdateGeminiSession` to use this helper.
- [x] Task: TDD - Extract detectGeminiModelRealTime
    - [x] Add unit tests for `detectGeminiModelRealTime`.
    - [x] Implement `detectGeminiModelRealTime` in `internal/session/instance.go`.
    - [x] Update `UpdateGeminiSession` to use this helper.
- [x] Task: TDD - Extract updateGeminiLatestPrompt
    - [x] Add unit tests for `updateGeminiLatestPrompt`.
    - [x] Implement `updateGeminiLatestPrompt` in `internal/session/instance.go`.
    - [x] Update `UpdateGeminiSession` to use this helper.
- [x] Task: Final Logic Verification
    - [x] Run full E2E suite (`tests/e2e`) to ensure no regressions in dynamic labeling or status detection.
- [x] Task: Conductor - User Manual Verification 'Phase 2: UpdateGeminiSession Refactoring' (Protocol in workflow.md)

## Phase 3: Final Verification & Cleanup [checkpoint: 782dfb5]

- [x] Task: Code style and documentation audit
    - [ ] Ensure all new standalone functions are documented.
    - [ ] Run `go fmt` and `go vet` on modified files.
- [x] Task: Final E2E Verification
    - [ ] Run all project tests to ensure system-wide stability.
- [x] Task: Conductor - User Manual Verification 'Phase 3: Final Verification & Cleanup' (Protocol in workflow.md)