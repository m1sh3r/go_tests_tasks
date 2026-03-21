package mocks

type Validator interface {
	ValidateCoeffs(a, b, c float64) error
}

type ValidatingSolver struct {
	validator Validator
	store     HistoryStore
}

func NewValidatingSolver(validator Validator, store HistoryStore) *ValidatingSolver {
	return &ValidatingSolver{
		validator: validator,
		store:     store,
	}
}

func (s *ValidatingSolver) Solve(eqID string, a, b, c float64) ([]float64, error) {
	if err := s.validator.ValidateCoeffs(a, b, c); err != nil {
		return nil, err
	}

	roots, err := solveQuadraticFromCoeffs(a, b, c)
	if err != nil {
		return nil, err
	}

	if s.store != nil {
		_ = s.store.Save(eqID, roots)
	}

	return roots, nil
}
