package fizzbuzz

import "testing"

func TestCheckNumber(t *testing.T) {
	var stdout string
	print := func(s string) {
		stdout += s + "\n"
	}
	Print(print)
	want := "1\n"
	if got := stdout; got != want {
		t.Errorf("Print(print): %q, want %q", got, want)
	}
}
