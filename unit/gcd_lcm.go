package unit

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(a, b int) int {
	a = absInt(a)
	b = absInt(b)

	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	return absInt(a/GCD(a, b)) * absInt(b)
}
