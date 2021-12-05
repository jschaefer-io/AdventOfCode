package day05

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "math"
    "regexp"
    "strconv"
    "strings"
)

func init() {
    orchestration.MainDispatcher.AddSolver("Day05", orchestration.NewSolver(func(data string, result *orchestration.Result) error {
        exp := regexp.MustCompile("(\\d+)")
        xMax := 0
        yMax := 0
        lines := make([]Line, 0)
        for _, line := range strings.Split(data, "\n") {
            if len(line) == 0 {
                continue
            }
            match := exp.FindAllStringSubmatch(line, -1)
            x1, _ := strconv.Atoi(match[0][0])
            y1, _ := strconv.Atoi(match[1][0])
            x2, _ := strconv.Atoi(match[2][0])
            y2, _ := strconv.Atoi(match[3][0])
            xMax = int(math.Max(float64(xMax), math.Max(float64(x1), float64(x2))))
            yMax = int(math.Max(float64(yMax), math.Max(float64(y1), float64(y2))))
            lines = append(lines, Line{
                From: Coordinate{x1, y1},
                To:   Coordinate{x2, y2},
            })
        }

        mA := NewMap(yMax+1, xMax+1)
        mB := NewMap(yMax+1, xMax+1)
        for _, line := range lines {
            if !line.Diagonal() {
                mA.DrawLine(line)
            }
            mB.DrawLine(line)
        }
        result.AddResult(strconv.Itoa(mA.CountMin(2)))
        result.AddResult(strconv.Itoa(mB.CountMin(2)))
        return nil
    }))
}
