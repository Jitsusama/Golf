package cli

import (
	"fmt"
	"io"
)

type Cli struct {
	env    map[string]string
	args   []string
	stdout io.Writer
}

func NewCli(env map[string]string, args []string, stdout io.Writer) *Cli {
	return &Cli{env: env, args: args, stdout: stdout}
}

func (c *Cli) Run() error {
	var message string
	if len(c.args) == 2 && c.args[1] == "begin" {
		message = "game started"
	} else {
		message = fmt.Sprintf("hello %s", c.env["NAME"])
	}
	_, err := fmt.Fprintln(c.stdout, message)
	return err
}
