package e2e

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestReportsOnOneHoleOnePersonGame(t *testing.T) {
	adapter, cleanup := startAdapter(t)
	defer cleanup()
	gateway, cleanup := startGateway(t, adapter)
	defer cleanup()

	runGolf(t, gateway, "begin")
	runGolf(t, gateway, "round", "5")
	stdout := runGolf(t, gateway, "end")

	if !strings.Contains(stdout, "5 points") {
		t.Errorf("%q does not contain 5 points", stdout)
	}
}

func startAdapter(t *testing.T) (string, func()) {
	t.Helper()

	stderr := &bytes.Buffer{}
	path := filepath.Join("..", "cmd", "adapter", "main.go")

	cmd := exec.Command("go", "run", path)
	cmd.Stdout = io.Discard
	cmd.Stderr = stderr
	cmd.Env = append(os.Environ(), "LISTENING_PORT=9012")

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to run adapter: %s", stderr)
	}

	return "http://localhost:9012", func() {
		_ = cmd.Process.Kill()
	}
}

func startGateway(t *testing.T, adapter string) (string, func()) {
	t.Helper()

	stderr := &bytes.Buffer{}
	path := filepath.Join("..", "cmd", "gateway", "main.go")

	cmd := exec.Command("go", "run", path)
	cmd.Stdout = io.Discard
	cmd.Stderr = stderr
	cmd.Env = append(
		os.Environ(),
		"LISTENING_PORT=9013",
		fmt.Sprintf("ADAPTER_HOST=%s", adapter),
	)

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to run adapter: %s", stderr)
	}

	return "http://localhost:9013", func() {
		_ = cmd.Process.Kill()
	}
}

func runGolf(t *testing.T, gateway string, args ...string) string {
	t.Helper()

	stdout := &bytes.Buffer{}
	path := filepath.Join("..", "cmd", "golf", "main.go")
	args = append([]string{"run", path}, args...)

	cmd := exec.Command("go", args...)
	cmd.Stdout = stdout
	cmd.Stderr = io.Discard
	cmd.Env = append(os.Environ(), fmt.Sprintf("GOLF_HOST=%s", gateway))

	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to run golf: %v", err)
	}

	return stdout.String()
}
