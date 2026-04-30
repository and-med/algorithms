package main

import (
	"fmt"
	"slices"
)

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	slices.Sort(nums)

	return true
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
}
