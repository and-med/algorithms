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

func to_number(digits []int) int {
	s := 0
	for _, d := range digits {
		s *= 10
		s += d
	}
	return s
}

func to_number_int64(digits []int) int64 {
	s := int64(0)
	for _, d := range digits {
		s *= 10
		s += int64(d)
	}
	return s
}

var rules map[int]int

func try_permutations(is_used []bool, curr_perm []int) int64 {
	length := len(curr_perm)
	rule, ok := rules[length]
	if ok {
		if (to_number(curr_perm[length - 3: length]) % rule) != 0 {
			return 0
		} else if length == 10 {
			return to_number_int64(curr_perm)
		}
	}

	s := int64(0)
	for i := 0; i < 10; i++ {
		if !is_used[i] {
			is_used[i] = true
			curr_perm = append(curr_perm, i)
			s += try_permutations(is_used, curr_perm)
			is_used[i] = false
			curr_perm = curr_perm[:len(curr_perm) - 1]
		}
	}
	return s
}

func main() {
	defer elapsed("main")()
	rules =  map[int] int { 4 : 2, 5 : 3, 6 : 5, 7 : 7, 8 : 11, 9 : 13, 10 : 17 }
	fmt.Println(try_permutations(make([]bool, 10), []int {}))
}