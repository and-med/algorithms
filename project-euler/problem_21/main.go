package main

import "fmt"

func main() {
	anum := make([]int, 2)
	limit := 10000

	for i := 2; i < limit; i++ {
		sum := 0
		for j := 1; j <= i / 2; j++ {
			if i % j == 0 {
				sum += j
			}
		}
		anum = append(anum, sum)
	}

	ami_nums := make(map[int]bool)
	for i := 2; i < limit; i++ {
		if anum[i] < limit && anum[anum[i]] == i && i != anum[i] {
			ami_nums[i] = true
			ami_nums[anum[i]] = true
		}
	}

	s := 0
	for key := range ami_nums {
		s += anum[key]
	}
	fmt.Println(s)
}