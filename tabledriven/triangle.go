package tabledriven

import "errors"

var ErrImpossibleTriangle = errors.New("невозможный треугольник")

func TriangleType(a, b, c float64) (string, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return "impossible", ErrImpossibleTriangle
	}
	if a+b <= c || a+c <= b || b+c <= a {
		return "impossible", ErrImpossibleTriangle
	}

	if a == b && b == c {
		return "equilateral", nil
	}
	if a == b || b == c || a == c {
		return "isosceles", nil
	}
	return "scalene", nil
}
