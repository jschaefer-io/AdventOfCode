package day12

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

func Solve(data string, result *orchestration.Result) error {
    system := NewSystem()

    for _, line := range strings.Split(data, "\n") {
        if len(line) == 0 {
            continue
        }
        parts := strings.Split(line, "-")
        start := system.AddCave(parts[0])
        start.AddConnection(parts[1])

        end := system.AddCave(parts[1])
        end.AddConnection(parts[0])
    }

    c := system.FindPaths("start", "end", make(visited), "")
    result.AddResult(strconv.Itoa(len(c)))

    variants := make(visited)
    for _, name := range system.Names() {
        if name == "start" || name == "end" {
            continue
        }
        paths := system.FindPaths("start", "end", make(visited), name)
        for _, path := range paths {
            pathSum := strings.Join(path, ",")
            variants.Visit(pathSum)
        }
    }
    result.AddResult(strconv.Itoa(len(variants)))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day12", orchestration.NewSolver(Solve))
}
