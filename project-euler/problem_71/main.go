package main

import (
	"fmt"
	"math"
	"time"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func solve() int64 {
	limit := int(1e6)
	max_nom := math.MinInt32
	max_denom := math.MinInt32

	for i := limit; i >= 2; i-- {
		nom, denom := 3 * i - 1, 7 * i
		divisor := gcd(nom, denom)
		new_nom, new_denom := nom / divisor, denom / divisor
		if new_denom <= limit && new_denom > max_denom {
			max_nom = new_nom
			max_denom = new_denom
		}
	}
	fmt.Println(max_nom, "/", max_denom)

	return -1
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}