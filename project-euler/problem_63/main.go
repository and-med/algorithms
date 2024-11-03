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

func solve() int64 {
	logs := []float64 {}
	for i := 1; i <= 9; i++ {
		logs = append(logs, math.Log10(float64(i)))
	}

	cnt := 0
	for n := 1; true; n++ {
		prev_cnt := cnt
		for _, log := range logs {
			if log >= float64(n - 1)/float64(n) {
				cnt++
			}
		}
		if prev_cnt == cnt {
			break
		}
	}

	return int64(cnt)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}

// 10 ^ (n - 1) <= x ^ n < 10 ^ n
// (n - 1) / n <= log10(x) < 1
// All integer x are solutions, obviously x is always less than 10