package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func length(n int) int {
	return int(math.Log10(float64(n)))
}

func solve() int64 {
	iterations, nom, denom, cnt := 1000, big.NewInt(0), big.NewInt(1), 0
	for i := 0; i < iterations; i++ {
		new_denom := big.NewInt(2)
		new_denom.Mul(new_denom, denom)
		new_denom.Add(new_denom, nom)
		nom, denom = denom, new_denom
		new_nom := big.NewInt(0)
		new_nom.Add(nom, denom)

		if len(new_nom.String()) > len(denom.String()) {
			cnt++
		}
	}
	return int64(cnt)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}