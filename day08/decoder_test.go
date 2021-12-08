package main

import (
	"reflect"
	"testing"
)

func Test_newDecoder(t *testing.T) {
	type args struct {
		patterns string
	}
	tests := []struct {
		name string
		args args
		want decoder
	}{
		{"EFG", args{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"}, decoder{'e', 'f', 'g'}},
		{"CFE", args{"gad cgdfab ag edacgb agbfedc facdb dcebaf adfbg agfc gbdef"}, decoder{'c', 'f', 'e'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDecoder(tt.args.patterns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decoder_decode(t *testing.T) {
	type args struct {
		digit string
	}
	tests := []struct {
		name   string
		fields decoder
		args   args
		want   int
	}{
		{"acedgfb: 8", newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"), args{"acedgfb"}, 8},
		{"cdfbe: 5", newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"), args{"cdfbe"}, 5},
		{"gcdfa: 2", newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"), args{"gcdfa"}, 2},
		{"fbcad: 3", newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"), args{"fbcad"}, 3},
		{"dab: 7", newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"), args{"dab"}, 7},
		{"cefabd: 9", newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"), args{"cefabd"}, 9},
		{"cdfgeb: 6", newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"), args{"cdfgeb"}, 6},
		{"eafb: 4", newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"), args{"eafb"}, 4},
		{"cagedb: 0", newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"), args{"cagedb"}, 0},
		{"ab: 1", newDecoder("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"), args{"ab"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				segmentB: tt.fields.segmentB,
				segmentD: tt.fields.segmentD,
				segmentE: tt.fields.segmentE,
			}
			if got := d.decode(tt.args.digit); got != tt.want {
				t.Errorf("decoder.decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
