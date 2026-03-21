package unit

import (
	"math"
	"testing"
)

func TestPowerZeroExponent(t *testing.T) {
	got := Power(7, 0)
	want := 1.0
	if got != want {
		t.Errorf("Power(7,0): ожидалось %v, получено %v", want, got)
	}
}

func TestPowerNegativeExponent(t *testing.T) {
	got := Power(2, -3)
	want := 0.125
	if math.Abs(got-want) > 1e-12 {
		t.Errorf("Power(2,-3): ожидалось %v, получено %v", want, got)
	}
}

func TestPowerZeroBase(t *testing.T) {
	got := Power(0, 5)
	want := 0.0
	if got != want {
		t.Errorf("Power(0,5): ожидалось %v, получено %v", want, got)
	}
}

func TestPowerPanicOnInvalidInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Power(0,-1): ожидалась паника, но ее не было")
		}
	}()

	_ = Power(0, -1)
}

func TestSqrtAccuracy(t *testing.T) {
	got, err := Sqrt(2)
	if err != nil {
		t.Errorf("Sqrt(2): неожиданная ошибка: %v", err)
	}

	want := math.Sqrt(2)
	if math.Abs(got-want) >= 1e-9 {
		t.Errorf("Sqrt(2): ожидалась точность < 1e-9, получена разница %v", math.Abs(got-want))
	}
}

func TestSqrtNegative(t *testing.T) {
	_, err := Sqrt(-1)
	if err == nil {
		t.Errorf("Sqrt(-1): ожидалась ошибка, получено nil")
	}
}
