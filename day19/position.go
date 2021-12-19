package main

import (
	"math"
	"strconv"
	"strings"
)

type Position struct {
	x, y, z int
}

func ParsePositions(lines []string) []Position {
	reading := []Position{}

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		reading = append(reading, Position{x, y, z})
	}

	return reading
}

func (p *Position) Distances(positions *[]Position) map[float64]bool {
	distances := make(map[float64]bool)

	for _, otherBeacon := range *positions {
		if p.x == otherBeacon.x && p.y == otherBeacon.y && p.z == otherBeacon.z {
			continue
		}
		distance := distance(*p, otherBeacon)
		distances[distance] = true
	}

	return distances
}

type cacheKey struct {
	from Position
	to   Position
}

var cache = make(map[cacheKey]float64)

func distance(from Position, to Position) float64 {

	result, found := cache[cacheKey{from, to}]
	if found {
		return result
	}
	result, found = cache[cacheKey{to, from}]
	if found {
		return result
	}
	result = math.Sqrt(
		float64(
			math.Pow(float64(from.x-to.x), 2) +
				math.Pow(float64(from.y-to.y), 2) +
				math.Pow(float64(from.z-to.z), 2),
		),
	)
	cache[cacheKey{from, to}] = result

	return result
}

func SameDistance(ps1 []Position, ps2 []Position) (Position, bool) {
	dist := Position{
		ps1[0].x - ps2[0].x,
		ps1[0].y - ps2[0].y,
		ps1[0].z - ps2[0].z,
	}

	for i := 1; i < len(ps1); i++ {
		if ps1[i].x-ps2[i].x != dist.x {
			return Position{}, true
		}
		if ps1[i].y-ps2[i].y != dist.y {
			return Position{}, true
		}
		if ps1[i].z-ps2[i].z != dist.z {
			return Position{}, true
		}
	}

	return dist, false
}

func Variations(reading *[]Position) [][]Position {
	variations := [][]int{
		{1, 1, 1},
		{1, 1, -1},
		{1, -1, 1},
		{1, -1, -1},
		{-1, 1, 1},
		{-1, 1, -1},
		{-1, -1, 1},
		{-1, -1, -1},
	}
	rotations := make([][]Position, 48)

	for _, beacon := range *reading {
		position := 0
		for i := 0; i < 6; i++ {
			rotation := beacon
			if i == 1 {
				rotation = Position{beacon.z, beacon.x, beacon.y}
			} else if i == 2 {
				rotation = Position{beacon.y, beacon.z, beacon.x}
			} else if i == 3 {
				rotation = Position{beacon.x, beacon.z, beacon.y}
			} else if i == 4 {
				rotation = Position{beacon.z, beacon.y, beacon.x}
			} else if i == 5 {
				rotation = Position{beacon.y, beacon.x, beacon.z}
			}

			for _, v := range variations {

				rotations[position] = append(rotations[position], Position{
					rotation.x * v[0],
					rotation.y * v[1],
					rotation.z * v[2],
				})
				position++
			}
		}

	}
	return rotations
}
