package unit

import (
	"fmt"
	"testing"
)

func runFactorialCases(t *testing.T, implName string, fn func(int) (int64, error)) {
	t.Helper()

	tests := []struct {
		n       int
		want    int64
		wantErr bool
	}{
		{n: 0, want: 1, wantErr: false},
		{n: 1, want: 1, wantErr: false},
		{n: 5, want: 120, wantErr: false},
		{n: 20, want: 2432902008176640000, wantErr: false},
		{n: -1, want: 0, wantErr: true},
		{n: 21, want: 0, wantErr: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%s_n_%d", implName, tc.n), func(t *testing.T) {
			got, err := fn(tc.n)
			if tc.wantErr {
				if err == nil {
					t.Errorf("Factorial(%d): ожидалась ошибка, получено nil", tc.n)
				}
				return
			}

			if err != nil {
				t.Errorf("Factorial(%d): неожиданная ошибка: %v", tc.n, err)
				return
			}

			if got != tc.want {
				t.Errorf("Factorial(%d): ожидалось %d, получено %d", tc.n, tc.want, got)
			}
		})
	}
}

func TestFactorialIterative(t *testing.T) {
	runFactorialCases(t, "iterative", factorialIterative)
}

func TestFactorialRecursive(t *testing.T) {
	runFactorialCases(t, "recursive", factorialRecursive)
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n    int
		want int64
	}{
		{n: 0, want: 1},
		{n: 1, want: 1},
		{n: 6, want: 720},
	}

	for _, tc := range tests {
		got, err := Factorial(tc.n)
		if err != nil {
			t.Errorf("Factorial(%d): неожиданная ошибка: %v", tc.n, err)
			continue
		}
		if got != tc.want {
			t.Errorf("Factorial(%d): ожидалось %d, получено %d", tc.n, tc.want, got)
		}
	}
}
