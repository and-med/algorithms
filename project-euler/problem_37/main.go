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

var primes map [int64]bool

func is_prime(n int64) bool {
	if _, ok := primes[n]; !ok {
		prime := true
		if n < 2 {
			prime = false
		}
		for i := int64(2); i <= int64(math.Sqrt(float64(n))); i++ {
			if n % i == 0 {
				prime = false
				break
			}
		}
		primes[n] = prime
	}

	return primes[n]
}

func is_truncatable_prime(digits []int) bool {
	left := int64(0)
	right := int64(0)
	power := int64(1)

	for i := 0; i < len(digits); i++ {
		left *= 10
		left += int64(digits[i])
		right += int64(digits[len(digits) - i - 1]) * power
		power *= 10
		if left == right && !is_prime(left) {
			return false
		} else if !is_prime(left) || !is_prime(right) {
			return false
		}
	}

	return true
}

func copy(arr []int) []int {
	res := make([]int, len(arr))
	for id, v := range arr {
		res[id] = v
	}
	return res
}

func try_possible_primes_recur(digits []int, last_digits []int, curr_perm []int, length int) ([][]int, int) {
	r := make([][]int, 0)
	tries := 0

	if length > 0 && len(curr_perm) != 1 {
		tries++
		if is_truncatable_prime(curr_perm) {
			r = append(r, copy(curr_perm))
		}
	} else if length == 0 {
		tries++
		if is_truncatable_prime(curr_perm) {
			return [][]int { copy(curr_perm) }, tries
		}
		return nil, tries
	}

	var res [][]int
	var total int
	if length == 1 {
		res, total = try_possible_primes_inner(last_digits, digits, last_digits, curr_perm, length)
	} else {
		res, total = try_possible_primes_inner(digits, digits, last_digits, curr_perm, length)
	}
	r = append(r, res...)
	tries += total

	return r, tries
}

func try_possible_primes_inner(to_try []int, digits []int, last_digits []int, curr_perm []int, length int) ([][]int, int) {
	r := make([][]int, 0)
	tries := 0

	for _, digit := range to_try {
		curr_perm = append(curr_perm, digit)
		res, total := try_possible_primes_recur(digits, last_digits, curr_perm, length - 1)
		tries += total
		if res != nil {
			r = append(r, res...)
		}
		curr_perm = curr_perm[:len(curr_perm) - 1]
	}

	return r, tries
}

func try_possible_primes(digits []int, first_digits []int, last_digits []int, length int) ([][]int, int) {
	return try_possible_primes_inner(first_digits, digits, last_digits, []int {}, length);
}

func to_number(digits []int) int64 {
	s := int64(0)

	for _, d := range digits {
		s *= 10
		s += int64(d)
	}

	return s
}

func to_numbers(arr [][]int) []int64 {
	res := []int64 {}
	for _, n := range arr {
		res = append(res, to_number(n))
	}
	return res
}

func sum(arr []int64) int64 {
	s := int64(0)
	for _, n := range arr {
		s += n
	}
	return s
}

func main() {
	defer elapsed("main")()
	primes = make(map [int64]bool)
	digits := []int { 1, 3, 7, 9 }
	first_digits := []int { 2, 3, 5, 7 }
	last_digits := []int { 3, 7 }
	permutations, tries := try_possible_primes(digits, first_digits, last_digits, 6)

	res := to_numbers(permutations)

	fmt.Println(res)
	fmt.Println("Length: ", len(res))
	fmt.Println("Sum: ", sum(res))
	fmt.Println("Total numbers tried: ", tries)
}