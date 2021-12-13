package day13

import "math"

type Grid struct {
    width  int
    height int
    dots   [][]bool
}

func (g Grid) Fold(t int) Grid {
    upperLen := t
    lowerLen := g.height - t - 1
    newHeight := int(math.Max(float64(upperLen), float64(lowerLen)))
    newGrid := NewGrid(g.width, newHeight)
    for y := 0; y < newHeight; y++ {
        for x := 0; x < g.width; x++ {
            newGrid.dots[y][x] = newGrid.dots[y][x] || g.Value(x, upperLen-newHeight+y)
            newGrid.dots[y][x] = newGrid.dots[y][x] || g.Value(x, t+newHeight-y)
        }
    }
    return newGrid
}

func (g Grid) Value(x, y int) bool {
    if y < 0 || y >= g.height || x < 0 || x >= g.width {
        return false
    }
    return g.dots[y][x]
}

func (g Grid) Flip() Grid {
    newGrid := NewGrid(g.height, g.width)
    for y := 0; y < g.width; y++ {
        for x := 0; x < g.height; x++ {
            newGrid.dots[y][x] = g.dots[x][y]
        }
    }
    return newGrid
}

func NewGrid(width, height int) Grid {
    grid := Grid{
        width:  width,
        height: height,
        dots:   make([][]bool, height),
    }
    for y := 0; y < height; y++ {
        grid.dots[y] = make([]bool, width)
    }
    return grid
}
