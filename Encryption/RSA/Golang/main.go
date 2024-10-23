// First try. Barley functional. Has some massive bugs like randomly producing negative cipher values. 

package main

import (
	"fmt"
	"math"
	"math/rand"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Generate p and q values
	p, q := generate_prime(10, 100)

	// Value of n
	n := p * q

	// Calculate the phi of p and q.
	phi_n := (p - 1) * (q - 1)

	// Find the value of e where the GCD of e and phi_n is 1.
	e := calculate_e(phi_n)

	// Find the value for d by getting the mod inverse.
	d := mod_inverse(e, phi_n)

	fmt.Println("Public Key:", e)
	fmt.Println("Private Key:", d)
	fmt.Println("N value:", n)
	fmt.Println("phi_n:", phi_n)
	fmt.Println("Q value:", q)
	fmt.Println("P value:", p)

	fmt.Println(cipher(e, n, "Hello, world!"))
}

// --------------------------------------------------------------------------------------------------------------------

func cipher(e, n int, message string) []int {
	// Encrypts a message.

	encrypted_string := []int{}

	for _, char := range message {
		encrypted_value := int(math.Pow(float64(char), float64(e))) % n
		encrypted_string = append(encrypted_string, encrypted_value)

	}
	return encrypted_string
}

// --------------------------------------------------------------------------------------------------------------------

func calculate_e(phi_n int) int {
	// Attempts to calculate the e value of phi_n.

	// Find the gcd == 1 for i and phi_n.
	i := 3

	for {
		gcd_value := gcd(i, phi_n)

		if gcd_value == 1 {
			return i
		}
		i++
	}
}

// --------------------------------------------------------------------------------------------------------------------

func gcd(x, y int) int {
	// Euclidean Algorithm to find the greatest common divisor.

	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// --------------------------------------------------------------------------------------------------------------------

func generate_prime(min, max int) (int, int) {
	// Sieve of Eratosthenes to generate prime numbers.

	// Generate an array of bools between 0 and
	primes := []bool{}

	// Populate the primes slice.
	for i := 0; i <= max; i++ {
		if i == 0 || i == 1 {
			primes = append(primes, false)
		} else {
			primes = append(primes, true)
		}
	}

	// Mark multiples of primes from 2 until the square root of the upper bound.
	for i := 2; i <= int(math.Sqrt(float64(max))); i++ {
		if primes[i] {
			for j := int(math.Pow(float64(i), 2)); j < max+1; j += i {
				primes[j] = false
			}
		}
	}

	// Convert the slice of bools to a slice of integers containing the prime numbers.
	prime_numbers := []int{}

	// Add primes between our min and max limits.
	for i := min; i <= max; i++ {
		if primes[i] {
			prime_numbers = append(prime_numbers, i)
		}
	}

	// Select 2 random unique prime numbers from the slice and return them.
	rand_index_1 := rand.Intn(len(prime_numbers))
	rand_index_2 := rand.Intn(len(prime_numbers))

	for rand_index_1 == rand_index_2 {
		rand_index_2 = rand.Intn(len(prime_numbers))
	}

	return prime_numbers[rand_index_1], prime_numbers[rand_index_2]
}

// --------------------------------------------------------------------------------------------------------------------

func mod_inverse(e, phi int) int {
	// Finds the mod inverse of e * d and phi.

	for d := 3; d < phi; d++ {
		if (d*e)%phi == 1 {
			return d
		}
	}
	return -1
}

// --------------------------------------------------------------------------------------------------------------------
