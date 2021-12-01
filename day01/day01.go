package day01

import (
    "fmt"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

func offSetCount(s int, offset int, list []int) int {
    sum := 0
    for i := s; i > s-offset; i-- {
        sum += list[i]
    }
    return sum
}

func traverseAndCompare(groupSize int, list []int) int {
    length := len(list)
    c := 0
    for i := groupSize; i < length; i++ {
        if offSetCount(i, groupSize, list) > offSetCount(i-1, groupSize, list) {
            c++
        }
    }
    return c
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day01", orchestration.NewSolver(func(data string, result *orchestration.Result) error {
        depths := make([]int, 0)
        for _, line := range strings.Split(data, "\n") {
            v, _ := strconv.Atoi(line)
            depths = append(depths, v)
        }
        result.AddResult(fmt.Sprintf("%d", traverseAndCompare(1, depths)))
        result.AddResult(fmt.Sprintf("%d", traverseAndCompare(3, depths)))
        return nil
    }))
}
