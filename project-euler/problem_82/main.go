package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	file, err := os.Open("p082_matrix.txt")
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

func min(params... int) int {
	min := math.MaxInt32
	for _, x := range params {
		if x < min {
			min = x
		}
	}
	return min
}

func solve() int64 {
	matrix := readMatrix()
	size := len(matrix)

	right_bottom_min := make([][]int, size)
	for i := 0; i < size; i++ {
		right_bottom_min[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i - 1 < 0 && j - 1 < 0 {
				right_bottom_min[i][j] = matrix[i][j]
			} else if i - 1 < 0 {
				right_bottom_min[i][j] = matrix[i][j] + right_bottom_min[i][j - 1]
			} else if j - 1 < 0 {
				right_bottom_min[i][j] = matrix[i][j] + right_bottom_min[i - 1][j]
			} else {
				right_bottom_min[i][j] = matrix[i][j] + min(right_bottom_min[i - 1][j],  right_bottom_min[i][j - 1])
			}
		}
	}

	right_top_min := make([][]int, size)
	for i := 0; i < size; i++ {
		right_top_min[i] = make([]int, size)
	}
	
	for i := size - 1; i >= 0; i-- {
		for j := 0; j < size; j++ {
			if i + 1 > size - 1 && j - 1 < 0 {
				right_top_min[i][j] = matrix[i][j]
			} else if i + 1 > size - 1 {
				right_top_min[i][j] = matrix[i][j] + right_top_min[i][j - 1]
			} else if j - 1 < 0 {
				right_top_min[i][j] = matrix[i][j] + right_top_min[i + 1][j]
			} else {
				right_top_min[i][j] = matrix[i][j] + min(right_top_min[i + 1][j],  right_top_min[i][j - 1])
			}
		}
	}

	sol := make([][]int, size)
	for i := 0; i < size; i++ {
		sol[i] = make([]int, size)
	}

	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			if j == 0 {
				sol[i][j] = matrix[i][j]
			} else if i == 0 {
				sol[i][j] = matrix[i][j] + min(sol[i][j - 1], right_top_min[i + 1][j])
			} else if i == size - 1 {
				sol[i][j] = matrix[i][j] + min(sol[i][j - 1], right_bottom_min[i - 1][j])
			} else if j == size - 1 {
				sol[i][j] = matrix[i][j] + sol[i][j - 1]
			} else {
				sol[i][j] = matrix[i][j] + min(sol[i][j - 1], right_top_min[i + 1][j], right_bottom_min[i - 1][j])
			}
		}
	}

	min := math.MaxInt32
	for i := 0; i < size; i++ {
		if sol[i][size - 1] < min {
			min = sol[i][size - 1]
		}
	}

	return int64(min)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}