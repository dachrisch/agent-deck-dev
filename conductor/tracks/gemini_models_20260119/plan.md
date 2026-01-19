# Implementation Plan: Gemini Model Management

## Phase 1: Submodule Scaffolding & API Integration
- [x] Task: Prepare feature branch in agent-deck submodule (3a53115)
    - [x] Checkout main in agent-deck
    - [x] Create branch `feature/gemini-model-management`
- [ ] Task: Implement Gemini Model Discovery
    - [ ] Add logic to fetch available models from the Gemini API within `internal/session/gemini.go`
    - [ ] Ensure API calls are cached or handled efficiently to prevent UI lag
- [ ] Task: Conductor - User Manual Verification 'Phase 1' (Protocol in workflow.md)

## Phase 2: UI Updates (Session List)
- [ ] Task: Update Session Item Rendering
    - [ ] Modify `internal/ui/home.go` to extract and display the Gemini model name in the session list
    - [ ] Apply appropriate Lipgloss styling to the model tag
- [ ] Task: Conductor - User Manual Verification 'Phase 2' (Protocol in workflow.md)

## Phase 3: Interactive Management
- [ ] Task: Update Action Menu
    - [ ] Add "Change Model" option to the menu in `internal/ui/home.go`
    - [ ] Implement a selection dialog (fuzzy list) to pick from discovered models
- [ ] Task: Implement Dynamic Model Switching
    - [ ] Add logic to communicate the model change to the running Gemini session
    - [ ] Verify the UI updates immediately after a successful switch
- [ ] Task: Conductor - User Manual Verification 'Phase 3' (Protocol in workflow.md)
