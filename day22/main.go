package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

type Step struct {
	Switch bool
	Cuboid Cuboid
}

func part1(in []string) int {
	steps := parseInput(in)
	g := make(map[Point]bool)

	initSteps := []Step{}
	for _, step := range steps {
		c := step.Cuboid
		if c.FromX < -50 || c.ToX > 50 || c.FromY < -50 || c.ToY > 50 || c.FromZ < -50 || c.ToZ > 50 {
			continue
		}
		initSteps = append(initSteps, step)
	}
	for _, step := range initSteps {
		c := step.Cuboid

		for x := c.FromX; x <= c.ToX; x++ {
			for y := c.FromY; y <= c.ToY; y++ {
				for z := c.FromZ; z <= c.ToZ; z++ {
					p := Point{x, y, z}
					if step.Switch {
						g[p] = true
					} else {
						delete(g, p)
					}
				}
			}
		}
	}
	return len(g)
}

func part2(in []string) int {
	steps := parseInput(in)

	processedCuboids := []Cuboid{}

	for _, step := range steps {
		stepCuboids := []Cuboid{}
		// For every Cuboid:
		// We try to intersect it with any previous Cuboid
		// If it does, we "remove" the current cuboid from the previous one,
		// leaving up to 6 kuboids that won't intersect instead.
		for _, nextCuboid := range processedCuboids {
			if step.Cuboid.IntersectsWith(nextCuboid) {
				diffCubes := nextCuboid.Remove(step.Cuboid)
				stepCuboids = append(stepCuboids, diffCubes...)
			} else {
				stepCuboids = append(stepCuboids, nextCuboid)
			}
		}

		// We can add the full Cuboid if it's a switch on because
		// all the previous Cuboids won't intersect with him.
		// If it's a switch off, we do nothing: we already removed
		// the affected points.
		if step.Switch {
			stepCuboids = append(stepCuboids, step.Cuboid)
		}

		processedCuboids = stepCuboids
	}

	volume := 0
	for _, c := range processedCuboids {
		volume += c.Volume()
	}

	return volume
}

func parseInput(in []string) []Step {
	steps := []Step{}

	for _, line := range in {
		parts := strings.Split(line, " ")
		t := true
		if parts[0] == "off" {
			t = false
		}
		xPart := strings.Split(strings.Split(strings.Split(parts[1], ",")[0], "=")[1], "..")
		yPart := strings.Split(strings.Split(strings.Split(parts[1], ",")[1], "=")[1], "..")
		zPart := strings.Split(strings.Split(strings.Split(parts[1], ",")[2], "=")[1], "..")

		fromX, _ := strconv.Atoi(xPart[0])
		toX, _ := strconv.Atoi(xPart[1])
		fromY, _ := strconv.Atoi(yPart[0])
		toY, _ := strconv.Atoi(yPart[1])
		fromZ, _ := strconv.Atoi(zPart[0])
		toZ, _ := strconv.Atoi(zPart[1])

		steps = append(steps, Step{t, Cuboid{fromX, toX, fromY, toY, fromZ, toZ}})
	}
	return steps
}
