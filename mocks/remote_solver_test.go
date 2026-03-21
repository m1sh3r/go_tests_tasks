package mocks

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestRemoteSolverSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockAPIClient(ctrl)
	solver := NewRemoteSolver(client, time.Second)

	client.EXPECT().FetchCoeffs("eq_ok").Return(1.0, -5.0, 6.0, nil)

	got, err := solver.Solve("eq_ok")
	if err != nil {
		t.Errorf("Solve(eq_ok): неожиданная ошибка: %v", err)
		return
	}

	want := []float64{3, 2}
	if !equalRoots(got, want) {
		t.Errorf("Solve(eq_ok): ожидались корни %v, получены %v", want, got)
	}
}

func TestRemoteSolverNetworkError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockAPIClient(ctrl)
	solver := NewRemoteSolver(client, time.Second)
	apiErr := errors.New("ошибка сети")

	client.EXPECT().FetchCoeffs("eq_err").Return(0.0, 0.0, 0.0, apiErr)

	_, err := solver.Solve("eq_err")
	if !errors.Is(err, apiErr) {
		t.Errorf("Solve(eq_err): ожидалась ошибка %v, получено %v", apiErr, err)
	}
}

func TestRemoteSolverTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockAPIClient(ctrl)
	solver := NewRemoteSolver(client, 50*time.Millisecond)

	client.EXPECT().FetchCoeffs("eq_slow").DoAndReturn(func(string) (float64, float64, float64, error) {
		time.Sleep(120 * time.Millisecond)
		return 1.0, -3.0, 2.0, nil
	})

	_, err := solver.Solve("eq_slow")
	if !errors.Is(err, ErrAPITimeout) {
		t.Errorf("Solve(eq_slow): ожидался таймаут %v, получено %v", ErrAPITimeout, err)
	}
}
