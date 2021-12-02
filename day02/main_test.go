package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		data []command
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Given Input", args{getInput()}, 2322630},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.data); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		data []command
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Given Input", args{getInput()}, 2105273490},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.data); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
