# Specification: Backend Performance Optimizations (Phase 1)

## Overview
This track focuses on the first set of high-priority backend performance optimizations identified during the critical fix validation. We will implement smart caching for high-latency tmux calls and rate limiting for log file monitoring to reduce CPU usage and improve TUI responsiveness.

## Functional Requirements
1.  **tmux Subprocess Caching**:
    -   Implement a thread-safe, TTL-based cache for tmux `capture-pane` output in the `tmux.Session` struct.
    -   Use a default TTL of **500ms** to align with the application's polling frequency.
    -   Ensure all calls to `CapturePane()` check this cache before executing a system command.
    -   Invalidate the cache immediately whenever `SendKeys()` or `Start()` is called on the session.

2.  **LogWatcher Rate Limiting**:
    -   Implement a token bucket or similar rate-limiting mechanism in the `LogWatcher` (or equivalent file monitoring component).
    -   Set the processing limit to **20 events/second** per session to shield the backend from high-frequency log bursts.
    -   Coalesce multiple rapid-fire events into a single update to maintain UI stability.

## Non-Functional Requirements
-   **Concurrency Safety**: All caching mechanisms must be fully thread-safe using `sync.Mutex` or `sync.RWMutex`.
-   **Minimal Latency**: The caching layer itself must add negligible overhead (sub-millisecond).
-   **Testability**: Both caching and rate-limiting logic must be verifiable via unit tests with mock time or manual clock advancement.

## Acceptance Criteria
-   `go test ./internal/tmux/...` and relevant UI tests pass.
-   CPU usage remains stable even when monitoring highly active sessions (e.g., sessions producing 100+ log lines per second).
-   The TUI reflects model changes and log updates within the expected 500ms window.
-   Manual verification confirms that terminal previews still update correctly and are not "stuck" on old content beyond the TTL.

## Out of Scope
-   Prometheus metrics or deep runtime observability.
-   Persistent File I/O caching (deferred to a later track).
-   LRU cache eviction for analytics data.
