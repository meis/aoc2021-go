package main

import (
	"fmt"
	"strings"
)

type Map [][]string

func NewMap(in []string) Map {
	m := make(Map, len(in))

	for y, row := range in {
		m[y] = make([]string, len(row))
		for x, col := range strings.Split(row, "") {
			m[y][x] = col
		}
	}

	return m
}

func (m Map) step() (Map, int) {
	moved := 0
	newMap1 := [][]string{}
	for _, col := range m {
		newCol := []string{}
		newCol = append(newCol, col...)
		newMap1 = append(newMap1, newCol)
	}

	for x, col := range m {
		for y, cucumber := range col {
			if cucumber == ">" {
				newY := y + 1
				if newY == len(col) {
					newY = 0
				}

				if m[x][newY] == "." {
					newMap1[x][y] = "."
					newMap1[x][newY] = ">"
					moved++
				}
			}
		}
	}
	newMap2 := [][]string{}
	for _, col := range newMap1 {
		newCol := []string{}
		newCol = append(newCol, col...)
		newMap2 = append(newMap2, newCol)
	}
	for x, col := range m {
		for y, cucumber := range col {
			if cucumber == "v" {
				newX := x + 1
				if newX == len(m) {
					newX = 0
				}
				if newMap1[newX][y] == "." {
					newMap2[x][y] = "."
					newMap2[newX][y] = "v"
					moved++
				}
			}
		}
	}

	return newMap2, moved
}

func (m *Map) print() {
	for _, col := range *m {
		for _, row := range col {
			fmt.Printf("%s", row)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}
