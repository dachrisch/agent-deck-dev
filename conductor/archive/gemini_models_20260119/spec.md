# Specification: Gemini Model Management

## Overview
Enhance the Agent Deck TUI to display and manage the active model for Gemini agent sessions. This allows users to see which model is currently driving an agent and switch models dynamically without leaving the deck.

## Goal
To provide real-time visibility and control over Gemini model selection directly from the Agent Deck interface.

## Functional Requirements
- **Model Display:**
    - Update the session list item rendering to include the active Gemini model name in the format: `agent-deck gemini(model)`.
    - Example: `▶└─ ○ agent-deck gemini(gemini-2.0-flash) [YOLO]`
- **Model Discovery:**
    - Implement dynamic discovery of available Gemini models using the Gemini API.
- **Model Management:**
    - Add a "Change Model" option to the existing session action menu.
    - Selecting "Change Model" should present a list of available Gemini models for selection.
- **Dynamic Application:**
    - When a new model is selected, attempt to apply the change immediately to the running session.

## Acceptance Criteria
- [ ] Gemini sessions in the list view display their active model name.
- [ ] The action menu for Gemini sessions includes a "Change Model" entry.
- [ ] Users can successfully select a new model from a dynamically generated list.
- [ ] The change is reflected in the UI and applied to the session.
