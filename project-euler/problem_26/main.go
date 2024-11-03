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

type LongDivision struct {
	Whole []int
	Quotient []int
	Cycle int
}

func newLongDivision() LongDivision {
	return LongDivision{
		Cycle: -1,
	}
}

func (ld LongDivision) cycleLength() int {
	if ld.Cycle == -1 {
		return 0
	}
	return len(ld.Quotient) - ld.Cycle
}

func (ld LongDivision) print() {
	for _, digit := range ld.Whole {
		fmt.Print(digit)
	}

	if (len(ld.Quotient) > 0) {
		fmt.Print(".")
	}

	for id, digit := range ld.Quotient {
		if id == ld.Cycle {
			fmt.Print("(")
		}
		fmt.Print(digit)
	}

	if ld.Cycle == -1 {
		fmt.Println()
	} else {
		fmt.Println("),", "cycle length:", len(ld.Quotient) - ld.Cycle)
	}

}

func divide(dividend int, divisor int) LongDivision {
	var whole []int
	var quot []int
	dividends := make(map[int]int)

	if dividend >= divisor {
		div := dividend / divisor;
		for div != 0 {
			whole = append(whole, div % 10)
			div /= 10
		}
		
		for i := 0; i < len(whole) / 2; i++ {
			whole[i], whole[len(whole) - i - 1] = whole[len(whole) - i - 1], whole[i]
		}

		dividend = dividend % divisor
	} else {
		whole = append(whole, 0)
	}

	for dividend != 0 {
		for dividend < divisor {
			elem, ok  := dividends[dividend]
			if ok {
				return LongDivision{whole, quot, elem}
			}

			dividends[dividend] = len(quot)
			quot = append(quot, 0)

			dividend *= 10
		}
	
		quot[len(quot) - 1] = dividend / divisor;
		dividend = dividend % divisor
	}
	
	return LongDivision{whole, quot, -1}
}

func main() {
	defer elapsed("main")()
	max := newLongDivision()
	max_digit := 0

	for i := 1; i < 1000; i++ {
		res := divide(1, i)
		
		if res.Cycle != -1 {
			if res.cycleLength() > max.cycleLength() {
				max_digit = i
				max = res
			}
		}
	}

	fmt.Println("1 /", max_digit)
	max.print()
}