package main

import (
	"sort"
	"strings"
)

// To decode a digit, we really only need:
// - Number of segments for {1, 4, 7, 8}
// - Segments B and E to disambiguate {2, 3, 5}
// - Segments D and E to disambiguate {0, 6, 9}
//
// Segments B and E are easy to spot in patterns because they will
// appear 6 and 4 times. No other signals will appear that number of times.
// There are two segments (D and G) that appear 6 times, to determine
// which is D, we look at the third-shortest pattern (number 4) which
// contains D but not G.

type decoder struct {
	segmentB rune
	segmentD rune
	segmentE rune
}

func newDecoder(in string) decoder {
	signals := make(map[rune]int)
	var segmentB, segmentD, segmentE rune

	patterns := strings.Split(in, " ")
	sort.Slice(patterns, func(i, j int) bool {
		return len(patterns[i]) < len(patterns[j])
	})

	for _, signal := range strings.Join(patterns, "") {
		signals[signal]++
	}

	for k, v := range signals {
		if v == 6 {
			segmentB = k
		} else if v == 7 && strings.ContainsRune(patterns[2], k) {
			segmentD = k
		} else if v == 4 {
			segmentE = k
		}
	}

	return decoder{segmentB, segmentD, segmentE}
}

func (d *decoder) decode(digit string) int {
	var num int

	if len(digit) == 2 {
		num = 1
	} else if len(digit) == 3 {
		num = 7
	} else if len(digit) == 4 {
		num = 4
	} else if len(digit) == 5 {
		if strings.ContainsRune(digit, d.segmentB) {
			num = 5
		} else if strings.ContainsRune(digit, d.segmentE) {
			num = 2
		} else {
			num = 3
		}
	} else if len(digit) == 6 {
		if !strings.ContainsRune(digit, d.segmentD) {
			num = 0
		} else if strings.ContainsRune(digit, d.segmentE) {
			num = 6
		} else {
			num = 9
		}
	} else {
		num = 8
	}

	return num
}
