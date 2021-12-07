package main

import (
    _ "github.com/jschaefer-io/aoc2021/day01"
    _ "github.com/jschaefer-io/aoc2021/day02"
    _ "github.com/jschaefer-io/aoc2021/day03"
    _ "github.com/jschaefer-io/aoc2021/day04"
    _ "github.com/jschaefer-io/aoc2021/day05"
    _ "github.com/jschaefer-io/aoc2021/day06"
    _ "github.com/jschaefer-io/aoc2021/day07"
    "github.com/jschaefer-io/aoc2021/orchestration"
    "os"
    "strings"
)

func getAocWorkload() orchestration.WorkLoad {
    dir, err := os.ReadDir("./_inputs")
    if err != nil {
        panic(err)
    }
    w := make(orchestration.WorkLoad)
    for _, entry := range dir {
        if !entry.IsDir() {
            name := strings.ReplaceAll(entry.Name(), ".txt", "")
            if !strings.Contains(entry.Name(), ".txt") {
                continue
            }
            content, err := os.ReadFile("./_inputs/" + entry.Name())
            if err != nil {
                panic(err)
            }
            w[name] = string(content)
        }
    }
    return w
}

func main() {
    orchestration.MainDispatcher.Start(getAocWorkload())
}
