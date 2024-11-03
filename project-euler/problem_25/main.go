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

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func add(a []int, b []int) []int {
	var res []int
	rem := 0
	
	for i := 0; i < max(len(a), len(b)); i++ {
		s := 0
		if i >= len(a) {
			s = b[i] + rem
		} else if i >= len(b) {
			s = a[i] + rem
		} else {
			s = a[i] + b[i] + rem
		}
		res = append(res, s % 10)
		rem = s / 10
	}

	if rem != 0 {
		res = append(res, rem)
	}

	return res
}

func main() {
	defer elapsed("main")()
	f1 := []int { 1 }
	f2 := []int { 1 }

	id := 2
	for ; len(f2) < 1000; id++ {
		f1, f2 = f2, add(f1, f2)
	}

	fmt.Println(id)
}
