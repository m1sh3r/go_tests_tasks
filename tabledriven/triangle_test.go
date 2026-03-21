package tabledriven

import "testing"

func TestTriangleType(t *testing.T) {
	tests := []struct {
		name     string
		sides    [3]float64
		wantType string
		wantErr  bool
	}{
		{name: "равносторонний", sides: [3]float64{3, 3, 3}, wantType: "equilateral", wantErr: false},
		{name: "равнобедренный", sides: [3]float64{3, 3, 5}, wantType: "isosceles", wantErr: false},
		{name: "разносторонний", sides: [3]float64{4, 5, 6}, wantType: "scalene", wantErr: false},
		{name: "невозможный", sides: [3]float64{1, 2, 10}, wantType: "impossible", wantErr: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotType, err := TriangleType(tc.sides[0], tc.sides[1], tc.sides[2])
			if tc.wantErr {
				if err == nil {
					t.Errorf("TriangleType(%v,%v,%v): ожидалась ошибка, получено nil", tc.sides[0], tc.sides[1], tc.sides[2])
				}
			} else if err != nil {
				t.Errorf("TriangleType(%v,%v,%v): неожиданная ошибка: %v", tc.sides[0], tc.sides[1], tc.sides[2], err)
			}

			if gotType != tc.wantType {
				t.Errorf("TriangleType(%v,%v,%v): ожидался тип %q, получено %q", tc.sides[0], tc.sides[1], tc.sides[2], tc.wantType, gotType)
			}
		})
	}
}
