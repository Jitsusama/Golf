package e2e

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"
)

func TestReportsOnOneHoleOnePersonGame(t *testing.T) {
	stdout := &bytes.Buffer{}

	adapter, cleanup := startAdapter(t)
	defer cleanup()
	gateway, cleanup := startGateway(t, adapter)
	defer cleanup()

	runGolf(t, gateway, stdout, "begin")
	runGolf(t, gateway, stdout, "round", "5")
	runGolf(t, gateway, stdout, "end")

	if matches, err := regexp.Match(
		"/game started.*5 strokes for hole.*5 strokes for game/im",
		stdout.Bytes(),
	); err != nil || !matches {
		t.Errorf(
			"%q lacks either the game's start, hole score or game score",
			stdout,
		)
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

func runGolf(t *testing.T, gateway string, stdout *bytes.Buffer, args ...string) {
	t.Helper()

	path := filepath.Join("..", "cmd", "golf", "main.go")
	args = append([]string{"run", path}, args...)

	cmd := exec.Command("go", args...)
	cmd.Stdout = stdout
	cmd.Stderr = io.Discard
	cmd.Env = append(os.Environ(), fmt.Sprintf("GOLF_HOST=%s", gateway))

	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to run golf: %v", err)
	}
}
