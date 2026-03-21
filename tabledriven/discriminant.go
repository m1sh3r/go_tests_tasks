package tabledriven

func AnalyzeDiscriminant(a, b, c float64) (d float64, rootsCount int) {
	d = b*b - 4*a*c
	if d > 0 {
		return d, 2
	}
	if d < 0 {
		return d, 0
	}
	return d, 1
}
