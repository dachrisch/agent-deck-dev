# Specification: Update Hardcoded Gemini Model List

## Overview
This track updates the hardcoded fallback list of Gemini models in the application to ensure that the latest preview and v2.5 models are selectable in the TUI when the dynamic API call is unavailable or disabled.

## Functional Requirements
1.  **Update Fallback Model List**:
    -   Modify the `geminiModelFallback` variable in `internal/session/gemini.go`.
    -   Replace the existing list with the following five models:
        1.  `gemini-3-pro-preview`
        2.  `gemini-3-flash-preview`
        3.  `gemini-2.5-pro`
        4.  `gemini-2.5-flash`
        5.  `gemini-2.5-flash-lite`

## Non-Functional Requirements
-   **Consistency**: Ensure the model names match the expected format for the Gemini CLI and API.
-   **Testability**: Verify that the updated list is correctly returned by `GetAvailableGeminiModels()` when no API key is provided and no cache exists.

## Acceptance Criteria
-   The "Gemini Model" selection dialog in the TUI correctly displays the new five models.
-   Selecting any of these models correctly updates the session configuration.
-   `go test ./internal/session/...` passes.

## Out of Scope
-   Updating pricing metadata for the new models.
-   Changes to the dynamic API discovery logic.
