package pri_queue

import "list"

type PriorityQueue struct {
    l *list.List
}

func New() *PriorityQueue {
    return &PriorityQueue{list.New()}
}

func (pq *PriorityQueue) Size() int {
    return pq.l.Length()
}
