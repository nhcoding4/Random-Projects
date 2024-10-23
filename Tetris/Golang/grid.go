package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand/v2"
)

type Grid struct {
	cells             [][]int32
	cellSize          int32
	rows              int32
	columns           int32
	colors            []rl.Color
	blockData         Blocks
	currentBlock      block
	rotation          int32
	currentColour     int32
	xOffset           int32
	yOffset           int32
	totalLinesCleared int32
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) init() {
	g.populateGrid()
	g.blockData.init()
	g.selectBlock()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) blockCoordinates() [][]int32 {
	var positions [][]int32
	for _, cell := range g.currentBlock[g.rotation] {
		var cellLocations []int32
		cellLocations = append(cellLocations, cell[0]+g.xOffset)
		cellLocations = append(cellLocations, cell[1]+g.yOffset)

		positions = append(positions, cellLocations)
	}
	return positions
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) checkFilledRow() {
	var rowsToRemove []int32

	for i, row := range g.cells {
		fullRow := func() bool {
			for _, cell := range row {
				if cell == 0 {
					return false
				}
			}
			return true
		}()
		if fullRow {
			rowsToRemove = append(rowsToRemove, int32(i))
		}
	}

	if len(rowsToRemove) > 0 {
		g.makeNewGrid(rowsToRemove)
		g.totalLinesCleared += int32(len(rowsToRemove))
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) clearPrevious() {
	for _, cell := range g.currentBlock[g.rotation] {
		x := cell[0] + g.xOffset
		y := cell[1] + g.yOffset

		g.cells[y][x] = 0
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) checkSelfY(x, y int32, positions [][]int32) bool {
	for _, cells := range positions {
		if x+g.xOffset != cells[0] {
			continue
		}
		if y+g.yOffset+1 == cells[1] {
			return true
		}
	}
	return false
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) checkSelfX(x, y, direction int32, positions [][]int32) bool {
	for _, cells := range positions {
		if y+g.yOffset != cells[1] {
			continue
		}
		if x+g.xOffset+direction == cells[0] {
			return true
		}
	}
	return false
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) drawGrid() {
	for i, row := range g.cells {
		for j, cell := range row {
			rl.DrawRectangle(
				int32(j)*g.cellSize+1,
				int32(i)*g.cellSize+1,
				g.cellSize-1,
				g.cellSize-1,
				g.colors[cell],
			)
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) makeNewGrid(rowToRemove []int32) {
	var newCells [][]int32

	for range rowToRemove {
		var newRow []int32
		for range g.columns {
			newRow = append(newRow, 0)
		}
		newCells = append(newCells, newRow)
	}

	for i, row := range g.cells {
		toSkip := func(currentRow int32) bool {
			for _, toSkip := range rowToRemove {
				if currentRow == toSkip {
					return true
				}
			}
			return false
		}(int32(i))

		if toSkip {
			continue
		}
		newCells = append(newCells, row)
	}

	g.cells = newCells
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) moveBlock(direction string) {
	g.clearPrevious()
	switch direction {
	case "left":
		g.xOffset--
	case "right":
		g.xOffset++
	case "down":
		g.yOffset++
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) placeCurrentBlock() {
	for _, cell := range g.currentBlock[g.rotation] {
		x := cell[0] + g.xOffset
		y := cell[1] + g.yOffset

		g.cells[y][x] = g.currentColour
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) populateGrid() {
	for range g.rows {
		var newRow []int32
		for range g.columns {
			newRow = append(newRow, 0)
		}
		g.cells = append(g.cells, newRow)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) selectBlock() {
	g.xOffset = 0
	g.yOffset = 0
	selection := rand.IntN(len(g.blockData.blocks))
	g.currentBlock = g.blockData.blocks[selection]
	g.currentColour = int32(selection + 1)
	g.checkFilledRow()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Grid) updateGrid() {
	g.placeCurrentBlock()
}

// ---------------------------------------------------------------------------------------------------------------------
