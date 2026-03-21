package golden

import "testing"

func TestResultsToCSVGolden(t *testing.T) {
	r1 := 1.0
	r2 := 2.0
	results := []CSVResult{
		{
			ID:     "eq2",
			A:      1,
			B:      2,
			C:      5,
			Root1:  nil,
			Root2:  nil,
			Status: "no_real_roots",
		},
		{
			ID:     "eq1",
			A:      1,
			B:      -3,
			C:      2,
			Root1:  &r1,
			Root2:  &r2,
			Status: "solved",
		},
		{
			ID:     "eq,3",
			A:      2,
			B:      0,
			C:      -8,
			Root1:  makePtr(2),
			Root2:  makePtr(-2),
			Status: "contains,comma",
		},
	}

	got := ResultsToCSV(results)
	assertGolden(t, "results_csv", got)
}
