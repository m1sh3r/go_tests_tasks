package tabledriven

import "errors"

var (
	ErrInvalidBase   = errors.New("основание системы счисления должно быть в диапазоне 2..36")
	ErrNegativeValue = errors.New("число должно быть неотрицательным")
)

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

func ToBase(n int, base int) (string, error) {
	if base < 2 || base > 36 {
		return "", ErrInvalidBase
	}
	if n < 0 {
		return "", ErrNegativeValue
	}
	if n == 0 {
		return "0", nil
	}

	buf := make([]byte, 0, 64)
	for n > 0 {
		rem := n % base
		buf = append(buf, digits[rem])
		n /= base
	}

	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf), nil
}
