package pri_queue

import "testing"
//import "fmt"

func TestNewPriorityQueue(t *testing.T) {
    pq := New()
    if pq.Size() != 0 {
        t.Errorf("Priority queue should be 0\n")
    }
}

func TestEnqueue(t *testing.T) {
    pq := New()
    pq.Enqueue(4)
    if pq.Size() != 1 {
        t.Errorf("pg.Size() whould be 1\n")
    }

    pq.Enqueue(5)
    if pq.Size() != 2 {
        t.Errorf("pg.Size() whould be 2\n")
    }

    pq.Enqueue(6)
    if pq.Size() != 3 {
        t.Errorf("pq.Size() should be 3\n")
    }
    
    pq.Dequeue()
}

