package main

import "fmt"

func nextPermutation(perm []int) bool {
	// find the rightmost i such that perm[i] < perm[i + 1]
	i := len(perm) - 2
	for ; i >= 0; i-- {
		if perm[i] < perm[i + 1] {
			break;
		}
	}

	// if no such i, means that the sequence is decreasing
	// hence there's no next permutation
	if i == -1 {
		return false
	}

	// find min(perm[j]) such that j > i and perm[i] < perm[j]
	// in other words: the smallest element on the right of i-th element that is greater than i-th element
	min_perm := i + 1
	for j := i + 2 ; j < len(perm); j++ {
		if perm[i] < perm[j] && perm[j] < perm[min_perm] {
			min_perm = j
		}
	}

	// swap i and j
	perm[i], perm[min_perm] = perm[min_perm], perm[i]

	// now sort everything from i + 1 to the end
	// since the sequence from i + 1 to the end is decreasing, we can just reverse it
	for j := i + 1; j < len(perm) - (len(perm) - i - 1) / 2; j++ {
		right := len(perm) - (j - i);
		perm[j], perm[right] = perm[right], perm[j]
	}

	return true
}

func main() {
	perm := []int { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }

	for i := 0; i < 999999; i++ {
		nextPermutation(perm)
	}

	for i := 0; i < len(perm); i++ {
		fmt.Print(perm[i])
	}
	fmt.Println()
}