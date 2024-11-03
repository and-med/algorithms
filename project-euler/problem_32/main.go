package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func elapsed(what string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", what, time.Since(start))
    }
}

func contains(list []int, n int) bool {
	for _, x := range list {
		if x == n {
			return true
		}
	}

	return false
}

func is_legit_perm(perm []int, prod int) bool {
	digits := make([]int, 0)
	for prod != 0 {
		digits = append(digits, prod % 10)
		prod /= 10
	}

	if len(digits) != 4 {
		return false
	}

	for _, dig := range digits {
		if contains(perm, dig) {
			return false
		}
	}

	sort.Ints(digits)

	if digits[0] == 0 {
		return false
	}

	for i := 1; i < len(digits); i++ {
		if digits[i] == digits[i - 1] {
			return false
		}
	}

	return true
}

func to_number(arr []int) int {
	s := 0
	for id, x := range arr {
		s += x * int(math.Pow10(len(arr) - id - 1))
	}
	return s
}

func try_permutations(curr_perm []int) []int {
	if len(curr_perm) == 5 {
		prod := to_number(curr_perm[0:2]) * to_number(curr_perm[2:])
		if is_legit_perm(curr_perm, prod) {
			return []int { prod }
		}

		prod = to_number(curr_perm[0:1]) * to_number(curr_perm[1:])
		if is_legit_perm(curr_perm, prod) {
			return []int { prod }
		}

		return nil
	}

	res := make([]int, 0)
	for i := 1; i <= 9; i ++ {
		if !contains(curr_perm, i) {
			curr_perm = append(curr_perm, i)
			perm := try_permutations(curr_perm)
			if perm != nil {
				res = append(res, perm...)
			}
			curr_perm = curr_perm[:len(curr_perm) - 1]
		}
	}

	if len(res) > 0 {
		return res
	}
	return nil
}

func main() {
	defer elapsed("main")()
	res := try_permutations([]int {})
	
	unique := make(map[int]bool)
	for _, num := range res {
		unique[num] = true
	}

	sum := 0
	for key := range unique {
		sum += key
	}
	fmt.Println(sum)
}
