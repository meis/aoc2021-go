package main

import (
	"strconv"
	"strings"
)

type signal []bool

func newSignal(length int) signal {
	return signal(make([]bool, length))
}

func parseSignal(in string) signal {
	s := signal(make([]bool, len(in)))

	for i, b := range in {
		if b == '0' {
			s[i] = false
		} else {
			s[i] = true
		}
	}

	return s
}

func (s *signal) intValue() int {
	asStrings := make([]string, len(*s))
	for i, b := range *s {
		if b {
			asStrings[i] = "1"
		} else {
			asStrings[i] = "0"
		}
	}
	value, _ := strconv.ParseInt(strings.Join(asStrings, ""), 2, 64)
	return int(value)
}

func splitByBitInPosition(signals []signal, position int) (falseInPosition []signal, trueInPosition []signal) {
	falseInPosition = make([]signal, 0)
	trueInPosition = make([]signal, 0)

	for _, line := range signals {
		if line[position] {
			trueInPosition = append(trueInPosition, line)
		} else {
			falseInPosition = append(falseInPosition, line)
		}
	}

	return falseInPosition, trueInPosition
}
