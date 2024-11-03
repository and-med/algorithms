package main

import "fmt"

func is_palindrom_ten(n int) bool {
	num := n
	reverse := 0
	for num != 0 {
		reverse *= 10
		reverse += num % 10
		num /= 10
	}
	return reverse == n
}

func is_palindrom_binary(n int) bool {
	num := n
	binary := make([]int, 0)
	for num != 0 {
		binary = append(binary, num % 2)
		num /= 2
	}

	for i := 0; i < len(binary) / 2; i++ {
		if binary[i] != binary[len(binary) - i - 1] {
			return false
		}
	}

	return true
}

func main() {
	s := 0
	for i := 1; i < 1000000; i++ {
		if is_palindrom_ten(i) && is_palindrom_binary(i) {
			s += i
		}
	}
	fmt.Println(s)
}
