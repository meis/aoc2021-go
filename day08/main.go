package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	data := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(data))
	fmt.Printf("Solution of part two: %d\n", part2(data))
}

func part1(in []string) int {
	entries := parseLines(in)
	counter := 0

	for _, entry := range entries {
		for _, digits := range entry.digits {
			if len(digits) == 2 || len(digits) == 3 || len(digits) == 4 || len(digits) == 7 {
				counter++
			}
		}
	}

	return counter
}

func part2(in []string) int {
	entries := parseLines(in)
	var counter int = 0

	for _, entry := range entries {
		decoder := newDecoder(entry.patterns)

		for i, digit := range entry.digits {
			num := decoder.decode(digit)
			counter += int(math.Pow(10, float64(-i+3))) * num
		}

	}

	return counter
}

type entry struct {
	patterns string
	digits   []string
}

func parseLines(lines []string) []entry {
	var entries []entry

	for _, e := range lines {
		parts := strings.Split(e, " | ")
		signals := parts[0]
		digits := strings.Split(parts[1], " ")
		entries = append(entries, entry{signals, digits})
	}

	return entries
}
