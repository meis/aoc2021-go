package main

import (
	"fmt"

	"github.com/meis/aoc2021-go/input"
)

func main() {
	in := getInput()

	fmt.Printf("Solution of part one: %d\n", part1(in))
	fmt.Printf("Solution of part two: %d\n", part2(in))
}

func getInput() []signal {
	var in []signal
	for _, c := range input.GetInputStrings() {
		in = append(in, parseSignal(c))
	}

	return in
}

func part1(signals []signal) int {
	signalLength := len(signals[0])
	var counter = make([]int, signalLength)

	for _, s := range signals {
		for i, b := range s {
			if b {
				counter[i]++
			} else {
				counter[i]--
			}
		}
	}

	var gammaRate = newSignal(signalLength)
	var epsilonRate = newSignal(signalLength)

	for i, c := range counter {
		if c > 1 {
			gammaRate[i] = true
			epsilonRate[i] = false
		} else {
			gammaRate[i] = false
			epsilonRate[i] = true
		}
	}

	return gammaRate.intValue() * epsilonRate.intValue()
}

func part2(signals []signal) int {
	signalLength := len(signals[0])

	oxygenGeneratorSignals := signals
	co2ScrubberSignals := signals

	for i := 0; i <= signalLength-1; i++ {
		falseInPosition, trueInPosition := splitByBitInPosition(oxygenGeneratorSignals, i)

		if len(trueInPosition) == 0 {
			oxygenGeneratorSignals = falseInPosition
		} else if len(falseInPosition) == 0 {
			oxygenGeneratorSignals = trueInPosition
		} else if len(falseInPosition) > len(trueInPosition) {
			oxygenGeneratorSignals = falseInPosition
		} else {
			oxygenGeneratorSignals = trueInPosition
		}
	}

	for i := 0; i <= signalLength-1; i++ {
		falseInPosition, trueInPosition := splitByBitInPosition(co2ScrubberSignals, i)

		if len(trueInPosition) == 0 {
			co2ScrubberSignals = falseInPosition
		} else if len(falseInPosition) == 0 {
			co2ScrubberSignals = trueInPosition
		} else if len(falseInPosition) > len(trueInPosition) {
			co2ScrubberSignals = trueInPosition
		} else {
			co2ScrubberSignals = falseInPosition
		}
	}

	return oxygenGeneratorSignals[0].intValue() * co2ScrubberSignals[0].intValue()
}
