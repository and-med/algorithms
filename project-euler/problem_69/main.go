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

const limit = 1000000
var primes[]int 

func init() {
	primes = get_primes(limit + 1)
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

func get_phi(n int) int {
	res := n
	for i := 0; i < len(primes) && primes[i] * primes[i] <= n; i++ {
		if n % primes[i] == 0 {
			for n % primes[i] == 0 {
				n /= primes[i]
			}

			res = res * (primes[i] - 1) / primes[i]
		}
	}

	if n > 1 {
		res = res * (n - 1) / n
	}

	return res
}

func solve() int64 {
	max := float64(0)
	max_n := 0
	for i := 2; i <= limit; i++ {
		new_phi := float64(i)/float64(get_phi(i))
		if new_phi > max {
			max = new_phi
			max_n = i
		}
	}
	return int64(max_n)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}