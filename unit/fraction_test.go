package unit

import "testing"

func TestFractionAdd(t *testing.T) {
	left := Fraction{Num: 1, Den: 2}
	right := Fraction{Num: 1, Den: 3}

	got := left.Add(right)
	want := Fraction{Num: 5, Den: 6}
	if got != want {
		t.Errorf("1/2 + 1/3: ожидалось %v, получено %v", want, got)
	}
}

func TestFractionReduce(t *testing.T) {
	got := (Fraction{Num: 4, Den: 8}).Reduce()
	want := Fraction{Num: 1, Den: 2}
	if got != want {
		t.Errorf("Reduce(4/8): ожидалось %v, получено %v", want, got)
	}
}

func TestFractionReduceNegativeDenominator(t *testing.T) {
	got := (Fraction{Num: 1, Den: -2}).Reduce()
	want := Fraction{Num: -1, Den: 2}
	if got != want {
		t.Errorf("Reduce(1/-2): ожидалось %v, получено %v", want, got)
	}
}

func TestFractionDenominatorInvariant(t *testing.T) {
	got := (Fraction{Num: -4, Den: -8}).Reduce()
	if got.Den <= 0 {
		t.Errorf("Reduce(-4/-8): нарушен инвариант, Den=%d", got.Den)
	}
}
