package unit

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{n: -5, want: false},
		{n: 0, want: false},
		{n: 1, want: false},
		{n: 2, want: true},
		{n: 3, want: true},
		{n: 4, want: false},
		{n: 5, want: true},
		{n: 7, want: true},
		{n: 11, want: true},
		{n: 12, want: false},
	}

	for _, tc := range tests {
		got := IsPrime(tc.n)
		if got != tc.want {
			t.Errorf("IsPrime(%d): ожидалось %t, получено %t", tc.n, tc.want, got)
		}
	}
}

func TestNextPrime(t *testing.T) {
	got := NextPrime(10)
	want := 11
	if got != want {
		t.Errorf("NextPrime(10): ожидалось %d, получено %d", want, got)
	}

	got = NextPrime(13)
	want = 17
	if got != want {
		t.Errorf("NextPrime(13): ожидалось %d, получено %d", want, got)
	}
}
