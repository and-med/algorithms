package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func get_continued_fraction_coefficients(n int) []int {
	a := int(math.Sqrt(float64(n)))
	sol := []int {a}

	if (a * a == n) {
		return sol
	}

	curr_d := 1
	curr_n := a
	for {
		curr_d = (n - curr_n * curr_n) / curr_d
		res := int(float64(curr_n + a) / float64(curr_d))
		curr_n = res * curr_d - curr_n
		sol = append(sol, res)
		if curr_d == 1 && curr_n == a {
			break
		}
	}

	return sol
}

func get_coefficient(coefficients []int, n int) int64 {
	return int64(coefficients[n % len(coefficients)])
}

func calculate_continued_fraction(coef []int, n int) (*big.Int, *big.Int) {
	p_0 := big.NewInt(get_coefficient(coef, 0))
	p_1 := big.NewInt(get_coefficient(coef, 1))
	p_1 = p_1.Mul(p_1, p_0).Add(p_1, big.NewInt(1))

	for i := 2; i <= n; i++ {
		a_n := big.NewInt(get_coefficient(coef, i))
		p_n := big.NewInt(1)
		p_n = p_n.Mul(a_n, p_1).Add(p_n, p_0)
		p_0, p_1 = p_1, p_n
	}

	q_0 := big.NewInt(1)
	q_1 := big.NewInt(get_coefficient(coef, 1))

	for i := 2; i <= n; i++ {
		a_n := big.NewInt(get_coefficient(coef, i))
		q_n := big.NewInt(1)
		q_n = q_n.Mul(a_n, q_1).Add(q_n, q_0)
		q_0, q_1 = q_1, q_n
	}

	return p_1, q_1
}

func solve_diophantine(d int) (*big.Int, *big.Int) {
	coefficients := get_continued_fraction_coefficients(d);

	if len(coefficients) == 1 {
		return big.NewInt(-1), big.NewInt(-1)
	}

	r := len(coefficients) - 2
	switch {
	case r % 2 == 1:
		return calculate_continued_fraction(coefficients, r)
	default:
		return calculate_continued_fraction(coefficients, 2 * r + 1)
	}
}

func solve() int64 {
	max_d := 5
	max_sol := big.NewInt(9)
	for d := 8; d <= 1000; d++ {
		x, _ := solve_diophantine(d)
		if max_sol.Cmp(x) < 0 {
			max_d = d
			max_sol = x
		}
	}
	fmt.Println("Solution: ", max_sol.Text(10))
	return int64(max_d)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}