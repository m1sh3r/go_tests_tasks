package fuzzing

import "errors"

var (
	ErrNonFiniteCoefficient = errors.New("коэффициенты должны быть конечными числами")
	ErrAllZeroCoefficients  = errors.New("хотя бы один коэффициент должен быть отличен от нуля")
)

func Validate(a, b, c float64) error {
	if !isFinite(a) || !isFinite(b) || !isFinite(c) {
		return ErrNonFiniteCoefficient
	}
	if a == 0 && b == 0 && c == 0 {
		return ErrAllZeroCoefficients
	}
	return nil
}
