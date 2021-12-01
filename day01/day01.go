package day01

import (
    "fmt"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

func init() {
    orchestration.MainDispatcher.AddSolver("Day01", orchestration.NewSolver(func(data string, result *orchestration.Result) error {
        depths := make([]int, 0)
        for _, line := range strings.Split(data, "\n") {
            v, _ := strconv.Atoi(line)
            depths = append(depths, v)
        }
        length := len(depths)

        aCount := 0
        for i := 1; i < length; i++ {
            if depths[i] > depths[i-1] {
                aCount++
            }
        }
        result.AddResult(fmt.Sprintf("%d", aCount))

        bCount := 0
        for i := 3; i < length; i++ {
            sumA := depths[i] + depths[i-1] + depths[i-2]
            sumB := depths[i-1] + depths[i-2] + depths[i-3]
            if sumA > sumB {
                bCount++
            }
        }
        result.AddResult(fmt.Sprintf("%d", bCount))
        return nil
    }))
}
