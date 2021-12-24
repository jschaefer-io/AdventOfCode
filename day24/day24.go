package day24

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "strconv"
    "strings"
)

type Group struct {
    start [3]int
    end   [3]int
}

func BuildStackPath(list [][3]int) []Group {
    path := make([]Group, 0)
    listLen := len(list)
    count := 0
    start := 0
    for start < listLen && count >= 0 {
        for i := start; i < listLen; i++ {
            if list[i][1] > 0 {
                count++
            } else {
                count--
            }
            if count == 0 {
                path = append(path, BuildStackPath(list[start+1:i])...)
                path = append(path, Group{
                    start: list[start],
                    end:   list[i],
                })
                start = i + 1
                break
            }
        }
    }
    return path
}

type MonadNumber [14]int

func (m MonadNumber) String() string {
    var sb strings.Builder
    for _, n := range m {
        sb.WriteString(strconv.Itoa(n))
    }
    return sb.String()
}

func Solve(data string, result *orchestration.Result) error {
    lines := strings.Split(data, "\n")
    groups := make([][3]int, 0)
    for i := 0; i < 14; i++ {
        a, err := strconv.Atoi(strings.Replace(lines[18*i+5], "add x ", "", -1))
        if err != nil {
            return err
        }
        b, err := strconv.Atoi(strings.Replace(lines[18*i+15], "add y ", "", -1))
        if err != nil {
            return err
        }
        groups = append(groups, [3]int{i, a, b})
    }

    maxNum := MonadNumber{}
    minNum := MonadNumber{}
    for _, k := range BuildStackPath(groups) {
        t := k.start[2] + k.end[1]
        if t > 0 {
            maxNum[k.start[0]] = 9 - t
            maxNum[k.end[0]] = 9

            minNum[k.start[0]] = 1
            minNum[k.end[0]] = t + 1
        } else {
            maxNum[k.start[0]] = 9
            maxNum[k.end[0]] = 9 + t

            minNum[k.start[0]] = (t * -1) + 1
            minNum[k.end[0]] = 1
        }
    }

    result.AddResult(maxNum.String())
    result.AddResult(minNum.String())
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day24", orchestration.NewSolver(Solve))
}
