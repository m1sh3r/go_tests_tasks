package mocks

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestRouterUsesFactoryAndSolver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	factory := NewMockSolverFactory(ctrl)
	solver := NewMockSolver(ctrl)
	router := NewRouter(factory)
	coeffs := []float64{1, -3, 2}

	factory.EXPECT().Create("quadratic").Return(solver, nil)
	solver.EXPECT().Solve(coeffs).Return([]float64{2, 1}, nil)

	got, err := router.RouteSolve("quadratic", coeffs)
	if err != nil {
		t.Errorf("RouteSolve(quadratic): неожиданная ошибка: %v", err)
		return
	}

	want := []float64{2, 1}
	if !equalRoots(got, want) {
		t.Errorf("RouteSolve(quadratic): ожидались корни %v, получены %v", want, got)
	}
}

func TestRouterFactoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	factory := NewMockSolverFactory(ctrl)
	router := NewRouter(factory)
	createErr := errors.New("тип не поддерживается")

	factory.EXPECT().Create("cubic").Return(nil, createErr)

	_, err := router.RouteSolve("cubic", []float64{1, 0, 0, 1})
	if !errors.Is(err, createErr) {
		t.Errorf("RouteSolve(cubic): ожидалась ошибка %v, получено %v", createErr, err)
	}
}
