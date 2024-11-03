package main

import (
	"fmt"
	"math"
	"sort"
)

func length(n int) int {
	if n < 10 {
		return 1
	} else if n < 100 {
		return 2
	}
	return 3
}

func to_digits(n int) []int {
	return []int { n / 10, n % 10 }
}

func contains(arr []int, x int) bool {
	for _, a := range arr {
		if a == x {
			return true
		}
	}
	return false
}

func contains_int(n int, x int) bool {
	if n / 10 == x || n % 10 == x {
		return true
	}
	return false 
}

func other_digit(n int, dig int) int {
	if dig == n / 10 {
		return n % 10
	}
	return n / 10
}

func is_curious(nom int, denom int, new_nom int, new_denom int) bool {
	if length(new_nom) == 1 && 
		length(new_denom) == 1 && 
		contains_int(nom, new_nom) && 
		contains_int(denom, new_denom) &&
		other_digit(nom, new_nom) == other_digit(denom, new_denom) &&
		other_digit(nom, new_nom) != 0 {
		return true
	}

	return false
}

func get_curious_fractions(nom int, denom int, div int) []int {
	factor := 1
	new_nom := (nom / div) * factor
	new_denom := (denom / div) * factor

	for new_denom < 10 {
		if is_curious(nom, denom, new_nom, new_denom) {
			return []int { new_nom, new_denom }
		}
		factor++
		new_nom = (nom / div) * factor
		new_denom = (denom / div) * factor
	}

	return nil
}

func get_divisors(nom int, denom int) []int {
	unique := make(map[int]bool)

	for i := 1; i <= int(math.Sqrt(float64(denom))); i++ {
		if nom % i == 0 && denom % i == 0 {
			unique[i] = true
		}
		if nom % i == 0 && denom % (nom / i) == 0 {
			unique[nom / i] = true
		}
		if denom % i == 0 && nom % (denom / i) == 0 {
			unique[denom / i] = true
		}
	}

	res := make([]int, 0)
	for key := range unique {
		res = append(res, key)
	}

	sort.Ints(res)

	return res
}

func main() {
	var res [][]int 
	for denom := 11; denom < 100; denom++ {
		for nom := 10; nom < denom; nom++ {
			divisors := get_divisors(nom, denom)

			for _, divisor := range divisors {
				fractions := get_curious_fractions(nom, denom, divisor)
				if fractions != nil {
					res = append(res, []int { nom, denom, fractions[0], fractions[1] })
				}
			}
		}
	}

	for _, fr := range res {
		fmt.Println(fr[0], "/", fr[1], "==", fr[2], "/", fr[3])
	}
}