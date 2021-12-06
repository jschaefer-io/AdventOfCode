package day05

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

func init() {
    orchestration.MainDispatcher.AddSolver("Day06", orchestration.NewSolver(func(data string, result *orchestration.Result) error {
        fish := make([]int8, 0)
        for _, f := range strings.Split(data, ",") {
            n, err := strconv.Atoi(strings.Trim(f, " \n"))
            if err != nil {
                return err
            }
            fish = append(fish, int8(n))
        }

        fL := NewFishList(fish)
        aLimit := 80
        bLimit := 256

        // A
        for s := 0; s < aLimit; s++ {
            fL.Tick()
        }
        result.AddResult(strconv.Itoa(fL.Count()))

        // B
        for s := 0; s < bLimit-aLimit; s++ {
            fL.Tick()
        }
        result.AddResult(strconv.Itoa(fL.Count()))
        return nil
    }))
}
