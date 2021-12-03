package main

import (
	"reflect"
	"testing"
)

func Test_signal_intValue(t *testing.T) {
	tests := []struct {
		name string
		s    signal
		want int
	}{
		{"00000", parseSignal("00000"), 0},
		{"10110", parseSignal("10110"), 22},
		{"01001", parseSignal("01001"), 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.intValue(); got != tt.want {
				t.Errorf("signal.intValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitByBitInPosition(t *testing.T) {
	type args struct {
		signals  []signal
		position int
	}
	tests := []struct {
		name                string
		args                args
		wantFalseInPosition []signal
		wantTrueInPosition  []signal
	}{
		{"0", args{[]signal{parseSignal("0"), parseSignal("1")}, 0}, []signal{parseSignal("0")}, []signal{parseSignal("1")}},
		{"0", args{[]signal{parseSignal("00"), parseSignal("10")}, 1}, []signal{parseSignal("00"), parseSignal("10")}, []signal{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFalseInPosition, gotTrueInPosition := splitByBitInPosition(tt.args.signals, tt.args.position)
			if !reflect.DeepEqual(gotFalseInPosition, tt.wantFalseInPosition) {
				t.Errorf("splitByBitInPosition() gotFalseInPosition = %v, want %v", gotFalseInPosition, tt.wantFalseInPosition)
			}
			if !reflect.DeepEqual(gotTrueInPosition, tt.wantTrueInPosition) {
				t.Errorf("splitByBitInPosition() gotTrueInPosition = %v, want %v", gotTrueInPosition, tt.wantTrueInPosition)
			}
		})
	}
}
