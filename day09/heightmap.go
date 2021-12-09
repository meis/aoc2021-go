package main

import (
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type heightmap map[point]int

func newHeightmap(in []string) heightmap {
	hm := make(heightmap, 0)

	for x, col := range in {
		for y, row := range strings.Split(col, "") {
			value, _ := strconv.Atoi(row)
			hm[point{x, y}] = value
		}
	}

	return hm
}

func (hm heightmap) getLowerPoints() []point {
	var lowerPoints []point

	for p, height := range hm {
		neighbors := hm.getNeighbors(p)

		isLowerPoint := true
		for _, n := range neighbors {
			if height >= hm[n] {
				isLowerPoint = false
			}
		}
		if isLowerPoint {
			lowerPoints = append(lowerPoints, p)
		}
	}

	return lowerPoints
}

func (hm heightmap) getNeighbors(p point) []point {
	var neighbors []point

	adjacent := []point{
		{p.x - 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y},
	}

	for _, n := range adjacent {
		_, found := hm[n]
		if found {
			neighbors = append(neighbors, n)
		}
	}

	return neighbors
}

func (hm heightmap) getBasinPoints(p point) []point {
	visited := map[point]bool{p: true}
	points := []point{p}
	points = append(points, hm.getBasinPointsRec(visited, p)...)

	return points
}

func (hm heightmap) getBasinPointsRec(visited map[point]bool, p point) []point {
	var points []point

	for _, n := range hm.getNeighbors(p) {
		found := visited[n]
		if !found && hm[n] != 9 {
			points = append(points, n)
			visited[n] = true
			points = append(points, hm.getBasinPointsRec(visited, n)...)
		}

	}

	return points
}
