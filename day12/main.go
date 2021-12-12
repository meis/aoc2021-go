package main

import (
	"fmt"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	data := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(data))
	fmt.Printf("Solution of part two: %d\n", part2(data))
}

func part1(in []string) int {
	g := newGraph(in)

	return len(g.getPathsFrom("start", 0, path{}))
}

func part2(in []string) int {
	g := newGraph(in)

	return len(g.getPathsFrom("start", 1, path{}))
}
