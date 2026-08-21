package main

import (
	"os"

	"github.com/twoBoots/bender/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
