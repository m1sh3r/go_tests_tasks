package fuzzing

import (
	"errors"
	"strconv"
	"strings"
)

var ErrNonFiniteFloat = errors.New("значение NaN/Inf не поддерживается")

func ParseFloatSafe(s string) (float64, error) {
	v, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0, err
	}
	if !isFinite(v) {
		return 0, ErrNonFiniteFloat
	}
	return v, nil
}
