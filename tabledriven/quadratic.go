package tabledriven

import (
	"errors"
	"math"
)

var ErrNoEquation = errors.New("уравнение не задано")

func QuadraticRoots(a, b, c float64) ([]float64, error) {
	if a == 0 {
		if b == 0 {
			return nil, ErrNoEquation
		}
		return []float64{-c / b}, nil
	}

	d := b*b - 4*a*c
	if d < 0 {
		return []float64{}, nil
	}
	if d == 0 {
		return []float64{-b / (2 * a)}, nil
	}

	sqrtD := math.Sqrt(d)
	x1 := (-b + sqrtD) / (2 * a)
	x2 := (-b - sqrtD) / (2 * a)
	return []float64{x1, x2}, nil
}
