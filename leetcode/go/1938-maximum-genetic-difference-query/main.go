package main

import (
	"fmt"
	"math"
)

func manualPow(a int32, pow int32) int32 {
	res := int32(1)
	for i := int32(0); i < pow; i++ {
		res *= a
	}
	return res
}

func manualPowSum(a int32, pow int32) int32 {
	res := int32(1)
	product := int32(1)

	for i := int32(0); i < pow; i++ {
		product *= a
		res += product
	}

	return res
}

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	fmt.Println(manualPow(2, 10), manualPowSum(2, 10))
	fmt.Println(manualPow(2, 30), manualPowSum(2, 30))
}
