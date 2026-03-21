package golden

import "fmt"

func GenerateReport(eqID string, roots []float64) string {
	reportType := "quadratic"
	status := "solved"
	rootsLine := "Roots: none"
	discriminant := 0.0

	switch len(roots) {
	case 2:
		rootsLine = fmt.Sprintf("Roots: x1 = %.2f, x2 = %.2f", roots[0], roots[1])
		delta := roots[0] - roots[1]
		discriminant = delta * delta
	case 1:
		rootsLine = fmt.Sprintf("Roots: x1 = %.2f", roots[0])
		discriminant = 0
	default:
		status = "no_real_roots"
	}

	return fmt.Sprintf(
		"Equation: %s\nType: %s\n%s\nDiscriminant: %.2f\nStatus: %s",
		eqID, reportType, rootsLine, discriminant, status,
	)
}
