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

func TestMaintenanceDisabledByDefault(t *testing.T) {
	// Build the binary
	binPath, err := BuildBinary()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	// Setup isolated test environment (no config file)
	tmpDir := t.TempDir()
	customHome := filepath.Join(tmpDir, "home")
	os.MkdirAll(customHome, 0755)

	// Start the app in a PTY
	cmd := exec.Command(binPath)
	cmd.Env = append(os.Environ(), "HOME="+customHome)
	
	f, err := pty.StartWithAttrs(cmd, &pty.Winsize{Rows: 40, Cols: 120}, nil)
	if err != nil {
		t.Fatalf("Failed to start app in PTY: %v", err)
	}
	defer f.Close()
	defer cmd.Process.Kill()

	// Buffer for reading output
	output := make(chan string, 100)
	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			output <- line
		}
	}()

	// 1. Wait for startup (splash screen then Deck)
	waitForString(t, output, "Deck", 15*time.Second)

	// 2. Wait for a few seconds and ensure NO 'Maintenance' message appears
	// Since maintenance is disabled by default, we should never see this banner.
	deadline := time.After(5 * time.Second)
	for {
		select {
		case line := <-output:
			cleanLine := StripANSI(line)
			if strings.Contains(strings.ToLower(cleanLine), "maintenance:") {
				t.Fatalf("ERROR: Maintenance message appeared but should be disabled by default: %q", cleanLine)
			}
		case <-deadline:
			// Success! No maintenance message seen.
			t.Log("Verified: Maintenance is disabled by default.")
			return
		}
	}
}
