package tabledriven

import (
	"math"
	"strconv"
	"testing"
)

func TestToBase(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		base    int
		want    string
		wantErr bool
	}{
		{name: "10_в_двоичную", n: 10, base: 2, want: "1010", wantErr: false},
		{name: "255_в_hex", n: 255, base: 16, want: "ff", wantErr: false},
		{name: "основание_меньше_2", n: 10, base: 1, want: "", wantErr: true},
		{name: "основание_больше_36", n: 10, base: 37, want: "", wantErr: true},
		{name: "отрицательное_число", n: -1, base: 10, want: "", wantErr: true},
		{name: "ноль", n: 0, base: 2, want: "0", wantErr: false},
		{name: "максимальный_int_в_10", n: math.MaxInt, base: 10, want: strconv.FormatInt(int64(math.MaxInt), 10), wantErr: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToBase(tc.n, tc.base)
			if tc.wantErr {
				if err == nil {
					t.Errorf("ToBase(%d,%d): ожидалась ошибка, получено nil", tc.n, tc.base)
				}
				return
			}

			if err != nil {
				t.Errorf("ToBase(%d,%d): неожиданная ошибка: %v", tc.n, tc.base, err)
				return
			}

			if got != tc.want {
				t.Errorf("ToBase(%d,%d): ожидалось %q, получено %q", tc.n, tc.base, tc.want, got)
			}
		})
	}
}
