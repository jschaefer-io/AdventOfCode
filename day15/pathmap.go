package day15

type PathMap map[Position]Position

func (p PathMap) GetPath(target Position) []Position {
    path := make([]Position, 0)
    current := target
    for {
        pre, ok := p[current]
        if !ok {
            break
        }
        path = append(path, pre)
        current = pre
    }
    return path
}

