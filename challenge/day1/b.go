package day1

import (
	"fmt"
	"github.com/oscarteg/advent-of-code-2020/util"

	"github.com/oscarteg/advent-of-code-2020/challenge"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.FromFile()))
		},
	}
}

func b(challenge *challenge.Input) int {
	var entries []int

	for v := range challenge.Lines() {
		entries = append(entries, util.MustAtoI(v))
	}

	elementMap := util.SliceToMap(entries)

	for element := range elementMap {
		if secondElement, exists := elementMap[2020-element]; exists {
			return element * secondElement
		}
	}

	panic("There is no answer")
}