package main

import (
	"strconv"
	"strings"

	"github.com/meis/aoc2021-go/input"
)

type Board struct {
	size  int
	tiles [][]BoardTile
}

type BoardTile struct {
	value int
	drawn bool
}

func MakeBoard(size int, r1 string, r2 string, r3 string, r4 string, r5 string) Board {
	return Board{size, [][]BoardTile{
		parseRow(size, r1),
		parseRow(size, r2),
		parseRow(size, r3),
		parseRow(size, r4),
		parseRow(size, r5),
	}}
}

func (b *Board) Draw(number int) {
	for i := range b.tiles {
		for j := range b.tiles[i] {
			if b.tiles[i][j].value == number {
				b.tiles[i][j].drawn = true
			}
		}
	}
}

func (b *Board) Finished() bool {
	for _, row := range b.tiles {
		if allDrawn(row) {
			return true
		}
	}

	for i := range b.tiles {
		col := make([]BoardTile, b.size)
		for j := 0; j < b.size; j++ {
			col[j] = b.tiles[j][i]
		}
		if allDrawn(col) {
			return true
		}
	}

	return false
}

func (b *Board) LeftToDraw() []int {
	var toDraw []int

	for i := range b.tiles {
		for j := range b.tiles[i] {
			if !b.tiles[i][j].drawn {
				toDraw = append(toDraw, b.tiles[i][j].value)
			}
		}
	}

	return toDraw
}

func parseRow(size int, in string) []BoardTile {
	row := make([]BoardTile, size)

	for i, s := range input.CleanEmpty(strings.Split(in, " ")) {
		n, _ := strconv.Atoi(s)

		row[i] = BoardTile{n, false}
	}

	return row
}

func allDrawn(tiles []BoardTile) bool {
	for _, tile := range tiles {
		if !tile.drawn {
			return false
		}
	}
	return true
}
