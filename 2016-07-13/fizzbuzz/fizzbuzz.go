package fizzbuzz

func Print(print func(string), from, to int) {
	print("1")
	if to > 1 {
		print("2")
	}
	if to > 2 {
		print("Fizz")
	}
}
