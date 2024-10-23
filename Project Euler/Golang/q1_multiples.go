package main

import "fmt"

// --------------------------------------------------------------------------------------------------------------------

func question_1_multiples() {
	// All multiples of 3 or 5 below 1000

	const limit = 1000

	// Find all multiples of 3 and 5
	multiples := []int{}

	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			multiples = append(multiples, i)
		}
	}

	// Sum them
	var total int
	for _, number := range multiples {
		total += number
	}

	fmt.Println("The sum for all multiples of 3 and 5 below 1000 is:", total)
}

// --------------------------------------------------------------------------------------------------------------------
