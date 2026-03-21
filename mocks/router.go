package mocks

type SolverFactory interface {
	Create(eqType string) (Solver, error)
}

type Solver interface {
	Solve(coeffs []float64) ([]float64, error)
}

type Router struct {
	factory SolverFactory
}

func NewRouter(factory SolverFactory) *Router {
	return &Router{factory: factory}
}

func (r *Router) RouteSolve(eqType string, coeffs []float64) ([]float64, error) {
	solver, err := r.factory.Create(eqType)
	if err != nil {
		return nil, err
	}

	return solver.Solve(coeffs)
}
