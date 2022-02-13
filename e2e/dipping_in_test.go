package e2e

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestSaysHello(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	path := filepath.Join("..", "cmd", "golf", "main.go")
	cmd := exec.Command("go", "run", path)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to run golf: %s", stderr.String())
	}
	if stdout.String() != "hello" {
		t.Errorf("stdout: got %q want %q", stdout.String(), "hello")
	}
}
