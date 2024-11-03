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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

func get_count_fractions(n int) int {
	start, end, cnt := n / 3 + 1, (n + 1) / 2, 0
	for i := start; i < end; i++ {
		if gcd(i, n) == 1 {
			cnt++
		}
	}
	return cnt
}

func solve() int64 {
	limit := 12000
	cnt := 0
	for i := 2; i <= limit; i++ {
		cnt += get_count_fractions(i)
	}
	return int64(cnt)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}