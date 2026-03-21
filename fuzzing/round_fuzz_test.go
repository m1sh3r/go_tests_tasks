package fuzzing

import (
	"math"
	"testing"
)

func FuzzRoundTo(f *testing.F) {
	f.Add(1.2345, 2)
	f.Add(-1.2345, 3)
	f.Add(math.Inf(1), 2)
	f.Add(math.NaN(), 2)
	f.Add(100.0, -1)
	f.Add(100.0, 1000)

	f.Fuzz(func(t *testing.T, x float64, precision int) {
		got := RoundTo(x, precision)
		if !isFinite(x) {

			if math.IsNaN(x) && !math.IsNaN(got) {
				t.Fatalf("RoundTo(NaN,%d): ожидался NaN, получено %v", precision, got)
			}
			return
		}

		if !isFinite(got) {
			t.Fatalf("RoundTo(%v,%d): ожидалось конечное значение, получено %v", x, precision, got)
		}

		p := clampPrecision(precision)
		maxDelta := 0.5 * math.Pow10(-p)
		ulpAllowance := 1e-12 * math.Max(1, math.Abs(x))
		if math.Abs(got-x) > maxDelta+ulpAllowance {
			t.Fatalf("RoundTo(%v,%d): отклонение %v больше допустимого %v", x, precision, math.Abs(got-x), maxDelta+ulpAllowance)
		}
	})
}
