package day08

import "strings"

type Wiring struct {
    Count   int
    Letters map[rune]struct{}
}

func (w Wiring) String() string {
    var sb strings.Builder
    for k := range w.Letters {
        sb.WriteString(string(k))
    }
    return sb.String()
}

func (w Wiring) Segments() []rune {
    list := make([]rune, w.Count)
    i := 0
    for k, _ := range w.Letters {
        list[i] = k
        i++
    }
    return list
}

func NewWiring(input string) Wiring {
    count := len(input)
    list := make(map[rune]struct{})
    for _, r := range input {
        list[r] = struct{}{}
    }
    return Wiring{
        Count:   count,
        Letters: list,
    }
}
