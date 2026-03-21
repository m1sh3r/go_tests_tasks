package tabledriven

import "testing"

func TestValidateExpr(t *testing.T) {
	tests := []struct {
		expr  string
		valid bool
		desc  string
	}{
		{expr: "x^2 + 2x + 1", valid: true, desc: "полное выражение"},
		{expr: "2x^2-3x", valid: true, desc: "без константы"},
		{expr: "x^3 + 1", valid: false, desc: "степень больше 2"},
		{expr: "x² + 2x + 1", valid: true, desc: "unicode степень 2"},
		{expr: "x³ + 1", valid: false, desc: "unicode степень 3"},
		{expr: "2x + + 3", valid: false, desc: "двойной оператор"},
		{expr: "", valid: false, desc: "пустая строка"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			err := ValidateExpr(tc.expr)
			if tc.valid && err != nil {
				t.Errorf("ValidateExpr(%q): ожидалось валидное выражение, получена ошибка: %v", tc.expr, err)
			}
			if !tc.valid && err == nil {
				t.Errorf("ValidateExpr(%q): ожидалась ошибка, получено nil", tc.expr)
			}
		})
	}
}
