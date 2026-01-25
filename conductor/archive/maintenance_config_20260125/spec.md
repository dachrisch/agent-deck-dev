# Specification: Maintenance Configuration and Feedback

## Overview
This track introduces user control and visibility for the automated maintenance system. Users will be able to enable/disable background maintenance via configuration or the UI, and will receive real-time feedback when a maintenance run completes.

## Functional Requirements
1.  **Configuration Settings**:
    -   Add a `[maintenance]` section to `config.toml`.
    -   Include an `enabled` boolean key, defaulting to `false`.
2.  **TUI Settings Integration**:
    -   Add a "Maintenance Enabled" toggle to the Settings panel (`Ctrl+S`).
    -   Toggling this setting should immediately update the background worker's state and persist to `config.toml`.
3.  **Visual Indicator**:
    -   Implement a temporary status line message in the bottom bar of the TUI.
    -   The message should appear when a maintenance run finishes (e.g., "Maintenance: Pruned 525 logs, Archived 2 sessions").
    -   The message should automatically clear after a short duration (e.g., 5-10 seconds).
4.  **Worker Control**:
    -   The `StartMaintenanceWorker` should respect the `enabled` setting.
    -   If disabled, the worker should either not start or remain idle without performing tasks.

## Non-Functional Requirements
-   **Low Overhead**: Maintenance runs and UI updates must not block the main TUI loop.
-   **Persistence**: Settings changes in the TUI must be saved to the profile's configuration.

## Acceptance Criteria
-   [ ] Maintenance does NOT run by default on a fresh installation.
-   [ ] Enabling maintenance in `config.toml` starts the worker on next launch.
-   [ ] Toggling maintenance in the Settings panel immediately starts/stops the background task logic.
-   [ ] A clear status message appears at the bottom of the TUI after a maintenance run.
-   [ ] The status message disappears after a reasonable timeout.

## Out of Scope
-   Manual triggering of maintenance from the UI (outside of the periodic worker).
-   Detailed maintenance history logs accessible via the UI.
