package day22

import (
    "fmt"
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
    list []Boundary
}

func (g *Grid) ApplyOperation(op Operation) {
    newList := make([]Boundary, 0)
    for _, box := range g.list {
        intersect, intersectBox := op.Boundary.Intersect(box)
        if !intersect {
            newList = append(newList, box)
        } else {
            newList = append(newList, box.Sub(intersectBox)...)
        }
    }
    if op.Switch {
        newList = append(newList, op.Boundary)
    }
    g.list = newList
}

func (g *Grid) CountActive(constraint *Boundary) int {
    sum := 0
    if constraint != nil {
        for _, box := range g.list {
            if ok, cBox := constraint.Intersect(box); ok {
                sum += cBox.Score()
            }
        }
    } else {
        for _, box := range g.list {
            sum += box.Score()
        }
    }
    return sum
}

type Boundary struct {
    X1 int
    X2 int
    Y1 int
    Y2 int
    Z1 int
    Z2 int
}

func (b *Boundary) Score() int {
    return (b.X2 - b.X1 + 1) * (b.Y2 - b.Y1 + 1) * (b.Z2 - b.Z1 + 1)
}

func (b Boundary) intersectRange(a1, a2, b1, b2 int) (bool, int, int) {
    found := false
    var intersect1 int
    var intersect2 int
    for x := a1; x <= a2; x++ {
        if x >= b1 && x <= b2 {
            if !found {
                found = true
                intersect1 = x
                intersect2 = x
            } else {
                intersect2 = x
            }
        }
    }
    return found, intersect1, intersect2
}

func (b Boundary) Intersect(a Boundary) (bool, Boundary) {
    x, x1, x2 := b.intersectRange(b.X1, b.X2, a.X1, a.X2)
    y, y1, y2 := b.intersectRange(b.Y1, b.Y2, a.Y1, a.Y2)
    z, z1, z2 := b.intersectRange(b.Z1, b.Z2, a.Z1, a.Z2)
    return x && y && z, Boundary{x1, x2, y1, y2, z1, z2}
}

func (b Boundary) Valid() bool {
    return b.X1 <= b.X2 && b.Y1 <= b.Y2 && b.Z1 <= b.Z2
}

func (b Boundary) Sub(a Boundary) []Boundary {
    list := make([]Boundary, 0)
    partials := []Boundary{
        {b.X1, b.X2, b.Y1, a.Y1 - 1, a.Z1, a.Z2},
        {b.X1, b.X2, a.Y2 + 1, b.Y2, a.Z1, a.Z2},
        {b.X1, a.X1 - 1, a.Y1, a.Y2, a.Z1, a.Z2},
        {a.X2 + 1, b.X2, a.Y1, a.Y2, a.Z1, a.Z2},
        {b.X1, b.X2, b.Y1, b.Y2, b.Z1, a.Z1 - 1},
        {b.X1, b.X2, b.Y1, b.Y2, a.Z2 + 1, b.Z2},
    }
    for _, partial := range partials {
        if partial.Valid() {
            list = append(list, partial)
        }
    }

    return list
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

    g := Grid{list: make([]Boundary, 0)}
    for i, operation := range operations {
        g.ApplyOperation(operation)
    }

    // a
    result.AddResult(strconv.Itoa(g.CountActive(&Boundary{-50, 50, -50, 50, -50, 50})))

    // b
    result.AddResult(strconv.Itoa(g.CountActive(nil)))
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day22", orchestration.NewSolver(Solve))
}
