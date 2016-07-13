package fizzbuzz

func Print(print func(string), from, to int) {
	if n := 1; from <= n && n <= to {
		print(fizzbuzz(1))
	}
	if n := 2; from <= n && n <= to {
		print(fizzbuzz(2))
	}
	if n := 3; from <= n && n <= to {
		print(fizzbuzz(3))
	}
	if n := 4; from <= n && n <= to {
		print(fizzbuzz(4))
	}
}

func fizzbuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}
	return string('0' + n)
}
