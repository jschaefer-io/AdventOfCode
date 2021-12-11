package day11

import (
    "strconv"
    "strings"
)

type Swarm struct {
    grid [10][10]int
}

func NewSwarm(data string) (Swarm, error) {
    s := Swarm{}
    for y, line := range strings.Split(data, "\n") {
        if y >= 10 {
            continue
        }
        for x, v := range strings.Split(line, "") {
            n, err := strconv.Atoi(v)
            if err != nil {
                return s, err
            }
            s.grid[y][x] = n
        }
    }
    return s, nil
}

func (s *Swarm) Step() int {
    flashed := make(map[int]struct{})
    flashQueue := Queue{}
    for y := 0; y < 10; y++ {
        for x := 0; x < 10; x++ {
            s.grid[y][x]++
            if s.grid[y][x] > 9 {
                flashQueue.Add([2]int{y, x})
            }
        }
    }
    flashes := 0
    for flashQueue.Length() > 0 {
        v := flashQueue.Pop()
        check := 10*v[0] + v[1]
        if _, ok := flashed[check]; ok {
            continue
        }
        flashed[check] = struct{}{}
        flashes++
        s.grid[v[0]][v[1]] = 0
        for oX := -1; oX <= 1; oX++ {
            for oY := -1; oY <= 1; oY++ {
                tX := v[1] + oX
                tY := v[0] + oY
                if oX == 0 && oY == 0 || tX < 0 || tX >= 10 || tY < 0 || tY >= 10 {
                    continue
                }
                if _, ok := flashed[10*tY+tX]; ok {
                    continue
                }
                s.grid[tY][tX]++
                if s.grid[tY][tX] > 9 {
                    flashQueue.Add([2]int{tY, tX})
                }
            }
        }
    }
    return flashes
}

func (s Swarm) String() string {
    var sb strings.Builder
    for y := 0; y < 10; y++ {
        for x := 0; x < 10; x++ {
            sb.WriteString(strconv.Itoa(s.grid[y][x]))
        }
        sb.WriteString("\n")
    }
    return sb.String()
}
