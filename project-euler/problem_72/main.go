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

const limit = 1e6 + 1

func get_totients(limit int) []int {
	totients := make([]int, limit)

	totients[0] = 0
	totients[1] = 0

	for i := 2; i < limit; i += 2 {
		totients[i] = i / 2
	}

	for i := 3; i < limit; i += 2 {
		totients[i] = i
	}

	for i := 3; i < limit; i += 2 {
		if totients[i] == i {
			for j := i; j < limit; j += i {
				totients[j] = totients[j] * (i - 1) / i
			}
		}
	}

	return totients
}

func solve() int64 {
	totients := get_totients(limit)
	sum := 0
	for _, t := range totients {
		sum += t
	}
	return int64(sum)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}