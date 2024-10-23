package calculate

import (
	"fmt"
	"sort"
	"strings"
)

// --------------------------------------------------------------------------------------------------------------------

// Gets a score from 2 entered strings.
func Calculate(name_1, name_2 string) string {

	if len(name_1) <= 0 || len(name_2) <= 0 {
		return "0"
	}

	// ---------------------------------------------------------------

	// Create a set of characters and total their frequency.

	completed_string := fmt.Sprintf("%vloves%v", name_1, name_2)
	char_count := make(map[string]int)

	for _, char := range completed_string {

		c := strings.ToLower(string(char))

		if rune(c[0]) < 'a' || rune(c[0]) > 'z' {
			continue
		}
		_, ok := char_count[c]
		if ok {
			char_count[c]++
		} else {
			char_count[c] = 1
		}
	}

	// ---------------------------------------------------------------

	/*
		Scoring algorithm. Pairs off integers from front and back, sums them and reinserts them
		into the slice until only 2 integers remain. Return the 2 integers as a string.
	*/

	integer_slice := []int{}
	for _, value := range char_count {
		integer_slice = append(integer_slice, value)
	}

	sort.Ints(integer_slice) // Needed to stop same inputs producing different results.

	for len(integer_slice) > 2 {
		i := 0
		for i < len(integer_slice) && len(integer_slice) > 2 {
			integer_slice[i] = integer_slice[i] + integer_slice[len(integer_slice)-1]
			integer_slice = integer_slice[:len(integer_slice)-1]
			i++
		}
	}

	return fmt.Sprintf("%v%v", integer_slice[0], integer_slice[1])
}

// --------------------------------------------------------------------------------------------------------------------
