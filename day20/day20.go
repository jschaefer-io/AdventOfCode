package day20

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Operation struct {
    X     int
    Y     int
    Value int
}

func Solve(data string, result *orchestration.Result) error {
    groups := strings.Split(data, "\n\n")
    mapping := groups[0]

    img := NewImage(mapping)
    for y, line := range strings.Split(groups[1], "\n") {
        if len(line) == 0 {
            continue
        }
        for x, c := range strings.Split(line, "") {
            img.AddPixel(x, y, c == "#")
        }
    }

    count := 0
    for count < 2 {
        img.Step()
        count++
    }
    result.AddResult(strconv.Itoa(img.CountPixels()))

    for count < 50 {
        img.Step()
        count++
    }
    result.AddResult(strconv.Itoa(img.CountPixels()))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day20", orchestration.NewSolver(Solve))
}
