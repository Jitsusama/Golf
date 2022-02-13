package e2e

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var path = filepath.Join("..", "cmd", "golf", "main.go")

func TestWelcomesUser(t *testing.T) {
	user := "Ariel"
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	cmd := exec.Command("go", "run", path)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Env = append(os.Environ(), fmt.Sprintf("NAME=%s", user))

	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to run cli: %s", stderr.String())
	}

	want := fmt.Sprintf("hello %s\n", user)
	if stdout.String() != want {
		t.Errorf("stdout: got %q want %q", stdout.String(), want)
	}
}
