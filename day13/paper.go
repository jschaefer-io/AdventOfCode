package day13

import "strings"

type Paper struct {
    grid Grid
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
