package main

type Grid map[Point]bool
type Point struct {
	x int
	y int
}

func (g *Grid) Lights() int {
	lights := 0
	for _, v := range *g {
		if v {
			lights++
		}
	}
	return lights
}

func (g *Grid) Enhance(al Algorithm, blip bool) Grid {
	newGrid := make(Grid)
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for p := range *g {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			bits := []bool{
				g.getPointBit(Point{x - 1, y - 1}, blip),
				g.getPointBit(Point{x, y - 1}, blip),
				g.getPointBit(Point{x + 1, y - 1}, blip),
				g.getPointBit(Point{x - 1, y}, blip),
				g.getPointBit(Point{x, y}, blip),
				g.getPointBit(Point{x + 1, y}, blip),
				g.getPointBit(Point{x - 1, y + 1}, blip),
				g.getPointBit(Point{x, y + 1}, blip),
				g.getPointBit(Point{x + 1, y + 1}, blip),
			}

			if al.GetBit(bits) {
				newGrid[Point{x, y}] = true
			} else {
				newGrid[Point{x, y}] = false
			}
		}
	}
	return newGrid
}

func (g *Grid) getPointBit(p Point, blip bool) bool {
	val, found := (*g)[p]
	if found {
		return val
	} else {
		return blip
	}
}
