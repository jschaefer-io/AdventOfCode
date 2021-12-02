package day02

import (
    "errors"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Submarine struct {
    Horizontal int
    Depth      int
    Aim        int
}

func (s *Submarine) Move(op Operation, b bool) error {
    switch op.Name {
    case "forward":
        s.Horizontal += op.Value
        if b {
            s.Depth += op.Value * s.Aim
        }
    case "down":
        if b {
            s.Aim += op.Value
        } else {
            s.Depth += op.Value
        }
    case "up":
        if b {
            s.Aim -= op.Value
        } else {
            s.Depth -= op.Value
        }
    default:
        return errors.New("op not defined")
    }
    return nil
}

type Operation struct {
    Name  string
    Value int
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day02", orchestration.NewSolver(func(data string, result *orchestration.Result) error {
        lines := strings.Split(data, "\n")
        ops := make([]Operation, 0)
        for _, line := range lines {
            lData := strings.Split(line, " ")
            if len(lData) != 2 {
                continue
            }
            value, _ := strconv.Atoi(lData[1])
            ops = append(ops, Operation{lData[0], value})
        }

        subs := [2]*Submarine{
            {},
            {},
        }
        for _, op := range ops {
            for i, sub := range subs {
                err := sub.Move(op, i != 0)
                if err != nil {
                    return err
                }
            }
        }

        for _, sub := range subs {
            result.AddResult(strconv.Itoa(sub.Depth * sub.Horizontal))
        }
        return nil
    }))
}
