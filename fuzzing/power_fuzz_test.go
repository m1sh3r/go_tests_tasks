package fuzzing

import (
	"math"
	"testing"
)

func FuzzPower(f *testing.F) {
	seeds := []int{0, 1, -1, 2, -2, 10, -10}
	for _, exp := range seeds {
		f.Add(2.0, exp)
	}
	f.Add(0.0, 1)
	f.Add(0.0, -1)
	f.Add(-2.0, 3)

	f.Fuzz(func(t *testing.T, base float64, exp int) {
		got := Power(base, exp)
		if math.IsNaN(got) || math.IsInf(got, 0) {
			t.Fatalf("Power(%v,%d): ожидалось конечное значение, получено %v", base, exp, got)
		}

		if exp >= 0 && base == 0 && got != 0 && exp != 0 {
			t.Fatalf("Power(0,%d): ожидалось 0, получено %v", exp, got)
		}
	})
}
