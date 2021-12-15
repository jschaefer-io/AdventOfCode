package day15

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Position struct {
    X    int
    Y    int
    Risk int
}

func SolveForField(field Field, to, from Position) int {
    paths := field.ShortestPaths(from)
    sum := 0
    for _, v := range paths.GetPath(to) {
        sum += v.Risk
    }
    return sum
}

func Solve(data string, result *orchestration.Result) error {
    field := Field{
        Position: make([][]Position, 0),
    }
    for y, line := range strings.Split(data, "\n") {
        if len(line) == 0 {
            continue
        }
        row := make([]Position, 0)
        for x, risk := range strings.Split(line, "") {
            n, err := strconv.Atoi(risk)
            if err != nil {
                return err
            }
            row = append(row, Position{x, y, n})
        }
        field.Position = append(field.Position, row)
    }
    field.Height = len(field.Position)
    field.Width = len(field.Position[0])

    a := SolveForField(field, field.Position[0][0], field.Position[field.Height-1][field.Width-1])
    result.AddResult(strconv.Itoa(a))

    largeField := field.Expand(5)
    b := SolveForField(largeField, largeField.Position[0][0], largeField.Position[largeField.Height-1][largeField.Width-1])
    result.AddResult(strconv.Itoa(b))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day15", orchestration.NewSolver(Solve))
}
