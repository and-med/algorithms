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
	for i := 3; i * i <= limit; i+=2 {
		if is_prime[i] {
			for j := i * i; j < limit; j += i {
				is_prime[j] = false
			}
		}
	}
	return is_prime
}

func log_ten(n int) int {
	return int(math.Log10(float64(n)))
}

func rotate(n int) int {
	num := n % 10
	pow := int(math.Pow10(log_ten(n)))
	n /= 10
	return n + pow * num
}

func main() {
	defer elapsed("main")()
	is_prime := get_primes(1000000)
	res := 0
	for num, prime := range is_prime {
		if prime {
			are_prime := true
			for i, n := 0, rotate(num); i < log_ten(num); i, n = i + 1, rotate(n) {
				if !is_prime[n] {
					are_prime = false
					break
				}
			}

			if are_prime {
				res++
			}
		}
	}

	fmt.Println(res)
}