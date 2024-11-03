package main

import "fmt"

func main() {
	limit := 1001
	fmt.Println((4 * limit * limit * limit + 3 * limit * limit + 8 * limit - 9) / 6)
}

// 73 74 75 76 77 78 79 80 81
// 72 43 44 45 46 47 48 49 50
// 71 42 21 22 23 24 25 26 51
// 70 41 20  7  8  9 10 27 52 
// 69 40 19  6  1  2 11 28 53
// 68 39 18  5  4  3 12 29 54 
// 67 38 17 16 15 14 13 30 55
// 66 37 36 35 34 33 32 31 56
// 65 64 63 62 61 60 59 58 57


// 4n - 3 + 16/3 * (n - 2) * (n - 1) * n + 10n * (n - 1)
// 10n^2 - 6n - 3 + 16/3 * (n - 2) * (n - 1) * n
// (16n^3 - 18n^2 + 14n - 9) / 3, where n = (spiralSize + 1) / 2
// or in other words: (4 * spiralSize^3 + 3 * spiralSize^2 + 8 * spiralSize - 9) / 6