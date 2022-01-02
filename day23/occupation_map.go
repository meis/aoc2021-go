package main

import (
	"fmt"
)

type OccupationMap []byte
type Move struct {
	To    int
	Steps int
}

func NewOccupationMap() OccupationMap {
	return OccupationMap{
		'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.',
		' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
	}
}

var ocuppationMapCache = make(map[string]int)

func (om OccupationMap) MinimalEnergyToSolve() int {
	key := string(om)

	energy, found := ocuppationMapCache[key]
	if found {
		return energy
	}

	if om.Solved() {
		ocuppationMapCache[key] = 0
		return 0
	}

	minEnergy := 1_000_000

	for space, occupant := range om {
		if occupant == '.' || occupant == ' ' {
			continue
		}

		for _, move := range om.ValidMoves(space) {
			energy := move.Steps
			if occupant == 'B' {
				energy *= 10
			} else if occupant == 'C' {
				energy *= 100
			} else if occupant == 'D' {
				energy *= 1000
			}

			newDiagram := make(OccupationMap, len(om))
			copy(newDiagram, om)
			newDiagram[space] = '.'
			newDiagram[move.To] = occupant

			bestNextConfiguration := newDiagram.MinimalEnergyToSolve()
			if bestNextConfiguration+energy < minEnergy {
				minEnergy = bestNextConfiguration + energy
			}

		}

	}

	ocuppationMapCache[key] = minEnergy
	return minEnergy
}

func (om OccupationMap) EmptyNeighbors(fromName int) int {
	empties := 0
	for _, space := range Neighbors[fromName] {
		if om[space] == '.' {
			empties++
		}
	}

	return empties
}

func (om OccupationMap) RoomHasForeigners(t byte) bool {
	switch t {
	case 'A':
		return !(om[RA01] == 'A' || om[RA01] == '.' || om[RA01] == ' ') ||
			!(om[RA02] == 'A' || om[RA02] == '.' || om[RA02] == ' ') ||
			!(om[RA03] == 'A' || om[RA03] == '.' || om[RA03] == ' ') ||
			!(om[RA04] == 'A' || om[RA04] == '.' || om[RA04] == ' ')
	case 'B':
		return !(om[RB01] == 'B' || om[RB01] == '.' || om[RB01] == ' ') ||
			!(om[RB02] == 'B' || om[RB02] == '.' || om[RB02] == ' ') ||
			!(om[RB03] == 'B' || om[RB03] == '.' || om[RB03] == ' ') ||
			!(om[RB04] == 'B' || om[RB04] == '.' || om[RB04] == ' ')
	case 'C':
		return !(om[RC01] == 'C' || om[RC01] == '.' || om[RC01] == ' ') ||
			!(om[RC02] == 'C' || om[RC02] == '.' || om[RC02] == ' ') ||
			!(om[RC03] == 'C' || om[RC03] == '.' || om[RC03] == ' ') ||
			!(om[RC04] == 'C' || om[RC04] == '.' || om[RC04] == ' ')
	case 'D':
		return !(om[RD01] == 'D' || om[RD01] == '.' || om[RD01] == ' ') ||
			!(om[RD02] == 'D' || om[RD02] == '.' || om[RD02] == ' ') ||
			!(om[RD03] == 'D' || om[RD03] == '.' || om[RD03] == ' ') ||
			!(om[RD04] == 'D' || om[RD04] == '.' || om[RD04] == ' ')
	default:
		return true
	}
}

func (om OccupationMap) ValidMoves(fromSpaceName int) []Move {
	occupant := om[fromSpaceName]
	// If we're already at the best position, don't leave
	if RoomType(fromSpaceName) == occupant && !om.RoomHasForeigners(occupant) {
		return []Move{}
	}
	return om.recValidMoves(fromSpaceName, fromSpaceName, fromSpaceName, 0)
}

func (om OccupationMap) recValidMoves(fromSpace int, currentSpace int, previousSpace int, stepCount int) []Move {
	moves := []Move{}

	for _, candidate := range Neighbors[currentSpace] {
		if candidate == previousSpace {
			continue
		}
		if om[candidate] != '.' {
			continue
		}
		neighborMoves := om.recValidMoves(fromSpace, candidate, currentSpace, stepCount+1)
		moves = append(moves, neighborMoves...)

		// Amphipods will never stop on the space immediately outside any room
		if len(Neighbors[candidate]) == 3 {
			continue
		}
		// Once an amphipod stops moving in the hallway, it will stay in that spot until
		// it can move into a room
		if !IsRoom(fromSpace) && !IsRoom(candidate) {
			continue
		}
		// Amphipods will never move from the hallway into a room unless that room is their
		// destination room and that room contains no amphipods which do not also have that
		// room as their own destination.
		if IsRoom(candidate) {
			if RoomType(candidate) != om[fromSpace] {
				continue
			}
			if om.RoomHasForeigners(RoomType(candidate)) {
				continue
			}
			// Go untill the end
			if om.EmptyNeighbors(candidate) > 1 {
				continue
			}
			if RoomType(fromSpace) == RoomType(candidate) {
				continue
			}
		}

		moves = append(moves, Move{candidate, stepCount + 1})
	}

	return moves

}

func (om OccupationMap) Solved() bool {
	for i := 0; i < RA01; i++ {
		if om[i] != '.' {
			return false
		}
	}
	for i := RA01; i < len(om); i++ {
		if om[i] != RoomType(i) && om[i] != ' ' {
			return false
		}
	}
	return true
}

func (om OccupationMap) print() {
	fmt.Printf("#############\n")
	fmt.Printf("#%s%s%s%s%s%s%s%s%s%s%s#\n", string(om[HW01]), string(om[HW02]), string(om[HW03]), string(om[HW04]), string(om[HW05]), string(om[HW06]), string(om[HW07]), string(om[HW08]), string(om[HW09]), string(om[HW10]), string(om[HW11]))
	fmt.Printf("###%s#%s#%s#%s###\n", string(om[RA01]), string(om[RB01]), string(om[RC01]), string(om[RD01]))
	fmt.Printf("  #%s#%s#%s#%s#\n", string(om[RA02]), string(om[RB02]), string(om[RC02]), string(om[RD02]))
	fmt.Printf("  #%s#%s#%s#%s#\n", string(om[RA03]), string(om[RB03]), string(om[RC03]), string(om[RD03]))
	fmt.Printf("  #%s#%s#%s#%s#\n", string(om[RA04]), string(om[RB04]), string(om[RC04]), string(om[RD04]))
	fmt.Printf("  #########\n")
}
