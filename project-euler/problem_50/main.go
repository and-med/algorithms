package main

import (
	"fmt"
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
	for i := 4; i < limit; i+=2 {
		is_prime[i] = false;
	}
	for i := 3; i * i < limit; i+=2 {
		if is_prime[i] {
			for j := i * i; j < limit; j += i {
				is_prime[j] = false
			}
		}
	}
	return is_prime
}

func main() {
	defer elapsed("main")()
	limit := 1_000_000
	is_prime := get_primes(limit)
	primes := []int {}

	for n, p := range is_prime {
		if p {
			primes = append(primes, n)
		}
	}

	res := 0
	cpr := 0
	for i := 0; i < len(primes); i++ {
		s := primes[i]
		for j := i + 1; s < limit && j < len(primes); j++ {
			if is_prime[s] && (j - i) > cpr {
				cpr = j - i
				res = s
			}
			s += primes[j]
		}
	}
	fmt.Println(res, ":", cpr)
}