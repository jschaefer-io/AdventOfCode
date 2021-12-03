package day03

import (
    "errors"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Input struct {
    length int
    value  uint
}

func (i Input) Value() uint {
    return i.value
}

func (i Input) Length() int {
    return i.length
}

func (i Input) CheckBit(position int) bool {
    return (i.value & (1 << position)) > 0
}

func NewInput(value string) (Input, error) {
    v, err := strconv.ParseInt(value, 2, 64)
    if err != nil {
        return Input{}, err
    }
    return Input{
        length: len(value),
        value:  uint(v),
    }, nil
}

func (i *Input) setValue(v uint) {
    var bitMap uint = 0
    for u := 0; u < i.length-1; u++ {
        bitMap |= 1 << u
    }
    i.value = v & bitMap
}

func (i Input) Transform(transformer func(base uint) uint) Input {
    value := transformer(i.value)
    i.setValue(value)
    return i
}

func DeriveInput(list []Input) Input {
    count := list[0].Length()
    countItems := len(list)
    var value uint = 0
    for pos := count - 1; pos >= 0; pos-- {
        oneCount := 0
        for _, input := range list {
            if input.CheckBit(pos) {
                oneCount++
            }
        }
        if oneCount*2 >= countItems {
            value |= 1 << pos
        }
    }
    return Input{count, value}
}

func ReduceDerive(list []Input, derive func(queue []Input) Input) (Input, error) {
    queue := list
    bitIndex := 0
    for len(queue) > 1 {
        i := derive(queue)
        newQueue := make([]Input, 0)
        for _, item := range queue {
            if item.CheckBit(item.Length()-bitIndex) == i.CheckBit(i.Length()-bitIndex) {
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

func applyNot(a uint) uint {
    return ^a
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day03", orchestration.NewSolver(func(data string, result *orchestration.Result) error {
        lines := strings.Split(data, "\n")
        list := make([]Input, 0)
        for _, line := range lines {
            if len(line) == 0 {
                continue
            }
            lInput, err := NewInput(line)
            if err != nil {
                return err
            }
            list = append(list, lInput)
        }

        // A
        a := DeriveInput(list)
        aInverse := a.Transform(applyNot)
        result.AddResult(strconv.Itoa(int(a.Value() * aInverse.Value())))

        // B
        bOxy, err := ReduceDerive(list, func(queue []Input) Input {
            return DeriveInput(queue)
        })
        if err != nil {
            return err
        }
        bCo2, err := ReduceDerive(list, func(queue []Input) Input {
            return DeriveInput(queue).Transform(applyNot)
        })
        if err != nil {
            return err
        }
        result.AddResult(strconv.Itoa(int(bOxy.Value() * bCo2.Value())))
        return nil
    }))
}
