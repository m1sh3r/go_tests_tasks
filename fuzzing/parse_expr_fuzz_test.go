package fuzzing

import (
	"math"
	"testing"
)

func requireFinite(t *testing.T, vals ...float64) {
	t.Helper()
	for _, v := range vals {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			t.Fatalf("ожидалось конечное число, получено %v", v)
		}
	}
}

func FuzzParseExpr(f *testing.F) {
	f.Add("x^2 + 2x + 1")
	f.Add("x² + 2x + 1")
	f.Add("2x^2-3x+4")
	f.Add("2x²-3x+4")

	f.Fuzz(func(t *testing.T, s string) {
		a, b, c, err := ParseExpr(s)
		if err == nil {
			requireFinite(t, a, b, c)
		}
	})
}
