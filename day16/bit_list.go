package main

import (
	"strconv"
	"strings"
)

type bitList []bool

func parseBitList(in string) bitList {
	bits := bitList{}

	for _, c := range in {
		if c == '0' {
			bits = append(bits, false, false, false, false)
		} else if c == '1' {
			bits = append(bits, false, false, false, true)
		} else if c == '2' {
			bits = append(bits, false, false, true, false)
		} else if c == '3' {
			bits = append(bits, false, false, true, true)
		} else if c == '4' {
			bits = append(bits, false, true, false, false)
		} else if c == '5' {
			bits = append(bits, false, true, false, true)
		} else if c == '6' {
			bits = append(bits, false, true, true, false)
		} else if c == '7' {
			bits = append(bits, false, true, true, true)
		} else if c == '8' {
			bits = append(bits, true, false, false, false)
		} else if c == '9' {
			bits = append(bits, true, false, false, true)
		} else if c == 'A' {
			bits = append(bits, true, false, true, false)
		} else if c == 'B' {
			bits = append(bits, true, false, true, true)
		} else if c == 'C' {
			bits = append(bits, true, true, false, false)
		} else if c == 'D' {
			bits = append(bits, true, true, false, true)
		} else if c == 'E' {
			bits = append(bits, true, true, true, false)
		} else if c == 'F' {
			bits = append(bits, true, true, true, true)
		}
	}

	return bits
}

func (bits *bitList) intValue() int {
	asStrings := make([]string, len(*bits))
	for i, b := range *bits {
		if b {
			asStrings[i] = "1"
		} else {
			asStrings[i] = "0"
		}
	}
	value, _ := strconv.ParseInt(strings.Join(asStrings, ""), 2, 64)
	return int(value)
}
