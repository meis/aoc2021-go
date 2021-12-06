package main

import (
	"fmt"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	data := input.GetInputIntsInOneLine()

	fmt.Printf("Solution of part one: %d\n", part1(data))
	fmt.Printf("Solution of part two: %d\n", part2(data))
}

func part1(in []int) int {
	counter := NewFishCounter(in)

	for it := 1; it <= 80; it++ {
		counter = counter.Iterate()
	}

	return counter.TotalFish()
}

func part2(in []int) int {
	counter := NewFishCounter(in)

	for it := 1; it <= 256; it++ {
		counter = counter.Iterate()
	}

	return counter.TotalFish()
}
