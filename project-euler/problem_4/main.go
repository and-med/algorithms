package main

import "fmt"

func reverse(n int) int {
	r := 0
	for n > 0 {
		r, n = r * 10 + n % 10, n / 10;
	}
	return r;
}

func main() {
	max := 0
	for i := 990; i >= 110; i -= 11 {
		for j := 999; j >= 100; j-- {
			p := i * j;
			if p == reverse(p) && p > max {
				max = p
			}
		}
	}
	fmt.Println(max)
}