package main

import (
	"testing"

	"github.com/meis/aoc2021-go/input"
)

var testInput = []int{3, 4, 3, 1, 2}

func Test_part1(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test Input", args{testInput}, 5934},
		{"Given Input", args{input.GetInputIntsInOneLine()}, 363101},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.in); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test Input", args{testInput}, 26984457539},
		{"Given Input", args{input.GetInputIntsInOneLine()}, 1644286074024},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.in); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
