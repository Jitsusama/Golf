package cli

import (
	"fmt"
	"io"
)

type Cli struct {
	env    map[string]string
	stdout io.Writer
}

func NewCli(env map[string]string, stdout io.Writer) *Cli {
	return &Cli{env: env, stdout: stdout}
}

func (c *Cli) Run() error {
	message := fmt.Sprintf("hello %s", c.env["NAME"])
	_, err := fmt.Fprintln(c.stdout, message)
	return err
}
