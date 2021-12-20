package main

import (
	"fmt"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func part1(in []string) int {
	al, g := parseInput(in)

	enchanced := enhance(al, g, 2)

	return enchanced.Lights()
}

func part2(in []string) int {
	al, g := parseInput(in)

	enchanced := enhance(al, g, 50)

	return enchanced.Lights()
}

func enhance(al Algorithm, g Grid, iterations int) Grid {
	blips := al.Blips()
	blip := false
	for i := 1; i <= iterations; i++ {
		if blips {
			g = g.Enhance(al, blip)
		} else {
			g = g.Enhance(al, false)
		}
		blip = !blip
	}

	return g

}

func parseInput(in []string) (Algorithm, Grid) {
	al := []bool{}
	g := make(Grid)

	for _, bit := range in[0] {
		if bit == '#' {
			al = append(al, true)
		} else {
			al = append(al, false)
		}
	}

	for y, row := range in {
		if y == 0 {
			continue
		}

		for x, col := range row {
			if col == '#' {
				g[Point{x, y - 1}] = true
			} else {
				g[Point{x, y - 1}] = false
			}
		}
	}

	return al, g
}
