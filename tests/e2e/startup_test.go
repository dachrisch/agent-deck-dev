package e2e

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/creack/pty"
)

func TestStartup(t *testing.T) {
	// Build the binary
	binPath, err := BuildBinary()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	// Use a temporary profile to avoid lock conflicts
	tmpDir, err := os.MkdirTemp("", "agent-deck-e2e-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Setup custom home for the app to keep it isolated
	customHome := filepath.Join(tmpDir, "home")
	os.MkdirAll(customHome, 0755)

	// Start the app in a PTY
	cmd := exec.Command(binPath, "--profile", "e2e-test")
	cmd.Env = append(os.Environ(), "HOME="+customHome)
	
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

	// Verify initial UI elements
	timeout := time.After(15 * time.Second)
	var capturedLines []string

	for {
		select {
		case line := <-output:
			capturedLines = append(capturedLines, line)
			// Look for 'Agent Deck' in the output
			if strings.Contains(line, "Agent Deck") {
				return // Success
			}
		case <-timeout:
			t.Logf("Captured %d lines of output before timeout", len(capturedLines))
			for i, l := range capturedLines {
				t.Logf("Line %d: %q", i, l)
			}
			t.Errorf("Timeout reached. Could not find 'Agent Deck' in UI output")
			return
		}
	}
}
