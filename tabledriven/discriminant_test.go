package tabledriven

import "testing"

func TestAnalyzeDiscriminant(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		c         float64
		wantD     float64
		wantRoots int
	}{
		{name: "D_больше_0", a: 1, b: -5, c: 6, wantD: 1, wantRoots: 2},
		{name: "D_равен_0", a: 1, b: -4, c: 4, wantD: 0, wantRoots: 1},
		{name: "D_меньше_0", a: 1, b: 2, c: 5, wantD: -16, wantRoots: 0},
		{name: "D_плюс_1e_15", a: 1, b: 0, c: -2.5e-16, wantD: 1e-15, wantRoots: 2},
		{name: "D_минус_1e_15", a: 1, b: 0, c: 2.5e-16, wantD: -1e-15, wantRoots: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotD, gotRoots := AnalyzeDiscriminant(tc.a, tc.b, tc.c)
			if !almostEqual(gotD, tc.wantD, 1e-18) {
				t.Errorf("AnalyzeDiscriminant(%v,%v,%v): дискриминант ожидался %v, получено %v", tc.a, tc.b, tc.c, tc.wantD, gotD)
			}
			if gotRoots != tc.wantRoots {
				t.Errorf("AnalyzeDiscriminant(%v,%v,%v): число корней ожидалось %d, получено %d", tc.a, tc.b, tc.c, tc.wantRoots, gotRoots)
			}
		})
	}
}
