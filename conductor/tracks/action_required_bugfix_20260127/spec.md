# Specification: Fix Claude "Action Required" Status Detection

## Overview
This track addresses a bug where Claude Code sessions waiting for user confirmation (e.g., "Action Required" dialogs) are incorrectly identified as "running" instead of "waiting."

## Functional Requirements
1.  **Enhance Claude Prompt Detection**:
    -   Update the `hasClaudePrompt` logic in `internal/tmux/detector.go`.
    -   Add the following patterns to the `permissionPrompts` list:
        -   `"Action Required"`
        -   `"Waiting for user confirmation"`
        -   `"Allow execution of"`
        -   `"Allow once"`
    -   Ensure these patterns correctly trigger the `StateWaiting` status.

2.  **Verify Robustness**:
    -   Ensure the detector continues to correctly identify "busy" states (e.g., when a spinner is present) even if these strings appear in the history. The current logic already checks busy indicators first, which should be maintained.

## Non-Functional Requirements
-   **Test Coverage**: Add new unit tests to `internal/tmux/tmux_test.go` covering the specific "Action Required" output format.
-   **Performance**: The additions must not significantly impact the performance of the O(n) ANSI stripping or pattern matching logic.

## Acceptance Criteria
-   A Claude session showing an "Action Required" dialog is correctly displayed as **waiting** (◐) in the TUI.
-   `go test ./internal/tmux/...` passes with the new test cases.
-   Manual verification confirms that once the user responds to the prompt, the status transitions correctly (e.g., back to **running** or **idle**).

## Out of Scope
-   Refactoring the entire status state machine.
-   Changes to non-Claude detectors (unless regressions are found).
