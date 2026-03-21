package mocks

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestValidatingSolverEarlyReturnOnValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	validator := NewMockValidator(ctrl)
	store := NewMockHistoryStore(ctrl)
	solver := NewValidatingSolver(validator, store)

	valErr := errors.New("невалидные коэффициенты")
	validator.EXPECT().ValidateCoeffs(0.0, 0.0, 5.0).Return(valErr)

	_, err := solver.Solve("eq_bad", 0, 0, 5)
	if !errors.Is(err, valErr) {
		t.Errorf("Solve(eq_bad): ожидалась ошибка %v, получено %v", valErr, err)
	}
}

func TestValidatingSolverContinuesWhenValidationOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	validator := NewMockValidator(ctrl)
	solver := NewValidatingSolver(validator, nil)

	validator.EXPECT().ValidateCoeffs(1.0, -3.0, 2.0).Return(nil)

	got, err := solver.Solve("eq_ok", 1, -3, 2)
	if err != nil {
		t.Errorf("Solve(eq_ok): неожиданная ошибка: %v", err)
		return
	}

	want := []float64{2, 1}
	if !equalRoots(got, want) {
		t.Errorf("Solve(eq_ok): ожидались корни %v, получены %v", want, got)
	}
}

func TestValidatingSolverCompositionValidatorAndStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	validator := NewMockValidator(ctrl)
	store := NewMockHistoryStore(ctrl)
	solver := NewValidatingSolver(validator, store)

	validator.EXPECT().ValidateCoeffs(1.0, -3.0, 2.0).Return(nil)
	store.EXPECT().Save(gomock.Eq("eq1"), gomock.Len(2)).Return(nil)

	got, err := solver.Solve("eq1", 1, -3, 2)
	if err != nil {
		t.Errorf("Solve(eq1): неожиданная ошибка: %v", err)
		return
	}

	want := []float64{2, 1}
	if !equalRoots(got, want) {
		t.Errorf("Solve(eq1): ожидались корни %v, получены %v", want, got)
	}
}
