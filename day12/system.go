package day12


type System struct {
    caves map[string]*Cave
}

func NewSystem() System {
    return System{
        caves: make(map[string]*Cave),
    }
}

func (s *System) Names() []string {
    list := make([]string, 0)
    for key := range s.caves {
        list = append(list, key)
    }
    return list
}

func (s *System) AddCave(name string) *Cave {
    c, ok := s.caves[name]
    if !ok {
        cave := NewCave(name)
        c = &cave
        s.caves[name] = &cave
    }
    return c
}

func (s *System) FindPaths(start, end string, excludes visited, allowTwice string) [][]string {
    excludes.Visit(start)
    connections := make([][]string, 0)
    if start == end {
        return [][]string{{end}}
    }
    sCave := s.caves[start]
    for _, connection := range sCave.connections {
        target := s.caves[connection]
        visitedCount := excludes.Visited(connection)
        if !target.isBig {
            if target.name == allowTwice {
                if visitedCount > 1 {
                    continue
                }
            } else if visitedCount > 0 {
                continue
            }
        }
        for _, subConnection := range s.FindPaths(connection, end, excludes.Copy(), allowTwice) {
            connections = append(connections, append([]string{start}, subConnection...))
        }
    }
    return connections
}
