package unit

import "testing"

func TestAdd(t *testing.T) {
	got := Add(-2, 2)
	want := 0.0
	if got != want {
		t.Errorf("Add(-2,2): ожидалось %v, получено %v", want, got)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(0, 5)
	want := -5.0
	if got != want {
		t.Errorf("Subtract(0,5): ожидалось %v, получено %v", want, got)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(-3, 0)
	want := 0.0
	if got != want {
		t.Errorf("Multiply(-3,0): ожидалось %v, получено %v", want, got)
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10,2): неожиданная ошибка: %v", err)
	}
	want := 5.0
	if got != want {
		t.Errorf("Divide(10,2): ожидалось %v, получено %v", want, got)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Errorf("Divide(10,0): ожидалась ошибка, получено nil")
	}
}
