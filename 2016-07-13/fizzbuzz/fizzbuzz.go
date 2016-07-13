package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func Print(out io.Writer, from, to int) {
	for n := from; n <= to; n++ {
		fmt.Fprintln(out, fizzbuzz(n))
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

func main() {
	Print(os.Stdout, 1, 100)
}
