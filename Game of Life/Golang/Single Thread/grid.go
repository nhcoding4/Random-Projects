package main

import rl "github.com/gen2brain/raylib-go/raylib"

// --------------------------------------------------------------------------------------------------------------------

type Grid struct {
	rows      int
	columns   int
	cell_size int
	cells     [][]int
}

// --------------------------------------------------------------------------------------------------------------------

func (g *Grid) Calculate_Size(height int, width int) {
	// Creates a grid heigth and width.

	g.rows = height / g.cell_size
	g.columns = width / g.cell_size
}

// --------------------------------------------------------------------------------------------------------------------

func (g *Grid) Populate() {
	// Populates grid with cells.

	for i := 0; i < g.rows; i++ {
		new_row := []int{}
		for j := 0; j < g.columns; j++ {
			new_row = append(new_row, 0)
		}
		g.cells = append(g.cells, new_row)
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (g *Grid) Draw() {
	// Draws the grid on the screen depending on the value of the cell.

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {
			// Choose a colour and draw rectangle.
			colour := rl.White
			if g.cells[i][j] == 0 {
				colour = rl.Black
			}
			rl.DrawRectangle(int32(j*g.cell_size)+1, int32(i*g.cell_size)+1, int32(g.cell_size-1), int32(g.cell_size-1), colour)
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (g *Grid) Set_Starting_Value() {
	// Sets a random starting value for all the cells in the grid. 1 / 5 chance by default.

	for i, row := range g.cells {
		for j := range row {
			random_value := int(rl.GetRandomValue(0, 5))
			if random_value == 5 {
				g.cells[i][j] = 1
			} else {
				g.cells[i][j] = 0
			}
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (g *Grid) Clear() {
	// Sets all the values in the grid to 0.

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {
			g.cells[i][j] = 0
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------
