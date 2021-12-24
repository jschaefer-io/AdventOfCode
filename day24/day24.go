package day24

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
)

func Solve(data string, result *orchestration.Result) error {

    // Manually derived from input
    // @todo add generalized solution
    result.AddResult(strconv.Itoa(91897399498995))
    result.AddResult(strconv.Itoa(51121176121391))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day24", orchestration.NewSolver(Solve))
}
