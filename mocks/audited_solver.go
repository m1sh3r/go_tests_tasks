package mocks

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, err error)
}

type AuditedSolver struct {
	logger Logger
}

func NewAuditedSolver(logger Logger) *AuditedSolver {
	return &AuditedSolver{logger: logger}
}

func (s *AuditedSolver) Solve(eqID string, a, b, c float64) ([]float64, error) {
	s.logger.Info("начало вычисления", map[string]any{
		"eq_id": eqID,
		"a":     a,
		"b":     b,
		"c":     c,
	})

	roots, err := solveQuadraticFromCoeffs(a, b, c)
	if err != nil {
		s.logger.Error("ошибка вычисления", err)
		return nil, err
	}

	s.logger.Info("вычисление завершено", map[string]any{
		"eq_id": eqID,
		"roots": roots,
	})
	return roots, nil
}
