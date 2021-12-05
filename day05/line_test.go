package main

import (
	"reflect"
	"testing"
)

func TestLine_GetPoints(t *testing.T) {
	type fields struct {
		start Point
		end   Point
	}
	tests := []struct {
		name   string
		fields fields
		want   []Point
	}{
		{
			"2,2 -> 2,1",
			fields{ParsePoint("2,2"), ParsePoint("2,1")},
			[]Point{ParsePoint("2,2"), ParsePoint("2,1")},
		},
		{
			"0,9 -> 3,9",
			fields{ParsePoint("0,9"), ParsePoint("3,9")},
			[]Point{ParsePoint("0,9"), ParsePoint("1,9"), ParsePoint("2,9"), ParsePoint("3,9")},
		},
		{
			"1,1 -> 3,3",
			fields{ParsePoint("1,1"), ParsePoint("3,3")},
			[]Point{ParsePoint("1,1"), ParsePoint("2,2"), ParsePoint("3,3")},
		},
		{
			"9,7 -> 7,9",
			fields{ParsePoint("9,7"), ParsePoint("7,9")},
			[]Point{ParsePoint("9,7"), ParsePoint("8,8"), ParsePoint("7,9")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Line{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := l.GetPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Line.GetPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_GetDirection(t *testing.T) {
	type fields struct {
		start Point
		end   Point
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  int
	}{
		{"1,1 -> 1,1", fields(ParseLine("1,1 -> 1,1")), 0, 0},
		{"1,1 -> 1,2", fields(ParseLine("1,1 -> 1,2")), 0, 1},
		{"1,1 -> 1,0", fields(ParseLine("1,1 -> 1,0")), 0, -1},
		{"1,1 -> 2,1", fields(ParseLine("1,1 -> 2,1")), 1, 0},
		{"1,1 -> 2,2", fields(ParseLine("1,1 -> 2,2")), 1, 1},
		{"1,1 -> 2,0", fields(ParseLine("1,1 -> 2,0")), 1, -1},
		{"1,1 -> 0,1", fields(ParseLine("1,1 -> 0,1")), -1, 0},
		{"1,1 -> 0,2", fields(ParseLine("1,1 -> 0,2")), -1, 1},
		{"1,1 -> 0,0", fields(ParseLine("1,1 -> 0,0")), -1, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Line{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			got, got1 := l.GetDirection()
			if got != tt.want {
				t.Errorf("Line.GetDirection() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Line.GetDirection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
