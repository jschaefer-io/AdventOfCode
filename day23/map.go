package day23

import (
    "math"
    "strings"
)

type Map struct {
    Tiles     map[Position]Tile
    Amphipods map[Position]Amphipod
    Forbidden map[Position]struct{}
    Score     int
    Depth     int
    cache     PathCache
    progress  []string
}

func NewMap(length, depth int) Map {
    m := Map{
        Tiles:     make(map[Position]Tile),
        Amphipods: make(map[Position]Amphipod),
        Forbidden: make(map[Position]struct{}),
        Depth:     depth,
        cache: PathCache{
            cache: make(map[Position]map[Position][]Path),
        },
        progress: make([]string, 0),
    }
    var offset rune = 0
    for i := 0; i < length; i++ {
        m.Tiles[Position{i, 0}] = Tile{}
        if i%2 == 0 && i != 0 && i != length-1 {
            m.Forbidden[Position{i, 0}] = struct{}{}
            for y := 0; y < depth; y++ {
                m.Tiles[Position{i, 1 + y}] = Tile{
                    Room: 'A' + offset,
                }
            }
            offset++
        }
    }
    return m
}

func (m *Map) Finished() bool {
    for pos, tile := range m.Tiles {
        if tile.Room == 0 {
            continue
        }
        amp, ok := m.Amphipods[pos]
        if !ok {
            return false
        }
        if amp.Destination != tile.Room {
            return false
        }
    }
    return true
}

func (m *Map) AddAmphipod(pos Position, amp Amphipod) {
    m.Amphipods[pos] = amp
}

func (m *Map) GetPaths(current, last Position) []Path {
    if ok, lookup := m.cache.Lookup(current, last); ok {
        return lookup
    }
    paths := make([]Path, 0)
    paths = append(paths, Path{})
    for _, n := range m.Neighbors(current) {
        if n == last {
            continue
        }
        paths = append(paths, m.GetPaths(n, current)...)
    }
    for i, path := range paths {
        nPath := make(Path, 0)
        nPath = append(nPath, current)
        nPath = append(nPath, path...)
        paths[i] = nPath
    }
    m.cache.AddLookup(current, last, paths)
    return paths
}

func (m *Map) Neighbors(pos Position) []Position {
    neighbors := make([]Position, 0)
    for xO := -1; xO <= 1; xO += 2 {
        target := Position{pos.X + xO, pos.Y}
        if _, ok := m.Tiles[target]; ok {
            neighbors = append(neighbors, target)
        }
    }
    for yO := -1; yO <= 1; yO += 2 {
        target := Position{pos.X, pos.Y + yO}
        if _, ok := m.Tiles[target]; ok {
            neighbors = append(neighbors, target)
        }
    }
    return neighbors
}

func (m *Map) Copy() Map {
    newMap := Map{
        Tiles:     make(map[Position]Tile),
        Amphipods: make(map[Position]Amphipod),
        Forbidden: make(map[Position]struct{}),
        Depth:     m.Depth,
        Score:     m.Score,
        cache:     m.cache,
        progress:  append(m.progress, m.String()),
    }

    for p, t := range m.Tiles {
        newMap.Tiles[p] = Tile{
            Room: t.Room,
        }
    }

    for p, amp := range m.Amphipods {
        newMap.Amphipods[p] = Amphipod{
            Destination: amp.Destination,
        }
    }

    for p, s := range m.Forbidden {
        newMap.Forbidden[p] = s
    }

    return newMap
}

func (m *Map) Step() []Map {
    variations := make([]Map, 0)
    for pos, amp := range m.Amphipods {
        for _, path := range amp.GetPaths(pos, m) {
            newMap := m.Copy()
            nAmp := newMap.Amphipods[pos]
            delete(newMap.Amphipods, pos)
            newMap.Amphipods[path.destination] = nAmp
            newMap.Score += nAmp.Energy(path.count)
            variations = append(variations, newMap)
        }
    }
    return variations
}

func (m Map) String() string {
    minX := math.MaxInt32
    minY := math.MaxInt32
    maxX := 0
    maxY := 0
    for pos, _ := range m.Tiles {
        if pos.X < minX {
            minX = pos.X
        }
        if pos.X > maxX {
            maxX = pos.X
        }
        if pos.Y < minY {
            minY = pos.Y
        }
        if pos.Y > maxY {
            maxY = pos.Y
        }
    }
    var sb strings.Builder
    for y := minY - 1; y <= maxY+1; y++ {
        for x := minX - 1; x <= maxX+1; x++ {
            pos := Position{x, y}
            if _, ok := m.Tiles[pos]; ok {
                if amp, ok := m.Amphipods[pos]; ok {
                    sb.WriteString(string(amp.Destination))
                } else {
                    sb.WriteString(".")
                }
            } else {
                sb.WriteString("#")
            }
        }
        sb.WriteString("\n")
    }
    return sb.String()
}
