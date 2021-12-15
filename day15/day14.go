package day15

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "math"
    "sort"
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
        list[v] = math.MaxInt32
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

    queue := PrioQueue{
        lookup: make(map[Position]struct{}),
        list:   make([]Position, 0),
        values: make(map[Position]int),
    }
    for p, _ := range leftover {
        queue.Add(p, distances[p])
    }
    queue.Sort()
    for queue.Count() > 0 {

        current := queue.Pop()
        delete(leftover, current)
        for _, neighbor := range f.neighbors(current) {
            if _, ok := leftover[neighbor]; !ok {
                continue
            }
            newDistance := distances[current] + neighbor.Risk
            if newDistance < distances[neighbor] {
                distances[neighbor] = newDistance
                queue.Update(neighbor, newDistance)
                predecessor[neighbor] = current
            }
        }
    }
    return predecessor
}

func (f Field) String(list []Position) string {
    pMap := make(map[Position]struct{})
    for _, v := range list {
        pMap[v] = struct{}{}
    }
    var sb strings.Builder
    for y := 0; y < f.Height; y++ {
        for x := 0; x < f.Width; x++ {
            if _, ok := pMap[f.Position[y][x]]; ok {
                sb.WriteString(".")
            } else {
                sb.WriteString(strconv.Itoa(f.Position[y][x].Risk))
            }
        }
        sb.WriteString("\n")
    }
    return sb.String()
}

func SolveForField(field Field, to, from Position) int {
    paths := field.ShortestPaths(from)

    //fmt.Println(field.String(paths.GetPath(to)))

    sum := 0
    for _, v := range paths.GetPath(to) {
        sum += v.Risk
    }
    return sum
}

type PrioQueue struct {
    count  int
    lookup map[Position]struct{}
    list   []Position
    values map[Position]int
}

func (q *PrioQueue) Count() int {
    return q.count
}

func (q *PrioQueue) Pop() Position {
    v := q.list[0]
    delete(q.lookup, v)
    delete(q.values, v)
    q.list = q.list[1:]
    q.count--
    return v
}

func (q *PrioQueue) Sort() {
    sort.Slice(q.list, func(i, j int) bool {
        return q.values[q.list[i]] < q.values[q.list[j]]
    })
}

func (q *PrioQueue) Update(pos Position, value int) {
    var index int
    for i, v := range q.list {
        if v == pos {
            index = i
            break
        }
    }
    var target int
    found := false
    newList := make([]Position, 0)
    if value == q.values[pos] {
        newList = q.list
    } else if value < q.values[pos] {
        for i := 0; i < index; i++ {
            target = i
            if q.values[q.list[i]] >= value {
                found = true
                break
            }
        }
        if found {
            newList = append(newList, q.list[:target]...)
            newList = append(newList, pos)
            newList = append(newList, q.list[target])
            newList = append(newList, q.list[target+1:index]...)
            newList = append(newList, q.list[index+1:]...)
        } else {
            newList = q.list
        }
    } else {
        for i := q.count - 1; i > index; i-- {
            target = i
            if q.values[q.list[i]] <= value {
                found = true
                break
            }
        }
        if found {
            newList = append(newList, q.list[:index]...)
            newList = append(newList, q.list[target])
            newList = append(newList, pos)
            newList = append(newList, q.list[index+1:target]...)
            newList = append(newList, q.list[target+1:]...)
        } else {
            newList = q.list
        }
    }
    q.list = newList
    q.values[pos] = value
}

func (q *PrioQueue) Add(pos Position, value int) {
    q.list = append(q.list, pos)
    q.values[pos] = value
    q.lookup[pos] = struct{}{}
    q.count++
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

    //largeField := field.Expand(5)
    //b := SolveForField(largeField, largeField.Position[0][0], largeField.Position[largeField.Height-1][largeField.Width-1])
    //result.AddResult(strconv.Itoa(b))

    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day15", orchestration.NewSolver(Solve))
}
