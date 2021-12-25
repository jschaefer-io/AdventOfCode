package main

import (
    _ "github.com/jschaefer-io/aoc2021/day01"
    _ "github.com/jschaefer-io/aoc2021/day02"
    _ "github.com/jschaefer-io/aoc2021/day03"
    _ "github.com/jschaefer-io/aoc2021/day04"
    _ "github.com/jschaefer-io/aoc2021/day05"
    _ "github.com/jschaefer-io/aoc2021/day06"
    _ "github.com/jschaefer-io/aoc2021/day07"
    _ "github.com/jschaefer-io/aoc2021/day08"
    _ "github.com/jschaefer-io/aoc2021/day09"
    _ "github.com/jschaefer-io/aoc2021/day10"
    _ "github.com/jschaefer-io/aoc2021/day11"
    _ "github.com/jschaefer-io/aoc2021/day12"
    _ "github.com/jschaefer-io/aoc2021/day13"
    _ "github.com/jschaefer-io/aoc2021/day14"
    _ "github.com/jschaefer-io/aoc2021/day15"
    _ "github.com/jschaefer-io/aoc2021/day16"
    _ "github.com/jschaefer-io/aoc2021/day17"
    _ "github.com/jschaefer-io/aoc2021/day18"
    _ "github.com/jschaefer-io/aoc2021/day19"
    _ "github.com/jschaefer-io/aoc2021/day20"
    _ "github.com/jschaefer-io/aoc2021/day21"
    _ "github.com/jschaefer-io/aoc2021/day22"
    _ "github.com/jschaefer-io/aoc2021/day23"
    _ "github.com/jschaefer-io/aoc2021/day24"
    _ "github.com/jschaefer-io/aoc2021/day25"
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
