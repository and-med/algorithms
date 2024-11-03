package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func readMatrix() [][]int {
	file, err := os.Open("p081_matrix.txt")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := [][]int {}
	for scanner.Scan() {
		row := []int {}
		for _, text := range strings.Split(scanner.Text(), ",") {
			value, _ := strconv.Atoi(text)
			row = append(row, value)
		}

		matrix = append(matrix, row)
	}
	return matrix
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve() int64 {
	matrix := readMatrix()
	size := len(matrix)
	sol := make([][]int, size)
	for i := 0; i < size; i++ {
		sol[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i - 1 < 0 && j - 1 < 0 {
				sol[i][j] = matrix[i][j]
			} else if i - 1 < 0 {
				sol[i][j] = matrix[i][j] + sol[i][j - 1]
			} else if j - 1 < 0 {
				sol[i][j] = matrix[i][j] + sol[i - 1][j]
			} else {
				sol[i][j] = matrix[i][j] + min(sol[i - 1][j],  sol[i][j - 1])
			}
		}
	}

	return int64(sol[size - 1][size - 1])
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}