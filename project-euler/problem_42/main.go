package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func parseWords(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		text := scanner.Text()
		quoted_words := strings.Split(text, ",")
		words := []string {}
		for _, word := range quoted_words {
			words = append(words, strings.Trim(word, "\""))
		}
		return words
	}
	return nil
}

func get_word_value(word string) int {
	s := 0
	for _, char := range word {
		s += int(char - 'A') + 1
	}
	return s
}

func is_triangle_number(n int) bool {
	root := int(math.Sqrt(float64(n * 2)))
	return (root * (root + 1)) / 2 == n
}

func main() {
	words := parseWords("words.txt")
	cnt := 0
	for _, word := range words {
		if is_triangle_number(get_word_value(word)) {
			cnt++
		}
	}
	fmt.Println(cnt)
}