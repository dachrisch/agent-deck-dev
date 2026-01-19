# Product Guidelines

## Visual Identity & TUI Aesthetic
- **Modern Terminal Styling:** Use high-contrast, vibrant color palettes (e.g., Tokyo Night, Catppuccin) to differentiate session states and components.
- **Rich Feedback:** Employ Lipgloss for sophisticated border styles, padding, and subtle animations to make the terminal feel responsive and modern.
- **Information Density vs. Clarity:** Prioritize a clean layout that avoids clutter while providing deep visibility into session status and agent activity.

## User Experience (UX) Principles
- **Keyboard-First Navigation:** Ensure every action is accessible via intuitive hotkeys, mirroring the efficiency of tools like `tmux` or `vim`.
- **Non-Blocking Operations:** Long-running tasks (like agent initialization or complex searches) should never freeze the UI; use Bubble Tea commands for async updates.
- **Proactive Notifications:** Use unobtrusive status bars or toast-style notifications within the TUI to inform users of state changes (e.g., "Session Forked", "MCP Server Connected").

## Prose & Communication Style
- **Technical & Concise:** Tool outputs, help menus, and error messages should be direct, actionable, and geared toward power users.
- **Consistent Terminology:** Strictly use defined terms like "Session", "Group", "MCP Pool", and "Fork" throughout the interface and documentation.

## Design Constraints
- **Responsibility:** The TUI must adapt gracefully to various terminal sizes, using adaptive layouts to hide or collapse panels when space is limited.
- **No Blocking Dialogs:** Favor non-intrusive status updates over modal alerts that interrupt the user's flow.
