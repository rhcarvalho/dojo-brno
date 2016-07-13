package main

import "testing"

type testPrinter struct {
	stdout string
}

func (p *testPrinter) Print(s string) {
	p.stdout += s + "\n"
}

func TestCheckNumber(t *testing.T) {
	tests := []struct {
		from, to int
		want     string
	}{
		{1, 1, "1\n"},
		{1, 2, "1\n2\n"},
		{1, 3, "1\n2\nFizz\n"},
		{4, 4, "4\n"},
		{4, 5, "4\nBuzz\n"},
		{11, 11, "11\n"},
	}
	for _, tt := range tests {
		p := &testPrinter{} // same as new(testPrinter)
		Print(p, tt.from, tt.to)
		if got := p.stdout; got != tt.want {
			t.Errorf("Print(print, %v, %v): %q, want %q", tt.from, tt.to, got, tt.want)
		}
	}
}
