package main

import (
	"fmt"
	"math"
	"time"
)

var limit int = 100000000
var limit_len int = 8
var limit_sum int = 100000
var solution_length int = 5
var is_prime_store []bool
var primes []int

func init() {
	is_prime_store = get_primes(limit)

	primes = []int{}
	for n, p := range is_prime_store {
		if p && n != 2 && n != 5 {
			primes = append(primes, n)
		}
	}
}

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

func is_prime(n int) bool {
	if length(n) <= limit_len {
		return is_prime_store[n]
	}

	if n < 2 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for _, p := range primes {
		if p > limit {
			break
		}
		if n%p == 0 {
			return false
		}
	}

	return true
}

func length(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func are_concatenated_primes(a int, b int) bool {
	ab := a * int(math.Pow10(length(b))) + b
	ba := b * int(math.Pow10(length(a))) + a

	return is_prime(ab) && is_prime(ba)
}

func are_concatenated_primes_multiple(primes []int) bool {
	for i := 0; i < len(primes) - 1; i++ {
		for j := i + 1; j < len(primes); j++ {
			if !are_concatenated_primes(primes[i], primes[j]) {
				return false
			}
		}
	}
	return true
}

func find_concatenated_multiples(perm []int, sum int, curr_prime int) {
	if len(perm) > 1 && !are_concatenated_primes_multiple(perm) {
		return
	}
	if len(perm) == solution_length {
		fmt.Println("Found sum:", sum)
		fmt.Println(perm)
		limit_sum = sum
		return
	}

	for i := curr_prime; i < len(primes); i++ {
		if sum + primes[i] > limit_sum {
			return
		}
		perm = append(perm, primes[i])
		find_concatenated_multiples(perm, sum + primes[i], i + 1)
		perm = perm[:len(perm) - 1]
	}
}

func solve() int64 {
	find_concatenated_multiples([]int {}, 0, 0)
	return -1
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}