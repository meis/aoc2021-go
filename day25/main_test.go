package main

import (
	"testing"

	"github.com/meis/aoc2021-go/input"
)

var testInput = []string{
	"v...>>.vv>",
	".vv>>.vv..",
	">>.>v>...v",
	">>v>>.>.v.",
	"v>v.vv.v..",
	">.>>..v...",
	".vv..>.>v.",
	"v.v..>>v.v",
	"....v..v.>",
}

func Test_part1(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test Input", args{testInput}, 58},
		{"Given Input", args{input.GetInputStrings()}, 412},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.in); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
