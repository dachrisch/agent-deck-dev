# Plan: Backend Performance Optimizations (Phase 1)

This plan outlines the implementation of high-priority backend performance optimizations, specifically tmux subprocess caching and LogWatcher rate limiting.

## Phase 1: tmux Subprocess Caching [checkpoint: a7c6b98]

- [x] Task: TDD - Implement Cache Infrastructure in tmux.Session (60d63cb)
    - [x] Add `cacheContent`, `cacheTime`, and `cacheMu` fields to the `Session` struct in `internal/tmux/tmux.go`.
    - [x] Create unit tests in `internal/tmux/tmux_test.go` to verify cache set/get/expiry logic.
- [x] Task: TDD - Integrate Cache into CapturePane (60d63cb)
    - [x] Modify `CapturePane()` to check the cache and respect the 500ms TTL.
    - [x] Write unit tests to verify that `CapturePane()` only calls the system command when the cache is stale or empty.
- [x] Task: TDD - Implement Cache Invalidation (60d63cb)
    - [x] Update `SendKeys()` and `Start()` to invalidate the cache.
    - [x] Write unit tests to verify immediate invalidation after interactive actions.
- [ ] Task: Conductor - User Manual Verification 'Phase 1: tmux Subprocess Caching' (Protocol in workflow.md)

## Phase 2: LogWatcher Rate Limiting

- [x] Task: TDD - Implement Token Bucket Rate Limiter (a20c92b)
    - [x] Create a standalone `RateLimiter` utility or integrate directly into the monitoring loop.
    - [x] Write unit tests to verify the 20 events/second limit and event coalescing.
- [x] Task: TDD - Integrate Rate Limiter into LogWatcher (e6a129f)
    - [x] Modify the log monitoring loop to use the rate limiter before triggering UI updates.
    - [x] Write unit tests simulating high-frequency log writes to verify backend shielding.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: LogWatcher Rate Limiting' (Protocol in workflow.md)

## Phase 3: Final Verification & Quality Gates

- [ ] Task: Performance Benchmarking
    - [ ] Measure CPU usage before and after optimizations with 5+ active sessions.
- [ ] Task: Full Project Verification
    - [ ] Run all unit tests (`go test ./...`) and E2E tests (`cd tests/e2e && go test -v .`).
    - [ ] Run linter and static analysis (`make lint` if available, otherwise `go vet ./...`).
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Final Verification & Quality Gates' (Protocol in workflow.md)
