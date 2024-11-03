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

func try_coins_dp(coins []int) int {
	var dp [8][201]int
	for j := 0; j < 8; j++ {
		dp[j][0] = 1
	}
	for i := 0; i <= 200; i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= 200; i++ {
		for j := 1; j < 8; j++ {
			if i >= coins[j] {
				dp[j][i] = dp[j-1][i] + dp[j][i - coins[j]]
			} else {
				dp[j][i] = dp[j-1][i]
			}
		}
	}

	return dp[7][200]
}

func try_coins(coins []int, sum int, curr int) int {
	if sum == 0 {
		return 1
	}

	if curr >= len(coins) {
		return 0
	}

	total := 0
	for i := 0; i <= sum / coins[curr]; i++ {
		total += try_coins(coins, sum - coins[curr] * i, curr + 1)
	}
	return total
}

func main() {
	defer elapsed("main")()
	coins := []int { 1, 2, 5, 10, 20, 50, 100, 200 }
	fmt.Println(try_coins(coins, 200, 0))
	fmt.Println(try_coins_dp(coins))
}