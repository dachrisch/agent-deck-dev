# Agent Deck - Dev Project

This repository serves as the **Conductor Management Hub** for the [Agent Deck](https://github.com/asheshgoplani/agent-deck) ecosystem. It uses the Conductor methodology to manage project requirements, architecture, workflow, and implementation tracks.

## Project Structure

- **Dev Project (Root):** Orchestrates project management via the `conductor/` directory.
- **External Project (`agent-deck/`):** The functional codebase for Agent Deck (linked as a git submodule).

## Conductor Context

The project is organized around the following core artifacts in the `conductor/` directory:

- **[Index](./conductor/index.md):** The entry point for all project context.
- **[Product Definition](./conductor/product.md):** Vision, features, and target audience.
- **[Workflow](./conductor/workflow.md):** Rules for branching, TDD, and dual-repo synchronization.
- **[Tracks](./conductor/tracks.md):** The high-level roadmap of work units.

## Development Workflow

This project follows a strict dual-repository synchronization protocol:

1.  **Organizational Changes:** All updates to plans, specs, and management files are committed directly to this repository.
2.  **Feature Development:** 
    - Always branch from the `main` branch of the `agent-deck` submodule.
    - Implement features, tests, and code fixes within the submodule.
    - Follow the `agent-deck/CONTRIBUTING.md` guidelines.
3.  **Synchronization:** This repository's submodule pointer is updated whenever functional changes are completed in the external project.

## Getting Started

To explore the current implementation state, refer to the [Tracks Registry](./conductor/tracks.md).
