package golden

import "testing"

func TestResultToJSONGolden(t *testing.T) {
	r := Result{
		ID:     "eq1",
		A:      1,
		B:      -3,
		C:      2,
		Roots:  []float64{1, 2},
		Status: "solved",
	}

	got := string(ResultToJSON(r))
	assertGolden(t, "result_json_eq1", got)
}

func TestResultToJSONStable(t *testing.T) {
	r := Result{
		ID:     "eq2",
		A:      1,
		B:      2,
		C:      5,
		Roots:  []float64{},
		Status: "no_real_roots",
	}

	first := string(ResultToJSON(r))
	second := string(ResultToJSON(r))
	if first != second {
		t.Errorf("ResultToJSON: ожидалась стабильная сериализация, но вывод отличается")
	}
}
