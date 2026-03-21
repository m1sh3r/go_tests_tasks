package fuzzing

import "math/rand"

func GenerateRandomEq(seed int64) (a, b, c float64) {
	r := rand.New(rand.NewSource(seed))

	a = randomCoeff(r)
	b = randomCoeff(r)
	c = randomCoeff(r)

	if a == 0 && b == 0 && c == 0 {
		a = 1
	}

	return a, b, c
}

func randomCoeff(r *rand.Rand) float64 {
	v := float64(r.Intn(41)-20) / 2.0
	return v
}
