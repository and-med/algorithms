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

func main() {
	defer elapsed("main")()
	limit := 28123
	sum_divisors := make([]int, limit + 1)

	for i := 1; i <= limit; i++ {
		for j := 2 * i; j <= limit; j += i {
			sum_divisors[j] += i
		}
	}

	var a_nums []int
	for i := 1; i <= limit; i++ {
		if (sum_divisors[i] > i) {
			a_nums = append(a_nums, i)
		}
	}

	sum := 0
	for i := 1; i <= limit; i++ {
		can_express := false
		for j := 0; j < len(a_nums) && a_nums[j] < i; j++ {
			num := i - a_nums[j]
			if sum_divisors[num] >num {
				can_express = true
				break;
			}
		}

		if !can_express {
			sum += i
		}
	}
	fmt.Println(sum)
}