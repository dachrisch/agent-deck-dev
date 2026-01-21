# Specification: Gemini Model Enhancements

## Overview
Enhance the Gemini session management to support new models, provide an "auto" selection option, and detect model changes that occur within the Gemini backend.

## Goal
To keep the Agent Deck TUI synchronized with the latest Gemini models and backend state.

## Functional Requirements
- **Expanded Model List:**
    - Add the following models to the discovery/selection list:
        - `gemini-3-pro-preview`
        - `gemini-3-flash-preview`
        - `gemini-2.5-pro`
        - `gemini-2.5-flash`
        - `gemini-2.5-flash-lite`
- **"Auto" Model Selection:**
    - Add an "auto" option to the model selection dialog.
    - If "auto" is selected, the `--model` flag should NOT be passed when starting/restarting the Gemini session.
- **Model Change Detection:**
    - Implement a mechanism to detect if the model has changed within the Gemini session (e.g., via backend feedback or tmux inspection) and update the TUI accordingly.

## Acceptance Criteria
- [ ] New Gemini models are available for selection in the Model Manager.
- [ ] Selecting "auto" results in the Gemini session running without a specific `--model` flag.
- [ ] If the model changes in the backend, the UI model tag (`gemini(model)`) updates to reflect the new state.
- [ ] E2E tests verify the "auto" selection and model detection logic.
