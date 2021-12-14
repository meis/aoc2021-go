package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/meis/aoc2021-go/input"
)

type polymer []rune
type pair struct {
	first  rune
	second rune
}
type rules map[pair]rune

func main() {
	in := input.GetInputStrings()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func part1(in []string) int {
	poly, rules := parseInput(in)

	return iterateNtimes(poly, rules, 10)
}

func part2(in []string) int {
	poly, rules := parseInput(in)

	return iterateNtimes(poly, rules, 40)
}

func iterateNtimes(poly polymer, rules rules, iterations int) int {
	pairs := make(map[pair]int)

	for i := 0; i < len(poly)-1; i++ {
		pairs[pair{poly[i], poly[i+1]}]++
	}

	for it := 0; it < iterations; it++ {
		newPairs := make(map[pair]int)
		for p, num := range pairs {
			next := rules[p]
			newPairs[pair{p.first, next}] += num
			newPairs[pair{next, p.second}] += num
		}

		pairs = newPairs
	}

	counter := make(map[rune]int)
	for p, num := range pairs {
		counter[p.first] += num
	}
	counter[poly[len(poly)-1]]++

	min := math.MaxInt
	max := 0
	for _, v := range counter {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}

func parseInput(in []string) (polymer, rules) {
	r := make(rules)
	var t polymer

	for _, line := range in {
		parts := strings.Split(line, " -> ")

		if len(parts) == 2 {
			first := rune(parts[0][0])
			second := rune(parts[0][1])
			r[pair{first, second}] = rune(parts[1][0])

		} else if len(line) > 1 {
			t = polymer(line)
		}
	}

	return t, r
}
