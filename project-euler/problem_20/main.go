package main

import "fmt"

func mult(arr []int, d int) []int {
	var res []int
	rem := 0
	for j := 0; j < len(arr); j++ {
		p := d * arr[j] + rem
		res = append(res, p % 10)
		rem = p / 10
	}
	if rem != 0 {
		res = append(res, rem)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func print(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i]);
	}
	fmt.Println()
}

func main() {
	fact := []int { 0, 0, 8, 8, 2, 6, 3 }

	for i := 11; i < 100; i++ {
		d1 := i % 10
		d2 := i / 10

		p1 := mult(fact, d1)
		p2 := mult(fact, d2)

		fact = make([]int, 0)
		rem := 0
		for j := 0; j < max(len(p1), len(p2) + 1); j++ {
			s := 0
			if j == 0 {
				s = p1[j] + rem
			} else if j >= len(p1) {
				s = p2[j - 1] + rem
			} else {
				s = p1[j] + p2[j - 1] + rem
			}
			fact = append(fact, s % 10)
			rem = s / 10
		}
		if rem != 0 {
			fact = append(fact, rem)
		}
	}

	s := 0
	for i := 0; i < len(fact); i++ {
		s += fact[i]
	}
	fmt.Println(s)
}