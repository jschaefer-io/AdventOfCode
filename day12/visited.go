package day12

type visited map[string]int

func (v *visited) Copy() visited {
    newV := make(visited)
    for k, v := range *v {
        newV[k] = v
    }
    return newV
}

func (v *visited) Visit(i string) {
    (*v)[i]++
}

func (v *visited) Visited(i string) int {
    return (*v)[i]
}
