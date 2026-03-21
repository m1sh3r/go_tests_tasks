package golden

import (
	"fmt"
	"math"
	"strings"
)

func FormatPolynomial(coeffs []float64) string {
	if len(coeffs) == 0 {
		return "0"
	}

	terms := make([]string, 0, len(coeffs))
	for degree := len(coeffs) - 1; degree >= 0; degree-- {
		coef := coeffs[degree]
		if coef == 0 {
			continue
		}
		terms = append(terms, formatTerm(coef, degree, len(terms) == 0))
	}

	if len(terms) == 0 {
		return "0"
	}
	return strings.Join(terms, "")
}

func formatTerm(coef float64, degree int, first bool) string {
	sign := " + "
	if coef < 0 {
		sign = " - "
	}
	if first {
		sign = ""
		if coef < 0 {
			sign = "-"
		}
	}

	absCoef := math.Abs(coef)
	coefStr := ""
	if degree == 0 || absCoef != 1 {
		coefStr = trimFloat(absCoef)
	}

	var variable string
	switch degree {
	case 0:
		variable = ""
	case 1:
		variable = "x"
	case 2:
		variable = "x^2"
	default:
		variable = fmt.Sprintf("x^%d", degree)
	}

	return sign + coefStr + variable
}

func trimFloat(v float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.10f", v), "0"), ".")
}
