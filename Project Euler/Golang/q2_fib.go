package main

import "fmt"

// --------------------------------------------------------------------------------------------------------------------

func question_2_fibonacci() {
	// Find the sum of all even fibonacci numbers under 4MM

	const limit = 4_000_000

	// Calculate all fib numbers <= 4MM
	fib_seq := []int{1, 2}

	var x int
	for x <= limit {
		x = fib_seq[len(fib_seq)-1] + fib_seq[len(fib_seq)-2]
		fib_seq = append(fib_seq, x)
	}

	var even_total int
	// Sum all the even numbers
	for _, number := range fib_seq {
		if number%2 == 0 {
			even_total += number
		}
	}

	fmt.Println(even_total)
}

// --------------------------------------------------------------------------------------------------------------------
