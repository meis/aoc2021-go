package main

import (
	"strings"
	"testing"

	"github.com/meis/aoc2021-go/input"
)

var testInput = strings.Split(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`, "\n")

func Test_part1(t *testing.T) {
	type args struct {
		lines []Line
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test Input", args{parseInput(testInput)}, 5},
		{"Given Input", args{parseInput(input.GetInputStrings())}, 5774},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		lines []Line
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test Input", args{parseInput(testInput)}, 12},
		{"Given Input", args{parseInput(input.GetInputStrings())}, 18423},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.lines); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
