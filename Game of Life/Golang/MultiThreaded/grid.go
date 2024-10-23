package main

import (
	"math/rand"
)

// ---------------------------------------------------------------------------------------------------------------------

type Pair struct {
	x int32
	y int32
}

var offset = []Pair{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

// ---------------------------------------------------------------------------------------------------------------------

type Grid struct {
	rows, columns int32
	grid          [][]int32
	ch            chan *[][]int32
	channelSize   int32
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) Init() {
	g.grid = g.newGrid()
	g.setStartingStatus()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) calculateAndSend() {
	for {
		if int32(len(g.ch)) < g.channelSize {
			newGrid := g.newGrid()
			g.calculateStatus(newGrid)
			g.grid = newGrid
			g.ch <- &g.grid
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) calculateStatus(grid [][]int32) {
	for i, row := range g.grid {
		for j, column := range row {
			live := g.countLiveNeighbors(int32(j), int32(i))
			if live == 3 {
				grid[i][j] = 1
			} else if live == 2 {
				grid[i][j] = column
			} else {
				grid[i][j] = 0
			}
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) countLiveNeighbors(x, y int32) int32 {
	var liveNeighbors int32
	for _, pair := range offset {
		neighborRow := (y + pair.y + g.rows) % g.rows
		neighborColumn := (x + pair.x + g.columns) % g.columns
		liveNeighbors += g.grid[neighborRow][neighborColumn]
	}
	return liveNeighbors
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) newGrid() [][]int32 {
	newGrid := make([][]int32, g.rows)
	for i := range g.rows {
		column := make([]int32, g.columns)
		for j := range g.columns {
			column[j] = 0
		}
		newGrid[i] = column
	}
	return newGrid
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) setStartingStatus() {
	for i, row := range g.grid {
		for j, _ := range row {
			choice := rand.Intn(10)
			if choice <= 2 {
				g.grid[i][j] = 1
			}
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------
