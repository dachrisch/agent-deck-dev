# Plan: Gemini and Agent Deck Maintenance & Cleanup

This plan outlines the implementation of an automated maintenance system to prune redundant log files and manage bloated Gemini session data.

## Phase 1: Core Maintenance Logic

- [ ] Task: TDD - Implement Gemini log pruning
    - [ ] Create `internal/session/maintenance_test.go` with tests for pruning `.txt` files in project directories.
    - [ ] Implement `pruneGeminiLogs` in `internal/session/maintenance.go`.
    - [ ] Verify that files in `chats/` are preserved while other `.txt` files are deleted.
- [ ] Task: TDD - Implement backup and temp file cleanup
    - [ ] Add tests for keeping only 3 most recent `sessions.json.bak.*` files.
    - [ ] Implement `cleanupDeckBackups` and `cleanupProjectTempFiles` in `internal/session/maintenance.go`.
- [ ] Task: TDD - Implement automatic session archiving
    - [ ] Add tests for moving `.json` files > 30MB to an `archive/` subdirectory.
    - [ ] Implement `archiveBloatedSessions` in `internal/session/maintenance.go`.
- [ ] Task: Conductor - User Manual Verification 'Phase 1: Core Maintenance Logic' (Protocol in workflow.md)

## Phase 2: Orchestration & Scheduling

- [ ] Task: TDD - Implement background maintenance scheduler
    - [ ] Create tests for the maintenance interval logic (15 minutes).
    - [ ] Implement `StartMaintenanceWorker` in `internal/session/maintenance.go` using a separate goroutine.
- [ ] Task: Integrate with application lifecycle
    - [ ] Update `cmd/agent-deck/main.go` or `internal/ui/home.go` to trigger maintenance on startup.
    - [ ] Ensure maintenance runs at the specified 15-minute interval during active use.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: Orchestration & Scheduling' (Protocol in workflow.md)

## Phase 3: Final Verification

- [ ] Task: Final E2E Verification
    - [ ] Run full E2E suite (`tests/e2e`) to ensure no regressions in session discovery or status detection.
    - [ ] Verify disk usage and file counts after a full maintenance run.
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Final Verification' (Protocol in workflow.md)
