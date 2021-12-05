package day05

type Line struct {
    From Coordinate
    To   Coordinate
}

func (l Line) Diagonal() bool {
    return l.From.X != l.To.X && l.From.Y != l.To.Y
}

type Coordinate struct {
    X int
    Y int
}
