package main

import "fmt"

func main() {
	var p []bool
	n := 2000000
	for i := 0; i < n; i++ {
		p = append(p, true)
	}

	p[0] = false
	p[1] = false

	for i := 2; i < n; i++ {
		if (p[i]) {
			for j := 2 * i; j < n; j += i {
				p[j] = false
			}
		}
	}

	s := 0
	for i := 0; i < n; i++ {
		if p[i] {
			s += i
		}
	}
	fmt.Println(s)
}