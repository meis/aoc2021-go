package main

import (
	"fmt"
	"math"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(in))
}

func part1(in []string) int {
	m := NewMap(in)

	moved := math.MaxInt
	steps := 0
	for moved != 0 {
		m, moved = m.step()
		steps++
	}

	return steps
}
