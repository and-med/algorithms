package main

import "fmt"

func main() {
	max_p := 0
	max_cnt := 0

	for p := 3; p <= 1000; p++ {
		cnt := 0
		for a := 1; a < p/3; a++ {
			nom := p * p - 2 * a * p
			denom := 2 * p - 2 * a
			if a >= nom/denom {
				break
			}
			if nom % denom == 0 {
				cnt++
			}
		}
		if cnt > max_cnt {
			max_cnt = cnt
			max_p = p
		}
	}
	
	fmt.Println(max_p)
}

// a^2 + b^2 = c^2
// a + b + c = p

// b := (p^2 - 2ap)/(2p - 2a)
// c = p - a - b

// a^2 + b^2 = c^2
// a + b < c
// a^2 + b^2 = (p - a - b)^2
// (p - a - b) * (p - a - b) = p^2 - pa - pb - pa + a^2 + ab - pb + ab + b^2
// = 
// a^2 + b^2 = p^2 + a^2 + b^2 - 2ap - 2bp + 2ab
// p^2 - 2pa - 2pb + 2ab = 0
// p^2 - 2pa = 2pb - 2ab
// p^2 - 2pa = 2b (p - a)
// b = (p^2 - 2pa)/2(p - a) = (2p^2 - 2pa - p^2)/(2p - 2a) = (p*(2p - 2a) - p^2)/(2p - 2a) = p - p^2/(2p - 2a)