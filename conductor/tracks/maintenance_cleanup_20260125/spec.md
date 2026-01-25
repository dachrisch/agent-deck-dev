# Spec: Gemini and Agent Deck Maintenance & Cleanup

## Overview
Implement an automated maintenance system to prune redundant log files and manage bloated session data. This will reduce disk I/O, lower the total file count (currently >10k), and improve TUI responsiveness.

## Functional Requirements
1.  **Scheduled Maintenance:**
    - Perform cleanup on application startup.
    - Run background maintenance every 15 minutes during active sessions.
    - Ensure maintenance tasks have low CPU/IO priority to avoid UI stutter.
2.  **Redundant File Pruning:**
    - **Gemini Logs:** Delete all `.txt` files in `~/.gemini/tmp/<hash>/` directories that are NOT within the `chats/` subdirectory.
    - **Deck Backups:** Remove old backup files (`sessions.json.bak.*`) in Agent Deck profiles, keeping only the 3 most recent backups.
    - **Temp Files:** Clean up any stale files in the Agent Deck project's temporary directory.
3.  **Automatic Session Archiving:**
    - Identify Gemini JSON session files exceeding 30MB.
    - Automatically move these files to `~/.gemini/tmp/<hash>/archive/` to remove them from the active scanning pool.
4.  **Status Reporting:**
    - Log the number of files cleaned and space reclaimed to the application log for transparency.

## Non-Functional Requirements
- **Branch Management:** **All work must be implemented in a new, isolated branch (e.g., `feature/maintenance-cleanup`) and NOT merged into current feature development branches until verified.**
- **Performance:** Maintenance must run in a separate goroutine.
- **Safety:** Verify file paths and extensions strictly before deletion to prevent data loss.
- **Scalability:** Handle directories with 10,000+ files efficiently using buffered directory reading.

## Acceptance Criteria
- File count in `~/.gemini/tmp` is significantly reduced after cleanup.
- Gemini sessions > 30MB are moved to an `archive` folder and no longer appear in the primary TUI list.
- Application startup and background polling remain smooth.
- Maintenance runs automatically at the specified 15-minute interval.

## Out of Scope
- Archiving Claude Code or OpenCode sessions.
- User-facing UI for manual file selection during cleanup.
