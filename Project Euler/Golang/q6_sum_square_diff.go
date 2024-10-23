package main

import (
	"fmt"
	"math"
)

func q6_square_difference() {
	// Finds the difference between the square of each number between 1 and X (added together)
	// and the square of the sum of those numbers

	// The amout of numbers we want in our calculation
	const limit = 100

	squared_total := 0.0
	i_value_total := 0.0

	for i := 1; i <= limit; i++ {
		squared_total += math.Pow(float64(i), 2)
		i_value_total += float64(i)
	}

	fmt.Println(math.Pow(i_value_total, 2) - squared_total)
}
