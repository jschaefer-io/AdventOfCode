package day01

import (
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

func TraverseAndCompare(groupSize int, list []int) int {
    length := len(list)
    c := 0
    for i := groupSize; i < length; i++ {
        if offSetCount(i, groupSize, list) > offSetCount(i-1, groupSize, list) {
            c++
        }
    }
    return c
}

func Solve(data string, result *orchestration.Result) error {
    depths := make([]int, 0)
    for _, line := range strings.Split(data, "\n") {
        v, _ := strconv.Atoi(line)
        depths = append(depths, v)
    }
    result.AddResult(strconv.Itoa(TraverseAndCompare(1, depths)))
    result.AddResult(strconv.Itoa(TraverseAndCompare(3, depths)))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day01", orchestration.NewSolver(Solve))
}
