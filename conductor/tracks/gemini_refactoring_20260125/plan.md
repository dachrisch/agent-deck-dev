# Plan: Gemini Refactoring and Code Cleanup

This plan outlines the steps to refactor the Gemini session management logic and consolidate utility functions.

## Phase 1: Utility Consolidation

- [ ] Task: TDD - Consolidate findNewestFile
    - [ ] Create `internal/session/utils_test.go` with tests for `findNewestFile`.
    - [ ] Create `internal/session/utils.go` and implement `findNewestFile`.
    - [ ] Update `internal/session/gemini.go` to remove duplicate implementations.
    - [ ] Verify all tests pass and coverage is maintained.
- [ ] Task: Conductor - User Manual Verification 'Phase 1: Utility Consolidation' (Protocol in workflow.md)

## Phase 2: UpdateGeminiSession Refactoring

- [ ] Task: TDD - Extract syncGeminiSessionFromTmux
    - [ ] Add unit tests for `syncGeminiSessionFromTmux` in `internal/session/gemini_refactor_test.go`.
    - [ ] Implement `syncGeminiSessionFromTmux` in `internal/session/instance.go`.
    - [ ] Update `UpdateGeminiSession` to use this helper.
- [ ] Task: TDD - Extract syncGeminiSessionFromFile
    - [ ] Add unit tests for `syncGeminiSessionFromFile`.
    - [ ] Implement `syncGeminiSessionFromFile` in `internal/session/instance.go`.
    - [ ] Update `UpdateGeminiSession` to use this helper.
- [ ] Task: TDD - Extract updateGeminiAnalytics
    - [ ] Add unit tests for `updateGeminiAnalytics`.
    - [ ] Implement `updateGeminiAnalytics` in `internal/session/instance.go`.
    - [ ] Update `UpdateGeminiSession` to use this helper.
- [ ] Task: TDD - Extract detectGeminiModelRealTime
    - [ ] Add unit tests for `detectGeminiModelRealTime`.
    - [ ] Implement `detectGeminiModelRealTime` in `internal/session/instance.go`.
    - [ ] Update `UpdateGeminiSession` to use this helper.
- [ ] Task: TDD - Extract updateGeminiLatestPrompt
    - [ ] Add unit tests for `updateGeminiLatestPrompt`.
    - [ ] Implement `updateGeminiLatestPrompt` in `internal/session/instance.go`.
    - [ ] Update `UpdateGeminiSession` to use this helper.
- [ ] Task: Final Logic Verification
    - [ ] Run full E2E suite (`tests/e2e`) to ensure no regressions in dynamic labeling or status detection.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: UpdateGeminiSession Refactoring' (Protocol in workflow.md)

## Phase 3: Final Verification & Cleanup

- [ ] Task: Code style and documentation audit
    - [ ] Ensure all new standalone functions are documented.
    - [ ] Run `go fmt` and `go vet` on modified files.
- [ ] Task: Final E2E Verification
    - [ ] Run all project tests to ensure system-wide stability.
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Final Verification & Cleanup' (Protocol in workflow.md)
