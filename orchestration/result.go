package orchestration

import "time"

type Result struct {
    Id       string
    Err      error
    Duration time.Duration
    Results  []string
}

func (r *Result) AddResult(result string) {
    r.Results = append(r.Results, result)
}
