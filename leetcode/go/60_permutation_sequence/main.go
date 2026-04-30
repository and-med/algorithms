package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSet int

// Creates the set of natural numbers from 1 to n (inclusive)
func NewIntSet(n int) *IntSet {
	available := 0
	for i := 0; i < n; i++ {
		available |= (1 << i)
	}
	set := IntSet(available)
	return &set
}

// Removes the element at position i from the set maintaining the order
func (s *IntSet) Rem(pos int, n int) int {
	underlying := *((*int)(s))
	curr := 0
	for i := 0; i < n; i++ {
		offset := 1 << i
		if (underlying & offset) > 0 {
			if pos == curr {
				underlying = underlying ^ offset
				*s = IntSet(underlying)
				return i + 1
			}
			curr++
		}
	}
	
	panic("element not found")
}

type Factorial []int

// Returns a type that contains cached factorials
func NewFactorial(n int) *Factorial {
	factorials := make([]int, 10)

	factorials[0] = 1
	factorials[1] = 1

	for i := 2; i <= 9; i++ {
		factorials[i] = factorials[i-1] * i
	}

	factorial := Factorial(factorials)

	return &factorial
}

// Returns the factorial from cache
func (f *Factorial) Of(i int) int {
	return (*f)[i]
}

type Permutation strings.Builder

func NewPermutation() *Permutation {
	permutation := Permutation(strings.Builder{})
	return &permutation
}

func (p *Permutation) Append(s int) {
	underlying := (*strings.Builder)(p)
	underlying.WriteString(strconv.Itoa(s))
}

func (p *Permutation) String() string {
	underlying := (*strings.Builder)(p)
	return underlying.String()
}

func getPermutation(n int, k int) string {
	factorial := NewFactorial(n)
	perm := NewPermutation()
	set := NewIntSet(n)

	k = k - 1

	for i := n - 1; i >= 1; i-- {
		perm.Append(set.Rem(k / factorial.Of(i), n))
		k = k % factorial.Of(i)
	}

	perm.Append(set.Rem(0, n))
	return perm.String()
}

func main() {
	fmt.Println(getPermutation(4, 9))
}

// 1.  1234
// 2.  1243
// 3.  1324
// 4.  1342
// 5.  1423
// 6.  1432

// 7.  2134
// 8.  2143
// 9.  2314
// 10. 2341
// 11. 2413
// 12. 2431

// 13. 3124
// 14. 3142
// 15. 3214
// 16. 3241
// 17. 3412
// 18. 3421

// 19. 4123
// 20. 4132
// 21. 4213
// 22. 4231
// 23. 4312
// 24. 4321
