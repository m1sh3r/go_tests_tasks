package golden

import "testing"

func TestGenerateReportGolden(t *testing.T) {
	got := GenerateReport("eq_001", []float64{2, 3})
	assertGolden(t, "report_eq_001", got)
}
