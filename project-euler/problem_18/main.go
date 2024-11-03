package main

import (
	"fmt"
	"strconv"
	"strings"
)

func max(a int, b int) int {
	if a > b { 
		return a
	}
	return b
}

func main() {
	t := `75
	95 64
	17 47 82
	18 35 87 10
	20 04 82 47 65
	19 01 23 75 03 34
	88 02 77 73 07 63 67
	99 65 04 28 06 16 70 92
	41 41 26 56 83 40 80 70 33
	41 48 72 33 47 32 37 16 94 29
	53 71 44 65 25 43 91 52 97 51 14
	70 11 33 28 77 73 17 78 39 68 17 57
	91 71 52 38 17 14 91 43 58 50 27 29 48
	63 66 04 68 89 53 67 30 73 16 69 87 40 31
	04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`
	var tr [][]int

	for _, row := range(strings.Split(t, "\n")) {
		var r []int
		for _, num := range(strings.Split(strings.TrimSpace(row), " ")) {
			a, _ := strconv.Atoi(strings.TrimSpace(num))
			r = append(r, a)
		}
		tr = append(tr, r)
	}
	
	for i := 1; i < len(tr); i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				tr[i][j] = tr[i][j] + tr[i - 1][j]
			} else if j == i {
				tr[i][j] = tr[i][j] + tr[i - 1][j - 1]
			} else {
				tr[i][j] = tr[i][j] + max(tr[i - 1][j - 1], tr[i - 1][j])
			}
		}
	}

	for i := 0; i < len(tr); i++ {
		for j := 0; j < (len(tr) - i - 1)*3; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print(tr[i][j], " ")
			if tr[i][j] < 1000 {
				fmt.Print("  ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}