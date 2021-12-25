package day25

type Position struct {
    X int
    Y int
}

type Cucumber struct {
    down     bool
    Position Position
}

func (c *Cucumber) NextPosition() Position {
    if c.down {
        return Position{c.Position.X, c.Position.Y + 1}
    } else {
        return Position{c.Position.X + 1, c.Position.Y}
    }
}
