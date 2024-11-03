package main

import "fmt"

func main() {
	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000 - a - b; b++ {
			c := 1000 - a - b
			if (a * a + b * b == c * c) {
				fmt.Println(a, b, c)
				fmt.Println(a*b*c)
				return
			}
		}
	}
}