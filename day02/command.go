package main

import (
	"strconv"
	"strings"
)

type command struct {
	direction string
	units     int
}

func newCommand(s string) command {
	components := strings.Split(s, " ")

	direction := components[0]
	distance, err := strconv.Atoi(components[1])
	if err != nil {
		panic(err)
	}

	return command{
		direction: direction,
		units:     distance,
	}
}
