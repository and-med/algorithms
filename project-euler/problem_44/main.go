package main

import (
	"fmt"
	"math"
	"time"
)

func elapsed(what string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", what, time.Since(start))
    }
}

func real_root_exist(d int64) (bool, int64) {
	r := int64(math.Sqrt(float64(d)))
	return r * r == d && ((1 + r) % 6) == 0, (1 + r) / 6
}

func check_pentagonal(k int64, n int64) (bool, int64) {
	exist_a, _ := real_root_exist(1 + 12*(3*(n*n + k*k) - (n + k)))
	exist_b, _ := real_root_exist(1 + 12*(3*(n*n - k*k) - (n - k)))
	
	return exist_a && exist_b, (3*(n*n - k*k) - (n - k))/2
}

func main() {
	defer elapsed("main")()
	limit := 10000
	min_diff := int64(math.MaxInt64)
	for i := 1; i <= limit; i++ {
		pentagonal, diff := false, int64(0)
		for j := i + 1; j <= limit && diff < min_diff; j++ {
			pentagonal, diff = check_pentagonal(int64(i), int64(j))
			if pentagonal && diff < min_diff {
				min_diff = diff
			}
		}
	}
	fmt.Println(min_diff)
}

// p = n(3n - 1)/2

// 0 = 3n^2 - n - 2p

// n = (1 + sqrt(1 + 24p))/6

// n(3n - 1)/2 + k(3k - 1)/2 = (3(k^2 + n^2) - k - n)/2

// n > k
// 3n^2 + 3k^2 - n - k = p(3p - 1)

// a > n > k
// n > b
// n, k, a, b > 0 && are natural
// 3n^2 + 3k^2 - n - k = 3a^2 - a
// 3n^2 - 3k^2 - n + k = 3b^2 - b

// 3a^2 - a - (3n^2 + 3k^2 - n - k) = 0
// 3b^2 - b - (3n^2 - 3k^2 - n + k) = 0

// a = (1 + sqrt(1 + 4*3*(3n^2 + 3k^2 - n - k)))/6
// b = (1 + sqrt(1 + 4*3*(3n^2 - 3k^2 - n + k)))/6

// a = (1 + sqrt(1 + 12*(3(n^2 + k^2) - (n + k))))/6
// b = (1 + sqrt(1 + 12*(3(n^2 - k^2) - (n - k))))/6

// min over diff = (3(n^2 - k^2) - (n - k)) / 2

// 1 + 12*(3(n^2 + k^2) - (n + k)) = 36(n^2 + 2nk + k^2) - 12(n + k) + 1 - 72nk = (6(n + k) - 1)^2 - 72n*k