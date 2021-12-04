package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputString()
	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func part1(in string) int {
	numbers, boards := parseInput(in)
	winner, lastNum := findWinner(numbers, boards)
	unmarkedSum := 0

	for _, num := range winner.LeftToDraw() {
		unmarkedSum += num
	}

	return unmarkedSum * lastNum
}

func part2(in string) int {
	numbers, boards := parseInput(in)
	looser, lastNum := findLooser(numbers, boards)
	unmarkedSum := 0

	for _, num := range looser.LeftToDraw() {
		unmarkedSum += num
	}

	return unmarkedSum * lastNum
}

func findWinner(numbers []int, boards []Board) (Board, int) {
	for _, number := range numbers {
		for _, board := range boards {
			board.Draw(number)
			if board.Finished() {
				return board, number
			}
		}
	}

	return Board{}, -1
}

func findLooser(numbers []int, boards []Board) (Board, int) {
	alreadyWonIndexes := make(map[int]bool)

	for _, number := range numbers {
		for i, board := range boards {
			if alreadyWonIndexes[i] {
				continue
			}

			board.Draw(number)
			if board.Finished() {
				alreadyWonIndexes[i] = true
			}

			// Was this the last board to win??
			if len(alreadyWonIndexes) == len(boards) {
				return board, number
			}
		}
	}

	return Board{}, -1
}

func parseInput(in string) ([]int, []Board) {
	var numbers []int
	var boards []Board

	parts := strings.Split(in, "\n")
	numberStrings := strings.Split(parts[0], ",")
	boardStrings := input.CleanEmpty(parts[1:])

	for _, s := range numberStrings {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	for i := 0; i < len(boardStrings); i++ {
		b := boardStrings[i : i+5]
		boards = append(boards, MakeBoard(5, b[0], b[1], b[2], b[3], b[4]))
		i += 4
	}

	return numbers, boards
}
