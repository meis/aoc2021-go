package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	data := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(data))
	fmt.Printf("Solution of part two: %s\n", part2(data))
}

func part1(in []string) int {
	g, instructions := parseInput(in)

	inst := instructions[0]
	if inst.direction == "up" {
		g = g.foldUp(inst.value)
	} else {
		g = g.foldLeft(inst.value)
	}

	return len(g)
}

func part2(in []string) string {
	g, instructions := parseInput(in)

	for _, inst := range instructions {
		if inst.direction == "up" {
			g = g.foldUp(inst.value)
		} else {
			g = g.foldLeft(inst.value)
		}
	}

	return g.asString()
}

type point struct {
	x int
	y int
}
type grid map[point]bool
type instruction struct {
	direction string
	value     int
}

func parseInput(in []string) (grid, []instruction) {
	g := make(grid)
	var instructions []instruction

	for _, line := range in {
		if len(line) > 13 {
			parts := strings.Split(strings.Split(line, " ")[2], "=")
			var direction string
			if parts[0] == "x" {
				direction = "left"
			} else {
				direction = "up"
			}

			value, _ := strconv.Atoi(parts[1])

			instructions = append(instructions, instruction{direction, value})
		} else if line != "\n" {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			g[point{x, y}] = true
		}
	}

	return g, instructions
}

func (g grid) foldUp(y int) grid {
	folded := make(grid)

	for p := range g {
		foldedX := p.x
		var foldedY int
		if p.y < y {
			foldedY = p.y
		} else if p.y > y {
			foldedY = y - (p.y - y)
		}

		folded[point{foldedX, foldedY}] = true
	}

	return folded
}
func (g grid) foldLeft(x int) grid {
	folded := make(grid)

	for p := range g {
		var foldedX int
		if p.x < x {
			foldedX = p.x
		} else if p.x > x {
			foldedX = x - (p.x - x)
		}
		foldedY := p.y

		folded[point{foldedX, foldedY}] = true
	}

	return folded
}

func (g grid) asString() string {
	var s strings.Builder
	maxX := 0
	maxY := 0
	for p := range g {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	fmt.Fprintf(&s, "\n")
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if g[point{j, i}] {
				fmt.Fprintf(&s, "#")
			} else {
				fmt.Fprintf(&s, " ")
			}
		}

		fmt.Fprintf(&s, "\n")
	}

	return s.String()
}
