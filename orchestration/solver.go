package orchestration

import "time"

type Solver struct {
    process func(string, *Result) error
}

func (s *Solver) Handle(id string, payload string, results chan<- Result) {
    start := time.Now()
    result := Result{id, nil, 0, make([]string, 0)}
    result.Err = s.process(payload, &result)
    result.Duration = time.Since(start)
    results <- result
}

func NewSolver(process func(string, *Result) error) Solver {
    return Solver{
        process: process,
    }
}
