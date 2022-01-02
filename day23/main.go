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
	om := LoadFoldedDiagram(ParseInput(in))

	return om.MinimalEnergyToSolve()
}

func part2(in []string) int {
	om := LoadUnfoldedDiagram(ParseInput(in))

	return om.MinimalEnergyToSolve()
}

type Input struct {
	A1, A2, B1, B2, C1, C2, D1, D2 byte
}

func ParseInput(in []string) Input {
	return Input{
		in[2][3],
		in[3][3],
		in[2][5],
		in[3][5],
		in[2][7],
		in[3][7],
		in[2][9],
		in[3][9],
	}
}

func LoadFoldedDiagram(in Input) OccupationMap {
	om := NewOccupationMap()

	om[RA01] = in.A1
	om[RA02] = in.A2
	om[RB01] = in.B1
	om[RB02] = in.B2
	om[RC01] = in.C1
	om[RC02] = in.C2
	om[RD01] = in.D1
	om[RD02] = in.D2

	return om
}

func LoadUnfoldedDiagram(in Input) OccupationMap {
	om := NewOccupationMap()

	om[RA01] = in.A1
	om[RA02] = 'D'
	om[RA03] = 'D'
	om[RA04] = in.A2
	om[RB01] = in.B1
	om[RB02] = 'C'
	om[RB03] = 'B'
	om[RB04] = in.B2
	om[RC01] = in.C1
	om[RC02] = 'B'
	om[RC03] = 'A'
	om[RC04] = in.C2
	om[RD01] = in.D1
	om[RD02] = 'A'
	om[RD03] = 'C'
	om[RD04] = in.D2

	return om
}
