// 4.6 seconds using int32

package main

import (
	"fmt"
)

func single_threaded() {

	total := int32(250_000)
	var foundPrimes []int32

	for i := int32(2); i <= total; i++ {
		if isNumberPrime(i) {
			foundPrimes = append(foundPrimes, i)
		}
	}
	fmt.Println("Found primes:", len(foundPrimes))	
}

func isNumberPrime(number int32) bool {
	for i := int32(2); i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
