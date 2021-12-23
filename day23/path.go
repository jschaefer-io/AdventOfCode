package day23

type Path []Position

type AmphipodPath struct {
    destination Position
    count       int
}

type PathCache struct {
    cache map[Position]map[Position][]Path
}

func (p *PathCache) Lookup(a, b Position) (bool, []Path) {
    if _, ok := p.cache[a]; !ok {
        return false, nil
    }
    path, ok := p.cache[a][b]
    return ok, path
}

func (p *PathCache) AddLookup(a, b Position, paths []Path) {
    if _, ok := p.cache[a]; !ok {
        p.cache[a] = make(map[Position][]Path)
    }
    p.cache[a][b] = paths
}
