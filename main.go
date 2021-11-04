package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/atomicgo/cursor"
)

type Vector2Int struct {
	X int
	Y int
}

const (
	underpopulation = 1
	overpopulation  = 4
	birthpopulation = 3
)

var (
	size     Vector2Int
	cells    [][]bool
	timeStep int

	neighbours = []Vector2Int{
		{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1},
		{X: -1, Y: 0}, {X: 1, Y: 0},
		{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
	}
)

func main() {
	fmt.Print("Choose size [x y]> ")
	fmt.Scanln(&size.X, &size.Y)
	cells = make([][]bool, size.Y)
	for i := range cells {
		cells[i] = make([]bool, size.X)
	}
	fmt.Print("Choose timestep [ms]> ")
	fmt.Scanln(&timeStep)

	if size.X == 0 || size.Y == 0 {
		size = Vector2Int{X: 9, Y: 9}
		cells = make([][]bool, size.Y)
		for i := range cells {
			cells[i] = make([]bool, size.X)
		}
		generateGlider(&cells)
	} else {
		generateCells(&cells)
	}
	showCells(cells, false)
	for {
		time.Sleep(time.Duration(timeStep) * time.Millisecond)
		cells = updateCells(cells, size)
		showCells(cells, true)
	}
}

func updateCells(cells [][]bool, size Vector2Int) [][]bool {
	newCells := make([][]bool, size.Y)
	for i := range newCells {
		newCells[i] = make([]bool, size.X)
	}
	for y, row := range cells {
		for x := range row {
			newCells[y][x] = updateCell(Vector2Int{X: x, Y: y}, cells)
		}
	}
	return newCells
}

func updateCell(pos Vector2Int, cells [][]bool) bool {
	neighbourCount := 0
	size := Vector2Int{X: len(cells[0]), Y: len(cells)}
	for _, neighbour := range neighbours {
		checkPos := Vector2Int{
			((pos.X + neighbour.X) + size.X) % size.X,
			((pos.Y + neighbour.Y) + size.Y) % size.Y,
		}
		if cells[checkPos.Y][checkPos.X] {
			neighbourCount++
		}
	}
	if cells[pos.Y][pos.X] {
		return underpopulation < neighbourCount && neighbourCount < overpopulation
	} else {
		return neighbourCount == birthpopulation
	}
}

func showCells(cells [][]bool, overwrite bool) {
	if overwrite {
		cursor.Up(size.Y)
	}
	str := ""
	for _, row := range cells {
		for _, cell := range row {
			if cell {
				str += "■ "
			} else {
				str += "  "
			}
		}
		str += "\n\r"
	}
	print(str)
}

func generateCells(cells *[][]bool) {
	c := *cells
	for y, row := range c {
		for x := range row {
			c[y][x] = rand.Intn(2) == 1
		}
	}
	cells = &c
}

func generateGlider(cells *[][]bool) {
	c := *cells
	for y, row := range c {
		for x := range row {
			c[y][x] = false
		}
	}
	c[3][5] = true
	c[4][5] = true
	c[5][5] = true
	c[5][4] = true
	c[4][3] = true
	cells = &c
}
