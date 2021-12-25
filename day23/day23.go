package day23

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Position struct {
    X int
    Y int
}

type Tile struct {
    Room rune
}

func FromString(str string) Map {
    depth := 2
    if len(str) == 90 {
        depth = 4
    }
    m := NewMap(11, depth)
    for y, line := range strings.Split(str, "\n") {
        for x, c := range strings.Split(line, "") {
            switch c[0] {
            case '#':
            case '.':
            case ' ':
                continue
            default:
                m.AddAmphipod(Position{x - 1, y - 1}, Amphipod{rune(c[0])})
            }
        }
    }
    return m
}

func FindMin(m Map) int {
    mem := make(map[string]int)
    mem[m.String()] = 0
    var winner *Map = nil
    activeMaps := []Map{m}
    iterations := 0
    for len(activeMaps) > 0 {
        newList := make([]Map, 0)
        for _, current := range activeMaps {
            for _, nMap := range current.Step() {
                checksum := nMap.String()
                if v, ok := mem[checksum]; ok {
                    if nMap.Score >= v {
                        continue
                    }
                }
                if winner != nil && winner.Score < nMap.Score {
                    continue
                }

                if nMap.Finished() {
                    c := nMap
                    winner = &c
                } else {
                    newList = append(newList, nMap)
                    mem[checksum] = nMap.Score
                }
            }
        }
        activeMaps = newList
        iterations++
    }
    return winner.Score
}

func Solve(data string, result *orchestration.Result) error {
    a := FindMin(FromString(data))
    result.AddResult(strconv.Itoa(a))

    bInput := make([]string, 0)
    lineRaw := strings.Split(data, "\n")
    bInput = append(bInput, lineRaw[:3]...)
    bInput = append(bInput, []string{"  #D#C#B#A#", "  #D#B#A#C#"}...)
    bInput = append(bInput, lineRaw[3:]...)
    b := FindMin(FromString(strings.Join(bInput, "\n")))
    result.AddResult(strconv.Itoa(b))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day23", orchestration.NewSolver(Solve))
}
