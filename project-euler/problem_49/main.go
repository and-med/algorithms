package main

import (
	"fmt"
	"sort"
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
	for i := 3; i*i < limit; i += 2 {
		if is_prime[i] {
			for j := i * i; j < limit; j += i {
				is_prime[j] = false
			}
		}
	}
	return is_prime
}

func sort_digits(n int) int {
	digits := []int {}
	for n != 0 {
		digits = append(digits, n % 10)
		n /= 10
	}
	sort.Ints(digits)
	res := 0
	for _, digit := range digits {
		res *= 10
		res += digit
	}
	return res
}

func main() {
	defer elapsed("main")()
	limit := 10001
	is_prime := get_primes(limit)
	
	same_digit_primes := map [int][] int {}
	for n, p := range is_prime {
		if p && n > 1000 {
			s := sort_digits(n)
			same_digit_primes[s] = append(same_digit_primes[s], n)
		}
	}

	for _, v := range same_digit_primes {
		if len(v) >= 3 {
			for i := 0; i < len(v); i++ {
				for j := i + 1; j < len(v); j++ {
					for k := j + 1; k < len(v); k++ {
						if v[k] - v[j] == v[j] - v[i] {
							fmt.Printf("%d%d%d", v[i], v[j], v[k])
							fmt.Println()
						}
					}
				}
			}
		}
	}
}