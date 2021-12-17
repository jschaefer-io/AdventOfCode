package day17

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "regexp"
    "strconv"
)

func SumTo(n int) int {
    return (n * (n + 1)) / 2
}

type Point struct {
    X int
    Y int
}

func Solve(data string, result *orchestration.Result) error {
    exp := regexp.MustCompile("(?:-|)\\d+")
    nums := make([]int, 4)
    for i, sn := range exp.FindAllStringSubmatch(data, -1) {
        n, err := strconv.Atoi(sn[0])
        if err != nil {
            return err
        }
        nums[i] = n
    }
    target := Target{nums[0], nums[1], nums[2], nums[3]}

    yMax := 0
    count := 0
    for x := target.MinX(); x <= 400; x++ {
        for y := -100; y < 400; y++ {
            t := NewTrajectory(x, y)
            for s := 0; s < 400; s++ {
                p := t.GetPoint(s)
                if target.CheckHit(p) {
                    if t.maxY > yMax {
                        yMax = t.maxY
                    }
                    count++
                    break
                }
                if p.X > target.X2 || p.Y < target.Y1 {
                    break
                }
            }
        }
    }
    result.AddResult(strconv.Itoa(yMax))
    result.AddResult(strconv.Itoa(count))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day17", orchestration.NewSolver(Solve))
}
