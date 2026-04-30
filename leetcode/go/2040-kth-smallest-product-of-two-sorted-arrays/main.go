package main

import "fmt"

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	spaceMin := int64(-100000 * 100000)
	spaceMax := int64(100000 * 100000)
	return binarySearchSpace(nums1, nums2, spaceMin, spaceMax, k)
}

func findPositiveAndNegativeBoundaries(nums []int) (int64, int64, int64, int64) {
	nEnd := int64(0)
	for idx, num := range nums {
		if num < 0 {
			nEnd = int64(idx + 1)
		}
	}

	return int64(0), nEnd, nEnd, int64(len(nums))
}

func binarySearchSpace(nums1 []int, nums2 []int, start int64, end int64, needle int64) int64 {
	var binarySearchSpaceInner func(start int64, end int64) int64
	var binarySearchNumOfProducts func(nums1 int, product int64, start int64, end int64) int64
	var binarySearchNumOfProductsNormalOrder func(nums1 int, product int64, start int64, end int64) int64
	var binarySearchNumOfProductsReversedOrder func(nums1 int, product int64, start int64, end int64) int64
	var binarySearchNumOfProductsLessThan2 func(product int64) int64
	var binarySearchLeftBoundary func(left int64, right int64, smallest int64, product int64) int
	var binarySearchRightBoundary func(left int64, right int64, smallest int64, product int64) int

	n1negativeStart, n1negativeEnd, n1positiveStart, n1positiveEnd := findPositiveAndNegativeBoundaries(nums1)
	n2negativeStart, n2negativeEnd, n2positiveStart, n2positiveEnd := findPositiveAndNegativeBoundaries(nums2)

	binarySearchSpaceInner = func(start int64, end int64) int64 {
		// looking for the leftmost such that numOfProducts(currProduct) <= needle
		if start == end {
			//fmt.Println(fmt.Sprintf("exiting space search %d", start))
			return start
		}

		adjustment := 0

		if start <= 0 && end <= 0 {
			adjustment = -1
		}

		currProduct := (start + end + int64(adjustment)) / 2

		// numOfProducts := binarySearchNumOfProductsLessThan(nums1, nums2, currProduct)
		numOfProducts := binarySearchNumOfProductsLessThan2(currProduct)

		//fmt.Println("searching space from %d to %d, currProduct=%d, numOfProducts=%d", start, end, currProduct, numOfProducts)

		if numOfProducts < needle {
			// if less -> find in the right, exclude current (impossible to return current)
			return binarySearchSpaceInner(currProduct+1, end)
		}

		// else if numOfProducts >= needle
		// go left, exclude current
		return binarySearchSpaceInner(start, currProduct)
	}

	binarySearchNumOfProducts = func(num1 int, product int64, start int64, end int64) int64 {
		if start == end {
			return start
		}

		middle := (start + end) / 2

		currProduct := int64(num1) * int64(nums2[middle])

		if num1 >= 0 {
			if currProduct > product {
				// exclude middle, go left -> but still possible to return middle
				return binarySearchNumOfProducts(num1, product, start, middle)
			}

			// if currProduct <= product
			// exclude middle, go right -> impossible to return middle
			return binarySearchNumOfProducts(num1, product, middle+1, end)
		}

		// if num1 < 0
		// order is reversed now

		if currProduct <= product {
			// go left -> but still possible to return middle
			return binarySearchNumOfProducts(num1, product, start, middle)
		}

		return binarySearchNumOfProducts(num1, product, middle+1, end)
	}

	binarySearchNumOfProductsNormalOrder = func(num1 int, product int64, start int64, end int64) int64 {
		if start == end {
			return start
		}

		middle := (start + end) / 2

		currProduct := int64(num1) * int64(nums2[middle])

		if currProduct > product {
			// exclude middle, go left -> but still possible to return middle
			return binarySearchNumOfProducts(num1, product, start, middle)
		}

		// if currProduct <= product
		// exclude middle, go right -> impossible to return middle
		return binarySearchNumOfProducts(num1, product, middle+1, end)
	}

	binarySearchNumOfProductsReversedOrder = func(num1 int, product int64, start int64, end int64) int64 {
		if start == end {
			return start
		}

		middle := (start + end) / 2

		currProduct := int64(num1) * int64(nums2[middle])

		if currProduct <= product {
			// go left -> but still possible to return middle
			return binarySearchNumOfProducts(num1, product, start, middle)
		}

		return binarySearchNumOfProducts(num1, product, middle+1, end)
	}

	binarySearchLeftBoundary = func(left int64, right int64, smallest int64, product int64) int {
		if left == right {
			return int(left)
		}

		middle := (left + right) / 2

		currProduct := int64(nums1[int(middle)]) * smallest

		if currProduct <= product {
			return binarySearchLeftBoundary(left, middle, smallest, product)
		}

		return binarySearchLeftBoundary(middle+1, right, smallest, product)
	}

	binarySearchRightBoundary = func(left int64, right int64, smallest int64, product int64) int {
		if left == right {
			return int(left)
		}

		middle := (left + right) / 2

		currProduct := int64(nums1[int(middle)]) * smallest
		if currProduct <= product {
			return binarySearchRightBoundary(middle+1, right, smallest, product)
		}

		return binarySearchRightBoundary(left, middle, smallest, product)
	}

	binarySearchNumOfProductsLessThan2 = func(product int64) int64 {
		count := int64(0)
		if product >= 0 {
			// negative num1 + positive num2 - all combinations
			count += (n1negativeEnd - n1negativeStart) * (n2positiveEnd - n2positiveStart)

			// negative num1 + negative num2 - find ranges for nums1 for each el in nums1 - binary search nums2
			if n1negativeEnd > 0 && n2negativeEnd > 0 {
				smallest := int64(nums2[n2negativeEnd-1])
				left := binarySearchLeftBoundary(n1negativeStart, n1negativeEnd, smallest, product)
				right := int(n1negativeEnd)

				for i := left; i < right; i++ {
					count += n2negativeEnd - binarySearchNumOfProductsReversedOrder(nums1[i], product, n2negativeStart, n2negativeEnd)
				}
			}

			// positive num1 + negative num2 - all combinations
			count += (n1positiveEnd - n1positiveStart) * (n2negativeEnd - n2negativeStart)

			// positive num1 + positive num2 - find ranges for nums1 - for each el in nums1 - binary search nums2
			if n1positiveStart < int64(len(nums1)) && n2positiveStart < int64(len(nums2)) {
				smallest := int64(nums2[n2positiveStart])
				left := int(n1positiveStart)
				right := binarySearchRightBoundary(n1positiveStart, n1positiveEnd, smallest, product)

				//initial := count
				for i := left; i < right; i++ {
					num1 := nums1[i]
					count += binarySearchNumOfProductsNormalOrder(num1, product, n2positiveStart, n2positiveEnd) - n2positiveStart
				}

				//fmt.Printf("case pos num1, pos num2 for product %d and smallest %d: from %d to %d total %d\n", product, smallest, left, right, count-initial)
			}
		} else {
			// if product < 0

			// positive num1 + positive num2 - no combinations
			// negative num1 + negative num2 - no combinations

			// positive num1 + negative num2 - find ranges for num1 - for each el in nums1 - binary search nums2
			if n1positiveStart < int64(len(nums1)) && n2negativeEnd > 0 {
				smallest := int64(nums2[n2negativeStart])
				left := binarySearchLeftBoundary(n1positiveStart, n1positiveEnd, smallest, product)
				right := int(n1positiveEnd)

				for i := left; i < right; i++ {
					count += binarySearchNumOfProductsNormalOrder(nums1[i], product, n2negativeStart, n2negativeEnd) - n2negativeStart
				}
			}

			// negative num1 + positive num2 - find ranges for num1 - for each el in nums1 - binary search nums2
			if n1negativeEnd > 0 && n2positiveStart < int64(len(nums2)) {
				smallest := int64(nums2[n2positiveEnd-1])
				left := int(n1negativeStart)
				right := binarySearchRightBoundary(n1negativeStart, n1negativeEnd, smallest, product)

				for i := left; i < right; i++ {
					count += n2positiveEnd - binarySearchNumOfProductsReversedOrder(nums1[i], product, n2positiveStart, n2positiveEnd)
				}
			}
		}

		return count
	}

	return binarySearchSpaceInner(start, end)
}

func main() {
	fmt.Println(kthSmallestProduct([]int{-10, -9, -4, -2, -1, 7, 7, 9, 10}, []int{-10, -2, -2, 6, 8, 8}, 29))
}
