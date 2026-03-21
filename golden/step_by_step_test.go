package golden

import "testing"

func TestStepByStepSolutionGolden(t *testing.T) {
	got := StepByStepSolution(1, -3, 2)
	assertGolden(t, "step_solution_1_-3_2", got)
}
