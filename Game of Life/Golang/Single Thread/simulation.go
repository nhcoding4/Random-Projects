package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Populate the offsets. Use a literal or it fries your computer.
var offset = []Pair{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

// --------------------------------------------------------------------------------------------------------------------

type Simulation struct {
	width     int
	height    int
	cell_size int
	grid      Grid
	run       bool
}

type Pair struct {
	x_value int
	y_value int
}

// --------------------------------------------------------------------------------------------------------------------

func (s *Simulation) Populate_Grid() {
	// A wrapper to help easily create a starting state for the simulation.

	s.grid.cell_size = s.cell_size
	s.grid.Calculate_Size(s.height, s.width)
	s.grid.Populate()
}

// --------------------------------------------------------------------------------------------------------------------

func (s *Simulation) Count_Live_Neighbors(row int, column int) int {
	// Counts the total live neighbors for a cell at X/Y.

	// Count the live neighbors by adding its value to the total (live = 1, everything else = 0).
	total_rows := s.grid.rows
	total_columns := s.grid.columns
	live_neighbors := 0

	// This sets up a grid that wraps around. The top connects to the bottom and left connects to right.
	for _, pair := range offset {
		neighbor_row := (row + pair.y_value + total_rows) % total_rows
		neighbor_column := (column + pair.x_value + total_columns) % total_columns
		live_neighbors += s.grid.cells[neighbor_row][neighbor_column]
	}
	return live_neighbors
}

// --------------------------------------------------------------------------------------------------------------------

func (s *Simulation) fps() {
	fps := rl.GetFPS()
	rl.DrawText(fmt.Sprintf("%v", fps), 0, 0, 40, rl.Green)
}

// --------------------------------------------------------------------------------------------------------------------

func (s *Simulation) Update() {
	// Updates the grid.

	if s.run {
		temp_grid := [][]int{}

		for i := 0; i < len(s.grid.cells); i++ {
			row := []int{}

			for j := 0; j < len(s.grid.cells[i]); j++ {
				// Set status of the cell for the next game tick.
				alive := s.Count_Live_Neighbors(i, j)
				if alive == 3 {
					row = append(row, 1)
				} else if alive == 2 {
					row = append(row, s.grid.cells[i][j])
				} else {
					row = append(row, 0)
				}
			}
			temp_grid = append(temp_grid, row)
		}
		// Change the game grid to have the same values of the newly calculated game state.
		s.grid.cells = temp_grid
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (s *Simulation) Start_Stop() {
	// Starts / stops the program.

	if s.run {
		s.run = false
	} else {
		s.run = true
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (s *Simulation) Mouse_Function() {
	// Allows the user to use the mouse to change the alive/dead status of the cells.

	// Convert the current mouse position into a valid grid coordinate.
	mouse_pos := rl.GetMousePosition()
	row := int(mouse_pos.Y) / s.cell_size
	column := int(mouse_pos.X) / s.cell_size

	// If the converted value is a valid coordinate then swap the cell status.
	if (row >= 0 && row < s.height) && (column >= 0 && column < s.width) {
		if s.grid.cells[row][column] == 0 {
			s.grid.cells[row][column] = 1
		} else {
			s.grid.cells[row][column] = 0
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (s *Simulation) Input_Actions(key int32, target_fps *int32) {

	switch key {

	// -------------------------------

	case rl.KeySpace:
		// Pause the simulation using the enter key.
		s.Start_Stop()
		if s.run {
			rl.SetWindowTitle("Game of Life is Running")
		} else {
			rl.SetWindowTitle("Game of Life is Paused")
		}

		// -------------------------------

	case rl.KeyLeft:
		// Decrease simulation speed.
		*target_fps -= 2
		rl.SetTargetFPS(*target_fps)

	case rl.KeyRight:
		// Increase simulation speed.
		*target_fps += 2
		rl.SetTargetFPS(*target_fps)

		// -------------------------------

	case rl.KeyR:
		// Reset the simulation.
		s.grid.Clear()
		s.grid.Set_Starting_Value()

	case rl.KeyC:
		// Clear the simulation grid.
		s.grid.Clear()

		// -------------------------------
	}
}

// --------------------------------------------------------------------------------------------------------------------
