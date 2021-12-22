package main

type Point struct {
	X, Y, Z int
}
type Cuboid struct {
	FromX, ToX, FromY, ToY, FromZ, ToZ int
}

func (c1 Cuboid) IntersectsWith(c2 Cuboid) bool {
	fromX := int(max((c1.FromX), (c2.FromX)))
	toX := int(min((c1.ToX), (c2.ToX)))
	fromY := int(max((c1.FromY), (c2.FromY)))
	toY := int(min((c1.ToY), (c2.ToY)))
	fromZ := int(max((c1.FromZ), (c2.FromZ)))
	toZ := int(min((c1.ToZ), (c2.ToZ)))

	c := Cuboid{fromX, toX, fromY, toY, fromZ, toZ}

	return c.Volume() > 0
}

func (c1 Cuboid) Volume() int {
	if c1 == (Cuboid{}) {
		return 0
	}
	x := c1.ToX - c1.FromX + 1
	if x < 0 {
		return 0
	}
	y := c1.ToY - c1.FromY + 1
	if y < 0 {
		return 0
	}
	z := c1.ToZ - c1.FromZ + 1
	if z < 0 {
		return 0
	}
	return x * y * z
}

func (c Cuboid) Remove(toRemove Cuboid) []Cuboid {
	totalCuboids := []Cuboid{
		{c.FromX, toRemove.FromX - 1, c.FromY, min(c.ToY, toRemove.ToY), max(c.FromZ, toRemove.FromZ), min(c.ToZ, toRemove.ToZ)},
		{c.FromX, min(c.ToX, toRemove.ToX), toRemove.ToY + 1, c.ToY, max(c.FromZ, toRemove.FromZ), min(c.ToZ, toRemove.ToZ)},
		{toRemove.ToX + 1, c.ToX, max(c.FromY, toRemove.FromY), c.ToY, max(c.FromZ, toRemove.FromZ), min(c.ToZ, toRemove.ToZ)},
		{max(c.FromX, toRemove.FromX), c.ToX, c.FromY, toRemove.FromY - 1, max(c.FromZ, toRemove.FromZ), min(c.ToZ, toRemove.ToZ)},
		{c.FromX, c.ToX, c.FromY, c.ToY, c.FromZ, toRemove.FromZ - 1},
		{c.FromX, c.ToX, c.FromY, c.ToY, toRemove.ToZ + 1, c.ToZ},
	}

	cuboidsWithVolume := []Cuboid{}
	for _, c := range totalCuboids {
		if c.Volume() > 0 {
			cuboidsWithVolume = append(cuboidsWithVolume, c)
		}
	}

	return cuboidsWithVolume
}
