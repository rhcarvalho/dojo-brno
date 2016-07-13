package fizzbuzz

import "testing"

func TestCheckNumber(t *testing.T) {
	tests := []struct {
		from, to int
		want     string
	}{
		{1, 1, "1\n"},
	}
	for _, tt := range tests {
		var stdout string
		print := func(s string) {
			stdout += s + "\n"
		}
		Print(print, tt.from, tt.to)
		if got := stdout; got != tt.want {
			t.Errorf("Print(print, %v, %v): %q, want %q", tt.from, tt.to, got, tt.want)
		}
	}
}
