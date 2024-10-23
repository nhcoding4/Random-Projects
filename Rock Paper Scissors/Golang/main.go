// Contains types, type functions and main function responsible for calling other functions.

package main

import (
	"fmt"
	"os"
)

// --------------------------------------------------------------------------------------------------------------------

// Custom type.
type Choice int

// -----------

// Enum of custom type.
const (
	Rock Choice = iota
	Paper
	Scissors
	Lizard
	Spock
)

// -----------

func (player_choice Choice) String() string {
	// Returns a string related the number of choice made.

	names := [...]string{
		"Rock",
		"Paper",
		"Scissors",
		"Lizard",
		"Spock",
	}

	return names[player_choice]
}

// -----------

func (player_choice Choice) Beats() []Choice {
	// Returns a slice of integers which contains the number of the enums a choice wins against.

	wins := []Choice{}

	// Return the choices X choice wins against.
	switch player_choice {
	case 0:
		wins = append(wins, 2, 3)
	case 1:
		wins = append(wins, 0, 4)
	case 2:
		wins = append(wins, 1, 3)
	case 3:
		wins = append(wins, 1, 4)
	case 4:
		wins = append(wins, 0, 2)
	}

	return wins
}

// -----------

func (player_choice Choice) Check_Winner(opponent Choice) bool {
	// Check if player X beats player Y.

	beats := player_choice.Beats()

	for _, value := range beats {
		if opponent == value {
			return true
		}
	}
	return false
}

// --------------------------------------------------------------------------------------------------------------------

func main() {

	for {

		const menu_items = 4

		reset_screen()

		// ----------

		// Menu
		fmt.Println("Select an option:")
		fmt.Println("1) Human vs Computer")
		fmt.Println("2) Human vs Human")
		fmt.Println("3) Rules")
		fmt.Println("4) Quit")

		// ----------

		// User selection
		user_choice := take_input(menu_items)
		switch user_choice {

		case 0:
			human_vs_computer_menu()

		case 1:
			human_vs_human()

		case 2:
			rules()

		case 3:
			reset_screen()
			os.Exit(1)
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------
