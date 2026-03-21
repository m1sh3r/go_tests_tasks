package golden

import (
	"fmt"
	"math"
)

func StepByStepSolution(a, b, c float64) string {
	d := b*b - 4*a*c
	step1 := fmt.Sprintf(
		"Step 1: Calculate discriminant D = b^2 - 4ac = %.0f - %.0f = %.0f",
		b*b, 4*a*c, d,
	)

	if d < 0 {
		return step1 + "\nStep 2: D < 0 -> no real roots\nAnswer: no real roots"
	}

	if d == 0 {
		x := -b / (2 * a)
		return fmt.Sprintf(
			"%s\nStep 2: D = 0 -> one real root\nStep 3: x1 = %.2f\nAnswer: x = {%.2f}",
			step1, x, x,
		)
	}

	sqrtD := math.Sqrt(d)
	x1 := (-b + sqrtD) / (2 * a)
	x2 := (-b - sqrtD) / (2 * a)
	return fmt.Sprintf(
		"%s\nStep 2: D > 0 -> two real roots\nStep 3: x1 = %.0f\nStep 4: x2 = %.0f\nAnswer: x = {%.0f, %.0f}",
		step1, x1, x2, math.Min(x1, x2), math.Max(x1, x2),
	)
}
