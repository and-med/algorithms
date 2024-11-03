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

func multiply_by_digit(a []byte, b byte) []byte {
	res := make([]byte, len(a) + 1)
	rem := byte(0)
	for idx, digit := range a {
		p := digit * b + rem
		res[idx] = p % 10
		rem = p / 10
	}
	if rem != 0 {
		res[len(a)] = rem
	} else {
		res = res[:len(a)]
	}
	return res
}

func multiply(a []byte, b []byte) []byte {
	if len(b) > len(a) {
		a, b = b, a
	}

	partial_res := make([][]byte, len(b))
	for idx, digit := range b {
		p := multiply_by_digit(a, digit)
		partial_res[idx] = p
	}

	max_length := 0
	for i := 0; i < len(partial_res); i++ {
		if len(partial_res[i]) + i > max_length {
			max_length = len(partial_res[i]) + i
		}
	}

	res := []byte {}
	rem := 0
	for i := 0; i < max_length; i++ {
		s := rem
		for j := 0; j < len(partial_res); j++ {
			if i - j >= 0 {
				if i - j < len(partial_res[j]) {
					s += int(partial_res[j][i - j])
				}
			} else {
				break
			}
		}
		res = append(res, byte(s % 10))
		rem = s / 10
	}

	if rem != 0 {
		res = append(res, byte(rem))
	}

	return res
}

func to_digits(n int) []byte {
	d := []byte {}
	for n != 0 {
		d = append(d, byte(n % 10))
		n /= 10
	}
	return d
}

func power(n int, pow int) []byte {
	res := []byte { 1 }
	curr := to_digits(n)

	for pow != 0 {
		if pow % 2 == 1 {
			res = multiply(res, curr)
		}
		curr = multiply(curr, curr)
		pow /= 2
	}

	return res
}

func sum(digits []byte) int {
	s := 0
	for _, digit := range digits {
		s += int(digit)
	}
	return s
}

func solve() int64 {
	max_sum := 0
	max_a := 1
	max_b := 1
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			s := sum(power(a, b))
			if s > max_sum {
				max_sum = s
				max_a = a
				max_b = b
			}
		}
	}
	fmt.Println(max_a, "^", max_b, "=", max_sum)
	return int64(max_sum)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}