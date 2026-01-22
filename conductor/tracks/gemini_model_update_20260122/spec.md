# Track Specification: Gemini Model Detection and Hardcoded List Update

## Overview
This track focuses on two key improvements for Gemini integration: updating the hardcoded model list to include the latest models and implementing a mechanism to detect when the active model changes from within a running Gemini session.

## Functional Requirements
- **Model List Update:**
    - The hardcoded fallback list in `internal/session/gemini.go` must be updated to include:
        - `gemini-3-pro-preview`
        - `gemini-3-flash-preview`
        - `gemini-2.5-pro`
        - `gemini-2.5-flash`
        - `gemini-2.5-flash-lite`
- **Dynamic Model Detection:**
    - The system must detect when the active model changes from within a running Gemini session.
    - Upon detection, the TUI label for the session must update immediately to reflect the new model name.

## Non-Functional Requirements
- **Accuracy:** The model detection must accurately identify the new model name.
- **Performance:** The detection mechanism should not introduce any noticeable lag to the TUI.

## Acceptance Criteria
- [ ] The model selection dialog in the TUI correctly lists the newly added Gemini models.
- [ ] If the model changes during a session (e.g., due to an internal Gemini CLI command), the model name in the TUI session list updates automatically to the new model.
- [ ] The model label update happens in real-time without requiring a manual refresh or restart of the application.

## Out of Scope
- Implementing dynamic model discovery from the Gemini API (this change relies on the hardcoded list).
- Displaying toast notifications or other UI elements beyond the label update.
