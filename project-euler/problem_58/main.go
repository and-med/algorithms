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

func get_primes(limit int) []int {
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

	primes := []int{}
	for n, p := range is_prime {
		if p {
			primes = append(primes, n)
		}
	}

	return primes
}

func is_prime(n int, primes []int) bool {
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

func solve() int64 {
	primes := get_primes(100000)
	for step, cnt := 2, 0; true; step += 2 {
		last := (step - 1) * (step - 1)
		for j := 1; j < 4; j++ {
			if is_prime(last+j*step, primes) {
				cnt++
			}
		}
		if float64(cnt)/float64(2*step+1) < 0.1 {
			return int64(step + 1)
		}
	}
	return -1
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}
