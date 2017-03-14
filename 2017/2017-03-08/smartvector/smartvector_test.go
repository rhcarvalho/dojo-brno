package smartvector

import "testing"

func TestSmartVectorSetGet(t *testing.T) {
	tests := [][]int{ // index, value pairs
		[]int{
			0, 10,
			1, 10,
		},
		[]int{
			1, 5,
			0, 5,
		},
		[]int{
			0, 10,
			1, 11,
			2, 12,
		},
	}
	for _, tt := range tests {
		v := NewSmartVector(len(tt) / 2)
		for i := 0; i < len(tt); i += 2 {
			index, value := tt[i], tt[i+1]

			v.Set(index, value)

			// check immediately after setting
			if got := v.Get(index); got != value {
				t.Errorf("v.Get(%d) = %d, want %d\nv = %v", index, got, value, v)
			}
		}
		// check after setting all
		for i := 0; i < len(tt); i += 2 {
			index, value := tt[i], tt[i+1]

			if got := v.Get(index); got != value {
				t.Errorf("v.Get(%d) = %d, want %d\nv = %v", index, got, value, v)
			}
		}
	}
}

func TestSmartVectorGetZeroValueEmpty(t *testing.T) {
	v := NewSmartVector(10)
	for i := 0; i < 10; i++ {
		if got := v.Get(i); got != 0 {
			t.Errorf("v.Get(%d) == %d, want %d", i, got, 0)
		}
	}
}

func TestSmartVectorGetZeroValuePartiallyInitialized(t *testing.T) {
	v := NewSmartVector(10)
	v.Set(9, -1)
	for i := 0; i < 9; i++ {
		if got := v.Get(i); got != 0 {
			t.Errorf("v.Get(%d) == %d, want %d", i, got, 0)
		}
	}
}
