package cli_test

import (
	"bytes"
	"github.com/Jitsusama/Golf/pkg/cli"
	"testing"
)

func TestBeginsAGame(t *testing.T) {
	var env []string
	stdout := &bytes.Buffer{}
	args := []string{"golf", "begin"}
	game := &mockGame{}

	golf := cli.NewCli(env, args, stdout, game)
	if err := golf.Run(); err != nil {
		t.Fatalf("cli failed to run: %v", err)
	}

	if game.beginCount != 1 {
		t.Errorf("g begin: got %d want %d", game.beginCount, 1)
	}
}

// TODO: handle g already in progress on start

type mockGame struct {
	beginCount int
}

func (g *mockGame) Begin() error {
	g.beginCount++
	return nil
}
