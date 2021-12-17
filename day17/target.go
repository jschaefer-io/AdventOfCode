package day17

type Target struct {
    X1 int
    X2 int
    Y1 int
    Y2 int
}

func (t *Target) CheckHit(p Point) bool {
    return p.X >= t.X1 && p.X <= t.X2 && p.Y >= t.Y1 && p.Y <= t.Y2
}

func (t *Target) MinX() int {
    minX := 0
    for {
        if t.X1 <= SumTo(minX) {
            break
        }
        minX++
    }
    return minX
}
