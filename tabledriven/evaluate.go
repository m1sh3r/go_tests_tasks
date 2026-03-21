package tabledriven

func Evaluate(coeffs []float64, x float64) float64 {
	if len(coeffs) == 0 {
		return 0
	}

	result := 0.0
	for i := len(coeffs) - 1; i >= 0; i-- {
		result = result*x + coeffs[i]
	}

	return result
}
