package main

import (
	"fmt"
	"math"
)

func to_digits(n int) []int {
	d := []int {}
	for n != 0 {
		d = append(d, n % 10)
		n /= 10
	}

	for i := 0; i < len(d) / 2; i++ {
		d[i], d[len(d) - i - 1] = d[len(d) - i - 1], d[i]
	}
	return d
}

func get_digit_at(pos int, n int) int {
	fract := float64(pos) / float64(n) + math.Pow10(n - 1)
	num := pos/n + int(math.Pow10(n - 1))
	f_part := fract - float64(num)
	i := int(math.Round(f_part * float64(n)))
	return to_digits(num)[i]
}

func get_number_at(pos int) int {
	n_digits, class_start, class_length := 1, 1, 9
	for true {
		next_class_start := class_start + class_length * n_digits
		if pos >= class_start && pos < next_class_start {
			break
		}
		n_digits, class_start, class_length = n_digits + 1, next_class_start, class_length * 10
	}
	return get_digit_at(pos - class_start, n_digits)
}

func main() {
	p := 1
	for i := 1; i <= 1000000; i *= 10 {
		p *= get_number_at(i)
	}
	fmt.Println(p)
}

// 0.1234567891011121314151617181920
// len(1 - 9) = 9 * 1 = 9
// len(10 - 99) = 90 * 2 = 180
// len(100 - 999) = 900 * 3 = 2700
// len(1000 - 9999) = 9000 * 4 = 36000
// len(10000 - 99999) = 90000 * 5 = 450000


// digit number: d % 2 + 1
// digit: (d - 10) / 2 + 10

// digit number: (d - 1) % 3 + 1
// digit: (d - 190) / 3 + 100

// digit number: (d - 2) % 4 + 1
// digit: (d - 2890) / 4 + 1000

// digit number: d % 5 + 1
// digit: (d - 38890) / 5 + 10000

// digit: (d - 488890) / 6 + 100000

// d1 = 1 - sol
// ...
// d10 = 1 - sol
// ...
// d100 = 5 - sol
// ...
// d188 = 9
// d189 = 9
// d190 = 1
// d191 = 0
// d192 = 0
// ...
// d1000 = 3 - sol
// ...
// d2887 = 9
// d2888 = 9
// d2889 = 9
// d2890 = 1
// d2891 = 0
// d2892 = 0
// d2893 = 0
// d2894 = 1
// ...
// d10000 = 7 - sol
// ...
// d38886 = 9
// d38887 = 9
// d38888 = 9
// d38889 = 9
// d38890 = 1
// d38891 = 0
// d38892 = 0
// d38893 = 0
// d38894 = 0
// ...
// d100000 = 2 - sol
// ...
// d488890 = 1
// d488891 = 0
// d488892 = 0
// d488893 = 0
// d488894 = 0
// d488895 = 0
// ...
// d1000000 = 1 - sol