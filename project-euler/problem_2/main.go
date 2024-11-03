package main

import "fmt"

func main() {
	fib1, fib2, s := 1, 2, 0
	for fib1 < 4000000 {
		if fib1 % 2 == 0 {
			s += fib1
		}

		fib1, fib2 = fib2, fib1 + fib2
	}
	
	fmt.Println(s)
}