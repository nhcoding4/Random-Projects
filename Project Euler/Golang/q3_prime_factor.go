package main

import "fmt"

// --------------------------------------------------------------------------------------------------------------------

func question_3_prime_factor() {

	factors := []int{}
	prime := 2
	n := 600851475143

	for n > 1 {
		for n%prime == 0 {
			factors = append(factors, prime)
			n /= prime
		}
		prime++
		if prime*prime > n {
			if n > 1 {
				factors = append(factors, n)
				break
			}
		}
	}

	var largest_factor int
	for i := 0; i < len(factors); i++ {
		if factors[i] > largest_factor {
			largest_factor = factors[i]
		}
	}

	fmt.Println(largest_factor)
}

// --------------------------------------------------------------------------------------------------------------------
