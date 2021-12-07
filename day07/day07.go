package day07

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "math"
    "sort"
    "strconv"
    "strings"
)

func Average(list []int) (int, int) {
    sum := 0
    for _, n := range list {
        sum += n
    }
    avg := float64(sum) / float64(len(list))
    return int(math.Floor(avg)), int(math.Ceil(avg))
}

func Median(list []int) int {
    count := len(list)
    sortedList := make([]int, count)
    for i, _ := range list {
        sortedList[i] = list[i]
    }
    sort.Ints(sortedList)
    return sortedList[int(math.Floor(float64(count-1)/2))]
}

func ShiftToPosition(list []int, target int) []int {
    offsets := make([]int, len(list))
    for i, c := range list {
        offsets[i] = int(math.Abs(float64(target - c)))
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
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }

    result.AddResult(strconv.Itoa(SumList(ShiftToPosition(crabs, Median(crabs)))))
    fAvg, cAvg := Average(crabs)
    b := math.Min(float64(SumListIncremental(ShiftToPosition(crabs, fAvg))), float64(SumListIncremental(ShiftToPosition(crabs, cAvg))))
    result.AddResult(strconv.Itoa(int(b)))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day07", orchestration.NewSolver(Solve))
}
