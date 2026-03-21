package tabledriven

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrEmptyExpr          = errors.New("пустое выражение")
	ErrInvalidExprSyntax  = errors.New("некорректный синтаксис выражения")
	ErrUnsupportedExprPow = errors.New("поддерживается только степень x^2")
)

var digitsOnly = regexp.MustCompile(`^\d+$`)

func ValidateExpr(expr string) error {
	s := strings.ReplaceAll(expr, " ", "")
	if s == "" {
		return ErrEmptyExpr
	}
	if strings.Contains(s, "x³") || strings.Contains(s, "x^3") {
		return ErrUnsupportedExprPow
	}

	seenLinear := false
	seenQuadratic := false
	seenConst := false

	i := 0
	for i < len(s) {
		if s[i] == '+' || s[i] == '-' {
			i++
			if i >= len(s) {
				return ErrInvalidExprSyntax
			}
		}

		start := i
		for i < len(s) && s[i] != '+' && s[i] != '-' {
			i++
		}
		term := s[start:i]
		if term == "" {
			return ErrInvalidExprSyntax
		}

		kind, ok := classifyExprTerm(term)
		if !ok {
			return ErrInvalidExprSyntax
		}

		switch kind {
		case "quadratic":
			if seenQuadratic {
				return ErrInvalidExprSyntax
			}
			seenQuadratic = true
		case "linear":
			if seenLinear {
				return ErrInvalidExprSyntax
			}
			seenLinear = true
		case "const":
			if seenConst {
				return ErrInvalidExprSyntax
			}
			seenConst = true
		}
	}

	return nil
}

func classifyExprTerm(term string) (kind string, ok bool) {
	if strings.Contains(term, "x") {
		if strings.Count(term, "x") != 1 {
			return "", false
		}

		if strings.HasSuffix(term, "x²") || strings.HasSuffix(term, "x^2") {
			coef := strings.TrimSuffix(term, "x²")
			if coef == term {
				coef = strings.TrimSuffix(term, "x^2")
			}
			if coef == "" || digitsOnly.MatchString(coef) {
				return "quadratic", true
			}
			return "", false
		}

		if strings.HasSuffix(term, "x") {
			coef := strings.TrimSuffix(term, "x")
			if coef == "" || digitsOnly.MatchString(coef) {
				return "linear", true
			}
			return "", false
		}

		return "", false
	}

	if digitsOnly.MatchString(term) {
		return "const", true
	}

	return "", false
}
