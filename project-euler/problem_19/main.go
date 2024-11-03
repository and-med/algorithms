package main

import "fmt"

func getDaysInMonth(month int, year int) int {
	thirtyDayMonths := []int { 4, 6, 9, 11 }
	days := 31

	for _, m := range thirtyDayMonths {
		if m == month {
			days = 30
		}
	}

	if month == 2 {
		if year % 4 == 0 {
			if year % 100 == 0 {
				if year % 400 == 0 {
					days = 29
				} else {
					days = 28
				}
			} else {
				days = 29
			}
		} else {
			days = 28
		}
	}

	return days
}

func main() {
	day := 7
	month := 1
	year := 1900
	days := getDaysInMonth(month, year)

	cnt := 0
	for year < 2001 {
		day += 7
		month = month + day / (days + 1)
		day = day % (days + 1) + day / (days + 1)
		year = year + month / 13
		month = month % 13 + month / 13
		days = getDaysInMonth(month, year)

		if year > 1900 && year < 2001 && day == 1 {
			cnt++
		}
	}

	fmt.Println(cnt)
}