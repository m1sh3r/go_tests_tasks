package unit

import "errors"

var (
	ErrNegativeFactorialInput = errors.New("факториал не определен для отрицательных чисел")
	ErrFactorialOverflow      = errors.New("факториал переполняет int64 при n > 20")
)

func Factorial(n int) (int64, error) {
	return factorialIterative(n)
}

func factorialIterative(n int) (int64, error) {
	if n < 0 {
		return 0, ErrNegativeFactorialInput
	}
	if n > 20 {
		return 0, ErrFactorialOverflow
	}

	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}

	return result, nil
}

func factorialRecursive(n int) (int64, error) {
	if n < 0 {
		return 0, ErrNegativeFactorialInput
	}
	if n > 20 {
		return 0, ErrFactorialOverflow
	}
	if n == 0 || n == 1 {
		return 1, nil
	}

	prev, err := factorialRecursive(n - 1)
	if err != nil {
		return 0, err
	}

	return int64(n) * prev, nil
}
