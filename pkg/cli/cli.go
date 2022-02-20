package cli

import (
	"fmt"
	"io"
	"strings"
)

type Cli struct {
	env    []string
	args   []string
	stdout io.Writer
}

func NewCli(env []string, args []string, stdout io.Writer) *Cli {
	return &Cli{env: env, args: args, stdout: stdout}
}

func (c *Cli) Run() error {
	var message string
	if len(c.args) == 2 && c.args[1] == "begin" {
		message = "game started"
	} else {
		message = fmt.Sprintf("hello %s", c.getVariable("NAME"))
	}
	_, err := fmt.Fprintln(c.stdout, message)
	return err
}

func (c *Cli) getVariable(key string) string {
	for _, e := range c.env {
		parts := strings.SplitN(e, "=", 2)
		if parts[0] == key {
			return parts[1]
		}
	}
	return ""
}
