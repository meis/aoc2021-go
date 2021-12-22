package main

import (
	"reflect"
	"testing"
)

func TestCuboid_Remove(t *testing.T) {
	type fields struct {
		FromX int
		ToX   int
		FromY int
		ToY   int
		FromZ int
		ToZ   int
	}
	type args struct {
		toRemove Cuboid
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Cuboid
	}{
		{
			"9, 11 minus 10, 10",
			fields{9, 11, 9, 11, 9, 11},
			args{Cuboid{10, 10, 10, 10, 10, 10}},
			[]Cuboid{
				{9, 9, 9, 10, 10, 10},
				{9, 10, 11, 11, 10, 10},
				{11, 11, 10, 11, 10, 10},
				{10, 11, 9, 9, 10, 10},
				{9, 11, 9, 11, 9, 9},
				{9, 11, 9, 11, 11, 11},
			},
		}, {
			"10, 12 minus 11, 13",
			fields{10, 12, 10, 12, 10, 12},
			args{Cuboid{11, 13, 11, 13, 11, 13}},
			[]Cuboid{
				{10, 10, 10, 12, 11, 12},
				{11, 12, 10, 10, 11, 12},
				{10, 12, 10, 12, 10, 10},
			},
		}, {
			"11, 13 minus 10, 12",
			fields{11, 13, 11, 13, 11, 13},
			args{Cuboid{10, 12, 10, 12, 10, 12}},
			[]Cuboid{
				{11, 12, 13, 13, 11, 12},
				{13, 13, 11, 13, 11, 12},
				{11, 13, 11, 13, 13, 13},
			},
		}, {
			"(2, 5)(0,3) minus (0,3)(2,4)",
			fields{2, 5, 0, 3, 1, 1},
			args{Cuboid{0, 3, 2, 4, 1, 1}},
			[]Cuboid{
				{4, 5, 2, 3, 1, 1},
				{2, 5, 0, 1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cuboid{
				FromX: tt.fields.FromX,
				ToX:   tt.fields.ToX,
				FromY: tt.fields.FromY,
				ToY:   tt.fields.ToY,
				FromZ: tt.fields.FromZ,
				ToZ:   tt.fields.ToZ,
			}
			if got := c.Remove(tt.args.toRemove); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cuboid.Remove() \n   got %v\n, want %v", got, tt.want)
			}
		})
	}
}
