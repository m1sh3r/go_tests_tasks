package fuzzing

import "math"

func Power(base float64, exp int) float64 {
	if !isFinite(base) {
		return 0
	}

	if exp == 0 {
		return 1
	}
	if base == 0 && exp < 0 {
		return 0
	}

	negativeExp := exp < 0
	n := absExp(exp)
	result := 1.0
	b := base

	for n > 0 {
		if n&1 == 1 {
			result *= b
			if !isFinite(result) {
				return 0
			}
		}
		n >>= 1
		if n > 0 {
			b *= b
			if !isFinite(b) {
				return 0
			}
		}
	}

	if negativeExp {
		if result == 0 {
			return 0
		}
		result = 1 / result
		if !isFinite(result) {
			return 0
		}
	}

	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0
	}
	return result
}

func absExp(exp int) uint {
	if exp >= 0 {
		return uint(exp)
	}

	return uint(-(exp + 1)) + 1
}
