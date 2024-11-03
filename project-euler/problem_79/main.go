package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func readKeys() []string {
	file, err := os.Open("p079_keylog.txt")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	keys := []string {}
	for scanner.Scan() {
		keys = append(keys, scanner.Text())
	}
	return keys
}

func intValue(b byte) int {
	return int(b - '0')
}

func solve() int64 {
	keys := readKeys()
	orders := [][]int {}
	for _, key := range keys {
		orders = append(orders, []int { intValue(key[0]), intValue(key[1]) })
		orders = append(orders, []int { intValue(key[1]), intValue(key[2]) })
	}
	sort.Slice(orders, func(i, j int) bool {
		a, b := orders[i], orders[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	unique := map[int] bool {}
	unique_orders := [][]int {}
	for _, order := range orders {
		code := order[0]*100 + order[1]
		if _, ok := unique[code]; !ok {
			unique[code] = true
			unique_orders = append(unique_orders, order)
		}
	}
	for _, order := range unique_orders {
		fmt.Println(order)
	}
	return -1
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}

// 73162890