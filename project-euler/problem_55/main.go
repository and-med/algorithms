package main

import (
	"fmt"
	"time"
)

const LychrelIterations int = 50

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func to_digits(n int) []byte {
	d := []byte {}
	for n != 0 {
		d = append(d, byte(n % 10))
		n /= 10
	}
	return d
}

func reverse(num []byte) []byte {
	res := make([]byte, len(num))
	for i := 0; i < len(num); i++ {
		res[i] = num[len(num) - i - 1]
	}
	return res
}

func is_palindrome(n []byte) bool {
	for i := 0; i < len(n) / 2; i++ {
		if n[i] != n[len(n) - i - 1] {
			return false
		}
	}
	return true
}

func add(a []byte, b []byte) []byte {
	res := []byte {}
	rem := byte(0)
	for i := 0; i < len(a); i++ {
		s := a[i] + b[i] + rem
		res = append(res, s % 10)
		rem = s / 10
	}
	if rem != 0 {
		res = append(res, rem)
	}
	return res
}

func lychrel_process(n int) int {
	num := to_digits(n)
	for i := 0; i < LychrelIterations; i++ {
		rev := reverse(num)
		sum := add(num, rev)
		if is_palindrome(sum) {
			return i + 1
		}
		num = sum
	}
	return -1
}

func solve() int64 {
	cnt := 0
	for i := 1; i < 10000; i++ {
		if lychrel_process(i) == -1 {
			cnt++
		}
	}
	return int64(cnt)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}