# Implementation Plan: Gemini Model Enhancements

## Phase 1: Model List Expansion & "Auto" Support [checkpoint: b4d5b63]
- [x] Task: Update hardcoded model list and implement "auto" logic (b4d5b63)
    - [x] Add new models to `internal/session/gemini.go`
    - [x] Add "auto" to the model list
    - [x] Update `internal/session/instance.go` to omit `--model` when "auto" is selected
- [x] Task: Update cost calculation for new models (b4d5b63)
    - [x] Add pricing for new models in `internal/session/gemini_analytics.go`

## Phase 2: Dynamic Model Detection [checkpoint: b4d5b63]
- [x] Task: Implement backend model detection (b4d5b63)
    - [x] Add logic to parse the actual running model from Gemini session output or environment
    - [x] Update `internal/session/instance.go` to periodically refresh the active model from the running process
- [x] Task: Reflect detected model in UI (b4d5b63)
    - [x] Ensure `GeminiModel` field in `Instance` is updated when detection occurs
    - [x] Verify UI tag updates automatically

## Phase 3: Verification & Testing [checkpoint: b47bb69]
- [x] Task: Update E2E tests for "auto" and model detection (b47bb69)
    - [x] Add test case for "auto" model selection
    - [x] Add test case for external model change detection
- [x] Task: Conductor - User Manual Verification 'Phase 3'
