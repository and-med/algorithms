package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	file, err := os.Open("p067_triangle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var tr [][]int
	for scanner.Scan() {
		var r []int
		for _, num := range(strings.Split(scanner.Text(), " ")) {
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

	r := 0
	for i := 0; i < len(tr); i++ {
		r = max(r, tr[len(tr) - 1][i])
	}
	fmt.Println(r)
}