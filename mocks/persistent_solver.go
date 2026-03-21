package mocks

type Record struct {
	EqID  string
	Roots []float64
}

type HistoryStore interface {
	Save(eqID string, roots []float64) error
	GetHistory(limit int) ([]Record, error)
}

type PersistentSolver struct {
	store  HistoryStore
	logger Logger
}

func NewPersistentSolver(store HistoryStore, logger Logger) *PersistentSolver {
	return &PersistentSolver{
		store:  store,
		logger: logger,
	}
}

func (s *PersistentSolver) Solve(eqID string, a, b, c float64) ([]float64, error) {
	roots, err := solveQuadraticFromCoeffs(a, b, c)
	if err != nil {
		return nil, err
	}

	if err := s.store.Save(eqID, roots); err != nil {
		s.logger.Info("предупреждение: не удалось сохранить историю", map[string]any{
			"eq_id": eqID,
			"err":   err.Error(),
		})
	}

	return roots, nil
}
