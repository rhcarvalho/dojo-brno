package fizzbuzz

func Print(print func(string), from, to int) {
	if n := 1; from <= n && n <= to {
		print("1")
	}
	if n := 2; from <= n && n <= to {
		print("2")
	}
	if n := 3; from <= n && n <= to {
		print("Fizz")
	}
	if n := 4; from <= n && n <= to {
		print("4")
	}
}
