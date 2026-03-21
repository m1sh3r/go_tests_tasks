package mocks

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAuditedSolverSuccessLogsInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := NewMockLogger(ctrl)
	solver := NewAuditedSolver(logger)

	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()

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

func TestAuditedSolverErrorCallsErrorLogInOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := NewMockLogger(ctrl)
	solver := NewAuditedSolver(logger)

	gomock.InOrder(
		logger.EXPECT().Info(gomock.Any(), gomock.Any()),
		logger.EXPECT().Error(gomock.Any(), gomock.Any()),
	)

	_, err := solver.Solve("eq_err", 0, 2, 1)
	if !errors.Is(err, ErrInvalidLinearTerm) {
		t.Errorf("Solve(eq_err): ожидалась ошибка %v, получено %v", ErrInvalidLinearTerm, err)
	}
}
