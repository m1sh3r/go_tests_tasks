package golden

import "testing"

func TestFormatPolynomialGolden(t *testing.T) {
	tests := []struct {
		name   string
		coeffs []float64
	}{
		{name: "linear_pos", coeffs: []float64{3, 2}},
		{name: "quadratic_neg", coeffs: []float64{-1, 4, -1}},
		{name: "constant_only", coeffs: []float64{5}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatPolynomial(tc.coeffs)
			assertGolden(t, tc.name, got)
		})
	}
}
