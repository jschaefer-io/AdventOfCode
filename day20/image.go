package day20

import "strings"

type Image struct {
    mapping   string
    edgeValue bool
    pixels    map[int]map[int]bool
    minX      int
    minY      int
    maxX      int
    maxY      int
}

func NewImage(mapping string) Image {
    return Image{
        pixels:  make(map[int]map[int]bool),
        mapping: mapping,
    }
}

func (i *Image) GetBoundPixel(x, y int) bool {
    if x < i.minX || x > i.maxX || y < i.minY || y > i.maxY {
        return i.edgeValue
    }
    return i.GetPixel(x, y)
}

func (i *Image) GetPixel(x, y int) bool {
    if _, ok := i.pixels[x]; !ok {
        return false
    }
    return i.pixels[x][y]
}

func (i *Image) AddPixel(x, y int, value bool) {
    if _, ok := i.pixels[x]; !ok {
        i.pixels[x] = make(map[int]bool)
    }
    i.pixels[x][y] = value
    if x < i.minX {
        i.minX = x
    }
    if x > i.maxX {
        i.maxX = x
    }
    if y < i.minY {
        i.minY = y
    }
    if y > i.maxY {
        i.maxY = y
    }
}

func (i *Image) GetOperation(x, y int) Operation {
    pos := 8
    n := 0
    for yO := -1; yO <= 1; yO++ {
        for xO := -1; xO <= 1; xO++ {
            if i.GetBoundPixel(x+xO, y+yO) {
                n |= 1 << pos
            }
            pos--
        }
    }
    return Operation{x, y, n}
}

func (i *Image) GetOperations() []Operation {
    op := make([]Operation, 0)
    for y := i.minY - 1; y <= i.maxY+1; y++ {
        for x := i.minX - 1; x <= i.maxX+1; x++ {
            op = append(op, i.GetOperation(x, y))
        }
    }
    return op
}

func (i *Image) ApplyOperation(op Operation) {
    i.AddPixel(op.X, op.Y, i.mapping[op.Value] == '#')
}

func (i *Image) CountPixels() int {
    count := 0
    for y := i.minY; y <= i.maxY; y++ {
        for x := i.minX; x <= i.maxX; x++ {
            if i.GetPixel(x, y) {
                count++
            }
        }
    }
    return count
}

func (i *Image) Step() {
    operations := i.GetOperations()
    edgeOP := i.GetOperation(i.minX-4, i.minY-4)
    i.edgeValue = i.mapping[edgeOP.Value] == '#'
    for _, op := range operations {
        i.ApplyOperation(op)
    }
}

func (i Image) String() string {
    var sb strings.Builder
    for y := i.minY; y <= i.maxY; y++ {
        for x := i.minX; x <= i.maxX; x++ {
            if i.GetPixel(x, y) {
                sb.WriteString("#")
            } else {
                sb.WriteString(".")
            }
        }
        sb.WriteString("\n")
    }
    return sb.String()
}
