# Spec: Gemini UI Refinement - Model Relocation

## Overview
Relocate the Gemini model display from the tree view (left panel) to a dedicated dropdown-style indicator in the session information section (right panel). This restores the visual simplicity of the session tree while centralizing model management visibility within the session details.

## Functional Requirements
1.  **Tree View Restoration:**
    - Update the session rendering logic to remove model-specific suffixes (e.g., `(auto)`, `(1.5-pro)`) from the tree view labels.
    - Gemini sessions should simply display `gemini` (plus any applicable YOLO badges).
2.  **Right Panel Model Indicator:**
    - Add a new "Model" field to the Gemini-specific information block in the right panel, positioned directly below the "Session ID".
    - Style this field as a "dropdown" (visual border/indicator) to signify it is a manageable setting.
3.  **Dynamic Auto-Labeling:**
    - If the model is set to `auto` and a specific model has been detected from the terminal output, the indicator should display `auto (detected-model)` (e.g., `auto (2.0-flash)`).
    - If no specific model is detected yet, it should simply show `auto`.
4.  **Interaction:**
    - The `ctrl+g` keyboard shortcut remains the primary method for opening the Model Manager dialog to change the selection.

## Non-Functional Requirements
- **Theme Consistency:** Maintain the Tokyo Night color palette and Lipgloss styling used in the existing panels.
- **Layout Stability:** Ensure the right panel uses fixed-width rendering (`ensureExactWidth`) to prevent text wrapping or layout "bleeding" when model names are long.

## Acceptance Criteria
- Gemini sessions in the left tree view show only the tool name `gemini`.
- Selecting a Gemini session displays a "Model" field in the right panel with the current selection.
- In "auto" mode, the detected model is shown in parentheses within the right-panel field.
- The `ctrl+g` dialog still correctly updates the model and the new UI field reflects the change immediately.

## Out of Scope
- Implementing mouse-click handlers to open the model dialog directly from the right panel.
- Modifying the underlying model detection or persistence logic.
