package day15

import (
    "container/heap"
    "fmt"
)

type QueuePosition struct {
    pos      Position
    priority int
    index    int
}

// PriorityQueue derived from https://pkg.go.dev/container/heap
type PriorityQueue []*QueuePosition

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*QueuePosition)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil
    item.index = -1
    *pq = old[0 : n-1]
    return item
}

func (pq *PriorityQueue) update(item *QueuePosition, priority int) {
    item.priority = priority
    heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) Print() {
    for pq.Len() > 0 {
        current := heap.Pop(pq).(*QueuePosition)
        fmt.Println(current.pos, current.priority)
    }
}
