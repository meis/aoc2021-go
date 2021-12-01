package main

import (
	"fmt"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	data := input.GetInputInts()

	fmt.Printf("Solution of part one: %d\n", part1(data))
	fmt.Printf("Solution of part two: %d\n", part2(data))
}

func part1(data []int) int {
	previous := 0
	larger := 0

	for _, num := range data {
		if num > previous {
			larger++
		}
		previous = num
	}
	return larger - 1
}

func part2(data []int) int {
	previous := 0
	larger := 0

	for i, _ := range data {
		if i+2 < len(data) {
			num := data[i] + data[i+1] + data[i+2]
			if num > previous {
				larger++
			}
			previous = num
		}
	}
	return larger - 1
}
