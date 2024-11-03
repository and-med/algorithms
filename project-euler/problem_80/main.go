package main

import (
	"fmt"
	"math/big"
	"time"
)

var limit int = 100

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func decimal_sum(n int) int {
	s := big.NewFloat(float64(n))
	s.SetPrec(1000)

	text := s.Sqrt(s).Text('f', 1000)
	sum := 0
	cnt := 0
	for _, b := range text {
		if b != '.' {
			sum += int(b - '0')
			cnt++
		}

		if cnt == 100 {
			break
		}
	}

	return sum
}

func solve() int64 {
	sum := 0
	for i := 2; i <= limit; i++ {
		dec_sum := decimal_sum(i)
		if dec_sum > 10 {
			sum += dec_sum
		}
	}
	return int64(sum)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}