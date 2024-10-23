package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Take names.
	name_1, err := take_input("Enter the first name:")
	if err != nil {
		log.Fatal(err)
	}

	name_2, err := take_input("Enter the second name:")
	if err != nil {
		log.Fatal(err)
	}

	// Count the frequency each char appears.
	completed_string := fmt.Sprintf("%vloves%v", name_1, name_2)
	char_count := make(map[string]int)

	for _, char := range completed_string {
		// Ignore non a-z chars
		c := strings.ToLower(string(char))

		if rune(c[0]) < 'a' || rune(c[0]) > 'z' {
			continue
		}
		// Add the character into the map.
		_, ok := char_count[c]
		if ok {
			char_count[c]++
		} else {
			char_count[c] = 1
		}
	}

	// Get all the values in a nice slice of integers
	integer_slice := []int{}
	for _, value := range char_count {
		integer_slice = append(integer_slice, value)
	}

	// Pair off the values until there are only 2 remaining.
	for len(integer_slice) > 2 {
		i := 0
		for i < len(integer_slice) && len(integer_slice) > 2 {
			integer_slice[i] = integer_slice[i] + integer_slice[len(integer_slice)-1]
			integer_slice = integer_slice[:len(integer_slice)-1]
			i++
		}
	}

	fmt.Println(integer_slice)
}

// --------------------------------------------------------------------------------------------------------------------

func take_input(prompt string) (string, error) {
	// Take input.

	// Prompt user for input.
	fmt.Println(prompt)

	// Create a new scanner object to look for input.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}

	// Return the input.
	return scanner.Text(), nil
}

// --------------------------------------------------------------------------------------------------------------------
