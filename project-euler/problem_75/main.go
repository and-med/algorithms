package main

import (
	"fmt"
	"sort"
	"time"
)

var triplets map[int][]int

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

func create_triplet(a, b, c int) {
	t := []int { a, b, c }
	sort.Ints(t)

	curr_triplets := triplets[a + b + c]
	exist := false
	for i := 0; i < len(curr_triplets); i += 3 {
		if curr_triplets[i] == t[0] && curr_triplets[i + 1] == t[1] && curr_triplets[i + 2] == t[2] {
			exist = true
			break
		}
	}
	if !exist {
		triplets[a + b + c] = append(triplets[a+b+c], t...)
	}
}

func generate_triples() {
	limit := 1500000
	triplets = map[int][]int {}
	for m := 2; 2 * m * m + 2 * m <= limit; m++ {
		for n := 1; n < m && 2 * m * m + 2 * m * n <= limit; n++ {
			if gcd(m, n) == 1 {
				a, b, c := m * m - n * n, 2 * m * n, m * m + n * n
				create_triplet(a, b, c)
				for i := 2; i*a + i*b + i*c <= limit; i++ {
					create_triplet(i*a, i*b, i*c)
				}
			}
		}
	}
}

func solve() int64 {
	generate_triples()
	cnt := 0
	for _, t := range triplets {
		if len(t) == 3 {
			cnt++
		}
	}
	return int64(cnt)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}