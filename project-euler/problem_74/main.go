package main

import (
	"fmt"
	"time"
)

var factorials []int

func init() {
	factorials = []int {1}
	for i := 1; i < 10; i++ {
		factorials = append(factorials, i * factorials[i - 1]);
	}
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func get_factorial_sum(n int) int {
	sum := 0
	for n != 0 {
		sum += factorials[n % 10]
		n /= 10
	}
	return sum
}

func get_chain_length(n int) int {
	chain := map[int]bool {}
	cnt := 0

	for !chain[n] {
		chain[n] = true
		n = get_factorial_sum(n)
		cnt++
	}

	return cnt
}

func solve() int64 {
	cnt := 0
	for i := 2; i < 1000000; i++ {
		if get_chain_length(i) == 60 {
			cnt++
		}
	}
	return int64(cnt)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}