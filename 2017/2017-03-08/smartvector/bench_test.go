package smartvector

import (
	"math/rand"
	"testing"
)

// NewRepeatingSequence retuns a sequence of n ints that repeat 1 to imax times.
// Each int is in the range [imin, imax]. The seed controls the number of
// adjacent repeated ints.
func NewRepeatingSequence(n int, imin, imax uint64, seed int64) []int {
	s := make([]int, n)
	g := rand.NewZipf(rand.New(rand.NewSource(seed)), 3, 1, imax-imin)
	for i := 0; i < n; {
		k := g.Uint64()
		for j := 0; j < int(k) && i+j < n; j++ {
			s[i+j] = int(imin + k)
		}
		i += int(k)
	}
	return s
}

// NewTestData returns test data of size n, meant to match b.N.
func NewTestData(n int) []int {
	imax := uint64(n / 10)
	if imax < 25 {
		imax = 25
	}
	return NewRepeatingSequence(n, 5, imax, 42)
}

func BenchmarkVector(b *testing.B) {
	b.Run("Slice", func(b *testing.B) {
		v := NewTestData(b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s := Slice.New(nil, b.N)
			for j := 0; j < b.N; j++ {
				s.Set(j, v[j])
			}
		}
	})
	b.Run("SmartVector", func(b *testing.B) {
		v := NewTestData(b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s := NewSmartVector(b.N)
			for j := 0; j < b.N; j++ {
				s.Set(j, v[j])
			}
		}
	})
}
