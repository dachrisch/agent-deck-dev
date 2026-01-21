# Initial Concept
Agent Deck is an AI agent command center designed to provide a unified, TUI-based dashboard for managing multiple AI agent sessions (like Claude Code and Gemini CLI) simultaneously. It simplifies the orchestration of complex, multi-agent workflows by providing a visual interface, shared MCP resources, and seamless session management.

# Product Definition

## Target Audience
- **AI Power Users:** Developers and researchers who frequently use AI CLI tools and need to manage multiple long-running sessions.
- **Agent Developers:** Engineers building and testing Model Context Protocol (MCP) servers and multi-agent systems.
- **Workflow Automators:** Professionals using AI agents to automate complex system tasks across various environments.

## Core Value Proposition
Agent Deck bridges the gap between raw CLI agent tools and the need for a cohesive workspace. It provides:
- **Visibility:** Real-time status of all active agent sessions in a single TUI view.
- **Efficiency:** Shared MCP pooling to reduce overhead and simplify configuration across agents.
- **Organization:** Grouping and persistence of agent sessions for long-term project management.
- **Control:** Centralized management of agent lifecycles, from initialization to forking and termination.

## Key Features (MVP)
- **Unified TUI Dashboard:** A terminal-based UI built with Bubble Tea for managing multiple agent sessions.
- **Session Management:** Create, fork, group, and delete agent sessions (Claude, Gemini).
- **MCP Resource Pooling:** Centralized management of Model Context Protocol servers to share tools and context across different agents.
- **Tmux Integration:** Utilization of tmux for persistent session management and visual tiling.
- **Global Search:** Find and navigate through session history and outputs.
- **Status Monitoring:** Visual indicators for agent activity, sleep states, and task progress.
- **Automated CI Pipeline:** Management-level GitHub Actions workflow in the Dev Project for continuous testing and quality reporting of the ecosystem.
- **E2E Test Framework:** Robust testing framework using `creack/pty` for automated TUI verification and startup checks.
- **Advanced E2E Automation:** Automated verification of complex TUI interactions, including session state persistence and dynamic model management.
- **Gemini Model Management:** Real-time visibility and control over Gemini model selection directly from the TUI, with dynamic switching and backend model change detection.

## Success Metrics
- **Context Switching Reduction:** Decrease in time spent manually managing multiple terminal windows for different agents.
- **Resource Efficiency:** Reduced CPU/Memory footprint through intelligent MCP pooling.
- **Session Persistence:** High reliability in recovering and continuing long-running agent tasks.
