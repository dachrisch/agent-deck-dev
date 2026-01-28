# Specification: Project Path Autocomplete in New Session Dialog

## Overview
This track adds standard shell-like path autocomplete to the "Project Path" input field in the "New Session" dialog. This will allow users to quickly specify project directories by typing a partial path and using the `Tab` key to cycle through matching directories.

## Functional Requirements
1.  **Tab-Triggered Autocomplete**:
    -   Pressing the `Tab` key in the "Project Path" field of the "New Session" dialog must trigger directory autocomplete.
    -   The autocomplete must look for directories in the local filesystem relative to the current input string.
    -   Both absolute paths (starting with `/` or `~`) and relative paths must be supported.

2.  **Match Cycling**:
    -   If a single match is found, it is immediately completed in the input field.
    -   If multiple matches are found, subsequent `Tab` presses must cycle through all available matches in alphabetical order.
    -   Non-directory files must be excluded from the matches.

3.  **Visual Feedback**:
    -   The input field must update its value to show the currently selected match.
    -   If no matches are found, the `Tab` key should have no effect (or provide a subtle visual/auditory hint if supported by the TUI).

## Non-Functional Requirements
-   **Responsiveness**: The filesystem scan for matches must be efficient and not block the UI thread significantly, even in directories with many subfolders.
-   **Cross-Platform**: The path resolution logic must work correctly on Linux and macOS (Windows support is secondary but recommended where possible using Go's `filepath` package).

## Acceptance Criteria
-   Typing `~/dev/pla` and pressing `Tab` correctly completes to `~/dev/playground/` (assuming it exists).
-   If `~/dev/playground/` contains `agent-deck` and `energy-consumption`, pressing `Tab` multiple times cycles through these two directories.
-   Relative paths like `./inte` correctly autocomplete to matching local directories.
-   The "New Session" dialog remains functional and doesn't crash during rapid autocomplete attempts.

## Out of Scope
-   Autocomplete for fields other than "Project Path".
-   Interactive dropdown menus or list-based selection.
-   File autocomplete (limited strictly to directories).
