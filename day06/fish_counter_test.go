package main

import (
	"reflect"
	"testing"
)

func TestNewFishCounter(t *testing.T) {
	type args struct {
		fishes []int
	}
	tests := []struct {
		name string
		args args
		want FishCounter
	}{
		{"Initial values", args{[]int{3, 4, 3, 1, 2}}, FishCounter{0, 1, 1, 2, 1, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFishCounter(tt.args.fishes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFishCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFishCounter_Iterate(t *testing.T) {
	tests := []struct {
		name string
		c    *FishCounter
		want FishCounter
	}{
		{"Iteration 1", &FishCounter{0, 1, 1, 2, 1, 0, 0, 0, 0}, FishCounter{1, 1, 2, 1, 0, 0, 0, 0, 0}},
		{"Iteration 2", &FishCounter{1, 1, 2, 1, 0, 0, 0, 0, 0}, FishCounter{1, 2, 1, 0, 0, 0, 1, 0, 1}},
		{"Iteration 3", &FishCounter{1, 2, 1, 0, 0, 0, 1, 0, 1}, FishCounter{2, 1, 0, 0, 0, 1, 1, 1, 1}},
		{"Iteration 4", &FishCounter{2, 1, 0, 0, 0, 1, 1, 1, 1}, FishCounter{1, 0, 0, 0, 1, 1, 3, 1, 2}},
		{"Iteration 5", &FishCounter{1, 0, 0, 0, 1, 1, 3, 1, 2}, FishCounter{0, 0, 0, 1, 1, 3, 2, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Iterate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FishCounter.Iterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFishCounter_TotalFish(t *testing.T) {
	tests := []struct {
		name string
		c    *FishCounter
		want int
	}{
		{"Final values", &FishCounter{424, 729, 558, 790, 739, 762, 991, 370, 571}, 5934},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.TotalFish(); got != tt.want {
				t.Errorf("FishCounter.TotalFish() = %v, want %v", got, tt.want)
			}
		})
	}
}
