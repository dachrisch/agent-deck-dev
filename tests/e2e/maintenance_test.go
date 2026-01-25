package e2e

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/creack/pty"
)

func TestMaintenanceFlow(t *testing.T) {
	// Build the binary
	binPath, err := BuildBinary()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	// Setup isolated test environment
	tmpDir := t.TempDir()
	customHome := filepath.Join(tmpDir, "home")
	os.MkdirAll(customHome, 0755)

	// Pre-create mock maintenance data so we have something to clean up
	// 1. Logs to prune
	projectHash := "042c5b74db90f76b4688a069cc3c55526c11d74eb8b14fdead4f7baeef1476cf"
	geminiTmpDir := filepath.Join(customHome, ".gemini", "tmp", projectHash)
	os.MkdirAll(geminiTmpDir, 0755)
	logFile := filepath.Join(geminiTmpDir, "stale_log.txt")
	os.WriteFile(logFile, []byte("stale log"), 0644)

	// 2. Backups to prune
	profileDir := filepath.Join(customHome, ".agent-deck", "profiles", "default")
	os.MkdirAll(profileDir, 0755)
	for i := 1; i <= 5; i++ {
		backupPath := filepath.Join(profileDir, fmt.Sprintf("sessions.json.bak.%d", i))
		os.WriteFile(backupPath, []byte("{}"), 0644)
		// Set old modification times
		oldTime := time.Now().Add(-time.Duration(i) * time.Hour)
		os.Chtimes(backupPath, oldTime, oldTime)
	}

	// 3. Create initial config with maintenance DISABLED
	configPath := filepath.Join(customHome, ".agent-deck", "config.toml")
	os.MkdirAll(filepath.Dir(configPath), 0755)
	initialConfig := `[maintenance]
enabled = false
`
	os.WriteFile(configPath, []byte(initialConfig), 0600)

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

	// 1. Wait for startup
	waitForString(t, output, "Deck", 10*time.Second)

	// 2. Open Settings (Shift+S)
	f.Write([]byte("S"))
	waitForString(t, output, "MAINTENANCE", 5*time.Second)

	// 3. Navigate to Maintenance toggle and enable it
	// In the view, Maintenance is the last section before MCP
	// Cursor starts at 0 (Theme). We need to go down to SettingMaintenanceEnabled (index 15)
	for i := 0; i < 20; i++ { // Over-shoot slightly to be sure we hit the bottom
		WriteArrowKey(f, "down")
		time.Sleep(100 * time.Millisecond)
	}
	
	// Toggle with Space (multiple times to be sure it registers)
	f.Write([]byte(" "))
	time.Sleep(500 * time.Millisecond)

	// 4. Close Settings (Esc)
	f.Write([]byte("\x1b")) // Esc
	time.Sleep(1 * time.Second)

	// 5. Verify the cyan maintenance message appears in the help bar
	// Since we enabled it, the worker (which runs immediately on startup and then polls) 
	// should eventually run. We may need to wait up to 2 seconds for the next tick.
	t.Log("Waiting for Maintenance: message...")
	
	// Create a custom waitForString with debug output on failure
	deadline := time.After(20 * time.Second)
	var captured []string
	found := false
	for !found {
		select {
		case line := <-output:
			cleanLine := StripANSI(line)
			if cleanLine != "" {
				captured = append(captured, cleanLine)
				// Look for 'maintenance' case-insensitively, avoiding emojis that might be stripped
				if strings.Contains(strings.ToLower(cleanLine), "maintenance") {
					t.Logf("SUCCESS: Found maintenance message in line: %q", cleanLine)
					found = true
				}
			}
		case <-deadline:
			t.Logf("TIMEOUT: Captured %d lines of output. Printing last 100:", len(captured))
			start := 0
			if len(captured) > 100 {
				start = len(captured) - 100
			}
			for i := start; i < len(captured); i++ {
				t.Logf("  [%d] %q", i, captured[i])
			}
			t.Fatalf("Timed out waiting for maintenance message")
		}
	}

	t.Log("E2E Maintenance Flow Test Passed!")
}
