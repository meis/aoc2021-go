package main

import (
	"fmt"
	"math"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func part1(in []string) int {
	individualRisks, end := newGrid(in)

	risksToEnd := getGridOfRisksTo(individualRisks, end)

	return risksToEnd[point{0, 0}] - individualRisks[point{0, 0}]
}

func part2(in []string) int {
	individualRisks, end := newExtendedGrid(in)

	risksToEnd := getGridOfRisksTo(individualRisks, end)

	return risksToEnd[point{0, 0}] - individualRisks[point{0, 0}]
}

func getGridOfRisksTo(individualRisks grid, to point) grid {
	riskToDestiny := grid{to: individualRisks[to]}
	q := queue{map[point]bool{}, []point{}}
	q.append(to)

	for !q.empty() {
		current := q.next()

		_, currentFound := riskToDestiny[current]

		if !currentFound {
			riskToDestiny[current] = math.MaxInt
		}
		for _, n := range individualRisks.getNeighbors(current) {
			neighborFromStart, neighborFound := riskToDestiny[n]
			if neighborFound {
				if riskToDestiny[current] > neighborFromStart+individualRisks[current] {
					riskToDestiny[current] = neighborFromStart + individualRisks[current]
					for _, x := range individualRisks.getNeighbors(n) {
						q.append(x)
					}
				}

			} else {
				q.append(n)
			}

		}
	}

	return riskToDestiny
}
