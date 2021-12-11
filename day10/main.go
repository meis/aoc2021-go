package main

import (
	"fmt"
	"sort"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func part1(in []string) int {
	score := 0

	for _, s := range in {
		ics, _ := getIllegalCharacterScore(s)
		score += ics
	}

	return score
}

func part2(in []string) int {
	var scores []int

	for _, s := range in {
		ics, opened := getIllegalCharacterScore(s)
		if ics == 0 {
			score := 0

			for i := len(opened) - 1; i >= 0; i-- {
				char := opened[i]

				score *= 5
				if char == '(' {
					score += 1
				} else if char == '[' {
					score += 2
				} else if char == '{' {
					score += 3
				} else if char == '<' {
					score += 4
				}
			}
			scores = append(scores, score)
		}
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	return scores[len(scores)/2]
}

func getIllegalCharacterScore(subsystem string) (int, stack) {
	opened := make(stack, 0)

	for _, char := range subsystem {
		if char == '(' {
			opened = opened.Push(char)
		} else if char == ')' {
			if opened.Top() == '(' {
				opened, _ = opened.Pop()
			} else {
				return 3, opened
			}

		} else if char == '[' {
			opened = opened.Push(char)
		} else if char == ']' {
			if opened.Top() == '[' {
				opened, _ = opened.Pop()
			} else {
				return 57, opened
			}

		} else if char == '{' {
			opened = opened.Push(char)
		} else if char == '}' {
			if opened.Top() == '{' {
				opened, _ = opened.Pop()
			} else {
				return 1197, opened
			}

		} else if char == '<' {
			opened = opened.Push(char)
		} else if char == '>' {
			if opened.Top() == '<' {
				opened, _ = opened.Pop()
			} else {
				return 25137, opened
			}
		}
	}

	return 0, opened
}
