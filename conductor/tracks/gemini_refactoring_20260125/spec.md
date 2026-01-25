# Spec: Gemini Refactoring and Code Cleanup

## Overview
This track addresses code quality findings from the Gemini UI refinement review. It focuses on deduplicating file-system helpers and refactoring the complex `UpdateGeminiSession` logic into smaller, testable components.

## Functional Requirements
1.  **Utility Consolidation:**
    - Create `internal/session/utils.go`.
    - Move the `findNewestFile(pattern string) string` helper to this file.
    - Remove all duplicate implementations of `findNewestFile` from `internal/session/gemini.go`.
2.  **Logic Refactoring:**
    - Refactor `(i *Instance) UpdateGeminiSession(...)` in `internal/session/instance.go`.
    - Extract logic into standalone private functions:
        - `syncGeminiSessionFromTmux(i *Instance)`: Handles ID and YOLO mode discovery from environment.
        - `syncGeminiSessionFromFile(i *Instance)`: Handles fallback file scanning.
        - `updateGeminiAnalytics(i *Instance)`: Manages token usage and cost updates.
        - `detectGeminiModelRealTime(i *Instance)`: Handles regex-based detection from tmux output.
        - `updateGeminiLatestPrompt(i *Instance)`: Manages the cached user prompt from session files.
3.  **Behavior Preservation:**
    - Ensure no changes to existing functionality, especially the `auto (detected-model)` UI logic and the `mtime` caching performance optimizations.
4.  **Test Coverage:**
    - **Ensure all newly created functions have basic unit test coverage.**

## Non-Functional Requirements
- **Maintainability:** Reduce the cyclomatic complexity of `UpdateGeminiSession`.
- **Testability:** Standalone functions should be easier to target with unit tests.

## Acceptance Criteria
- Code compiles without warnings.
- `findNewestFile` exists only in `utils.go`.
- `UpdateGeminiSession` is simplified to a series of helper calls.
- **Unit tests are provided for each of the new standalone functions.**
- All E2E tests (`tests/e2e`) pass, confirming no regression in session discovery or status detection.

## Out of Scope
- Changing underlying logic for session ID capture.
- Modifying the UI/Lipgloss rendering code.
