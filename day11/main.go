package main

import (
	"fmt"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func part1(in []string) int {
	g := newGrid(in)
	totalFlashed := 0

	for i := 0; i < 100; i++ {
		totalFlashed += g.nextStep()
	}

	return totalFlashed
}

func part2(in []string) int {
	g := newGrid(in)
	i := 0

	for {
		octopusFlashed := g.nextStep()

		i++
		if octopusFlashed == len(g) {
			break
		}
	}

	return i
}
