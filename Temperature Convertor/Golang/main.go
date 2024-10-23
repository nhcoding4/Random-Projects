package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/inancgumus/screen"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	for {

		// Display avalible options to the user.
		fmt.Println("\nOptions\n-------")
		fmt.Println("1) Convert Celsius to Fahrenheit")
		fmt.Println("2) Convert Fahrenheit to Celsius")
		fmt.Println("3) Exit Program")

		// Get user input.
		option := take_input("Enter a option:")

		switch option {

		case 1:
			// Convert celsius for fahrenheit.
			fmt.Println("=", celsius_fahrenheit(take_input("\nEnter a temperature in Celsius:")), "F")
			press_continue()
			screen_clear()

		case 2:
			// Convert fahrenheit to celsius.
			fmt.Println("=", fahrenheit_celsius(take_input("\nEnter a temperature in Fahrenheit:")), "C")
			press_continue()
			screen_clear()

		case 3:
			// Exit program.
			screen_clear()
			os.Exit(0)

		default:
			// Tell the user they did something stupid.
			fmt.Println("\nInvalid input")
			press_continue()
			screen_clear()
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func celsius_fahrenheit(temperature float64) float64 {
	// Conversion formuale, celsius to fahrenheit.

	return (temperature * 1.8) + 32.0
}

// --------------------------------------------------------------------------------------------------------------------

func fahrenheit_celsius(temperature float64) float64 {
	// Conversion formulae, fahrenheit to celsius.

	return (temperature - 32) / 1.8
}

// --------------------------------------------------------------------------------------------------------------------

func take_input(prompt string) float64 {
	// Keeps prompting use for input until a valid number has been entered.

	for {
		fmt.Println(prompt)

		// Create a scanner to read input.
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		// Attempt to convert input into a valid integer.
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Invalid input")
		} else {
			return number
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func press_continue() {
	// Pauses a program until the user presses the enter key.

	fmt.Println("\nPress 'Enter' to continue")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// --------------------------------------------------------------------------------------------------------------------

func screen_clear() {
	// Clears the terminal and moves the cursor to the top left of the terminal.

	screen.Clear()
	screen.MoveTopLeft()
}

// --------------------------------------------------------------------------------------------------------------------
