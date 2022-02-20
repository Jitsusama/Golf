package main

import (
	"fmt"
	"github.com/Jitsusama/Golf/pkg/cli"
	"os"
)

func main() {
	golf := cli.NewCli(os.Environ(), os.Args, os.Stdout)

	if err := golf.Run(); err != nil {
		fmt.Printf("golf failed to run: %v\n", err)
		os.Exit(1)
	}
}
