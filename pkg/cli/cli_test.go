package cli_test

import (
	"bytes"
	"github.com/Jitsusama/Golf/pkg/cli"
	"testing"
)

func TestPrintsHello(t *testing.T) {
	stdout := &bytes.Buffer{}

	err := cli.NewCli(stdout).Run()

	if err != nil {
		t.Fatalf("cli failed to run: %v", err)
	}
	if stdout.String() != "hello\n" {
		t.Errorf("got %q want %q", stdout.String(), "hello")
	}
}
