package e2e

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
)

// BuildBinary builds the agent-deck binary and returns the path to it.
func BuildBinary() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get caller information")
	}

	root := filepath.Join(filepath.Dir(filename), "..", "..")
	submodulePath := filepath.Join(root, "agent-deck", "cmd", "agent-deck")
	binaryPath := filepath.Join(filepath.Dir(filename), "agent-deck-test-bin")

	cmd := exec.Command("go", "build", "-o", binaryPath, ".")
	cmd.Dir = submodulePath

	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("failed to build binary: %v, output: %s", err, string(output))
	}

	return binaryPath, nil
}

// TmuxSession represents a managed tmux session for testing.
type TmuxSession struct {
	Name string
}

// NewTmuxSession starts a new tmux session with the given name and command.
func NewTmuxSession(name string, binPath string) (*TmuxSession, error) {
	// Kill session if it exists
	_ = exec.Command("tmux", "kill-session", "-t", name).Run()

	// Start new session
	cmd := exec.Command("tmux", "new-session", "-d", "-s", name, binPath)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to start tmux session: %v", err)
	}

	return &TmuxSession{Name: name}, nil
}

// Cleanup kills the tmux session.
func (s *TmuxSession) Cleanup() {
	_ = exec.Command("tmux", "kill-session", "-t", s.Name).Run()
}
