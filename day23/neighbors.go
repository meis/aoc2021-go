package main

const (
	HW01 = iota
	HW02
	HW03
	HW04
	HW05
	HW06
	HW07
	HW08
	HW09
	HW10
	HW11
	RA01
	RA02
	RA03
	RA04
	RB01
	RB02
	RB03
	RB04
	RC01
	RC02
	RC03
	RC04
	RD01
	RD02
	RD03
	RD04
)

var Neighbors = [][]int{
	HW01: {HW02},
	HW02: {HW01, HW03},
	HW03: {HW02, HW04, RA01},
	HW04: {HW03, HW05},
	HW05: {HW04, HW06, RB01},
	HW06: {HW05, HW07},
	HW07: {HW06, HW08, RC01},
	HW08: {HW07, HW09},
	HW09: {HW08, HW10, RD01},
	HW10: {HW09, HW11},
	HW11: {HW10},
	RA01: {HW03, RA02},
	RA02: {RA01, RA03},
	RA03: {RA02, RA04},
	RA04: {RA03},
	RB01: {HW05, RB02},
	RB02: {RB01, RB03},
	RB03: {RB02, RB04},
	RB04: {RB03},
	RC01: {HW07, RC02},
	RC02: {RC01, RC03},
	RC03: {RC02, RC04},
	RC04: {RC03},
	RD01: {HW09, RD02},
	RD02: {RD01, RD03},
	RD03: {RD02, RD04},
	RD04: {RD03},
}

func IsRoom(space int) bool {
	return space >= RA01
}

func RoomType(space int) byte {
	switch space {
	case RA01, RA02, RA03, RA04:
		return 'A'
	case RB01, RB02, RB03, RB04:
		return 'B'
	case RC01, RC02, RC03, RC04:
		return 'C'
	case RD01, RD02, RD03, RD04:
		return 'D'
	default:
		return ' '
	}
}
