package main

import (
	"fmt"
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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) string {
	if n < 0 {
		return "-"
	}
	return "+"
}

func test_consecutives(is_prime []bool, a int, b int) int {
	cnt := 0

	for true {
		if is_prime[abs(cnt * cnt + a * cnt + b)] {
			cnt++
		} else {
			break
		}
	}

	return cnt
}

func main() {
	is_prime := get_primes(100000)
	coef := []int { -1, 0, 1 }
	for i := 0; i <= 1000; i++ {
		if is_prime[i] {
			coef = append(coef, i)
			coef = append(coef, -i)
		}
	}

	max := 0
	maxa := 0
	maxb := 0
	for _, a := range coef {
		for _, b := range coef {
			c := test_consecutives(is_prime, a, b)

			if c > max {
				max = c
				maxa = a
				maxb = b
			}
		}
	}

	fmt.Printf("n^2 %s %d * n %s %d", sign(maxa), abs(maxa), sign(maxb), abs(maxb))
	fmt.Println()
	fmt.Println(test_consecutives(is_prime, maxa, maxb), "consecutive primes")
	fmt.Println("a * b =", maxa * maxb)
}