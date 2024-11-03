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

func get_just_primes(is_prime []bool) []int {
	primes := []int {}
	for n, p := range is_prime {
		if p {
			primes = append(primes, n)
		}
	}
	return primes
}

func get_distinct_digits(n int) []int {
	digits := map[int]bool {}
	for n != 0 {
		digits[n % 10] = true
		n /= 10
	}
	unique_digits := []int {}
	for key := range digits {
		unique_digits = append(unique_digits, key)
	}
	return unique_digits
}

func replace_digit(n int, digit int, new_digit int) int {
	new_n := 0
	for i := 0; n != 0; i++ {
		if n % 10 == digit {
			new_n += new_digit * int(math.Pow10(i))
		} else {
			new_n += (n % 10) * int(math.Pow10(i))
		}
		n /= 10
	}
	return new_n
}

func same_length(a int, b int) bool {
	return int(math.Log10(float64(a))) == int(math.Log10(float64(b)))
}

func main() {
	defer elapsed("main")()
	limit := 1000000
	is_prime := get_primes(limit)
	primes := get_just_primes(is_prime)
	for _, prime := range primes {
		digits := get_distinct_digits(prime)
		found := false
		if len(digits) > 1 {
			for _, digit := range digits {
				cnt := 0
				for i := 0; i < 10; i++ {
					if i != digit {
						n := replace_digit(prime, digit, i)
						if is_prime[n] && same_length(n, prime) {
							cnt++
						}
					}
				}
				if cnt > 6 {
					fmt.Println(prime, digit)
					found = true
					break
				}
			}
		}
		if found {
			break
		}
	}
}