package day02

import (
    "errors"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Operation struct {
    Name  string
    Value int
}

type Submarine struct {
    Horizontal int
    Depth      int
    Aim        int
    opHandler  func(op Operation, s *Submarine) error
}

func (s *Submarine) Move(op Operation) error{
    return s.opHandler(op, s)
}

func MoveA(op Operation, s *Submarine) error {
    switch op.Name {
    case "forward":
        s.Horizontal += op.Value
    case "down":
        s.Depth += op.Value
    case "up":
        s.Depth -= op.Value
    default:
        return errors.New("op not defined")
    }
    return nil
}

func MoveB(op Operation, s *Submarine) error {
    switch op.Name {
    case "forward":
        s.Horizontal += op.Value
        s.Depth += op.Value * s.Aim
    case "down":
        s.Aim += op.Value
    case "up":
        s.Aim -= op.Value
    default:
        return errors.New("op not defined")
    }
    return nil
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

        subs := make([]*Submarine, 0)
        subs = append(subs, &Submarine{opHandler: MoveA})
        subs = append(subs, &Submarine{opHandler: MoveB})
        for _, op := range ops {
            for _, sub := range subs {
                err := sub.Move(op)
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
