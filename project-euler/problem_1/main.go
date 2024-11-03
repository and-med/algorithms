package main

import "fmt"

func main() {
	s := 0
	for n := 3; n < 1000; n++ {
		if (n % 3 == 0) {
			s += n
		} else if (n % 5 == 0) {
			s += n
		}
	}
	fmt.Println(s)
}