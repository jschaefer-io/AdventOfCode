package day14

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

func Solve(data string, result *orchestration.Result) error {
    mappings := make(map[string]string)
    groups := strings.Split(data, "\n\n")
    for _, line := range strings.Split(groups[1], "\n") {
        if len(line) == 0 {
            continue
        }
        parts := strings.Split(line, " -> ")
        mappings[parts[0]] = parts[1]
    }

    str := groups[0]
    lastChar := rune(str[0])
    firstChar := rune(str[len(str)-1])
    pairs := NewPairCounter(str)
    for i := 0; i < 10; i++ {
        pairs = pairs.ApplyPolymer(mappings)
    }
    result.AddResult(strconv.Itoa(pairs.Score(firstChar, lastChar)))

    for i := 0; i < 30; i++ {
        pairs = pairs.ApplyPolymer(mappings)
    }
    result.AddResult(strconv.Itoa(pairs.Score(firstChar, lastChar)))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day14", orchestration.NewSolver(Solve))
}
