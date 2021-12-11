package day11

type Queue struct {
    list  [][2]int
    count int
}

func (q *Queue) Length() int {
    return q.count
}

func (q *Queue) Add(v [2]int) {
    q.list = append(q.list, v)
    q.count++
}

func (q *Queue) Pop() [2]int {
    q.count--
    v := q.list[0]
    q.list = q.list[1:]
    return v
}
