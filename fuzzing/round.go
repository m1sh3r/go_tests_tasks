package fuzzing

import "math"

func RoundTo(x float64, precision int) float64 {
	if !isFinite(x) {
		return x
	}

	p := clampPrecision(precision)
	factor := math.Pow10(p)
	if !isFinite(factor) || factor == 0 {
		return x
	}

	v := math.Round(x*factor) / factor
	if !isFinite(v) {
		return x
	}
	return v
}

func clampPrecision(precision int) int {
	if precision < 0 {
		return 0
	}
	if precision > 15 {
		return 15
	}
	return precision
}
