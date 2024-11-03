package main

import (
	"fmt"
	"time"
)

var cache map[int]int
var cache_2 map[int]int

func init() {
	cache = map[int]int {}
	cache_2 = map[int]int {}
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func possible_sums_rec(last_number int, sum int) int {
	last_number = min(last_number, sum)
	if v, ok := cache[last_number * 1000 + sum]; ok {
		return v
	}
	if sum == 0 {
		return 1
	}
	
	cnt := 0
	for i := last_number; i > 0; i-- {
		cnt += possible_sums_rec(i, sum - i)
	}
	cache[last_number * 1000 + sum] = cnt
	return cnt
}

func possible_sums_rec_2(sum int, n int) int {
	if v, ok := cache_2[sum * 1000 + n]; ok {
		return v
	}
	if n > sum / 2 {
		return 0
	}
	cnt := possible_sums_rec_2(sum - n, n) + possible_sums_rec_2(sum, n + 1) + 1
	cache_2[sum * 1000 + n] = cnt
	return cnt;
}

func possible_sums(sum int) int {
	defer elapsed("possible_sums")()
	return possible_sums_rec(sum - 1, sum)
}

func possible_sums_2(sum int) int {
	defer elapsed("possible_sums_2")()
	return possible_sums_rec_2(sum, 1)
}

func possible_sums_dynamic(sum int) int {
	defer elapsed("possible_sums_dynamic")()
	dp := make([][]int, sum)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, sum)
	}

	for i := 0; i < sum; i++ {
		dp[i][0] = 1
		dp[i][1] = 1
	}

	for l := 1; l < sum; l++ {
		for s := 2; s < sum; s++ {
			cnt := 0 
			for curr := min(l, s); curr > 0; curr -- {
				cnt += dp[curr][s - curr]
			}
			dp[l][s] = cnt
		}
	}

	res := 0
	for i := 1; i < sum; i++ {
		res += dp[i][sum - i]
	}

	return res
}

func possible_sums_dynamic_2(sum int) int {
	defer elapsed("possible_sums_dynamic_2")()
	dp := make([][]int, sum)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, sum + 1)
	}

	for s := 0; s < sum + 1; s++ {
		dp[1][s] = 1
		if s != sum {
			dp[s][s] = 1
		}
	}

	for l := 2; l < sum; l++ {
		for s := l + 1; s < sum + 1; s++ {
			cnt := 0 
			for curr := min(l, s - l); curr > 0; curr -- {
				cnt += dp[curr][s - l]
			}
			dp[l][s] = cnt
		}
	}

	res := 0
	for i := 1; i < sum; i++ {
		res += dp[i][sum]
	}

	return res
}

func possible_sums_dynamic_3(sum int) int {
	defer elapsed("possible_sums_dynamic_3")()
	dp := make([][]int, sum/2 + 2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, sum + 1)
	}

	for s := 2; s < sum + 1; s++ {
		for n := s / 2; n > 0; n-- {
			dp[n][s] = dp[n][s - n] + dp[n + 1][s] + 1
		}
	}

	return dp[1][sum]
}

func solve() int64 {
	limit := 5
	fmt.Println(possible_sums(limit))
	fmt.Println(possible_sums_2(limit))
	fmt.Println(possible_sums_dynamic(limit))
	fmt.Println(possible_sums_dynamic_2(limit))
	res := possible_sums_dynamic_3(limit)
	return int64(res)
}

func main() {
	fmt.Println(solve())
}