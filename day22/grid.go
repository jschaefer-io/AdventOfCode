package day22

type Grid struct {
    list []Area
}

type Operation struct {
    Switch bool
    Area
}

func (g *Grid) ApplyOperation(op Operation) {
    newList := make([]Area, 0)
    for _, box := range g.list {
        intersect, intersectBox := op.Area.Intersect(box)
        if !intersect {
            newList = append(newList, box)
        } else {
            newList = append(newList, box.Sub(intersectBox)...)
        }
    }
    if op.Switch {
        newList = append(newList, op.Area)
    }
    g.list = newList
}

func (g *Grid) CountActive(constraint *Area) int {
    sum := 0
    if constraint != nil {
        for _, box := range g.list {
            if ok, cBox := constraint.Intersect(box); ok {
                sum += cBox.Score()
            }
        }
    } else {
        for _, box := range g.list {
            sum += box.Score()
        }
    }
    return sum
}
