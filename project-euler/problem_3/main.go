package main

import (
	"fmt"
	"math"
)

func main() {
	n := 600851475143
	var prime_factors []int

	for n % 2 == 0 {
		prime_factors = append(prime_factors, 2)
		n = n / 2
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		for n % i == 0 {
			prime_factors = append(prime_factors, i)
			n = n / i
		}
	}

	if n > 2 {
		prime_factors = append(prime_factors, n)
	}

	fmt.Println(prime_factors)
}