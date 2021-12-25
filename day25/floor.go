package day25

import "strings"

type Floor struct {
    width     int
    height    int
    cucumbers []*Cucumber
    field     map[Position]*Cucumber
}

func (f Floor) String() string {
    var sb strings.Builder
    for y := 0; y < f.height; y++ {
        for x := 0; x < f.width; x++ {
            c, ok := f.field[Position{x, y}]
            if ok {
                if c.down {
                    sb.WriteString("v")
                } else {
                    sb.WriteString(">")
                }
            } else {
                sb.WriteString(".")
            }
        }
        sb.WriteString("\n")
    }
    return sb.String()
}

func (f *Floor) updatePositions(down bool) bool {
    updates := make([][2]Position, 0)
    for _, cuc := range f.cucumbers {
        if cuc.down != down {
            continue
        }
        nextPos := f.FitPosition(cuc.NextPosition())
        if _, ok := f.field[nextPos]; !ok {
            updates = append(updates, [2]Position{cuc.Position, nextPos})
        }
    }
    for _, update := range updates {
        cuc := f.field[update[0]]
        cuc.Position = update[1]
        delete(f.field, update[0])
        f.field[update[1]] = cuc
    }
    return len(updates) > 0
}

func (f *Floor) Step() bool {
    east := f.updatePositions(false)
    south := f.updatePositions(true)
    return east || south
}

func (f *Floor) FitPosition(pos Position) Position {
    return Position{
        X: pos.X % f.width,
        Y: pos.Y % f.height,
    }
}
