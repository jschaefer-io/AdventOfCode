package day15

import (
    "container/heap"
    "math"
)

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

func (f *Field) ShortestPaths(start Position) PathMap {
    list, leftover := f.Positions()
    distances := f.distanceList()
    predecessor := make(PathMap)
    distances[start] = 0

    queue := make(PriorityQueue, len(list))
    items := make(map[Position]*QueuePosition)
    for i, v := range list {
        item := &QueuePosition{
            pos:      v,
            priority: distances[v],
            index:    i,
        }
        items[v] = item
        queue[i] = item
    }
    heap.Init(&queue)
    step := 0
    for queue.Len() > 0 {
        current := heap.Pop(&queue).(*QueuePosition)
        delete(leftover, current.pos)
        for _, neighbor := range f.neighbors(current.pos) {
            if _, ok := leftover[neighbor]; !ok {
                continue
            }
            newDistance := distances[current.pos] + neighbor.Risk
            if newDistance < distances[neighbor] {
                distances[neighbor] = newDistance
                queue.update(items[neighbor], newDistance)
                predecessor[neighbor] = current.pos
            }
        }
        step++
    }
    return predecessor
}
