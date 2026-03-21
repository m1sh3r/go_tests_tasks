package mocks

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Cache interface {
	Get(key string) (string, bool)
	Set(key, value string, ttl time.Duration)
}

var (
	ErrInvalidCoeffsFormat = errors.New("некорректный формат коэффициентов")
	ErrInvalidLinearTerm   = errors.New("коэффициент a не должен быть равен нулю")
)

type CachedSolver struct {
	cache Cache
	ttl   time.Duration
}

func NewCachedSolver(cache Cache, ttl time.Duration) *CachedSolver {
	return &CachedSolver{
		cache: cache,
		ttl:   ttl,
	}
}

func (s *CachedSolver) Solve(eqID string, coeffs string) ([]float64, error) {
	if cached, ok := s.cache.Get(eqID); ok {
		return solveFromCSV(cached)
	}

	roots, err := solveFromCSV(coeffs)
	if err != nil {
		return nil, err
	}

	s.cache.Set(eqID, coeffs, s.ttl)
	return roots, nil
}

func solveFromCSV(csv string) ([]float64, error) {
	parts := strings.Split(csv, ",")
	if len(parts) != 3 {
		return nil, ErrInvalidCoeffsFormat
	}

	a, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidCoeffsFormat, err)
	}
	b, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidCoeffsFormat, err)
	}
	c, err := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidCoeffsFormat, err)
	}

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
