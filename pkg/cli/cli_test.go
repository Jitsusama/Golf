package cli_test

import (
	"bytes"
	"fmt"
	"github.com/Jitsusama/Golf/pkg/cli"
	"strings"
	"testing"
)

func TestBeginsAGame(t *testing.T) {
	var env []string
	stdout := &bytes.Buffer{}
	args := []string{"golf", "begin"}

	golf := cli.NewCli(env, args, stdout)
	if err := golf.Run(); err != nil {
		t.Fatalf("cli failed to run: %v", err)
	}

	if !strings.Contains(stdout.String(), "game started") {
		t.Errorf("%q does not contain %q", stdout.String(), "game started`")
	}
}

// TODO: handle game already in progress on start

func TestPrintsHello(t *testing.T) {
	stdout := &bytes.Buffer{}

	if err := cli.NewCli(nil, nil, stdout).Run(); err != nil {
		t.Fatalf("cli failed to run: %v", err)
	}

	if !strings.Contains(stdout.String(), "hello") {
		t.Errorf("%q does not contain %q", stdout.String(), "hello")
	}
}

func TestPrintsName(t *testing.T) {
	for _, name := range []string{
		"Joel",
		"Gailyn",
	} {
		t.Run(fmt.Sprintf("prints %s", name), func(t *testing.T) {
			stdout := &bytes.Buffer{}
			env := []string{fmt.Sprintf("NAME=%s", name)}

			if err := cli.NewCli(env, nil, stdout).Run(); err != nil {
				t.Fatalf("cli failed to run: %v", err)
			}

			if !strings.Contains(stdout.String(), name) {
				t.Errorf("%q does not contain %q", stdout.String(), name)
			}
		})
	}
}
