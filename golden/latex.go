package golden

import (
	"math"
	"strings"
)

func ToLaTeX(a, b, c float64) string {
	terms := make([]string, 0, 3)
	terms = appendTermLaTeX(terms, a, "x^{2}")
	terms = appendTermLaTeX(terms, b, "x")
	terms = appendTermLaTeX(terms, c, "")

	if len(terms) == 0 {
		return "0 = 0"
	}
	return strings.Join(terms, "") + " = 0"
}

func appendTermLaTeX(terms []string, coef float64, variable string) []string {
	if coef == 0 {
		return terms
	}

	first := len(terms) == 0
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
	if variable == "" || absCoef != 1 {
		coefStr = trimFloat(absCoef)
	}

	return append(terms, sign+escapeLaTeX(coefStr)+variable)
}

func escapeLaTeX(s string) string {
	replacer := strings.NewReplacer(
		`_`, `\_`,
		`{`, `\{`,
		`}`, `\}`,
	)
	return replacer.Replace(s)
}
