package unit

import "testing"

func TestGCDZeroOperand(t *testing.T) {
	got := GCD(0, 18)
	want := 18
	if got != want {
		t.Errorf("GCD(0,18): ожидалось %d, получено %d", want, got)
	}
}

func TestGCDSameValues(t *testing.T) {
	got := GCD(21, 21)
	want := 21
	if got != want {
		t.Errorf("GCD(21,21): ожидалось %d, получено %d", want, got)
	}
}

func TestGCDSymmetry(t *testing.T) {
	a, b := 48, 18
	left := GCD(a, b)
	right := GCD(b, a)
	if left != right {
		t.Errorf("симметричность GCD нарушена: GCD(%d,%d)=%d, GCD(%d,%d)=%d", a, b, left, b, a, right)
	}
}

func TestGCDLCMRelation(t *testing.T) {
	a, b := 21, 6
	g := GCD(a, b)
	l := LCM(a, b)
	got := g * l
	want := absInt(a * b)
	if got != want {
		t.Errorf("GCD(%d,%d)*LCM(%d,%d): ожидалось %d, получено %d", a, b, a, b, want, got)
	}
}
