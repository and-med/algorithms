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

func to_digit_map(n int64, res []int) {
	for i := 0; i < 10; i++ {
		res[i] = 0
	}
	for n != 0 {
		res[int(n % 10)]++
		n /= 10
	}
}

func equals(d_1 [10]int, d_2 [10]int) bool {
	for idx, value := range d_1 {
		if d_2[idx] != value {
			return false
		}
	}
	return true
}

func solve() int64 {
	digits := [10]int {}
	multiple := [10]int {}

	for start := int64(100000); true; start+=10 {
		end := int64(start * 10 / 6 + 1)
		for i := start; i < end; i++ {
			to_digit_map(i, digits[:])
			same := true
			for j := 2; j < 7; j++ {
				to_digit_map(i * int64(j), multiple[:])
				if !equals(digits, multiple) {
					same = false
					break
				}
			}
			if same {
				return i
			}
		}
	}
	return -1
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}