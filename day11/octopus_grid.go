package main

import (
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type octopus struct {
	level   int
	flashed bool
}

type octopusGrid map[point]octopus

func newGrid(in []string) octopusGrid {
	g := make(octopusGrid)

	for x, col := range in {
		for y, row := range strings.Split(col, "") {
			value, _ := strconv.Atoi(row)
			g[point{x, y}] = octopus{value, false}
		}
	}

	return g
}

func (g octopusGrid) nextStep() int {
	octopusFlashed := 0

	for p := range g {
		g.increaseEnergy(p)
	}

	for p, octopus := range g {
		if octopus.flashed {
			octopusFlashed++
			octopus.level = 0
			octopus.flashed = false
			g[p] = octopus
		}
	}

	return octopusFlashed
}

func (g octopusGrid) increaseEnergy(p point) {
	octopus := g[p]
	octopus.level++
	g[p] = octopus

	if octopus.level > 9 && !octopus.flashed {
		octopus.flashed = true
		g[p] = octopus

		for _, np := range g.getNeighbors(p) {
			g.increaseEnergy(np)
		}
	}
}

func (g octopusGrid) getNeighbors(p point) []point {
	var neighbors []point

	adjacent := []point{
		{p.x - 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y - 1},
		{p.x + 1, p.y},
		{p.x + 1, p.y + 1},
	}

	for _, n := range adjacent {
		_, found := g[n]
		if found {
			neighbors = append(neighbors, n)
		}
	}

	return neighbors
}
