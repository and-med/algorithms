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

func contains(arr [][]int, a, b int) bool {
	for _, x := range arr {
		if x[0] == a && x[1] == b {
			return true
		}
	}
	return false
}

// curr_d / (sqrt(n) - curr_n) = curr_d * (sqrt(n) + curr_n) / (n - curr_n^2)
// new_d = (n - curr_n^2) / curr_d
// (sqrt(n) + curr_n) / new_d
// curr_n = res * new_d - new_n
// (sqrt(n) + res * new_d - new_n) / new_d = res + (sqrt(n) - new_n) / new_d
func transform_root(n int) int {
	a := int(math.Sqrt(float64(n)))
	cnt := 0

	if (a * a == n) {
		return cnt
	}

	curr_d := 1
	curr_n := a
	for {
		curr_d = (n - curr_n * curr_n) / curr_d
		res := int(float64(curr_n + a) / float64(curr_d))
		curr_n = res * curr_d - curr_n
		cnt++
		if curr_d == 1 && curr_n == a {
			break
		}
	}

	return cnt
}

func solve() int64 {
	cnt := 0
	for i := 2; i <= 10000; i++ {
		res := transform_root(i)
		if res % 2 == 1 {
			cnt++
		}
	}
	return int64(cnt)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}