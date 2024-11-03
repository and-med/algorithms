package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func get_score(word string) int {
	score := 0
	for _, letter := range word {
		score += int(letter) - int('A') + 1
	}
	return score
}

func main() {
	file, err := os.Open("names.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	var names []string
	if scanner.Scan() {
		names_str := scanner.Text()
		for _, name := range strings.Split(names_str, ",") {
			names = append(names, strings.Trim(name, "\""))
		}
	}

	sort.Strings(names)
	res := 0
	for i := 0; i < len(names); i++ {
		res += (i + 1) * get_score(names[i])
	}
	fmt.Println(res)
}