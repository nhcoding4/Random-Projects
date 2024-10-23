package main

import (
	"fmt"
	"math/big"
)

func q7_find_prime() {
	// Uses a probablistic approach to find X prime number.

	const limit = 10_001 // The amount of primes we want to find
	const checks = 20 // The amount of checks ran against that number to decide if it is prime.

	// Use big int to allow for BIG prime numbers to be found
	primes_found := 0
	number := big.NewInt(1)

	for primes_found < limit {
		
		// Passes check, print and increase the primes found.
		if number.ProbablyPrime(checks) {
			primes_found++
			fmt.Printf("Prime %v: %v\n", primes_found, number)
		}

		// Add 1 to the number we are working with.
		to_add := big.NewInt(1)
		number.Add(number, to_add)
	}
}
