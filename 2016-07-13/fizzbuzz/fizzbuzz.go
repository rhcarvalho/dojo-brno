package fizzbuzz

import "strconv"

func Print(print func(string), from, to int) {
	for n := from; n <= to; n++ {
		print(fizzbuzz(n))
	}
}

func fizzbuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
