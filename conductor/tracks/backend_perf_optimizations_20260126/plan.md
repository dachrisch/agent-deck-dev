# Plan: Backend Performance Optimizations (Phase 1)

This plan outlines the implementation of high-priority backend performance optimizations, specifically tmux subprocess caching and LogWatcher rate limiting.

## Phase 1: tmux Subprocess Caching

- [ ] Task: TDD - Implement Cache Infrastructure in tmux.Session
    - [ ] Add `cacheContent`, `cacheTime`, and `cacheMu` fields to the `Session` struct in `internal/tmux/tmux.go`.
    - [ ] Create unit tests in `internal/tmux/tmux_test.go` to verify cache set/get/expiry logic.
- [ ] Task: TDD - Integrate Cache into CapturePane
    - [ ] Modify `CapturePane()` to check the cache and respect the 500ms TTL.
    - [ ] Write unit tests to verify that `CapturePane()` only calls the system command when the cache is stale or empty.
- [ ] Task: TDD - Implement Cache Invalidation
    - [ ] Update `SendKeys()` and `Start()` to invalidate the cache.
    - [ ] Write unit tests to verify immediate invalidation after interactive actions.
- [ ] Task: Conductor - User Manual Verification 'Phase 1: tmux Subprocess Caching' (Protocol in workflow.md)

## Phase 2: LogWatcher Rate Limiting

- [ ] Task: TDD - Implement Token Bucket Rate Limiter
    - [ ] Create a standalone `RateLimiter` utility or integrate directly into the monitoring loop.
    - [ ] Write unit tests to verify the 20 events/second limit and event coalescing.
- [ ] Task: TDD - Integrate Rate Limiter into LogWatcher
    - [ ] Modify the log monitoring loop to use the rate limiter before triggering UI updates.
    - [ ] Write unit tests simulating high-frequency log writes to verify backend shielding.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: LogWatcher Rate Limiting' (Protocol in workflow.md)

## Phase 3: Final Verification & Quality Gates

- [ ] Task: Performance Benchmarking
    - [ ] Measure CPU usage before and after optimizations with 5+ active sessions.
- [ ] Task: Full Project Verification
    - [ ] Run all unit tests (`go test ./...`) and E2E tests (`cd tests/e2e && go test -v .`).
    - [ ] Run linter and static analysis (`make lint` if available, otherwise `go vet ./...`).
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Final Verification & Quality Gates' (Protocol in workflow.md)
