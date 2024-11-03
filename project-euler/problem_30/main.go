package main

import (
	"fmt"
	"math"
)

func pow(a int, p int) int {
	return int(math.Pow(float64(a), float64(p)))
}

func try_digits(cnt_digits int, power int) []int {
	limit := pow(10, cnt_digits)
	res := make([]int, 0)

	for n := 10; n < limit; n++ {
		num := n
		s := 0
		for num != 0 {
			s += pow(num % 10, power)
			num /= 10
		}

		if s == n {
			res = append(res, n)
		}
	}

	return res
}

func main() {
	digits := 6
	power := 5

	results := try_digits(digits, power)

	sum := 0
	for _, res := range results {
		sum += res
	}

	fmt.Println(results)
	fmt.Println(sum)
}