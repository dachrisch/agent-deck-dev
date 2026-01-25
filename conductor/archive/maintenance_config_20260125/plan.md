# Plan: Maintenance Configuration and Feedback

This plan outlines the implementation of user-controlled maintenance settings and visual feedback within the Agent Deck TUI.

## Phase 1: Configuration & Backend Control [checkpoint: 9e6d9c9]

- [x] Task: TDD - Implement Maintenance Configuration (4f4c275)
    - [ ] Add `MaintenanceSettings` struct to `internal/session/userconfig.go`.
    - [ ] Update `UserConfig` to include the `[maintenance]` section.
    - [ ] Add `GetMaintenanceSettings()` helper.
    - [ ] Write unit tests verifying default value is `false` and TOML parsing works.
- [x] Task: TDD - Integrate Settings with Worker (f109a85)
    - [ ] Modify `StartMaintenanceWorker` to poll/check the `enabled` setting before each run.
    - [ ] Update `Maintenance()` function to return a structured result for UI consumption.
    - [ ] Write unit tests verifying the worker stays idle when disabled.
- [ ] Task: Conductor - User Manual Verification 'Phase 1: Configuration & Backend Control' (Protocol in workflow.md)

## Phase 2: TUI Settings & Feedback [checkpoint: 97720e5]

- [x] Task: TDD - Add Settings UI Toggle (3f6997e)
    - [ ] Add "Maintenance" toggle to `internal/ui/settings_panel.go`.
    - [ ] Ensure toggling updates the user configuration and persists to disk.
    - [ ] Write unit tests for the settings panel state change.
- [x] Task: TDD - Implement Status Message Logic (7e3ad26)
    - [ ] Add `maintenanceMsg` state to `internal/ui/home.go`.
    - [ ] Implement a `clearMaintenanceMsg` timer/command.
    - [ ] Update `renderPreviewPane` (or bottom bar) to display the message when active.
    - [ ] Write unit tests for message appearance and auto-clearing.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: TUI Settings & Feedback' (Protocol in workflow.md)

## Phase 3: Final Integration & E2E [checkpoint: 21438c8]

- [x] Task: Final E2E Verification (8376713, 642d59a)
    - [ ] Create an E2E test in `tests/e2e` that enables maintenance, mocks a run, and verifies the UI message appears.
    - [ ] Verify that maintenance remains disabled by default on a fresh run.
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Final Integration & E2E' (Protocol in workflow.md)
