package fuzzing

import "testing"

func FuzzValidate(f *testing.F) {
	f.Add(1.0, -3.0, 2.0)
	f.Add(0.0, 0.0, 0.0)
	f.Add(0.0, 2.0, -4.0)

	f.Fuzz(func(t *testing.T, a, b, c float64) {
		_ = Validate(a, b, c)
	})
}
