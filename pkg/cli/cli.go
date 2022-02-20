package cli

import (
	"fmt"
	g "github.com/Jitsusama/Golf/pkg/game"
	"io"
)

type Cli struct {
	args []string
	game g.Game
}

func NewCli(env []string, args []string, stdout io.Writer, game g.Game) *Cli {
	return &Cli{args, game}
}

func (c *Cli) Run() error {
	if len(c.args) == 2 && c.args[1] == "begin" {
		if err := c.game.Begin(); err != nil {
			return fmt.Errorf("failed to begin a g: %v", err)
		}
	}
	return nil
}
