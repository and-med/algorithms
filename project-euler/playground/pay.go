package main

import (
	"fmt"
	"time"
)

func main() {
	start, _ := time.Parse("01-02-2006", "09-20-2021")
	end, _ := time.Parse("01-02-2006", "11-22-2021")
	sum := 5656 + 490

	for start.Before(end) {
		fmt.Println(start.Format("01-02-2006"), sum)
		sum += 490
		start = start.Add(time.Hour * 24 * 7)
	}
}