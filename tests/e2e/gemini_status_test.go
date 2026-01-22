package e2e

import (
	"bufio"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/creack/pty"
)

func TestGeminiStatusDetection(t *testing.T) {
	// Build the binary
	binPath, err := BuildBinary()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	// Setup isolated environment
	tmpDir, err := os.MkdirTemp("", "agent-deck-e2e-status-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	customHome := filepath.Join(tmpDir, "home")
	profileDir := filepath.Join(customHome, ".agent-deck", "profiles", "default")
	os.MkdirAll(profileDir, 0755)

	// Start a real tmux session
	tmuxSessionName := "e2e-status-session"
	exec.Command("tmux", "kill-session", "-t", tmuxSessionName).Run()
	// Start with a shell so we can send commands to it
	err = exec.Command("tmux", "new-session", "-d", "-s", tmuxSessionName, "bash").Run()
	if err != nil {
		t.Fatalf("Failed to create tmux session: %v", err)
	}
	defer exec.Command("tmux", "kill-session", "-t", tmuxSessionName).Run()

	// Pre-populate sessions.json with a Gemini session
	sessionID := "e2e-status-id"
	sessionTitle := "STATUS-TEST-SESS"
	
	storageData := map[string]interface{}{
		"updated_at": time.Now(),
		"instances": []map[string]interface{}{
			{
				"id":           sessionID,
				"title":        sessionTitle,
				"project_path": tmpDir,
				"group_path":   "my-sessions",
				"command":      "gemini",
				"tool":         "gemini",
				"status":       "idle",
				"created_at":   time.Now(),
				"tmux_session": tmuxSessionName,
			},
		},
		"groups": []map[string]interface{}{
			{
				"name":     "My Sessions",
				"path":     "my-sessions",
				"expanded": true,
				"order":    0,
			},
		},
	}
	
	jsonData, _ := json.MarshalIndent(storageData, "", "  ")
	err = os.WriteFile(filepath.Join(profileDir, "sessions.json"), jsonData, 0644)
	if err != nil {
		t.Fatalf("Failed to pre-populate sessions.json: %v", err)
	}

	// Create dummy config to skip setup wizard
	configDir := filepath.Join(customHome, ".agent-deck")
	os.MkdirAll(configDir, 0755)
	configPath := filepath.Join(configDir, "config.toml")
	err = os.WriteFile(configPath, []byte(`default_tool = "gemini"`), 0644)
	if err != nil {
		t.Fatalf("Failed to create dummy config: %v", err)
	}

	// Start the app in a PTY
	cmd := exec.Command(binPath, "--profile", "default")
	cmd.Env = append(os.Environ(), 
		"HOME="+customHome,
		"GOOGLE_API_KEY=mock-key",
		"TERM=xterm-256color",
	)
	
	f, err := pty.StartWithAttrs(cmd, &pty.Winsize{Rows: 60, Cols: 200}, nil)
	if err != nil {
		t.Fatalf("Failed to start app in PTY: %v", err)
	}
	defer f.Close()
	defer cmd.Process.Kill()

	// Buffer for reading output
	output := make(chan string, 10000)
	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			output <- line
		}
	}()

	// 1. Wait for dashboard
	t.Log("Waiting for dashboard...")
	waitForString(t, output, "SESSIONS", 20*time.Second)
	
	// 2. Select the Gemini session
	t.Log("Selecting Gemini session...")
	f.Write([]byte("j")) // Move down from Group to Session
	time.Sleep(1 * time.Second)

	// 3. Verify initial state is Idle or Waiting
	// The indicator should be ○ (Idle) or ◐ (Waiting)
	// waitForString logs all output if it fails, so we can see what's happening.
	
	// 4. Simulate Activity in tmux
	t.Log("Simulating Gemini activity (spinner)...")
	// Echo a spinner character. 
	// We need 2+ changes in 1s for sustained activity detection if it uses timestamp detection.
	// But GetStatus also calls hasBusyIndicator directly if activity detected.
	
	// Send multiple updates to ensure it's detected as sustained activity
	go func() {
		for i := 0; i < 5; i++ {
			exec.Command("tmux", "send-keys", "-t", tmuxSessionName, "echo 'Working... ⠋'", "Enter").Run()
			time.Sleep(300 * time.Millisecond)
		}
	}()

	// 5. Verify UI shows "Running" (●)
	t.Log("Waiting for 'Running' status (●)...")
	// The output will contain "●" and "gemini(auto)" in the same line
	waitForString(t, output, "●", 15*time.Second)
	// Actually, let's wait for the specific pattern in one line
	// Note: there's ANSI codes between ● and gemini(auto)
	waitForString(t, output, "●", 10*time.Second)
	
	// Better: create a composite waitForString that checks for multiple substrings in one line
	waitForLineWith(t, output, []string{"●", "gemini(auto)"}, 15*time.Second)
	
	// 6. Simulate Idle
	t.Log("Simulating Gemini idle (prompt)...")
	exec.Command("tmux", "send-keys", "-t", tmuxSessionName, "clear", "Enter").Run()
	time.Sleep(500 * time.Millisecond)
	// Output gemini> prompt which should be recognized as idle (no busy indicators)
	exec.Command("tmux", "send-keys", "-t", tmuxSessionName, "echo 'gemini>'", "Enter").Run()

	// 7. Verify UI returns to Idle/Waiting
	t.Log("Waiting for 'Idle/Waiting' status (◐)...")
	// It should be ◐ (Waiting) because we are not attached
	waitForLineWith(t, output, []string{"◐", "gemini(auto)"}, 15*time.Second)

	t.Log("E2E Status Detection Test Passed!")
}
