# Implementation Plan: Gemini Model Enhancements

## Phase 1: Model List Expansion & "Auto" Support
- [ ] Task: Update hardcoded model list and implement "auto" logic
    - [ ] Add new models to `internal/session/gemini.go`
    - [ ] Add "auto" to the model list
    - [ ] Update `internal/session/instance.go` to omit `--model` when "auto" is selected
- [ ] Task: Update cost calculation for new models
    - [ ] Add pricing for new models in `internal/session/gemini_analytics.go`

## Phase 2: Dynamic Model Detection
- [ ] Task: Implement backend model detection
    - [ ] Add logic to parse the actual running model from Gemini session output or environment
    - [ ] Update `internal/session/instance.go` to periodically refresh the active model from the running process
- [ ] Task: Reflect detected model in UI
    - [ ] Ensure `GeminiModel` field in `Instance` is updated when detection occurs
    - [ ] Verify UI tag updates automatically

## Phase 3: Verification & Testing
- [ ] Task: Update E2E tests for "auto" and model detection
    - [ ] Add test case for "auto" model selection
    - [ ] Add test case for external model change detection
- [ ] Task: Conductor - User Manual Verification 'Phase 3'
