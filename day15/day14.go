package day15

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "math"
    "strconv"
    "strings"
)

type Position struct {
    X    int
    Y    int
    Risk int
}

type Field struct {
    Width    int
    Height   int
    Position [][]Position
}

func (f *Field) Expand(amount uint) Field {
    mul := int(amount)
    newField := Field{
        Width:  f.Width * mul,
        Height: f.Height * mul,
    }
    newField.Position = make([][]Position, newField.Height)

    for y := 0; y < newField.Height; y++ {
        newField.Position[y] = make([]Position, newField.Height)
        for x := 0; x < newField.Width; x++ {
            t := ((f.Position[y%f.Height][x%f.Width].Risk - 1) + (y / f.Width) + (x / f.Height)) % 9
            newField.Position[y][x] = Position{
                X:    x,
                Y:    y,
                Risk: t + 1,
            }
        }
    }
    return newField
}

func getMinIndex(distances map[Position]int, positions map[Position]struct{}) Position {
    i := 0
    var min Position
    for v := range positions {
        if i == 0 || distances[v] < distances[min] {
            min = v
        }
        i++
    }
    return min
}

func (f *Field) Positions() ([]Position, map[Position]struct{}) {
    pos := make([]Position, f.Width*f.Height)
    index := make(map[Position]struct{})
    for y := 0; y < f.Height; y++ {
        for x := 0; x < f.Width; x++ {
            pos[y*f.Width+x] = f.Position[y][x]
            index[f.Position[y][x]] = struct{}{}
        }
    }
    return pos, index
}

func (f *Field) distanceList() map[Position]int {
    length := f.Width * f.Height
    list := make(map[Position]int, length)
    positions, _ := f.Positions()
    for _, v := range positions {
        list[v] = math.MaxInt64
    }
    return list
}

func (f *Field) neighbors(pos Position) []Position {
    neighbors := make([]Position, 0)
    var tX int
    var tY int
    for oX := -1; oX <= 1; oX += 2 {
        tX = pos.X + oX
        if tX < 0 || tX >= f.Width {
            continue
        }
        neighbors = append(neighbors, f.Position[pos.Y][tX])
    }
    for oY := -1; oY <= 1; oY += 2 {
        tY = pos.Y + oY
        if tY < 0 || tY >= f.Height {
            continue
        }
        neighbors = append(neighbors, f.Position[tY][pos.X])
    }
    return neighbors
}

type PathMap map[Position]Position

func (p PathMap) GetPath(target Position) []Position {
    path := make([]Position, 0)
    current := target
    for {
        pre, ok := p[current]
        if !ok {
            break
        }
        path = append(path, pre)
        current = pre
    }
    return path
}

func (f *Field) ShortestPaths(start Position) PathMap {
    _, leftover := f.Positions()
    distances := f.distanceList()
    predecessor := make(PathMap)
    distances[start] = 0

    for len(leftover) > 0 {
        current := getMinIndex(distances, leftover)
        delete(leftover, current)
        for _, neighbor := range f.neighbors(current) {
            if _, ok := leftover[neighbor]; !ok {
                continue
            }

            newDistance := distances[current] + neighbor.Risk
            if newDistance < distances[neighbor] {
                distances[neighbor] = newDistance
                predecessor[neighbor] = current
            }
        }
    }
    return predecessor
}

func (f Field) String() string {
    var sb strings.Builder
    for y := 0; y < f.Height; y++ {
        for x := 0; x < f.Width; x++ {
            sb.WriteString(strconv.Itoa(f.Position[y][x].Risk))
        }
        sb.WriteString("\n")
    }
    return sb.String()
}

func SolveForField(field Field, to, from Position) int {
    paths := field.ShortestPaths(from)

    sum := 0
    for _, v := range paths.GetPath(to) {
        sum += v.Risk
    }
    return sum
}

func Solve(data string, result *orchestration.Result) error {
    field := Field{
        Position: make([][]Position, 0),
    }
    for y, line := range strings.Split(data, "\n") {
        if len(line) == 0 {
            continue
        }
        row := make([]Position, 0)
        for x, risk := range strings.Split(line, "") {
            n, err := strconv.Atoi(risk)
            if err != nil {
                return err
            }
            row = append(row, Position{x, y, n})
        }
        field.Position = append(field.Position, row)
    }
    field.Height = len(field.Position)
    field.Width = len(field.Position[0])

    a := SolveForField(field, field.Position[0][0], field.Position[field.Height-1][field.Width-1])
    result.AddResult(strconv.Itoa(a))

    largeField := field.Expand(5)
    b := SolveForField(largeField, largeField.Position[0][0], largeField.Position[largeField.Height-1][largeField.Width-1])
    result.AddResult(strconv.Itoa(b))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day15", orchestration.NewSolver(Solve))
}
