package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

const limit = 1e7

func get_totients(limit int) []int {
	totients := make([]int, limit)

	totients[0] = 0
	totients[1] = 1

	for i := 2; i < limit; i += 2 {
		totients[i] = i / 2
	}

	for i := 3; i < limit; i += 2 {
		totients[i] = i
	}

	for i := 3; i < limit; i += 2 {
		if totients[i] == i {
			for j := i; j < limit; j += i {
				totients[j] = totients[j] * (i - 1) / i
			}
		}
	}

	return totients
}

func get_sorted_digits(n int) int {
	digits := []int {}
	for n != 0 {
		digits = append(digits, n % 10)
		n /= 10
	}

	sort.Ints(digits)
	r := 0
	for _, d := range digits {
		r *= 10
		r += d
	}
	return r
}

func solve() int64 {
	min := math.MaxFloat64
	min_n := math.MaxInt32
	for n, totient := range get_totients(limit) {
		if n == 0 || n == 1 {
			continue
		}
		ratio := float64(n) / float64(totient)
		if get_sorted_digits(totient) == get_sorted_digits(n) && ratio < min {
			min, min_n = ratio, n
		}
	}
	return int64(min_n)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}