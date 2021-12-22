package day22

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "regexp"
    "strconv"
    "strings"
)

type Cube bool

type Coordinate struct {
    X int
    Y int
    Z int
}

type Grid struct {
    field map[Coordinate]Cube
}

func (g *Grid) Set(value bool, pos Coordinate) {
    g.field[pos] = Cube(value)
}

func (g *Grid) Get(pos Coordinate) Cube {
    return g.field[pos]
}

func (g *Grid) ApplyOperation(op Operation, boundary *Boundary) {
    for x := op.X1; x <= op.X2; x++ {
        if boundary != nil && (x < boundary.X1 || x > boundary.X2) {
            continue
        }
        for y := op.Y1; y <= op.Y2; y++ {
            if boundary != nil && (y < boundary.Y1 || y > boundary.Y2) {
                continue
            }
            for z := op.Z1; z <= op.Z2; z++ {
                if boundary != nil && (z < boundary.Z1 || z > boundary.Z2) {
                    continue
                }
                g.Set(op.Switch, Coordinate{x, y, z})
            }
        }
    }
}

func (g *Grid) CountActive() int {
    active := 0
    for _, s := range g.field {
        if s {
            active++
        }
    }
    return active
}

type Boundary struct {
    X1 int
    X2 int
    Y1 int
    Y2 int
    Z1 int
    Z2 int
}

type Operation struct {
    Switch bool
    Boundary
}

func Solve(data string, result *orchestration.Result) error {
    exp := regexp.MustCompile("(?:|-)\\d+")
    operations := make([]Operation, 0)
    for _, line := range strings.Split(data, "\n") {
        if len(line) == 0 {
            continue
        }
        group := strings.Split(line, " ")
        coords := exp.FindAllStringSubmatch(group[1], -1)
        x1, err := strconv.Atoi(coords[0][0])
        if err != nil {
            return err
        }
        x2, err := strconv.Atoi(coords[1][0])
        if err != nil {
            return err
        }
        y1, err := strconv.Atoi(coords[2][0])
        if err != nil {
            return err
        }
        y2, err := strconv.Atoi(coords[3][0])
        if err != nil {
            return err
        }
        z1, err := strconv.Atoi(coords[4][0])
        if err != nil {
            return err
        }
        z2, err := strconv.Atoi(coords[5][0])
        if err != nil {
            return err
        }
        operations = append(operations, Operation{group[0] == "on", Boundary{x1, x2, y1, y2, z1, z2}})
    }

    // a
    g := Grid{field: make(map[Coordinate]Cube)}
    aBound := Boundary{-50, 50, -50, 50, -50, 50}
    for _, operation := range operations {
        g.ApplyOperation(operation, &aBound)
    }
    result.AddResult(strconv.Itoa(g.CountActive()))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day22", orchestration.NewSolver(Solve))
}
