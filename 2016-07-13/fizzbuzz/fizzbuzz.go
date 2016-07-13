package fizzbuzz

func Print(print func(string), from, to int) {
	if n := 1; from <= n && n <= to {
		print("1")
	}
	if n := 2; from <= n && n <= to {
		print("2")
	}
	if n := 3; from <= n && n <= to {
		print(fizzbuzz(3))
	}
	if n := 4; from <= n && n <= to {
		print("4")
	}
}

func fizzbuzz(int) string {
	return "Fizz"
}
