package main

import (
	"testing"

	"github.com/meis/aoc2021-go/input"
)

var testInput = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func Test_part1(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test Input", args{testInput}, 37},
		{"Given Input", args{input.GetInputIntsInOneLine()}, 343468},
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
		{"Test Input", args{testInput}, 168},
		{"Given Input", args{input.GetInputIntsInOneLine()}, 96086265},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.in); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleDistance(t *testing.T) {
	type args struct {
		orig int
		dest int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Move from 16 to 2: 14 fuel", args{16, 2}, 14},
		{"Move from 1 to 2: 1 fuel", args{1, 2}, 1},
		{"Move from 2 to 2: 0 fuel", args{2, 2}, 0},
		{"Move from 0 to 2: 2 fuel", args{0, 2}, 2},
		{"Move from 4 to 2: 2 fuel", args{4, 2}, 2},
		{"Move from 7 to 2: 5 fuel", args{7, 2}, 5},
		{"Move from 14 to 2: 12 fuel", args{14, 2}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleDistance(tt.args.orig, tt.args.dest); got != tt.want {
				t.Errorf("singleDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_triangularDistance(t *testing.T) {
	type args struct {
		orig int
		dest int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Move from 16 to 5: 66 fuel", args{16, 5}, 66},
		{"Move from 1 to 5: 10 fuel", args{1, 5}, 10},
		{"Move from 2 to 5: 6 fuel", args{2, 5}, 6},
		{"Move from 0 to 5: 15 fuel", args{0, 5}, 15},
		{"Move from 4 to 5: 1 fuel", args{4, 5}, 1},
		{"Move from 7 to 5: 3 fuel", args{7, 5}, 3},
		{"Move from 14 to 2: 45 fuel", args{14, 5}, 45},
		{"Move from 5 to 5: 0 fuel", args{5, 5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangularDistance(tt.args.orig, tt.args.dest); got != tt.want {
				t.Errorf("triangularDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
