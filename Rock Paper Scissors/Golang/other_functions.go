// Other functions used

package main

import (
	"fmt"

	"github.com/inancgumus/screen"
)

// --------------------------------------------------------------------------------------------------------------------

func reset_screen() {
	// Clears the terminal

	screen.Clear()
	screen.MoveTopLeft()
}

// --------------------------------------------------------------------------------------------------------------------

func rules() {
	// Displays the rules of the game to the user

	reset_screen()

	fmt.Println("Rock beats Scissors and Lizard")
	fmt.Println("Paper beats Rock and Spock")
	fmt.Println("Scissors beats Paper and Lizard")
	fmt.Println("Lizard beats Spock and Paper")
	fmt.Println("Spock beats Scissors and Rock")

	fmt.Println("\nPress 'ENTER' to exit...")
	exit := ""
	fmt.Scanf("%v", &exit)

}

// --------------------------------------------------------------------------------------------------------------------


