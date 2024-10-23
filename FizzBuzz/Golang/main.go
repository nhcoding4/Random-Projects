package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Take user input
	largest_number := take_input("Enter a positive number:")

	// Fizzbuzzing
	for i := 1; i <= largest_number; i++ {

		var print_string string

		if i%5 == 0 {
			if i%3 == 0 {
				print_string += "FizzBuzz"
			} else {
				print_string += "Buzz"
			}
		} else if i%3 == 0 {
			print_string += "Fizz"
		} else {
			print_string += fmt.Sprintf("%v", i)
		}
		fmt.Println(print_string)
	}
}

// --------------------------------------------------------------------------------------------------------------------

func take_input(prompt string) int {
	// Takes input from the user until they enter a valid integer.

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
		number, err := strconv.Atoi(scanner.Text())
		if err != nil || number <= 0 {
			fmt.Println("Invalid input. Please enter a positive whole number.")
		} else {
			return number
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------
