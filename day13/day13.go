package day12

import (
    "github.com/jschaefer-io/aoc2021/orchestration"
    "math"
    "strconv"
    "strings"
)

type Paper struct {
    grid Grid
}

type Grid struct {
    width  int
    height int
    dots   [][]bool
}
type Fold struct {
    Position int
    Axis     bool
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

func NewPaper(dots [][2]int) Paper {
    maxWidth := 0
    maxHeight := 0
    paper := Paper{}
    for _, p := range dots {
        if maxHeight < p[1] {
            maxHeight = p[1]
        }
        if maxWidth < p[0] {
            maxWidth = p[0]
        }
    }
    paper.grid = NewGrid(maxWidth+1, maxHeight+1)
    for _, p := range dots {
        paper.grid.dots[p[1]][p[0]] = true
    }
    return paper
}

func (p Paper) String() string {
    var sb strings.Builder
    for y := 0; y < p.grid.height; y++ {
        for x := 0; x < p.grid.width; x++ {
            if p.grid.dots[y][x] {
                sb.WriteString("#")
            } else {
                sb.WriteString(" ")
            }
        }
        sb.WriteString("\n")
    }
    return sb.String()
}

func (p *Paper) Fold(op Fold) Paper {
    var newGrid Grid
    if !op.Axis {
        newGrid = p.grid.Fold(op.Position)
    } else {
        newGrid = p.grid.Flip().Fold(op.Position).Flip()
    }
    return Paper{
        grid: newGrid,
    }
}

func (p *Paper) Points() int {
    counter := 0
    for y := 0; y < p.grid.height; y++ {
        for x := 0; x < p.grid.width; x++ {
            if p.grid.dots[y][x] {
                counter++
            }
        }
    }
    return counter
}

func Solve(data string, result *orchestration.Result) error {
    groups := strings.Split(data, "\n\n")
    dotList := make([][2]int, 0)
    for _, dots := range strings.Split(groups[0], "\n") {
        values := strings.Split(dots, ",")
        x, err := strconv.Atoi(values[0])
        if err != nil {
            return err
        }
        y, err := strconv.Atoi(values[1])
        if err != nil {
            return err
        }
        dotList = append(dotList, [2]int{x, y})
    }

    operations := make([]Fold, 0)
    for _, op := range strings.Split(groups[1], "\n") {
        if len(op) == 0 {
            continue
        }
        values := strings.Split(op, "=")
        n, err := strconv.Atoi(values[1])
        if err != nil {
            return err
        }
        operations = append(operations, Fold{
            Axis:     strings.Contains(values[0], "x"),
            Position: n,
        })
    }

    paper := NewPaper(dotList)
    firstFold := paper.Fold(operations[0])
    result.AddResult(strconv.Itoa(firstFold.Points()))

    for _, op := range operations {
        paper = paper.Fold(op)
    }
    result.AddResult("\n" + paper.String())
    return nil
}

func init() {
    orchestration.MainDispatcher.AddSolver("Day13", orchestration.NewSolver(Solve))
}
