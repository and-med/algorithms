package main

import (
	"fmt"
	"time"
)

var limit = int(1e6)
var is_prime []bool
var primes []int
var cache map[int]int

func init() {
	is_prime = get_primes(limit)

	primes = []int{}
	for n, p := range is_prime {
		if p {
			primes = append(primes, n)
		}
	}

	cache = map[int]int {}
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

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func find_first_prime_by_condition(cond func (int) bool) int {
	first := 0
	for id, p := range primes {
		if cond(p) {
			first = id
		} else {
			break
		}
	}
	return first

}

func find_first_prime_less_than_or_equal(n int) int {
	return find_first_prime_by_condition(func (p int) bool { return p <= n })
}

func find_first_prime_less_than(n int) int {
	return find_first_prime_by_condition(func (p int) bool { return p < n })
}

func find_permutations_rec(prime_index int, sum int) int {
	if primes[prime_index] > sum {
		prime_index = find_first_prime_less_than_or_equal(sum)
	}

	if v, ok := cache[prime_index * 1000 + sum]; ok {
		return v
	}
	if sum == 0 {
		return 1
	}
	if sum % 2 == 1 && prime_index == 0 {
		return 0
	}

	cnt := 0
	for i := prime_index; i >= 0; i-- {
		cnt += find_permutations_rec(i, sum - primes[i])
	}
	cache[prime_index * 1000 + sum] = cnt
	return cnt
}

func find_permutations(sum int) int {
	first_less_than := find_first_prime_less_than(sum)

	return find_permutations_rec(first_less_than, sum)
}

func solve() int64 {
	res := 0
	for i := 6; i < 1000; i++ {
		if find_permutations(i) > 5000 {
			res = i
			break
		}
	}
	return int64(res)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}