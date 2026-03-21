package golden

import "testing"

func TestToLaTeXGolden(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		c    float64
	}{
		{name: "latex_standard", a: 1, b: -3, c: 2},
		{name: "latex_no_linear", a: 1, b: 0, c: 4},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ToLaTeX(tc.a, tc.b, tc.c)
			assertGolden(t, tc.name, got)
		})
	}
}
