package main

import (
	"os"

	"github.com/jimmyfielding/maps-api-project/cmd/cli/cmd"
)

func main() {
	cmd := cmd.NewTitlesCommand(os.Stdin, os.Stdout, os.Stderr)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
