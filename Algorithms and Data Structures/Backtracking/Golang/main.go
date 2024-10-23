package main

// --------------------------------------------------------------------------------------------------------------------

func main() {
	sodoku_board := board{data: test_data()}
	sodoku_board.Print()
	solve(&sodoku_board)
	sodoku_board.Print()
}

// --------------------------------------------------------------------------------------------------------------------

// A recursive function that uses backtracking to solve a sodoku board.

func solve(sodoku_board *board) bool {

	// Base case. No more empty cells in the board.
	empty_cell := find_empty_cell(sodoku_board)
	if empty_cell == nil {
		return true
	}

	// Solve the board. Undo changes that lead to an invalid state using backtracking.
	for i := 1; i <= len(sodoku_board.data); i++ {
		if valid(sodoku_board, i, empty_cell) {
			sodoku_board.data[empty_cell[0]][empty_cell[1]] = i
			if solve(sodoku_board) {
				return true
			}
			sodoku_board.data[empty_cell[0]][empty_cell[1]] = 0
		}
	}
	return false
}

// --------------------------------------------------------------------------------------------------------------------

// Find the row / column co-ordinate for an empty cell.

func find_empty_cell(sodoku_board *board) []int {
	for i, row := range sodoku_board.data {
		for j, number := range row {
			if number == 0 {
				return []int{i, j}
			}
		}
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// Checks for valid placement of the current number. - Tested as working on for placing a number in [0][2] and [0][3]

func valid(sodoku_board *board, number int, positions []int) bool {

	// Check if a number exists on the X axis.
	for i := range sodoku_board.data[0] {
		if sodoku_board.data[positions[0]][i] == number && positions[1] != i {
			return false
		}
	}

	// Check if a number exists on the Y axis.
	for i := range sodoku_board.data {
		if sodoku_board.data[i][positions[1]] == number && positions[0] != i {
			return false
		}
	}

	// Check the 3 x 3 grid.
	box_x := int(positions[1] / 3)
	box_y := int(positions[0] / 3)

	for i := box_y * 3; i < box_y+3; i++ {
		for j := box_x * 3; j < box_x+3; j++ {
			if (sodoku_board.data[i][j] == number) && (i != positions[0] && j != positions[1]) {
				return false
			}
		}
	}
	return true
}

// --------------------------------------------------------------------------------------------------------------------
