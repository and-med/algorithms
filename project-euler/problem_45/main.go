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

func find_t_and_p(h int64) (bool, int64, int64) {
	d_t := 1 + 4*(4*h*h-2*h)
	d_p := 1 + 12*(4*h*h-2*h)

	root_d_t := int64(math.Sqrt(float64(d_t)))
	root_d_p := int64(math.Sqrt(float64(d_p)))

	exist_t, t := root_d_t*root_d_t == d_t && ((root_d_t-1)%2 == 0), (root_d_t-1)/2
	exist_p, p := root_d_p*root_d_p == d_p && ((1+root_d_p)%6 == 0), (1+root_d_p)/6

	return exist_t && exist_p, t, p
}

func main() {
	defer elapsed("main")()
	for h := int64(144); true; h++ {
		exist, t, p := find_t_and_p(h)
		if exist {
			fmt.Println(t, p, h)
			fmt.Println((t * (t + 1)) / 2)
			break
		}
	}
}

// t = n(n+1)/2
// p = n(3n-1)/2
// h = n(2n-1)

// t(t+1)/2 = h(2h-1)
// p(3p-1)/2 = h(2h-1)

// t^2+t = 4h^2-2h = 4(h^2 + h) - 6h
// 3p^2-p = 4h^2-2h

// t^2 + t - (4h^2-2h) = 0
// 3p^2 - p - (4h^2-2h) = 0

// t = (-1 + sqrt(1 + 4(4h^2-2h)))/2
// p = (1 + sqrt(1 + 12(4h^2-2h)))/6

// t^2+t = 4h^2-2h = 4(h^2 + h) - 6h
// 0.25(t^2 + t) = h^2 + h - 1.5h
