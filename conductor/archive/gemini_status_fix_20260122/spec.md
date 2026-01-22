# Specification: Fix Gemini Status Detection

## Overview
Gemini sessions in Agent Deck consistently show as "Idle" (Gray) or "Waiting" (Yellow) even when they are actively generating text. This track aims to fix the status detection logic to correctly identify when a Gemini session is in a "Running" (Green) state.

## Functional Requirements
1. **Gemini Busy Detection:** Update the `hasBusyIndicator` logic to reliably detect Gemini's activity.
    - Expand the spinner detection range (currently limited to the last 5 lines).
    - Ensure it handles Gemini's specific spinner and streaming output format.
2. **Status Transition:** Ensure that when activity is detected, the session transitions to "Running" (Green) and stays there until the generation is complete.
3. **Prompt Awareness:** Use Gemini's prompt (`gemini>`) to help determine when a session has returned to an idle state.

## Non-Functional Requirements
- **Performance:** Maintain low overhead for the status worker (which polls tmux frequently).
- **Stability:** Avoid status "flickering" (rapidly switching between Green and Gray).

## Acceptance Criteria
- A Gemini session must show "Running" (Green) in the TUI while it is actively generating a response.
- The session must return to "Waiting" (Yellow) or "Idle" (Gray) once the response is finished.
- E2E tests must pass, specifically verifying the status transitions for Gemini.

## Out of Scope
- Adding forking support for Gemini (not supported by the CLI).
- Changes to Claude or OpenCode status detection (unless required for common logic refactoring).
