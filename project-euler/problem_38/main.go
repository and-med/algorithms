package main

import "fmt"

func to_digits(n int) []int {
	d := []int {}
	for n != 0 {
		d = append(d, n % 10)
		n /= 10
	}

	for i := 0; i < len(d) / 2; i++ {
		d[i], d[len(d) - i - 1] = d[len(d) - i - 1], d[i]
	}
	return d
}

func is_pandigital(a []int) bool {
	if len(a) != 9 {
		return false
	}
	d := make([]bool, 9)

	for _, x := range a {
		if x == 0 || d[x - 1] {
			return false
		}
		d[x - 1] = true
	}

	return true
}

func get_product(n int, v int) []int {
	p := []int {}
	for i := 1; i <= v; i++ {
		p = append(p, to_digits(n * i)...)
	}
	return p;
}

func solve() []int {
	for i := 9999; i > 5000; i-- {
		p := get_product(i, 2)
		if is_pandigital(p) {
			return p
		}
	}
	return nil
}

func main() {
	res := solve()
	for _, d := range res {
		fmt.Print(d)
	}
	fmt.Println()
}

// 1				x (1, 2, 3, 4, 5, 6, 7, 8, 9)
// 3				x (1, 2, 3, 4, 5, 6)
// (5 - 9)			x (1, 2, 3, 4, 5)
// (25 - 33)		x (1, 2, 3, 4)
// (100 - 333)		x (1, 2, 3)
// (5000 - 9999)	x (1, 2)