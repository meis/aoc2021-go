package main

import (
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}
type Line struct {
	start Point
	end   Point
}
type LineMap map[Point]int

func NewLineMap() LineMap {
	return LineMap(make(map[Point]int))
}
func ParseLine(s string) Line {
	c := strings.Split(s, " -> ")
	start := ParsePoint(c[0])
	end := ParsePoint(c[1])

	return Line{start, end}
}

func ParsePoint(s string) Point {
	dimensions := strings.Split(s, ",")
	x, _ := strconv.Atoi(dimensions[0])
	y, _ := strconv.Atoi(dimensions[1])

	return Point{x, y}
}

func (l *Line) IsVertical() bool {
	return l.start.y == l.end.y
}
func (l *Line) IsHorizontal() bool {
	return l.start.x == l.end.x
}
func (l *Line) GetPoints() []Point {
	var points []Point

	directionX, directionY := l.GetDirection()

	currentX := l.start.x
	currentY := l.start.y
	points = append(points, Point{currentX, currentY})

	for currentX != l.end.x || currentY != l.end.y {
		currentX += directionX
		currentY += directionY
		points = append(points, Point{currentX, currentY})
	}

	return points
}

func (l *Line) GetDirection() (int, int) {
	directionX := 0
	directionY := 0

	if l.start.x > l.end.x {
		directionX = -1
	} else if l.start.x < l.end.x {
		directionX = 1
	}

	if l.start.y > l.end.y {
		directionY = -1
	} else if l.start.y < l.end.y {
		directionY = 1
	}

	return directionX, directionY
}

func (m *LineMap) DangerousPoints() int {
	dangerous := 0

	for _, v := range *m {
		if v > 1 {
			dangerous++
		}
	}

	return dangerous
}
func (m *LineMap) AddLine(line Line) {
	for _, p := range line.GetPoints() {
		(*m)[p] += 1
	}
}
