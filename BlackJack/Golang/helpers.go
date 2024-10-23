package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------------------------------------------------------------------

// Unassociated helper functions.

// ---------------------------------------------------------------------------------------------------------------------

func takeInputInt(prompt string, min, max int) int {
	fmt.Println(prompt)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		converted, err := strconv.Atoi(text)

		if err != nil || (converted < min || converted > max) {
			fmt.Printf("Invalid input. Please enter an integer between %v and %v\n", min, max)
			continue
		}
		return converted
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func takeInputString(prompt string) string {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}

// ---------------------------------------------------------------------------------------------------------------------

func takeInputFloat(prompt string, min, max float64) float64 {
	fmt.Println(prompt)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		converted, err := strconv.ParseFloat(text, 64)

		if err != nil || (converted < min || converted > max) {
			fmt.Printf("Invalid input. Please enter an integer between %v and %v\n", min, max)
			continue
		}
		return converted
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func printSpacer() {
	fmt.Printf("-------------------------------------------------------------------------\n\n")
}

// ---------------------------------------------------------------------------------------------------------------------
