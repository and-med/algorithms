package main

import (
	"fmt"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func pow_mod(n int64, power int, mod int64) int64 {
	res := int64(1)
	for power != 0 {
		res = (res * n) % mod
		power -= 1
	}
	return res
}

func main() {
	defer elapsed("main")()
	mod := int64(10000000000)
	res := int64(0)
	for i := 1; i <= 1000; i++ {
		if (i % 10 != 0) {
			res = (res + pow_mod(int64(i), i, mod)) % mod
		}
	}
	fmt.Println(res)
}
