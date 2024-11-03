package main

import "fmt"

func to_digits(n int) []int {
	digits := make([]int, 0)
	for n != 0 {
		digits = append(digits, n % 10)
		n = n / 10
	}
	return digits
}

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return factorial(n - 1) * n
}

func calculate_factorials(n int) []int {
	factorials := make([]int, 0)
	for i := 0; i < n; i++ {
		factorials = append(factorials, factorial(i))
	}
	return factorials
}

func sum_digit_factorials(factorials []int, digits []int) int {
	s := 0
	for _, digit := range digits {
		s += factorials[digit]
	}
	return s
}

func main() {
	factorials := calculate_factorials(10)
	sum := 0
	limit := factorials[9] * 7
	for i := 10; i < limit; i++ {
		if sum_digit_factorials(factorials, to_digits(i)) == i {
			sum += i
		}
	}
	fmt.Println(sum)
}