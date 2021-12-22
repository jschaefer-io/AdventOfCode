package day22

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "regexp"
    "strconv"
    "strings"
)

func Solve(data string, result *orchestration.Result) error {
    exp := regexp.MustCompile("(?:|-)\\d+")
    operations := make([]Operation, 0)
    for _, line := range strings.Split(data, "\n") {
        if len(line) == 0 {
            continue
        }
        group := strings.Split(line, " ")
        coords := exp.FindAllStringSubmatch(group[1], -1)
        x1, err := strconv.Atoi(coords[0][0])
        if err != nil {
            return err
        }
        x2, err := strconv.Atoi(coords[1][0])
        if err != nil {
            return err
        }
        y1, err := strconv.Atoi(coords[2][0])
        if err != nil {
            return err
        }
        y2, err := strconv.Atoi(coords[3][0])
        if err != nil {
            return err
        }
        z1, err := strconv.Atoi(coords[4][0])
        if err != nil {
            return err
        }
        z2, err := strconv.Atoi(coords[5][0])
        if err != nil {
            return err
        }
        operations = append(operations, Operation{group[0] == "on", Area{x1, x2, y1, y2, z1, z2}})
    }

    g := Grid{list: make([]Area, 0)}
    for _, operation := range operations {
        g.ApplyOperation(operation)
    }

    // a
    result.AddResult(strconv.Itoa(g.CountActive(&Area{-50, 50, -50, 50, -50, 50})))

    // b
    result.AddResult(strconv.Itoa(g.CountActive(nil)))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day22", orchestration.NewSolver(Solve))
}
