package pri_queue

import "testing"

func TestNewPriorityQueue(t *testing.T) {
    pq := New()
    if pq.Size() != 0 {
        t.Errorf("Priority queue should be 0\n")
    }
}
