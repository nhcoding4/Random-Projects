package main

import "fmt"

// --------------------------------------------------------------------------------------------------------------------

type board struct {
	data        [][]int
	board_state string
}

// --------------------------------------------------------------------------------------------------------------------

// Updates the state of the sodoku board using the data variable.

func (b *board) Update() {

	// Divider.
	dashes := func() string {

		var dash string
		length := len(b.data[0]) + (len(b.data[0]) / 3)

		for i := range length {
			if i == 0 {
				dash += " -"
			}
			dash += "---"
		}
		return dash
	}()

	// Create the string representing the board state.
	var current_state string

	for i, row := range b.data {
		if i%3 == 0 {
			current_state += dashes + "\n"
		}

		for j, number := range row {
			if j%3 == 0 {
				current_state += " | "
			}
			if number == 0 {
				current_state += " - "
			} else {
				current_state += fmt.Sprintf(" %v ", number)
			}
		}
		current_state += " |\n"
	}
	current_state += dashes
	b.board_state = current_state
}

// --------------------------------------------------------------------------------------------------------------------

// Updates and prints the board. A frontend function to cut down on repetition.

func (b *board) Print() {
	b.Update()
	fmt.Println(b.board_state)
}

// --------------------------------------------------------------------------------------------------------------------
