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

func read_bytes() []rune {
	file, err := os.Open("p059_cipher.txt")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bytes := []rune {}
	if scanner.Scan() {
		byte_str := strings.Split(scanner.Text(), ",")
		for _, b := range byte_str {
			conv, _ := strconv.Atoi(b)
			bytes = append(bytes, rune(conv))
		}
	}
	return bytes
}

func apply_cipher(text []rune, cipher []rune) []rune {
	new_text := []rune {}
	for i := 0; i < len(text); {
		for j := 0; j < len(cipher) && i < len(text); j, i = j + 1, i + 1 {
			new_text = append(new_text, text[i] ^ cipher[j])
		}
	}
	return new_text
}

func solve() int64 {
	bytes := read_bytes()
	decr := apply_cipher(bytes, []rune { 'e', 'x', 'p' })
	fmt.Println(string(decr))
	sum := 0
	for _, b := range decr {
		sum += int(b)
	}
	return int64(sum)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}


// res = 'e' 'x' 'p'