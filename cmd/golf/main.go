package main

import (
	"fmt"
	"github.com/Jitsusama/Golf/pkg/cli"
	"os"
	"strings"
)

func main() {
	env := parseEnvironment(os.Environ())
	golf := cli.NewCli(env, os.Args, os.Stdout)

	if err := golf.Run(); err != nil {
		fmt.Printf("golf failed to run: %v\n", err)
		os.Exit(1)
	}
}

func parseEnvironment(env []string) map[string]string {
	result := make(map[string]string)
	for _, e := range env {
		parts := strings.SplitN(e, "=", 2)
		result[parts[0]] = parts[1]
	}
	return result
}
