# Track Specification: Gemini Model Label Shortening and Dynamic Detection

## Overview
This track aims to improve the visual clarity of Gemini sessions in the TUI by shortening the model labels (removing the redundant `gemini-` prefix) and enhancing the dynamic detection of model changes during a running session by parsing the session's output stream.

## Functional Requirements
- **Label Shortening:**
    - In the main session list, the `gemini-` prefix must be stripped from the model name (e.g., `gemini-2.0-flash` becomes `2.0-flash`).
    - If the session model is set to `auto` and a specific model is detected, the label should display as `auto(model-name)` (e.g., `auto(2.0-flash)`).
- **Dynamic Model Detection:**
    - Implement real-time detection of model changes by parsing the session's standard output stream for model name patterns.
    - When a change is detected, the session's `GeminiModel` field must be updated immediately.
    - The TUI must reflect the updated model name in the session list without requiring a manual refresh or restart.

## Non-Functional Requirements
- **Performance:** Parsing the output stream should be efficient and non-blocking to avoid TUI lag.
- **Robustness:** Model detection patterns should be resilient to varied output formats from the Gemini CLI.

## Acceptance Criteria
- [ ] Gemini sessions in the TUI list display models without the `gemini-` prefix.
- [ ] Sessions started with `auto` model selection display as `auto(detected-model)` once detected.
- [ ] Manually changing the model within a running Gemini session (if supported by the CLI) or the CLI switching models internally is reflected in the TUI label.
- [ ] Existing functionality for model switching via the "Change Model" menu remains intact.

## Out of Scope
- Removing the `gemini-` prefix from other TUI locations (headers, dialogs).
- Implementing model switching logic that isn't already supported by the underlying Gemini CLI.
