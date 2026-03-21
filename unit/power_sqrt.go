package unit

import (
	"errors"
	"math"
)

var ErrNegativeSqrt = errors.New("квадратный корень из отрицательного числа")

func Power(base float64, exp int) float64 {
	if base == 0 && exp < 0 {
		panic("ноль нельзя возводить в отрицательную степень")
	}

	if exp == 0 {
		return 1
	}

	if exp < 0 {
		return 1 / Power(base, -exp)
	}

	result := 1.0
	for i := 0; i < exp; i++ {
		result *= base
	}

	return result
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt
	}

	return math.Sqrt(x), nil
}
