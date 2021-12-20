package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputString()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

type area struct {
	minX int
	maxX int
	minY int
	maxY int
}
type velocity struct {
	x int
	y int
}

func part1(in string) int {
	target := parseInput(in)

	maxHeight := 0
	passedX := false

	for vx := 0; !passedX; vx++ {
		x := (vx * (vx + 1)) / 2
		if x >= target.minX && x <= target.maxX {
			// TODO: Really?
			for vy := 0; vy < 200; vy++ {
				willHit, max := willHit(velocity{vx, vy}, target)
				if willHit {
					if max > maxHeight {
						maxHeight = max
					}
				}
			}
		}
		if x > target.maxX {
			passedX = true
		}
	}
	return maxHeight
}

func part2(in string) int {
	target := parseInput(in)
	hits := 0
	// TODO: Really?
	for vx := -500; vx < 500; vx++ {
		for vy := -500; vy < 500; vy++ {

			willHit, _ := willHit(velocity{vx, vy}, target)
			if willHit {
				hits++
			}
		}
	}

	return hits
}

func willHit(v velocity, a area) (bool, int) {
	passed := false
	max := 0

	vx := v.x
	vy := v.y
	x := 0
	y := 0
	for t := 1; !passed; t++ {
		x += vx
		y += vy

		if vx > 0 {
			vx--
		} else if vx < 0 {
			vx++
		}
		vy--

		if y > max {
			max = y
		}

		if x >= a.minX && x <= a.maxX && y >= a.minY && y <= a.maxY {
			return true, max
		}

		if x >= a.maxX || y <= a.minY {
			passed = true
		}
	}

	return false, 0
}

func parseInput(in string) area {
	in = strings.TrimSuffix(in, "\n")
	parts := strings.Split(in, " ")
	x := strings.Split(strings.Split(strings.Split(parts[2], "=")[1], ",")[0], "..")
	y := strings.Split(strings.Split(parts[3], "=")[1], "..")

	minX, _ := strconv.Atoi(x[0])
	maxX, _ := strconv.Atoi(x[1])
	minY, _ := strconv.Atoi(y[0])
	maxY, _ := strconv.Atoi(y[1])

	return area{minX, maxX, minY, maxY}
}
