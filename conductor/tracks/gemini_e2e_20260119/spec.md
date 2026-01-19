# Specification: E2E Test for Gemini Model Switching

## Overview
Automate the end-to-end verification of the Gemini model switching feature using the established `creack/pty` framework. This ensures that the TUI interaction correctly triggers a session restart with the updated model parameters.

## Goal
To programmatically verify that selecting a new Gemini model via the TUI accurately restarts the agent session with the corresponding `--model` flag.

## Functional Requirements
- **Test Automation:** 
    - Create a new E2E test file `tests/e2e/model_switch_test.go`.
    - Simulate creating a Gemini session.
    - Send `Ctrl+G` to open the Model Manager.
    - Simulate navigation and selection of a non-default model.
- **Mocking:**
    - Use a mock Gemini API environment (simulated via environment variables or local overrides) to ensure predictable model lists.
- **Verification:**
    - Use `tmux list-panes` or similar command-level inspection to verify that the Gemini process was restarted with the expected `--model` flag.
    - Verify that the TUI continues to render correctly after the restart.

## Acceptance Criteria
- [ ] `tests/e2e/model_switch_test.go` exists and passes consistently.
- [ ] Successful verification of the process command-line change after a model switch.
