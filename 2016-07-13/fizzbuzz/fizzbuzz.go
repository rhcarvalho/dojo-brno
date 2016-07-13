package main

import (
	"fmt"
	"strconv"
)

type printer interface {
	Print(string)
}

func Print(p printer, from, to int) {
	for n := from; n <= to; n++ {
		p.Print(fizzbuzz(n))
	}
}

func fizzbuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}
	if n == 5 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

type stdoutPrinter struct{}

func (p stdoutPrinter) Print(s string) {
	fmt.Println(s)
}
func main() {
	Print(stdoutPrinter{}, 1, 100)
}
