package main

import (
	"fmt"
	"math"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func cpy(arr []int) []int {
	x := make([]int, len(arr))
	copy(x, arr)
	return x
}

func permutations_of(n int, k int, curr_perm []int, start int) [][]int {
	if len(curr_perm) == k {
		return [][]int { cpy(curr_perm) }
	}

	res := [][]int {}
	for i := start; i <= n; i++ {
		curr_perm = append(curr_perm, i)
		permutations := permutations_of(n, k, curr_perm, i + 1)
		res = append(res, permutations...)
		curr_perm = curr_perm[:len(curr_perm) - 1]
	}

	return res
}

func get_ring_sum_range(ring_rank int) (int, int) {
	ring_numbers_count := ring_rank * 2
	permutations := permutations_of(ring_numbers_count, 3, []int {}, 1)
	count_permutations_by_sum := map[int]int {}
	for _, perm := range permutations {
		sum := 0
		for _, x := range perm {
			sum += x
		}
		count_permutations_by_sum[sum] += 1
	}

	start := math.MaxInt32
	end := 0
	for sum, count := range count_permutations_by_sum {
		if count >= ring_rank && sum < start {
			start = sum
		}
		if count >= ring_rank && sum > end {
			end = sum
		}
	}

	return start, end
}

func arr_sum(arr []int) int {
	sum := 0
	for _, x := range arr {
		sum += x
	}
	return sum
}

func get_rings_by_sum(ring_rank int, sum int, curr_ring []int, is_used []bool) [][]int {
	curr_len := len(curr_ring)
	n := ring_rank * 2

	if curr_len == ring_rank * 3 {
		return [][]int { curr_ring }
	}

	res := [][]int {}
	if arr_sum(curr_ring[curr_len - curr_len % 3 : curr_len]) >= sum {
		return res
	}

	if curr_len / 3 == 0 {
		if curr_len == 0 || curr_len == 1 {
			for i := 1; i <= n; i++ {
				if arr_sum(curr_ring) >= sum {
					return res
				}
				if !is_used[i] {
					curr_ring = append(curr_ring, i)
					is_used[i] = true
					res = append(res, get_rings_by_sum(ring_rank, sum, curr_ring, is_used)...)
					is_used[i] = false
					curr_ring = curr_ring[:len(curr_ring) - 1]
				}
			}

			return res
		} else if curr_len == 2 {
			set_1 := curr_ring[0]
			set_2 := curr_ring[1]
			set_3 := sum - set_1 - set_2
			if set_3 > 0 && set_3 <= n && !is_used[set_3] {
				curr_ring = append(curr_ring, set_3)
				is_used[set_3] = true
				res = append(res, get_rings_by_sum(ring_rank, sum, curr_ring, is_used)...)
				is_used[set_3] = false
				curr_ring = curr_ring[:len(curr_ring) - 1]
			}

			return res
		}
	} else if curr_len / 3 > 0 {
		if curr_len % 3 == 0 {
			for i := 1; i <= n; i++ {
				if !is_used[i] {
					curr_ring = append(curr_ring, i)
					is_used[i] = true
					res = append(res, get_rings_by_sum(ring_rank, sum, curr_ring, is_used)...)
					is_used[i] = false
					curr_ring = curr_ring[:len(curr_ring) - 1]
				}
			}

			return res
		} else if curr_len % 3 == 1 {
			i := curr_ring[curr_len - 2]
			curr_ring = append(curr_ring, i)
			return append(res, get_rings_by_sum(ring_rank, sum, curr_ring, is_used)...)
		} else if curr_len % 3 == 2 {
			set_1 := curr_ring[curr_len - 2]
			set_2 := curr_ring[curr_len - 1]
			set_3 := sum - set_1 - set_2
			if set_3 > 0 && set_3 <= n {
				if curr_len + 1 < ring_rank * 3 && !is_used[set_3] {
					curr_ring = append(curr_ring, set_3)
					is_used[set_3] = true
					res = append(res, get_rings_by_sum(ring_rank, sum, curr_ring, is_used)...)
					is_used[set_3] = false
				} else if curr_len + 1 == ring_rank * 3 && set_3 == curr_ring[1] {
					curr_ring = append(curr_ring, set_3)
					res = append(res, get_rings_by_sum(ring_rank, sum, curr_ring, is_used)...)
				}
			}

			return res
		}
	}

	return res
}

func filter_rings(rings [][]int) [][]int {
	res := [][]int {}
	for _, ring := range rings {
		is_sol := true
		for i := 3; i < len(ring); i+=3 {
			if ring[i] < ring[0] {
				is_sol = false
			}
		}
		if is_sol {
			res = append(res, ring)
		}
	}
	return res
}

func get_rings(ring_rank int) [][]int {
	start, end := get_ring_sum_range(ring_rank)
	res := [][]int {}
	for sum := start; sum <= end; sum++ {
		is_used := make([]bool, ring_rank * 2 + 1)
		temp_res := filter_rings(get_rings_by_sum(ring_rank, sum, []int {}, is_used))
		res = append(res, temp_res...)
	}

	return res
}

func solve() int64 {
	ring_rank := 5
	rings := get_rings(ring_rank)
	fmt.Println(len(rings))
	for _, ring := range rings {
		fmt.Println(ring)
	}

	return -1
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}