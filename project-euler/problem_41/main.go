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

func is_prime(n int64) bool {
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

	return prime
}

func to_number(digits []int) int64 {
	s := int64(0)

	for _, d := range digits {
		s *= 10
		s += int64(d)
	}

	return s
}

func try_all_pandigital(is_used []bool, curr []int, length int) (int64, bool) {
	if len(curr) == length {
		num := to_number(curr)
		if is_prime(num) {
			return num, true
		}
		return 0, false
	}

	for i := length; i >= 1; i-- {
		if !is_used[i - 1] {
			is_used[i - 1] = true
			curr = append(curr, i)
			res, found := try_all_pandigital(is_used, curr, length)
			if found {
				return res, found
			}
			is_used[i - 1] = false
			curr = curr[:len(curr) - 1]
		}
	}

	return 0, false
}

func main() {
	defer elapsed("main")()
	for length := 9; length >= 1; length-- {
		max, found := try_all_pandigital(make([]bool, length), []int {}, length)
		if found {
			fmt.Println(max)
			break
		}
	}
}