package fuzzing

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var (
	ErrEmptyExpression      = errors.New("пустое выражение")
	ErrInvalidExpression    = errors.New("некорректное выражение")
	ErrUnsupportedPowerExpr = errors.New("поддерживается только степень 2")
)

func ParseExpr(s string) (a, b, c float64, err error) {
	s = strings.ReplaceAll(s, " ", "")
	if s == "" {
		return 0, 0, 0, ErrEmptyExpression
	}

	s = strings.ReplaceAll(s, "x²", "x^2")
	s = strings.ReplaceAll(s, "X²", "x^2")
	s = strings.ReplaceAll(s, "X", "x")
	if strings.Contains(s, "x^3") || strings.Contains(s, "x³") {
		return 0, 0, 0, ErrUnsupportedPowerExpr
	}

	if s[0] != '+' && s[0] != '-' {
		s = "+" + s
	}

	for i := 0; i < len(s); {
		sign := 1.0
		switch s[i] {
		case '+':
			sign = 1
		case '-':
			sign = -1
		default:
			return 0, 0, 0, ErrInvalidExpression
		}
		i++
		if i >= len(s) {
			return 0, 0, 0, ErrInvalidExpression
		}

		start := i
		for i < len(s) && s[i] != '+' && s[i] != '-' {
			i++
		}
		term := s[start:i]
		if term == "" {
			return 0, 0, 0, ErrInvalidExpression
		}

		value, kind, ok := parseTerm(term)
		if !ok {
			return 0, 0, 0, ErrInvalidExpression
		}

		switch kind {
		case "a":
			a += sign * value
		case "b":
			b += sign * value
		case "c":
			c += sign * value
		default:
			return 0, 0, 0, ErrInvalidExpression
		}
	}

	if !isFinite(a) || !isFinite(b) || !isFinite(c) {
		return 0, 0, 0, ErrInvalidExpression
	}

	return a, b, c, nil
}

func parseTerm(term string) (value float64, kind string, ok bool) {
	switch {
	case strings.HasSuffix(term, "x^2"):
		coef := strings.TrimSuffix(term, "x^2")
		if coef == "" {
			return 1, "a", true
		}
		v, err := strconv.ParseFloat(coef, 64)
		if err != nil || !isFinite(v) {
			return 0, "", false
		}
		return v, "a", true
	case strings.HasSuffix(term, "x"):
		coef := strings.TrimSuffix(term, "x")
		if coef == "" {
			return 1, "b", true
		}
		v, err := strconv.ParseFloat(coef, 64)
		if err != nil || !isFinite(v) {
			return 0, "", false
		}
		return v, "b", true
	default:
		v, err := strconv.ParseFloat(term, 64)
		if err != nil || !isFinite(v) {
			return 0, "", false
		}
		return v, "c", true
	}
}

func isFinite(v float64) bool {
	return !math.IsNaN(v) && !math.IsInf(v, 0)
}
