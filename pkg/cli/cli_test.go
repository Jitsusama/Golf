package cli_test

import (
	"bytes"
	"fmt"
	"github.com/Jitsusama/Golf/pkg/cli"
	"strings"
	"testing"
)

func TestPrintsHello(t *testing.T) {
	stdout := &bytes.Buffer{}

	if err := cli.NewCli(nil, stdout).Run(); err != nil {
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
			env := map[string]string{"NAME": name}

			if err := cli.NewCli(env, stdout).Run(); err != nil {
				t.Fatalf("cli failed to run: %v", err)
			}

			if !strings.Contains(stdout.String(), name) {
				t.Errorf("%q does not contain %q", stdout.String(), name)
			}
		})
	}
}
