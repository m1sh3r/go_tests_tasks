package mocks

import (
	"errors"
	"math"
	"time"
)

type APIClient interface {
	FetchCoeffs(eqID string) (a, b, c float64, err error)
}

var ErrAPITimeout = errors.New("таймаут запроса к внешнему API")

type RemoteSolver struct {
	client  APIClient
	timeout time.Duration
}

func NewRemoteSolver(client APIClient, timeout time.Duration) *RemoteSolver {
	return &RemoteSolver{
		client:  client,
		timeout: timeout,
	}
}

func (s *RemoteSolver) Solve(eqID string) ([]float64, error) {
	type result struct {
		a   float64
		b   float64
		c   float64
		err error
	}

	ch := make(chan result, 1)
	go func() {
		a, b, c, err := s.client.FetchCoeffs(eqID)
		ch <- result{a: a, b: b, c: c, err: err}
	}()

	select {
	case r := <-ch:
		if r.err != nil {
			return nil, r.err
		}
		return solveQuadraticFromCoeffs(r.a, r.b, r.c)
	case <-time.After(s.timeout):
		return nil, ErrAPITimeout
	}
}

func solveQuadraticFromCoeffs(a, b, c float64) ([]float64, error) {
	if a == 0 {
		return nil, ErrInvalidLinearTerm
	}

	d := b*b - 4*a*c
	if d < 0 {
		return []float64{}, nil
	}
	if d == 0 {
		return []float64{-b / (2 * a)}, nil
	}

	sqrtD := math.Sqrt(d)
	x1 := (-b + sqrtD) / (2 * a)
	x2 := (-b - sqrtD) / (2 * a)
	return []float64{x1, x2}, nil
}
