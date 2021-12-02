package main

import (
	"fmt"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := getInput()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func getInput() []command {
	var in []command
	for _, c := range input.GetInputStrings() {
		in = append(in, newCommand(c))
	}

	return in
}

func part1(commands []command) int {
	depth := 0
	horizontalPosition := 0

	for _, c := range commands {
		switch c.direction {
		case "forward":
			horizontalPosition += c.units
		case "up":
			depth -= c.units
		case "down":
			depth += c.units
		}
	}

	return depth * horizontalPosition
}

func part2(commands []command) int {
	depth := 0
	horizontalPosition := 0
	aim := 0

	for _, c := range commands {
		switch c.direction {
		case "forward":
			horizontalPosition += c.units
			depth += aim * c.units
		case "up":
			aim -= c.units
		case "down":
			aim += c.units
		}
	}

	return depth * horizontalPosition
}
