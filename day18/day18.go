package day18

import (
    "fmt"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Formula interface {
    fmt.Stringer
    split() (bool, Formula)
    explode(depth int, parents ExplodeTrace) bool
    add(f Formula) Formula
    Magnitude() int
    Copy() Formula
}

func ParseNumber(str string) Formula {
    strLen := len(str)
    if strLen == 1 {
        n, _ := strconv.Atoi(str)
        return &Number{n}
    }
    group := 0
    for i, r := range str[1 : strLen-1] {
        switch r {
        case '[':
            group++
        case ']':
            group--
        case ',':
            if group == 0 {
                return &Compound{
                    left:  ParseNumber(str[1 : i+1]),
                    right: ParseNumber(str[i+2 : strLen-1]),
                }
            }
        }
    }
    panic("invalid number format")
}

func ExecuteAddition(a, b Formula) Formula {
    res := a.add(b)
    for {
        if res.explode(0, make(ExplodeTrace, 0)) {
            continue
        }
        split, newRes := res.split()
        if split {
            res = newRes
            continue
        }
        break
    }
    return res
}

func Solve(data string, result *orchestration.Result) error {
    numbers := make([]Formula, 0)
    for _, line := range strings.Split(data, "\n") {
        if len(line) == 0 {
            continue
        }
        numbers = append(numbers, ParseNumber(line))
    }

    // A
    res := numbers[0]
    for i := 1; i < len(numbers); i++ {
        res = ExecuteAddition(res, numbers[i])
    }
    result.AddResult(strconv.Itoa(res.Magnitude()))

    // B
    max := -1
    for iA, a := range numbers {
        for iB, b := range numbers {
            if iA == iB {
                continue
            }
            res = ExecuteAddition(a, b)
            mag := res.Magnitude()
            if max == -1 || mag > max {
                max = mag
            }
        }
    }
    result.AddResult(strconv.Itoa(max))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day18", orchestration.NewSolver(Solve))
}
