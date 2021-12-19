package main

import (
	"fmt"
	"math"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputStrings()
	result := assemble(in)

	fmt.Printf("Solution of part one: %d\n", part1(result))
	fmt.Printf("Solution of part two: %d\n", part2(result))
}

type result struct {
	scanners []Position
	beacons  map[Position]bool
}

func part1(result result) int {

	return len(result.beacons)
}

func part2(result result) int {
	maxDistance := 0
	for _, s1 := range result.scanners {
		for _, s2 := range result.scanners {
			distance := int(math.Abs(float64(s1.x-s2.x)) + math.Abs(float64(s1.y-s2.y)) + math.Abs(float64(s1.z-s2.z)))
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return maxDistance
}

func assemble(in []string) result {
	readings := parseInput(in)

	scanners := []Position{{0, 0, 0}}
	beacons := make(map[Position]bool)
	for _, b := range readings[0] {
		beacons[b] = true
	}

	pending := NewIndexQueue()
	for i := 1; i < len(readings); i++ {
		pending.add(i)
	}

	for !pending.empty() {
		currentIndex := pending.next()
		current := readings[currentIndex]

		accumulated := []Position{}
		for p := range beacons {
			accumulated = append(accumulated, p)
		}
		overlaps := OverlapingIndexes(&accumulated, &current, 12)

		if len(overlaps) >= 12 {
			indexes1 := []int{}
			indexes2 := []int{}
			for k, v := range overlaps {
				indexes1 = append(indexes1, k)
				indexes2 = append(indexes2, v)
			}

			beaconsAccumulated := []Position{}
			for _, i := range indexes1 {
				beaconsAccumulated = append(beaconsAccumulated, accumulated[i])
			}
		out:
			for _, variation := range Variations(&current) {
				beaconsVariation := []Position{}
				for _, i := range indexes2 {
					beaconsVariation = append(beaconsVariation, variation[i])
				}
				distance, err := SameDistance(beaconsAccumulated, beaconsVariation)
				if !err {
					scanners = append(scanners, distance)
					for i, beacon := range variation {
						if !has(indexes2, i) {
							beacons[Position{
								beacon.x + distance.x,
								beacon.y + distance.y,
								beacon.z + distance.z,
							}] = true
						}
					}

					break out
				}
			}
		} else {
			pending.add(currentIndex)
		}

	}

	return result{scanners, beacons}
}

func has(a []int, el int) bool {
	for _, x := range a {
		if x == el {
			return true
		}
	}
	return false
}

func parseInput(in []string) [][]Position {
	readings := [][]Position{}

	var current []string

	for _, line := range in {
		if line[1] == '-' {
			if len(current) > 0 {
				readings = append(readings, ParsePositions(current))
			}
			current = []string{}
		} else {
			current = append(current, line)
		}

	}
	readings = append(readings, ParsePositions(current))

	return readings
}

func OverlapingIndexes(r1 *[]Position, r2 *[]Position, confidence int) map[int]int {
	overlaps := make(map[int]int)
	for i1, b1 := range *r1 {
		d1 := b1.Distances(r1)
		for i2, b2 := range *r2 {
			join := make(map[float64]bool)
			d2 := b2.Distances(r2)

			for v := range d1 {
				join[v] = true
			}
			for v := range d2 {
				join[v] = true
			}

			if len(d1)+len(d2)-len(join) >= confidence-1 {
				overlaps[i1] = i2
			}
		}
	}

	return overlaps
}
