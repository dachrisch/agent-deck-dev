package e2e

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/creack/pty"
)

func TestGeminiModelSwitch(t *testing.T) {
	// Build the binary
	binPath, err := BuildBinary()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	// Setup isolated environment
	tmpDir, err := os.MkdirTemp("", "agent-deck-e2e-switch-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	customHome := filepath.Join(tmpDir, "home")
	os.MkdirAll(customHome, 0755)

	// Mock models via env var
	mockModels := "gemini-2.0-flash,gemini-1.5-pro,mock-model-custom"
	
	// Start the app in a PTY
	cmd := exec.Command(binPath, "--profile", "e2e-switch")
	cmd.Env = append(os.Environ(), 
		"HOME="+customHome,
		"GEMINI_MODELS_OVERRIDE="+mockModels,
		"GOOGLE_API_KEY=mock-key", // Enable API discovery path
	)
	
	f, err := pty.StartWithAttrs(cmd, &pty.Winsize{Rows: 40, Cols: 120}, nil)
	if err != nil {
		t.Fatalf("Failed to start app in PTY: %v", err)
	}
	defer f.Close()
	defer cmd.Process.Kill()

	// Buffer for reading output
	output := make(chan string)
	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			output <- line
		}
	}()

	// TODO: Phase 2 logic goes here
	t.Log("Scaffolding complete. Phase 2 will implement interaction logic.")
}