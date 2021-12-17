package day17

type Trajectory struct {
    x    int
    y    int
    maxX int
    maxY int
}

func NewTrajectory(x, y int) Trajectory {
    return Trajectory{
        x:    x,
        y:    y,
        maxX: SumTo(x),
        maxY: SumTo(y),
    }
}

func (t *Trajectory) GetPoint(x int) Point {
    tY := t.maxY - SumTo(t.y-x)
    tX := t.maxX
    if t.x-x >= 0 {
        tX -= SumTo(t.x - x)
    }
    return Point{tX, tY}
}
