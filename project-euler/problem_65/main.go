package main

import (
	"fmt"
	"math/big"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func get_euler_seq(i int) int {
	switch {
	case i == 1:
		return 2
	case i%3 == 0:
		return (i / 3) * 2
	}
	return 1
}

func solve() int64 {
	limit := 100
	ratio := big.NewRat(1, int64(get_euler_seq(limit)))
	for i := limit - 1; i > 0; i-- {
		ratio.Add(ratio, big.NewRat(int64(get_euler_seq(i)), 1)).Inv(ratio)
	}
	sum := 0
	for _, d := range ratio.Denom().Text(10) {
		sum += int(d - '0')
	}
	return int64(sum)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}
