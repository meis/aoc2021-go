package main

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}
type grid map[point]int

func newGrid(in []string) (grid, point) {
	g := make(grid)

	for x, col := range in {
		for y, row := range strings.Split(col, "") {
			value, _ := strconv.Atoi(row)
			g[point{x, y}] = value
		}
	}

	return g, point{len(in) - 1, len(in[0]) - 1}
}

func newExtendedGrid(in []string) (grid, point) {
	g := make(grid)

	originalDimension := len(in)

	for x, col := range in {
		for y, row := range strings.Split(col, "") {
			value, _ := strconv.Atoi(row)
			g[point{x, y}] = value

			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if i == 0 && j == 0 {
						continue
					}
					replicaX := originalDimension*i + x
					replicaY := originalDimension*j + y
					g[point{replicaX, replicaY}] = ((value + i + j - 1) % 9) + 1
				}
			}
		}
	}

	return g, point{(originalDimension * 5) - 1, (originalDimension * 5) - 1}
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
			fmt.Fprintf(&s, strconv.FormatInt(int64(g[point{i, j}]), 10))
		}

		fmt.Fprintf(&s, "\n")
	}

	return s.String()
}

func (g grid) getNeighbors(p point) []point {
	var neighbors []point

	adjacent := []point{
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
	}

	for _, n := range adjacent {
		_, found := g[n]
		if found {
			neighbors = append(neighbors, n)
		}
	}

	return neighbors
}
