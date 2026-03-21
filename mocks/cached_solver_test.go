package mocks

import (
	"math"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func equalRoots(got, want []float64) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if math.Abs(got[i]-want[i]) > 1e-9 {
			return false
		}
	}
	return true
}

func TestCachedSolverFromCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cache := NewMockCache(ctrl)
	solver := NewCachedSolver(cache, time.Minute)

	cache.EXPECT().Get("eq1").Return("1,-3,2", true)

	got, err := solver.Solve("eq1", "100,100,100")
	if err != nil {
		t.Errorf("Solve(eq1): неожиданная ошибка: %v", err)
		return
	}

	want := []float64{2, 1}
	if !equalRoots(got, want) {
		t.Errorf("Solve(eq1): ожидались корни %v, получены %v", want, got)
	}
}

func TestCachedSolverSetOnMiss(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cache := NewMockCache(ctrl)
	ttl := 5 * time.Minute
	solver := NewCachedSolver(cache, ttl)

	cache.EXPECT().Get("eq2").Return("", false)
	cache.EXPECT().Set("eq2", "1,-3,2", ttl).Times(1)

	got, err := solver.Solve("eq2", "1,-3,2")
	if err != nil {
		t.Errorf("Solve(eq2): неожиданная ошибка: %v", err)
		return
	}

	want := []float64{2, 1}
	if !equalRoots(got, want) {
		t.Errorf("Solve(eq2): ожидались корни %v, получены %v", want, got)
	}
}
