package day10

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "sort"
    "strconv"
    "strings"
)

var BracketMapping = map[rune]rune{
    '(': ')',
    '[': ']',
    '{': '}',
    '<': '>',
}

var ErrorPoints = map[rune]int{
    ')': 3,
    ']': 57,
    '}': 1197,
    '>': 25137,
}

var CompletionPoints = map[rune]int{
    ')': 1,
    ']': 2,
    '}': 3,
    '>': 4,
}

func CheckLine(str string) (bool, Stack, *rune) {
    stack := NewStack()
    for _, c := range str {
        v, ok := BracketMapping[c]
        if ok {
            stack.Push(v)
        } else {
            check := stack.Peek()
            if check != c {
                return false, stack, &c
            }
            stack.Pop()
        }
    }
    return stack.length == 0, stack, nil
}

func Solve(data string, result *orchestration.Result) error {
    lines := make([]string, 0)
    for _, line := range strings.Split(data, "\n") {
        if len(line) == 0 {
            continue
        }
        lines = append(lines, line)
    }

    // a
    stacks := make([]Stack, 0)
    pointSum := 0
    for _, line := range lines {
        completed, stack, err := CheckLine(line)
        if err != nil {
            p := ErrorPoints[*err]
            pointSum += p
        } else if !completed {
            stacks = append(stacks, stack)
        }
    }
    result.AddResult(strconv.Itoa(pointSum))

    // b
    stackPointList := make([]int, 0)
    for _, stack := range stacks {
        stackPointSum := 0
        for stack.Length() > 0 {
            p := CompletionPoints[stack.Pop()]
            stackPointSum *= 5
            stackPointSum += p
        }
        stackPointList = append(stackPointList, stackPointSum)
    }
    sort.Ints(stackPointList)
    result.AddResult(strconv.Itoa(stackPointList[len(stackPointList)/2]))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day10", orchestration.NewSolver(Solve))
}
