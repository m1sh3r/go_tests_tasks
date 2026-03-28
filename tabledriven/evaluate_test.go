package tabledriven

import "testing"

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name   string
		coeffs []float64
		x      float64
		want   float64
	}{
		{name: "константа", coeffs: []float64{5}, x: 10, want: 5},
		{name: "линейный_2x_плюс_3", coeffs: []float64{3, 2}, x: 4, want: 11},
		{name: "x2_при_1e6", coeffs: []float64{0, 0, 1}, x: 1e6, want: 1e12},
		{name: "пустой_срез", coeffs: []float64{}, x: 5, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Evaluate(tc.coeffs, tc.x)
			if !almostEqual(got, tc.want, 1e-9) {
				t.Errorf("Evaluate(%v,%v): ожидалось %v, получено %v", tc.coeffs, tc.x, tc.want, got)
			}
		})
	}
}
