// Functions used for taking inputs

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --------------------------------------------------------------------------------------------------------------------

func get_name(players int) []string {
	// Take user name and return it as part of a slice.

	var names []string

	for i := 0; i < players; i++ {
		var name string
		fmt.Printf("Player %v, enter your name: ", i+1)
		fmt.Scanf("%v", &name)
		names = append(names, name)
	}

	return names
}

// --------------------------------------------------------------------------------------------------------------------

func take_action(name string) Choice {

	// User and computer actions.
	fmt.Println(name, " - Select an option:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v) %v\n", i+1, Choice(i))
	}

	return take_input(total_options)
}

// --------------------------------------------------------------------------------------------------------------------

func take_input(limit int) Choice {
	// Take input from the user.

	const minimum = 1

	// -----------

	for {

		// Create a reader and read input from the user.
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		// -----------

		// Check for errors and valid input
		number, err := strconv.Atoi(scanner.Text())
		if (err != nil) || (number < minimum || number > limit) {
			fmt.Println("Invalid input. Please enter a number between 1 and", limit)
			continue
		}

		// -----------

		return Choice(number - minimum)
	}
}

// --------------------------------------------------------------------------------------------------------------------
