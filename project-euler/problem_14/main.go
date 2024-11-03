package main

import "fmt"

func main() {
	var maxColl int64 = 0
	max := 0
	for n := int64(5); n < 1000000; n++ {
		coll := n
		len := 0
		for coll != 1 {
			len++
			if coll % 2 == 0 {
				coll = coll / 2
			} else {
				coll = 3 * coll + 1
			}
		}
		if len > max {
			max = len
			maxColl = n
		}
	}
	fmt.Println(maxColl)
}