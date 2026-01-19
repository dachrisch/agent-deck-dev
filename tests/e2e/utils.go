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
