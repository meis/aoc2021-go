package main

import (
	"testing"

	"github.com/meis/aoc2021-go/input"
)

func Test_part1(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"8A004A801A8002F478", args{"8A004A801A8002F478"}, 16},
		{"620080001611562C8802118E34", args{"620080001611562C8802118E34"}, 12},
		{"C0015000016115A2E0802F182340", args{"C0015000016115A2E0802F182340"}, 23},
		{"A0016C880162017C3686B18A3D4780", args{"A0016C880162017C3686B18A3D4780"}, 31},
		{"Given Input", args{input.GetInputString()}, 873},
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
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"C200B40A82", args{"C200B40A82"}, 3},
		{"04005AC33890", args{"04005AC33890"}, 54},
		{"880086C3E88112", args{"880086C3E88112"}, 7},
		{"CE00C43D881120", args{"CE00C43D881120"}, 9},
		{"D8005AC2A8F0", args{"D8005AC2A8F0"}, 1},
		{"F600BC2D8F", args{"F600BC2D8F"}, 0},
		{"9C005AC2F8F0", args{"9C005AC2F8F0"}, 0},
		{"9C0141080250320F1802104A08", args{"9C0141080250320F1802104A08"}, 1},
		{"Given Input", args{input.GetInputString()}, 402817863665},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.in); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
