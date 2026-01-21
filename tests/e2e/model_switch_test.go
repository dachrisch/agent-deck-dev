package e2e

import (
	"bufio"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

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
	profileDir := filepath.Join(customHome, ".agent-deck", "profiles", "default")
	os.MkdirAll(profileDir, 0755)

	// Start a real tmux session to match our pre-populated session
	tmuxSessionName := "e2e-gemini-session"
	exec.Command("tmux", "kill-session", "-t", tmuxSessionName).Run()
	err = exec.Command("tmux", "new-session", "-d", "-s", tmuxSessionName, "sleep 100").Run()
	if err != nil {
		t.Fatalf("Failed to create tmux session: %v", err)
	}
	defer exec.Command("tmux", "kill-session", "-t", tmuxSessionName).Run()

	// Pre-populate sessions.json with a Gemini session
	sessionID := "e2e-test-session-id"
	sessionTitle := "MODEL-TEST-SESS"
	
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

	// Create dummy config
	configPath := filepath.Join(customHome, ".agent-deck", "config.toml")
	err = os.WriteFile(configPath, []byte(`default_tool = "gemini"`), 0644)
	if err != nil {
		t.Fatalf("Failed to create dummy config: %v", err)
	}

	// Mock models via env var
	mockModels := "gemini-1.5-flash,mock-model-v1"
	
	// Start the app in a PTY
	cmd := exec.Command(binPath, "--profile", "default")
	cmd.Env = append(os.Environ(), 
		"HOME="+customHome,
		"GEMINI_MODELS_OVERRIDE="+mockModels,
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
	waitForString(t, output, "My Sessions", 10*time.Second)

	// 2. Select the Gemini session
	t.Log("Selecting Gemini session...")
	f.Write([]byte("j")) // Move down from Group to Session
	time.Sleep(1 * time.Second)

	// 3. Open Model Manager
	t.Log("Opening Model Manager (Ctrl+G)...")
	err = WriteCtrlKey(f, 'g')
	if err != nil {
		t.Fatalf("Failed to send Ctrl+G: %v", err)
	}

	// 4. Verify and Switch
	t.Log("Verifying dialog content...")
	waitForString(t, output, "Gemini Model", 10*time.Second)
	
	// Wait for models to load
	waitForString(t, output, "gemini-1.5-flash", 15*time.Second)

	t.Log("Selecting gemini-1.5-flash...")
	f.Write([]byte("\r"))
	time.Sleep(1 * time.Second)

	// 5. Verify UI reflects the change in the list
	t.Log("Verifying UI model tag update...")
	waitForString(t, output, "gemini(gemini-1.5-flash)", 30*time.Second)
	
	t.Log("E2E Test Passed!")
}

func TestGeminiAutoModelAndDetection(t *testing.T) {
	// Build the binary
	binPath, err := BuildBinary()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}

	// Setup isolated environment
	tmpDir, err := os.MkdirTemp("", "agent-deck-e2e-auto-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	customHome := filepath.Join(tmpDir, "home")
	profileDir := filepath.Join(customHome, ".agent-deck", "profiles", "default")
	os.MkdirAll(profileDir, 0755)

	// Pre-populate sessions.json with a Gemini session
	sessionID := "e2e-auto-session-id"
	sessionTitle := "AUTO-TEST-SESS"
	
	// Start a real tmux session to match
	tmuxSessionName := "e2e-auto-session"
	exec.Command("tmux", "kill-session", "-t", tmuxSessionName).Run()
	err = exec.Command("tmux", "new-session", "-d", "-s", tmuxSessionName, "sleep 100").Run()
	if err != nil {
		t.Fatalf("Failed to create tmux session: %v", err)
	}
	defer exec.Command("tmux", "kill-session", "-t", tmuxSessionName).Run()

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
				"gemini_model": "auto", // Set to auto
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

	// Create dummy config
	configPath := filepath.Join(customHome, ".agent-deck", "config.toml")
	err = os.WriteFile(configPath, []byte(`default_tool = "gemini"`), 0644)
	if err != nil {
		t.Fatalf("Failed to create dummy config: %v", err)
	}

	// Start the app in a PTY
	cmd := exec.Command(binPath, "--profile", "default")
	cmd.Env = append(os.Environ(), 
		"HOME="+customHome,
		"GEMINI_MODELS_OVERRIDE=auto,gemini-2.0-flash",
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
	waitForString(t, output, "My Sessions", 10*time.Second)
	
	// 2. Select the Gemini session
	t.Log("Selecting Gemini session...")
	f.Write([]byte("j")) // Move down from Group to Session
	time.Sleep(1 * time.Second)

	// 3. Verify UI shows "auto" initially
	t.Log("Verifying initial 'auto' tag...")
	waitForString(t, output, "gemini(auto)", 20*time.Second)

	// 3. Simulate Gemini creating a session file with a detected model
	// We need to match the hash of tmpDir
	projectHash := HashProjectPath(tmpDir)
	chatsDir := filepath.Join(customHome, ".gemini", "tmp", projectHash, "chats")
	os.MkdirAll(chatsDir, 0755)
	
	sessionFile := filepath.Join(chatsDir, "session-2026-01-21T12-00-"+sessionID[:8]+".json")
	sessionContent := map[string]interface{}{
		"sessionId": sessionID,
		"startTime": time.Now().Format(time.RFC3339),
		"lastUpdated": time.Now().Format(time.RFC3339),
		"messages": []map[string]interface{}{
			{
				"type": "gemini",
				"content": "I am running on gemini-3-pro-preview",
				"model": "gemini-3-pro-preview", // The detected model
				"tokens": map[string]int{"input": 10, "output": 20},
			},
		},
	}
	contentData, _ := json.Marshal(sessionContent)
	err = os.WriteFile(sessionFile, contentData, 0644)
	if err != nil {
		t.Fatalf("Failed to create mock Gemini session file: %v", err)
	}
	t.Logf("Created mock session file: %s", sessionFile)

	// 4. Set the GEMINI_SESSION_ID in tmux env so the app picks up the file
	err = exec.Command("tmux", "set-environment", "-t", tmuxSessionName, "GEMINI_SESSION_ID", sessionID).Run()
	if err != nil {
		t.Fatalf("Failed to set GEMINI_SESSION_ID in tmux: %v", err)
	}

	// 5. Verify UI updates from "auto" to "gemini-3-pro-preview"
	t.Log("Waiting for dynamic model detection...")
	waitForString(t, output, "gemini(gemini-3-pro-preview)", 30*time.Second)
	
	t.Log("E2E Detection Test Passed!")
}

func waitForString(t *testing.T, output chan string, target string, timeout time.Duration) {
	deadline := time.After(timeout)
	var captured []string
	for {
		select {
		case line := <-output:
			cleanLine := StripANSI(line)
			if cleanLine != "" {
				captured = append(captured, cleanLine)
				if strings.Contains(strings.ToLower(cleanLine), strings.ToLower(target)) {
					t.Logf("Found target %q in line: %q", target, cleanLine)
					return
				}
			}
		case <-deadline:
			t.Logf("Recent output lines (%d):", len(captured))
			start := 0
			if len(captured) > 50 {
				start = len(captured) - 50
			}
			for i := start; i < len(captured); i++ {
				t.Logf("  %q", captured[i])
			}
			t.Fatalf("Timed out waiting for string: %q", target)
		}
	}
}