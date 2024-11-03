package main

import "fmt"

func main() {
	s := 0;
	n := 100;
	for i := 1; i <= n; i++ {
		for j := 1; j<= n; j++ {
			if (i != j) {
				s += i * j;
			}
		}
	}
	fmt.Println(s)
}
