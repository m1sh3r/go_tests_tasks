package tabledriven

import (
	"math"
	"testing"
)

func almostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) <= eps
}

func equalFloatSlices(got, want []float64, eps float64) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if !almostEqual(got[i], want[i], eps) {
			return false
		}
	}
	return true
}

func TestQuadraticRoots(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		c       float64
		want    []float64
		wantErr bool
	}{
		{name: "два_корня", a: 1, b: -5, c: 6, want: []float64{3, 2}, wantErr: false},
		{name: "один_корень", a: 1, b: -4, c: 4, want: []float64{2}, wantErr: false},
		{name: "нет_корней", a: 1, b: 2, c: 5, want: []float64{}, wantErr: false},
		{name: "линейное", a: 0, b: 2, c: -4, want: []float64{2}, wantErr: false},
		{name: "не_уравнение", a: 0, b: 0, c: 5, want: nil, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := QuadraticRoots(tc.a, tc.b, tc.c)
			if tc.wantErr {
				if err == nil {
					t.Errorf("QuadraticRoots(%v,%v,%v): ожидалась ошибка, получено nil", tc.a, tc.b, tc.c)
				}
				return
			}

			if err != nil {
				t.Errorf("QuadraticRoots(%v,%v,%v): неожиданная ошибка: %v", tc.a, tc.b, tc.c, err)
				return
			}

			if !equalFloatSlices(got, tc.want, 1e-9) {
				t.Errorf("QuadraticRoots(%v,%v,%v): ожидалось %v, получено %v", tc.a, tc.b, tc.c, tc.want, got)
			}
		})
	}
}
