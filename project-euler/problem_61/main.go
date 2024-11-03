package main

import (
	"fmt"
	"sort"
	"time"
)

const (
	start_with_rank = 10
	end_with_rank   = 100
)

var number_keys []int
var numbers map[int][]int
var numbers_by_digits map[int]map[int][]int

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func init() {
	generators := map[int]func(int) int{3: triangle, 4: square, 5: pentagonal, 6: hexagonal, 7: heptagonal, 8: octagonal}
	numbers = map[int][]int{}
	for key, value := range generators {
		numbers[key] = generate_numbers(value)
	}
	numbers_by_digits = map[int]map[int][]int{}
	number_keys = []int {}
	for key, value := range numbers {
		numbers_by_digits[key*start_with_rank] = to_map_by(value, get_first_two_digits)
		numbers_by_digits[key*end_with_rank] = to_map_by(value, get_last_two_digits)
		number_keys = append(number_keys, key)
	}
	sort.Ints(number_keys)
}

func generate_numbers(generator func(int) int) []int {
	res := []int{}
	for i := 1; true; i++ {
		n := generator(i)
		if n > 9999 {
			break
		}
		if n > 999 {
			res = append(res, n)
		}
	}
	return res
}

func triangle(n int) int   { return n * (n + 1) / 2 }
func square(n int) int     { return n * n }
func pentagonal(n int) int { return n * (3*n - 1) / 2 }
func hexagonal(n int) int  { return n * (2*n - 1) }
func heptagonal(n int) int { return n * (5*n - 3) / 2 }
func octagonal(n int) int  { return n * (3*n - 2) }

func to_map_by(numbers []int, key_provider func(int) int) map[int][]int {
	res := map[int][]int{}
	for _, n := range numbers {
		key := key_provider(n)
		res[key] = append(res[key], n)
	}
	return res
}

func get_first_two_digits(n int) int {
	return n / 100
}

func get_last_two_digits(n int) int {
	return n % 100
}

func contains(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}

func build_solutions(key_1 int, key_2 int) [][]int {
	sol := [][]int {}
	numbers_1 := numbers[key_1]
	for _, n1 := range numbers_1 {
		last_digits_1 := get_last_two_digits(n1)
		numbers_2, ok := numbers_by_digits[key_2 * start_with_rank][last_digits_1]
		if ok {
			for _, n2 := range numbers_2 {
				sol = append(sol, []int { n1, n2 })
			}
		}
	}
	return sol
}

func append_solutions(solutions [][]int, new_key int) [][]int {
	new_solutions := [][]int {}
	for _, sol := range solutions {
		last_n := sol[len(sol) - 1]
		last_digits := get_last_two_digits(last_n)
		numbers, ok := numbers_by_digits[new_key * start_with_rank][last_digits]
		if ok {
			for _, n := range numbers {
				new_solution := append(sol, n)
				new_solutions = append(new_solutions, new_solution)
			}
		}
	}
	return new_solutions
}

func build_all_solutions(keys []int) [][]int {
	if len(keys) < 2 {
		return [][]int {}
	}
	solutions := build_solutions(keys[0], keys[1])
	for i := 2; i < len(keys); i++ {
		solutions = append_solutions(solutions, keys[i])
	}
	return solutions
}

func filter_solutions(solutions [][]int) [][]int {
	new_solutions := [][]int {}
	for _, sol := range solutions {
		first_n := sol[0]
		last_n := sol[len(sol) - 1]
		if get_first_two_digits(first_n) == get_last_two_digits(last_n) {
			new_solutions = append(new_solutions, sol)
		}
	}
	return new_solutions
}

func find_solution(key_perm[]int, solutions[][]int) [][]int {
	if len(key_perm) > 2 && len(solutions) == 0 {
		return solutions
	}
	if len(key_perm) == 6 {
		return filter_solutions(solutions)
	}

	res := [][]int {}
	for _, new_key := range number_keys {
		if !contains(key_perm, new_key) {
			key_perm = append(key_perm, new_key)
			solutions := build_all_solutions(key_perm)
			new_solutions := find_solution(key_perm, solutions)
			res = append(res, new_solutions...)
			key_perm = key_perm[:len(key_perm) - 1]
		}
	}
	return res
}

func solve() int64 {
	solutions := find_solution([]int {}, [][]int {})
	fmt.Println(solutions)
	sum := 0
	for _, n := range solutions[0] {
		sum += n
	}
	return int64(sum)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}
