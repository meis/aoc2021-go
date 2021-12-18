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
	numbers := parseSnailfishNumbers(in)
	current := numbers[0]
	for i := 1; i < len(numbers); i++ {
		current = current.add(numbers[i])
	}

	return current.magnitude()
}

func part2(in []string) int {
	numbers := parseSnailfishNumbers(in)
	max := 0

	for i, first := range numbers {
		for j, second := range numbers {
			if i != j {
				mag := first.add(second).magnitude()

				if mag > max {
					max = mag
				}
			}
		}
	}
	return max
}
