package day09

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "sort"
    "strconv"
    "strings"
)

func Solve(data string, result *orchestration.Result) error {
    heights := make([][]Point, 0)
    for y, line := range strings.Split(data, "\n") {
        if len(line) == 0 {
            continue
        }
        col := make([]Point, 0)
        for x, sN := range strings.Split(line, "") {
            n, err := strconv.Atoi(sN)
            if err != nil {
                return err
            }
            col = append(col, Point{x, y, n})
        }
        heights = append(heights, col)
    }
    height := len(heights)
    width := len(heights[0])
    heightMap := Heightmap{
        Width:  width,
        Height: height,
        Map:    heights,
    }

    // A
    sum := 0
    lowPoints := heightMap.FindLowPoints()
    for _, lp := range lowPoints {
        sum += 1 + lp.Value
    }
    result.AddResult(strconv.Itoa(sum))

    // B
    sizes := make([]int, 0)
    for _, lp := range lowPoints {
        basins := heightMap.ResolveBasinFromLowPoint(lp)
        sizes = append(sizes, len(basins))
    }
    sort.Ints(sizes)
    sizesCount := len(sizes)
    prod := 1
    for i := 0; i < 3; i++ {
        prod *= sizes[sizesCount-1-i]
    }
    result.AddResult(strconv.Itoa(prod))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day09", orchestration.NewSolver(Solve))
}
