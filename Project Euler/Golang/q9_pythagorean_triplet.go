package main

import (
	"fmt"
	"math"
)

func q9_pythagorean_triplet() {
	for a := 3; a < 1000; a++ {
		for b := a + 1; b < 999; b++ {
			csquared := math.Pow(float64(a), 2) + math.Pow(float64(b), 2)
			c := math.Pow(csquared, 0.5)

			if float64(a)+float64(b)+c == 1000 {
				fmt.Println(int(float64(a) * float64(b) * c))
				break
			}
		}
	}
}
