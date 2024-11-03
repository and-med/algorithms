package main

import (
	"fmt"
	"math"
)

func main() {
	n := 21
	i := 7
	div := 0
	for div < 500 {
		n += i
		i++
		limit := int(math.Sqrt(float64(n)))
		div = 0
		for j := 1; j <= limit; j++ {
			if (n % j == 0) {
				if j != n / j {
					div += 2
				} else {
					div += 1
				}
			}
		}
	}

	fmt.Println(n)
}