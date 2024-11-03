package main

import (
	"fmt"
	"math"
)

func get_primes(limit int) []bool {
	is_prime := make([]bool, limit)

	for i := 0; i < limit; i++ {
		is_prime[i] = true
	}

	is_prime[0] = false
	is_prime[1] = false
	for i := 4; i < limit; i+=2 {
		is_prime[i] = false;
	}
	for i := 3; i * i <= limit; i+=2 {
		if is_prime[i] {
			for j := i * i; j < limit; j += i {
				is_prime[j] = false
			}
		}
	}
	return is_prime
}

func is_perfect_square(n int) bool {
	root := int(math.Sqrt(float64(n)))
	return root * root == n
}

func main() {
	limit := int(1e4)
	is_prime := get_primes(limit)
	primes := []int {}

	for n, prime := range is_prime {
		if prime {
			primes = append(primes, n)
		}
	}

	for i := 33; i < limit; i+=2 {
		if is_prime[i] {
			continue
		}
		exists := false
		for j := 0; j < len(primes) && primes[j] < i; j++ {
			if (i - primes[j]) % 2 == 0 && is_perfect_square((i - primes[j]) / 2) {
				exists = true
			}
		}
		if !exists {
			fmt.Println(i)
			break
		}
	}
}
