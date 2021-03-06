package main

import (
	"os"

	"github.com/oscarteg/advent-of-code-2020/challenge/cmd"
)

//go:generate go run ./gen
func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}