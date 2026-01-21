package e2e

import (
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"crypto/sha256"
	"encoding/hex"
)

// HashProjectPath generates SHA256 hash of absolute project path.
func HashProjectPath(projectPath string) string {
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		return ""
	}
	realPath, err := filepath.EvalSymlinks(absPath)
	if err != nil {
		realPath = absPath
	}
	hash := sha256.Sum256([]byte(realPath))
	return hex.EncodeToString(hash[:])
}

var ansi = regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[-a-zA-Z\\d;#?]*)*)?[\u0007\u001B\\w])|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PR-TZcf-ntqry=><~]))")

// StripANSI removes ANSI escape codes from a string.
func StripANSI(str string) string {
	return ansi.ReplaceAllString(str, "")
}

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

// WriteArrowKey sends an arrow key to the PTY.
func WriteArrowKey(f io.Writer, key string) error {
	var code []byte
	switch key {
	case "up":
		code = []byte("\x1b[A")
	case "down":
		code = []byte("\x1b[B")
	case "right":
		code = []byte("\x1b[C")
	case "left":
		code = []byte("\x1b[D")
	}
	_, err := f.Write(code)
	return err
}

// WriteTab sends a tab key to the PTY.
func WriteTab(f io.Writer) error {
	_, err := f.Write([]byte{'\t'})
	return err
}

// WriteCtrlKey sends a control key combination to the PTY.
func WriteCtrlKey(f io.Writer, key rune) error {
	// Ctrl+A is 1, Ctrl+B is 2, etc.
	code := byte(key - 'a' + 1)
	_, err := f.Write([]byte{code})
	return err
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
