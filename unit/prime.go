package unit

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for d := 3; d*d <= n; d += 2 {
		if n%d == 0 {
			return false
		}
	}

	return true
}

func NextPrime(n int) int {
	candidate := n + 1
	if candidate <= 2 {
		return 2
	}

	for !IsPrime(candidate) {
		candidate++
	}

	return candidate
}
