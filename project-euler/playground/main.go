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
	x := make([]int, 3, 5)
	x[0], x[1], x[2] = 1, 2, 3
	y := append(x[1:3], 4)
	z := append(x, 5)
	fmt.Println(x, y, z)
	return -1
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}