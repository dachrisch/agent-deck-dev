# Technology Stack

## Core Language & Runtime
- **Go 1.24+:** High-performance, statically typed language ideal for system tools and TUI applications.

## TUI Framework (Charmbracelet)
- **Bubble Tea:** The Elm Architecture for Go; used for the core UI loop and state management.
- **Lipgloss:** Used for sophisticated TUI styling and layout.
- **Bubbles:** Pre-built TUI components (inputs, lists, viewports) for rapid UI development.

## Infrastructure & Platform Integration
- **tmux:** Backend for persistent session management, allowing agents to run in the background.
- **pty (creack/pty):** Go interface for pseudo-terminal allocation and control.
- **fsnotify:** Cross-platform file system notifications for monitoring session storage and config changes.
- **Gemini API:** Used for dynamic discovery of available AI models.

## Data & Configuration
- **TOML (BurntSushi/toml):** Primary format for user configuration and agent profiles.
- **JSON:** Used for internal state persistence and communication with AI agents.
- **SQLite (planned/test-mode):** For structured storage of session metadata and history.

## Development & Testing
- **Testify:** Comprehensive testing toolkit for unit and integration tests.
- **creack/pty:** Pseudo-terminal interface for TUI verification and E2E testing.
- **Make:** Used for build orchestration and development workflows.
