package main

import "fmt"

func main() {
	num := make([]int, 1)
	num[0] = 2
	l := 1
	pow := 63

	for i := 1; i < pow; i++ {
		p := 1
		rem := 0
		for j := 0; j < l; j++ {
			p = num[j] * 2 + rem
			num[j] = p % 10
			rem = p / 10
		}
		if rem != 0 {
			num = append(num, rem)
			l++
		}
	}

	s := 0
	for i := l - 1; i >= 0; i-- {
		fmt.Print(num[i])
		s += num[i]
	}
	fmt.Println()
	fmt.Println(s)
}
