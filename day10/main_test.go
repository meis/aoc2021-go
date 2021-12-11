package main

import (
	"testing"

	"github.com/meis/aoc2021-go/input"
)

var testInput = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
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
		{"Test Input", args{testInput}, 26397},
		{"Given Input", args{input.GetInputStrings()}, 411471},
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
		in []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test Input", args{testInput}, 288957},
		{"Given Input", args{input.GetInputStrings()}, 3122628974},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.in); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
