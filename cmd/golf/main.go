package main

import (
	"fmt"
	"github.com/Jitsusama/Golf/pkg/cli"
	"github.com/Jitsusama/Golf/pkg/game"
	"os"
)

func main() {
	env := os.Environ()
	args := os.Args
	stdout := os.Stdout
	var g game.Game

	if err := cli.NewCli(env, args, stdout, g).Run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
