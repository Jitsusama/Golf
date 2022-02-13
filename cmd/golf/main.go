package main

import (
	"fmt"
	"github.com/Jitsusama/Golf/pkg/cli"
	"os"
)

func main() {
	if err := cli.NewCli(os.Stdout).Run(); err != nil {
		fmt.Printf("golf failed to run: %v\n", err)
		os.Exit(1)
	}
}
