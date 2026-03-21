package fuzzing

import "testing"

func FuzzGenerateRandomEq(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(1))
	f.Add(int64(-1))
	f.Add(int64(42))

	f.Fuzz(func(t *testing.T, seed int64) {
		a1, b1, c1 := GenerateRandomEq(seed)
		a2, b2, c2 := GenerateRandomEq(seed)

		if !isFinite(a1) || !isFinite(b1) || !isFinite(c1) {
			t.Fatalf("GenerateRandomEq(%d): коэффициенты должны быть конечными", seed)
		}

		if a1 == 0 && b1 == 0 && c1 == 0 {
			t.Fatalf("GenerateRandomEq(%d): хотя бы один коэффициент должен быть отличен от нуля", seed)
		}

		if a1 != a2 || b1 != b2 || c1 != c2 {
			t.Fatalf("GenerateRandomEq(%d): нарушен детерминизм", seed)
		}
	})
}
