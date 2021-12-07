package day07

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "math"
    "strconv"
    "strings"
)

func ShiftToPosition(list []int, target int) []int {
    offsets := make([]int, 0)
    for _, c := range list {
        offsets = append(offsets, int(math.Abs(float64(target-c))))
    }
    return offsets
}

func SumList(list []int) int {
    sum := 0
    for _, n := range list {
        sum += n
    }
    return sum
}

func SumListIncremental(list []int) int {
    sum := 0
    for _, n := range list {
        sum += (n * (n + 1)) / 2
    }
    return sum
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day07", orchestration.NewSolver(func(data string, result *orchestration.Result) error {
        crabs := make([]int, 0)
        for _, crab := range strings.Split(data, ",") {
            n, err := strconv.Atoi(strings.Trim(crab, " \n"))
            if err != nil {
                return err
            }
            crabs = append(crabs, n)
        }

        crabCount := len(crabs)
        minFuelA := SumList(ShiftToPosition(crabs, 0))
        minFuelB := SumListIncremental(ShiftToPosition(crabs, 0))
        for p := 1; p < crabCount; p++ {
            shift := ShiftToPosition(crabs, p)
            sFuelA := SumList(shift)
            if sFuelA < minFuelA {
                minFuelA = sFuelA
            }

            sFuelB := SumListIncremental(shift)
            if sFuelB < minFuelB {
                minFuelB = sFuelB
            }
        }
        result.AddResult(strconv.Itoa(minFuelA))
        result.AddResult(strconv.Itoa(minFuelB))
        return nil
    }))
}
