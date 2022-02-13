package cli

import (
	"fmt"
	"io"
)

type Cli struct {
	stdout io.Writer
}

func NewCli(stdout io.Writer) *Cli {
	return &Cli{stdout: stdout}
}

func (c *Cli) Run() error {
	_, err := fmt.Fprintln(c.stdout, "hello")
	return err
}
