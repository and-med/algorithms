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

func solve() int64 {
	cnt := 0
	limit := 1000000
	trian := [101][101]int {}

	for n := 0; n <= 100; n++ {
		trian[n][0] = 1
	}
	
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			trian[n][r] = trian[n - 1][r - 1] + trian[n - 1][r]
			if trian[n][r] > limit {
				cnt++
				trian[n][r] = limit
			}
		}
	}

	return int64(cnt)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}