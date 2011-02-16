package pri_queue

import "list"
import "fmt"

type PriorityQueue struct {
    l *list.List
}

func New() *PriorityQueue {
    return &PriorityQueue{list.New()}
}

func (pq *PriorityQueue) Size() int {
    return pq.l.Length()
}

func (pq *PriorityQueue) Enqueue(val interface{}) {
    pq.l.Add(val)
    pq.swim(pq.l.Length()-1)
}

func (pq *PriorityQueue) swim(index int) {
    if index == 0 {
        return
    }
    parent := (index -1)/2

    if pq.l.Get(index).Value.(int) > pq.l.Get(parent).Value.(int) {
        pq.swap(index, parent)
        pq.swim(parent)
    }
}

func (pq *PriorityQueue) swap(index1, index2 int) {
    temp := pq.l.Get(index1).Value.(int)
    fmt.Printf("temp = %d, index1 = %d, index2 = %d\n", temp, index1, index2)
    pq.l.Set(index1, pq.l.Get(index2).Value.(int))
    pq.l.Set(index2, temp)
}

func (pq *PriorityQueue) Dequeue() interface{} {
    if pq.IsEmpty() {
        return nil
    }

    result := pq.l.Get(0).Value.(int)
    if pq.l.Length() > 1 {
        pq.l.Set(0, pq.l.Get(pq.l.Length()-1).Value.(int))
        pq.sink(0)
    }
    pq.l.Delete(pq.l.Length()-1)

    return result
}

func (pq *PriorityQueue) sink(index int) {
    left := index * 2 + 1
    right := index * 2 + 2

    if left >= pq.l.Length() {
        return
    }

    largest_child := left
    if right < pq.l.Length() {
        if pq.l.Get(left).Value.(int) < pq.l.Get(right).Value.(int) {
            largest_child = right
        }
    }
    fmt.Println(index)
    fmt.Println(pq.l.Get(index))
    fmt.Println(pq.l.Get(largest_child))

    if pq.l.Get(index).Value.(int) < pq.l.Get(largest_child).Value.(int) {
        pq.swap(index, largest_child)
        pq.sink(largest_child)
    }
}

func (pq *PriorityQueue) Clear() {
    pq.l.Clear()
}

func (pq *PriorityQueue) IsEmpty() bool {
    return pq.l.IsEmpty()
}

func (pq *PriorityQueue) Peak(index int) {
    fmt.Println(pq.l.Get(index))
}


