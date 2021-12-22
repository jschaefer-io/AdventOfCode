package day22

type Area struct {
    X1 int
    X2 int
    Y1 int
    Y2 int
    Z1 int
    Z2 int
}

func (b *Area) Score() int {
    return (b.X2 - b.X1 + 1) * (b.Y2 - b.Y1 + 1) * (b.Z2 - b.Z1 + 1)
}

func (b Area) intersectRange(a1, a2, b1, b2 int) (bool, int, int) {
    found := false
    var intersect1 int
    var intersect2 int
    for x := a1; x <= a2; x++ {
        if x >= b1 && x <= b2 {
            if !found {
                found = true
                intersect1 = x
                intersect2 = x
            } else {
                intersect2 = x
            }
        }
    }
    return found, intersect1, intersect2
}

func (b Area) Intersect(a Area) (bool, Area) {
    if b.X2 < a.X1 || b.Y2 < a.Y1 || b.Z2 < a.Z1 || b.X1 > a.X2 || b.Y1 > a.Y2 || b.Z1 > a.Z2 {
        return false, Area{}
    }
    x, x1, x2 := b.intersectRange(b.X1, b.X2, a.X1, a.X2)
    y, y1, y2 := b.intersectRange(b.Y1, b.Y2, a.Y1, a.Y2)
    z, z1, z2 := b.intersectRange(b.Z1, b.Z2, a.Z1, a.Z2)
    return x && y && z, Area{x1, x2, y1, y2, z1, z2}
}

func (b Area) Valid() bool {
    return b.X1 <= b.X2 && b.Y1 <= b.Y2 && b.Z1 <= b.Z2
}

func (b Area) Sub(a Area) []Area {
    list := make([]Area, 0)
    partials := []Area{
        {b.X1, b.X2, b.Y1, a.Y1 - 1, a.Z1, a.Z2},
        {b.X1, b.X2, a.Y2 + 1, b.Y2, a.Z1, a.Z2},
        {b.X1, a.X1 - 1, a.Y1, a.Y2, a.Z1, a.Z2},
        {a.X2 + 1, b.X2, a.Y1, a.Y2, a.Z1, a.Z2},
        {b.X1, b.X2, b.Y1, b.Y2, b.Z1, a.Z1 - 1},
        {b.X1, b.X2, b.Y1, b.Y2, a.Z2 + 1, b.Z2},
    }
    for _, partial := range partials {
        if partial.Valid() {
            list = append(list, partial)
        }
    }

    return list
}
