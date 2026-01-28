# Plan: Project Path Autocomplete in New Session Dialog

This plan outlines the implementation of shell-like tab-completion for the Project Path field in the New Session dialog.

## Phase 1: Path Autocomplete Logic (TDD)

- [ ] Task: TDD - Implement Directory Listing Utility
    - [ ] Create a utility function `GetDirectoryCompletions(input string) ([]string, error)` in `internal/session/utils.go`.
    - [ ] Implement logic to handle absolute, relative, and tilde-prefixed (`~`) paths.
    - [ ] Write unit tests in `internal/session/utils_test.go` to verify correct filtering of non-directory files and path resolution.
- [ ] Task: TDD - Implement Completion Cycler
    - [ ] Add a `CompletionCycler` struct to manage the state of active completion (matches, current index).
    - [ ] Write unit tests to verify cycling logic (wrapping around, reset on input change).
- [ ] Task: Conductor - User Manual Verification 'Phase 1: Path Autocomplete Logic (TDD)' (Protocol in workflow.md)

## Phase 2: TUI Integration

- [ ] Task: Integrate Cycler into NewSessionDialog
    - [ ] Update `internal/ui/new_dialog.go` to intercept `Tab` keys in the project path field.
    - [ ] Hook the `CompletionCycler` into the input update loop.
    - [ ] Ensure that typing any other key resets the completion state.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: TUI Integration' (Protocol in workflow.md)

## Phase 3: Final Verification & Quality Gates

- [ ] Task: Full Project Verification
    - [ ] Run all unit tests (`go test ./...`) and E2E tests (`cd tests/e2e && go test -v .`).
    - [ ] Run linter and static analysis (`go vet ./...`).
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Final Verification & Quality Gates' (Protocol in workflow.md)
