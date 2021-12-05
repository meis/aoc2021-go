package main

import (
	"fmt"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	lines := parseInput(input.GetInputStrings())
	fmt.Printf("Solution of part one: %d\n", part1(lines))
	fmt.Printf("Solution of part two: %d\n", part2(lines))
}

func parseInput(in []string) []Line {
	var lines []Line

	for _, l := range in {
		lines = append(lines, ParseLine(l))
	}

	return lines
}

func part1(lines []Line) int {
	lm := NewLineMap()

	for _, line := range lines {
		if line.IsVertical() || line.IsHorizontal() {
			lm.AddLine(line)
		}
	}

	return lm.DangerousPoints()
}

func part2(lines []Line) int {
	lm := NewLineMap()

	for _, line := range lines {
		lm.AddLine(line)
	}

	return lm.DangerousPoints()
}
