package main

import "fmt"

func cnt(m map[int]string, num int) (int) {
	d := num % 10
	num /= 10
	t := num % 10
	num /= 10
	h := num % 10
	num /= 10
	th := num % 10

	l := 0
	if th != 0 {
		l += len(m[th]) + len("thousand")
	}
	if h != 0 {
		l += len(m[h]) + len("hundred")
		if t != 0 || d != 0 {
			l+= len("and")
		}
	}
	if t != 0 {
		if t == 1 {
			l += len(m[d + 10])
		} else {
			l += len(m[t * 10])
		}
	}
	if d != 0 && t != 1 {
		l += len(m[d])
	}
	return l
}

func main() {
	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	m[4] = "four"
	m[5] = "five"
	m[6] = "six"
	m[7] = "seven"
	m[8] = "eight"
	m[9] = "nine"
	m[10] = "ten"
	m[11] = "eleven"
	m[12] = "twelve"
	m[13] = "thirteen"
	m[14] = "fourteen"
	m[15] = "fifteen"
	m[16] = "sixteen"
	m[17] = "seventeen"
	m[18] = "eighteen"
	m[19] = "nineteen"
	m[20] = "twenty"
	m[30] = "thirty"
	m[40] = "forty"
	m[50] = "fifty"
	m[60] = "sixty"
	m[70] = "seventy"
	m[80] = "eighty"
	m[90] = "ninety"

	l := 0
	for i := 1; i <= 1000; i++ {
		l += cnt(m, i)
	}

	fmt.Println(l)
}