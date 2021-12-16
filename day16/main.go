package main

import (
	"fmt"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputString()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func part1(in string) int {
	packet := parseInput(in)

	return packet.versionSum()
}

func part2(in string) int {
	packet := parseInput(in)

	return packet.value()

}

func parseInput(in string) packet {
	bits := parseBitList(in)
	packet, _ := parsePacket(bits)

	return packet
}
