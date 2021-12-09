package main

import (
	"fmt"
	"sort"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func part1(in []string) int {
	hm := newHeightmap(in)
	sum := 0

	for _, p := range hm.getLowerPoints() {
		sum += 1 + hm[p]
	}

	return sum
}

func part2(in []string) int {
	hm := newHeightmap(in)
	var basins [][]point

	for _, p := range hm.getLowerPoints() {
		basins = append(basins, hm.getBasinPoints(p))
	}
	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	return len(basins[0]) * len(basins[1]) * len(basins[2])
}
