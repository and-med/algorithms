package main

import (
	"fmt"
	"math"
	"strconv"
)

func pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func mult(a []int, d int) []int {
	var res []int
	rem := 0
	for j := 0; j < len(a); j++ {
		p := d * a[j] + rem
		res = append(res, p % 10)
		rem = p / 10
	}
	if rem != 0 {
		res = append(res, rem)
	}
	return res
}

func add(arr [][]int) []int {
	max_len := len(arr[0])
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > max_len {
			max_len = len(arr[i])
		}
	}

	rem := 0
	var sum []int
	for j := 0; j < max_len; j++ {
		s := 0
		for i := 0; i < len(arr); i++ {
			if j < len(arr[i]) {
				s = s + arr[i][j]
			}
		}
		s += rem
		sum = append(sum, s % 10)
		rem = s / 10
	}

	if rem > 0 {
		sum = append(sum, rem)
	}
	return sum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func add_two(a []int, b []int) []int {
	var res []int
	rem := 0
	
	for i := 0; i < max(len(a), len(b)); i++ {
		s := 0
		if i >= len(a) {
			s = b[i] + rem
		} else if i >= len(b) {
			s = a[i] + rem
		} else {
			s = a[i] + b[i] + rem
		}
		res = append(res, s % 10)
		rem = s / 10
	}

	if rem != 0 {
		res = append(res, rem)
	}

	return res
}

func multiply(a []int, b []int) []int {
	var prods [][]int
	for id, d := range b {
		res := mult(a, d)
		for i := 0; i < id; i++ {
			res = append([]int{0}, res...)
		}
		prods = append(prods, res)
	}
	return add(prods)
}

func reverse_print(a []int) {
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Print(a[i])
	}
	fmt.Println()
}

func to_string(a []int) string {
	s := ""
	for i := len(a) - 1; i >= 0; i-- {
		s += strconv.Itoa(a[i])
	}
	return s
}

func to_number(a []int) int {
	s := 0
	for i := len(a) - 1; i >= 0; i-- {
		s += a[i] * int(math.Pow(10, float64(i)))
	}
	return s
}

func brute_force() {
	one := []int { 1 }
	res := make(map[string][]string)
	for a := []int { 2 }; to_number(a) != 101; a = add_two(a, one) {
		prod := multiply(a, a)
		for b := []int { 2 }; to_number(b) != 101; b = add_two(b, one) {
			res[to_string(prod)] = append(res[to_string(prod)], to_string(a) + "^" + to_string(b))
			prod = multiply(prod, a)
		}
	}

	cnt := 0
	for _, value := range res {
		if len(value) > 1 {
			fmt.Print("Same group: ")
			for _, v := range value {
				fmt.Print(v, " ")
			}
			fmt.Println()
		}
		cnt++
	}

	fmt.Println(cnt)
}

func smart() {
	var visited [99][99]bool
	cnt := 0
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			if !visited[a - 2][b - 2] {
				visited[a - 2][b - 2] = true
				cnt++

				for div := 2; div <= 10; div++ {
					p := pow(a, div)
					if p > 100 {
						break;
					} else if b % div == 0 && b / div >= 2 {
						visited[p - 2][b / div - 2] = true
					}
				}
			}
		}
	}
	fmt.Println(cnt)
}

func main() {
	brute_force()
	smart()
}