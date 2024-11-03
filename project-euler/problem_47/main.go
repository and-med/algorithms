package main

import (
	"fmt"
	"math"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func get_primes(limit int) []bool {
	is_prime := make([]bool, limit)

	for i := 0; i < limit; i++ {
		is_prime[i] = true
	}

	is_prime[0] = false
	is_prime[1] = false
	for i := 4; i < limit; i += 2 {
		is_prime[i] = false
	}
	for i := 3; i*i <= limit; i += 2 {
		if is_prime[i] {
			for j := i * i; j < limit; j += i {
				is_prime[j] = false
			}
		}
	}
	return is_prime
}

func test_consecutive(is_prime []bool, n int, length int) bool {
	limit := int(math.Sqrt(float64(n + length - 1)))
	distinct_factors := map[int]int{}
	for j := 2; j <= limit; j++ {
		for i := n; i < n+length; i++ {
			if i%j == 0 {
				if is_prime[j] {
					distinct_factors[i]++
				}
				if is_prime[i/j] {
					distinct_factors[i]++
				}
			}
		}
	}

	if len(distinct_factors) != length {
		return false
	}
	for _, value := range distinct_factors {
		if value != length {
			return false
		}
	}
	return true
}

func main() {
	defer elapsed("main")()
	limit := int(1e6)
	length := 4
	is_prime := get_primes(limit)

	res := 2
	for !test_consecutive(is_prime, res, length) {
		res++
	}
	fmt.Println(res)
}
