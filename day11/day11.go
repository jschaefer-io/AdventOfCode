package day11

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
)

func Solve(data string, result *orchestration.Result) error {
    swarm, err := NewSwarm(data)
    if err != nil {
        return err
    }

    sum := 0
    stepCount := 0
    for stepCount < 100 {
        stepCount++
        sum += swarm.Step()
    }
    result.AddResult(strconv.Itoa(sum))

    for {
        stepCount++
        if swarm.Step() == 100 {
            break
        }
    }
    result.AddResult(strconv.Itoa(stepCount))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day11", orchestration.NewSolver(Solve))
}
