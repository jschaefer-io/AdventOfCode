package day23

type Amphipod struct {
    Destination rune
}

func (amp *Amphipod) GetPaths(current Position, m *Map) []AmphipodPath {
    paths := make([]AmphipodPath, 0)
    neighbors := m.Neighbors(current)
    for _, n := range neighbors {
        for _, path := range m.GetPaths(n, current) {
            if amp.ValidatePath(current, path, m) {
                count := len(path)
                paths = append(paths, AmphipodPath{
                    destination: path[count-1],
                    count:       count,
                })
            }
        }
    }
    return paths
}

func (amp *Amphipod) Energy(count int) int {
    switch amp.Destination {
    case 'A':
        return count
    case 'B':
        return count * 10
    case 'C':
        return count * 100
    case 'D':
        return count * 1000
    }
    panic("undefined amphipod type")
}

func (amp *Amphipod) ValidatePath(current Position, p Path, m *Map) bool {
    destination := p[len(p)-1]
    destinationTile := m.Tiles[destination]
    currentTile := m.Tiles[current]

    if _, ok := m.Forbidden[destination]; ok {
        return false
    }

    if _, ok := m.Amphipods[destination]; ok {
        return false
    }

    if destinationTile.Room != 0 && destinationTile.Room != amp.Destination {
        return false
    }

    if currentTile.Room == 0 && destinationTile.Room != amp.Destination {
        return false
    }

    if currentTile.Room == amp.Destination {
        if current.Y == m.Depth {
            return false
        }

        done := true
        for y := current.Y; y <= m.Depth; y++ {
            nPos := Position{current.X, y}
            nAmp, _ := m.Amphipods[nPos]
            if nAmp.Destination != amp.Destination {
                done = false
            }
        }
        if done {
            return false
        }
    }

    if destinationTile.Room == amp.Destination {
        for y := destination.Y + 1; y <= m.Depth; y++ {
            nPos := Position{destination.X, y}
            nAmp, ok := m.Amphipods[nPos]
            if !ok {
                return false
            }
            if nAmp.Destination != destinationTile.Room {
                return false
            }
        }
    }

    // cant cross other amp
    for _, pos := range p {
        if _, ok := m.Amphipods[pos]; ok {
            return false
        }
    }

    return true
}
