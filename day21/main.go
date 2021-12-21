package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func part1(in []string) int {
	p1, p2 := parseInput(in)

	result := PlayDeterministicDice(p1, p2, 1000)

	lessPoints := result.Player1Points
	if result.Player2Points < lessPoints {
		lessPoints = result.Player2Points
	}
	return lessPoints * result.Rolls
}

func part2(in []string) int {
	player1, player2 := parseInput(in)

	result := PlayDiracDice(player1, player2, 21)

	if result.Player1 > result.Player2 {
		return result.Player1
	}
	return result.Player2
}

func parseInput(in []string) (Player, Player) {
	stringP1 := strings.Split(in[0], ": ")[1]
	stringP2 := strings.Split(in[1], ": ")[1]
	p1, _ := strconv.Atoi(stringP1)
	p2, _ := strconv.Atoi(stringP2)

	return Player{p1, 0}, Player{p2, 0}
}
