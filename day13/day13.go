package day13

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Fold struct {
    Position int
    Axis     bool
}

func Solve(data string, result *orchestration.Result) error {
    groups := strings.Split(data, "\n\n")
    dotList := make([][2]int, 0)
    for _, dots := range strings.Split(groups[0], "\n") {
        values := strings.Split(dots, ",")
        x, err := strconv.Atoi(values[0])
        if err != nil {
            return err
        }
        y, err := strconv.Atoi(values[1])
        if err != nil {
            return err
        }
        dotList = append(dotList, [2]int{x, y})
    }

    operations := make([]Fold, 0)
    for _, op := range strings.Split(groups[1], "\n") {
        if len(op) == 0 {
            continue
        }
        values := strings.Split(op, "=")
        n, err := strconv.Atoi(values[1])
        if err != nil {
            return err
        }
        operations = append(operations, Fold{
            Axis:     strings.Contains(values[0], "x"),
            Position: n,
        })
    }

    paper := NewPaper(dotList)
    firstFold := paper.Fold(operations[0])
    result.AddResult(strconv.Itoa(firstFold.Points()))

    for _, op := range operations {
        paper = paper.Fold(op)
    }
    result.AddResult("\n" + paper.String())
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day13", orchestration.NewSolver(Solve))
}
