package day05

import (
    "math"
    "strconv"
    "strings"
)

type Map struct {
    width   int
    height  int
    diagram [][]int
}

func NewMap(width, height int) Map {
    diagram := make([][]int, width)
    for i, _ := range diagram {
        diagram[i] = make([]int, height)
    }
    return Map{width, height, diagram}
}

func (m *Map) CountMin(minValue int) int {
    count := 0
    for y := 0; y < m.width; y++ {
        for x := 0; x < m.height; x++ {
            if m.diagram[y][x] >= minValue {
                count++
            }
        }
    }
    return count
}

func (m Map) DrawLine(line Line) {
    minPoint := line.From
    maxPoint := line.To
    if line.To.X <= line.From.X && line.To.X <= line.From.Y {
        minPoint = line.To
        maxPoint = line.From
    }

    xDelta := maxPoint.X - minPoint.X
    xDiv := int(math.Abs(float64(maxPoint.X) - float64(minPoint.X)))
    xOffset := 0
    if xDiv != 0 {
        xOffset = xDelta / xDiv
    }

    yDelta := maxPoint.Y - minPoint.Y
    yDiv := int(math.Abs(float64(maxPoint.Y) - float64(minPoint.Y)))
    yOffset := 0
    if yDiv != 0 {
        yOffset = yDelta / yDiv
    }

    x := minPoint.X
    y := minPoint.Y
    for x != maxPoint.X+xOffset || y != maxPoint.Y+yOffset {
        m.diagram[y][x]++
        x += xOffset
        y += yOffset
    }
}

func (m Map) String() string {
    var sb strings.Builder
    for y := 0; y < m.width; y++ {
        for x := 0; x < m.height; x++ {
            if m.diagram[y][x] == 0 {
                sb.WriteString(".")
            } else {
                sb.WriteString(strconv.Itoa(m.diagram[y][x]))
            }
        }
        sb.WriteString("\n")
    }
    return sb.String()
}
