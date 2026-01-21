# Implementation Plan: Gemini Model Management

## Phase 1: Submodule Scaffolding & API Integration [checkpoint: 8737891]
- [x] Task: Prepare feature branch in agent-deck submodule (3a53115)
    - [x] Checkout main in agent-deck
    - [x] Create branch `feature/gemini-model-management`
- [x] Task: Implement Gemini Model Discovery (f72ec90)
    - [x] Add logic to fetch available models from the Gemini API within `internal/session/gemini.go`
    - [x] Ensure API calls are cached or handled efficiently to prevent UI lag
- [x] Task: Conductor - User Manual Verification 'Phase 1' (Protocol in workflow.md)

## Phase 2: UI Updates (Session List) [checkpoint: 7154894]
- [x] Task: Update Session Item Rendering (f6efea6)
    - [x] Modify `internal/ui/home.go` to extract and display the Gemini model name in the session list
    - [x] Apply appropriate Lipgloss styling to the model tag
- [x] Task: Conductor - User Manual Verification 'Phase 2'

## Phase 3: Interactive Management [checkpoint: 03f068a]
- [x] Task: Update Action Menu (f72ec90)
    - [x] Add "Change Model" option to the menu in `internal/ui/home.go`
    - [x] Implement a selection dialog (fuzzy list) to pick from discovered models
- [x] Task: Implement Dynamic Model Switching (f72ec90)
    - [x] Add logic to communicate the model change to the running Gemini session
    - [x] Verify the UI updates immediately after a successful switch
- [x] Task: Conductor - User Manual Verification 'Phase 3'

## Phase 4: Upstream Synchronization & Maintenance
- [x] Task: Sync submodule with upstream
    - [x] Fetch from upstream in `agent-deck`
    - [x] Merge `upstream/main` into local `main`
    - [x] Merge `main` into `feature/gemini-model-management`
- [x] Task: Verify feature still running
    - [x] Run E2E tests in `tests/e2e`
- [x] Task: Conductor - User Manual Verification 'Phase 4'
