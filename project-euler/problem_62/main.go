package main

import (
	"fmt"
	"sort"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func sort_digits(n int64) int64 {
	digits := []int{}
	for n != 0 {
		digits = append(digits, int(n%10))
		n /= 10
	}
	sort.Slice(digits, func(a, b int) bool {
		return digits[a] > digits[b]
	})
	res := int64(0)
	for _, d := range digits {
		res *= 10
		res += int64(d)
	}
	return res
}

func solve() int64 {
	cubes := map[int64][]int64{}
	for i := int64(1); true; i++ {
		cube := i * i * i
		sorted := sort_digits(cube)
		cubes[sorted] = append(cubes[sorted], cube)
		if len(cubes[sorted]) == 5 {
			return cubes[sorted][0]
		}
	}
	return -1
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}
