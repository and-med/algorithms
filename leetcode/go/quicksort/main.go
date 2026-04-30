package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func quicksort_internal(arr []int, start int, end int, compare func(a, b int) int) {
	if end-start <= 12 {
		for i := start + 1; i < end; i++ {
			for j := i; j > start && compare(arr[j-1], arr[j]) < 0; j-- {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		return
	}

	pivotIdx := rand.Intn(end-start) + start
	arr[start], arr[pivotIdx] = arr[pivotIdx], arr[start]

	left, right := start+1, end-1

	for left < right {
		if compare(arr[start], arr[left]) < 0 {
			// swap left and right
			left++
		} else {
			arr[left], arr[right] = arr[right], arr[left]
			right--
		}
	}

	if compare(arr[start], arr[left]) < 0 {
		arr[start], arr[left] = arr[left], arr[start]
	}

	quicksort_internal(arr, start, left, compare)
	quicksort_internal(arr, left, end, compare)
}

func quicksort(arr []int) {
	quicksort_internal(arr, 0, len(arr), func(a int, b int) int {
		return b - a
	})
}

func main() {
	n := 1000000
	max_value := math.MaxInt32
	test_cases := 100

	cnt_passed := 0
	for i := 0; i < test_cases; i++ {
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			nums[j] = rand.Intn(max_value)
		}

		nums_copy := make([]int, n)
		copy(nums_copy, nums)

		startLibrary := time.Now()
		sort.Ints(nums)
		libraryElapsed := time.Since(startLibrary)

		startMy := time.Now()
		quicksort(nums_copy)
		myElapsed := time.Since(startMy)

		equal := true
		for j := 0; j < n; j++ {
			if nums[j] != nums_copy[j] {
				equal = false
				break
			}
		}

		if equal {
			fmt.Println("Test passed", i, "library time:", libraryElapsed, "my time:", myElapsed)
			cnt_passed++
		}
	}

	fmt.Println("Total tests:", test_cases, "passed tests:", cnt_passed)
}
