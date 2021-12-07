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

func Solve(data string, result *orchestration.Result) error {
    crabs := make([]int, 0)
    min := math.MaxInt64
    max := 0
    for _, crab := range strings.Split(data, ",") {
        n, err := strconv.Atoi(strings.Trim(crab, " \n"))
        if err != nil {
            return err
        }
        crabs = append(crabs, n)
        if n < min{
            min = n
        }
        if n > max {
            max = n
        }
    }

    minFuelA := SumList(ShiftToPosition(crabs, min))
    minFuelB := SumListIncremental(ShiftToPosition(crabs, min))
    for p := min + 1; p < max; p++ {
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
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day07", orchestration.NewSolver(Solve))
}
