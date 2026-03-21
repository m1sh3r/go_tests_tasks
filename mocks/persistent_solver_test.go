package mocks

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestPersistentSolverSaveArgs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := NewMockHistoryStore(ctrl)
	logger := NewMockLogger(ctrl)
	solver := NewPersistentSolver(store, logger)

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

func TestPersistentSolverReturnsRootsOnSaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := NewMockHistoryStore(ctrl)
	logger := NewMockLogger(ctrl)
	solver := NewPersistentSolver(store, logger)

	saveErr := errors.New("ошибка БД")
	store.EXPECT().Save(gomock.Eq("eq1"), gomock.Len(2)).Return(saveErr)
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).Times(1)

	got, err := solver.Solve("eq1", 1, -3, 2)
	if err != nil {
		t.Errorf("Solve(eq1): ожидалось решение без ошибки, получено %v", err)
		return
	}

	want := []float64{2, 1}
	if !equalRoots(got, want) {
		t.Errorf("Solve(eq1): ожидались корни %v, получены %v", want, got)
	}
}
