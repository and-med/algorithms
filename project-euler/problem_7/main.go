package main

import (
	"fmt"
	"math"
)

func main() {
	count := 1
	i := 3
	for count < 10001 {
		factors := 0
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if (i % j == 0) {
				factors++
				break;
			}
		}
		if (factors == 0) {
			count += 1
		}
		i += 2
	}
	fmt.Println(i - 2)
}