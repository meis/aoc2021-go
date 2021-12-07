package main

import (
	"fmt"
	"math"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputIntsInOneLine()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func part1(in []int) int {
	return smallestAlignCost(in, singleDistance)
}

func part2(in []int) int {
	return smallestAlignCost(in, triangularDistance)
}

func smallestAlignCost(crabPositions []int, distanceMethod func(int, int) int) int {
	min := 0
	max := crabPositions[0]

	for _, num := range crabPositions {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	minFuel := math.MaxInt

	for i := min; i <= max; i++ {
		currentFuel := 0

		for _, num := range crabPositions {
			currentFuel += distanceMethod(num, i)
		}

		if currentFuel < minFuel {
			minFuel = currentFuel
		}
	}

	return minFuel
}

func singleDistance(orig int, dest int) int {
	dist := orig - dest
	if dist < 0 {
		dist = -dist
	}

	return dist
}

func triangularDistance(orig int, dest int) int {
	singleDistance := singleDistance(orig, dest)

	triangularDistance := 0
	step := 1
	for i := 1; i <= singleDistance; i++ {
		triangularDistance += step
		step++
	}

	return triangularDistance
}
