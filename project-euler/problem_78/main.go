package main

import (
	"fmt"
	"math/big"
	"time"
)

var limit = 100
var zero *big.Int
var one *big.Int
var minus_one *big.Int
var modulo *big.Int
var cache map[int]*big.Int

func init() {
	zero = big.NewInt(0)
	one = big.NewInt(1)
	minus_one = big.NewInt(-1)
	modulo = big.NewInt(int64(1e6))
	cache = map[int]*big.Int{}
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func g_k(k int) int {
	return k * (3*k - 1) / 2
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func mod_2(n int) int {
	return n % 2
}

func recurr_p_n(n int) (res *big.Int, solution int) {
	if n == 0 || n == 1 {
		return big.NewInt(1), -1
	}
	if val, ok := cache[n]; ok {
		return val, -1
	}

	res, solution = big.NewInt(0), -1
	for j := 1; true; j++ {
		k := (j+1)/2*mod_2(j) + (-j/2)*mod_2(j+1)
		g := g_k(k)
		if n-g >= 0 {
			p_n_g, solution := recurr_p_n(n - g)

			if solution != -1 {
				return p_n_g, solution
			}

			if mod_2(abs(k))-mod_2(abs(k)+1) == -1 {
				p_n_g = big.NewInt(0).Neg(p_n_g)
			}
			res.Add(res, p_n_g)
		} else {
			break
		}
	}
	if one.Mod(res, modulo).Cmp(zero) == 0 {
		return res, n
	}
	cache[n] = res
	return res, solution
}

func loop_p_n(n int) (*big.Int, int) {
	p_n := make([]*big.Int, n)
	g_k := make([]int, n)

	for j, k := 0, 1; j < n; j++ {
		g_k[j] = k * (3*k - 1) / 2

		if k < 0 {
			k = -k + 1
		} else {
			k = -k
		}
	}

	p_n[0] = big.NewInt(1)
	p_n[1] = big.NewInt(1)

	fmt.Println(0, ":", p_n[0])
	fmt.Println(1, ":", p_n[1])

	for i := 2; i < n; i++ {
		res := big.NewInt(0)

		for j, g := range g_k {
			if i < g {
				break
			}

			if mod_2(j/2 + 1) == 0 {
				res.Sub(res, p_n[i - g])
			} else {
				res.Add(res, p_n[i - g])
			}
		}

		p_n[i] = res
		fmt.Println(i, ":", p_n[i])

		if one.Mod(res, modulo).Cmp(zero) == 0 {
			return res, i
		}
	}

	return p_n[n - 1], -1
}

func solve() int64 {
	p_n, n := loop_p_n(limit)
	fmt.Println(p_n)
	return int64(n)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}
