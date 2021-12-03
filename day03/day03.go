package day03

import (
    "errors"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Input struct {
    S     string
    Value uint
}

func (i Input) Transform(transformer func(base uint) uint) Input {
    var mask uint = 0
    for i := range i.S {
        mask |= 1 << i
    }
    v := transformer(i.Value) & mask
    s := strconv.FormatInt(int64(v), 2)
    deltaLength := len(i.S) - len(s)
    var sb strings.Builder
    for i := 0; i < deltaLength; i++ {
        sb.WriteString("0")
    }
    sb.WriteString(s)
    return Input{sb.String(), v}
}

func DeriveInput(list []Input) (Input, error) {
    var sb strings.Builder
    count := len(list[0].S)
    countItems := len(list)
    for pos := 0; pos < count; pos++ {
        oneCount := 0
        for _, input := range list {
            if input.S[pos] == '1' {
                oneCount++
            }
        }
        if oneCount*2 >= countItems {
            sb.WriteString("1")
        } else {
            sb.WriteString("0")
        }
    }

    s := sb.String()
    v, err := strconv.ParseInt(s, 2, 64)
    if err != nil {
        return Input{}, nil
    }
    return Input{s, uint(v)}, nil
}

func applyNot(a uint) uint {
    return ^a
}

func ReduceDerive(list []Input, derive func(queue []Input) (Input, error)) (Input, error) {
    queue := list
    bitIndex := 0
    for len(queue) > 1 {
        i, err := derive(queue)
        if err != nil {
            return Input{}, err
        }
        newQueue := make([]Input, 0)
        for _, item := range queue {
            if item.S[bitIndex] == i.S[bitIndex] {
                newQueue = append(newQueue, item)
            }
        }
        queue = newQueue
        bitIndex++
    }
    if len(queue) != 1 {
        return Input{}, errors.New("no single item found")
    }
    return queue[0], nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day03", orchestration.NewSolver(func(data string, result *orchestration.Result) error {
        lines := strings.Split(data, "\n")
        list := make([]Input, 0)
        for _, line := range lines {
            if len(line) == 0 {
                continue
            }
            v, err := strconv.ParseInt(line, 2, 64)
            if err != nil {
                return err
            }
            list = append(list, Input{line, uint(v)})
        }

        // A
        a, err := DeriveInput(list)
        if err != nil {
            return err
        }
        aInverse := a.Transform(applyNot)
        result.AddResult(strconv.Itoa(int(a.Value * aInverse.Value)))

        //B
        bOxy, err := ReduceDerive(list, func(queue []Input) (Input, error) {
            return DeriveInput(queue)
        })
        if err != nil {
            return err
        }
        bCo2, err := ReduceDerive(list, func(queue []Input) (Input, error) {
            l, err := DeriveInput(queue)
            if err != nil {
                return l, err
            }
            return l.Transform(applyNot), nil
        })
        if err != nil {
            return err
        }
        result.AddResult(strconv.Itoa(int(bOxy.Value * bCo2.Value)))
        return nil
    }))
}
