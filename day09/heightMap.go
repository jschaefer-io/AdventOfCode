package day09

type Point struct {
    X     int
    Y     int
    Value int
}

type Heightmap struct {
    Width  int
    Height int
    Map    [][]Point
}

func (h *Heightmap) ResolveBasinFromLowPoint(lp Point) []Point {
    pointer := 0
    basin := []Point{lp}
    visited := map[Point]struct{}{lp: {}}
    for pointer != len(basin) {
        current := basin[pointer]
        for _, checkPoint := range h.AdjacentPoints(current) {
            if _, ok := visited[checkPoint]; ok {
                continue
            }
            if checkPoint.Value > current.Value && checkPoint.Value != 9 {
                basin = append(basin, checkPoint)
                visited[checkPoint] = struct{}{}
            }
        }
        pointer++
    }
    return basin
}

func (h *Heightmap) AdjacentPoints(p Point) []Point {
    points := make([]Point, 0)
    for xO := -1; xO <= 1; xO += 2 {
        cX := p.X + xO
        if cX < 0 || cX >= h.Width {
            continue
        }
        points = append(points, h.Map[p.Y][cX])
    }
    for yO := -1; yO <= 1; yO += 2 {
        cY := p.Y + yO
        if cY < 0 || cY >= h.Height {
            continue
        }
        points = append(points, h.Map[cY][p.X])
    }
    return points
}

func (h *Heightmap) PointIsLowPoint(x, y int) bool {
    value := h.Map[y][x]
    for _, check := range h.AdjacentPoints(value) {
        if check.Value <= value.Value {
            return false
        }
    }
    return true
}

func (h *Heightmap) FindLowPoints() []Point {
    list := make([]Point, 0)
    for y := 0; y < h.Height; y++ {
        for x := 0; x < h.Width; x++ {
            if h.PointIsLowPoint(x, y) {
                list = append(list, h.Map[y][x])
            }
        }
    }
    return list
}
