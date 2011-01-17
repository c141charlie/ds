package stack

import (
    "os"
    "container/list"
)

var EmptyStack = os.NewError("Empty list. Nothing to pop.")

type Stack struct {
    l *list.List 
}

func New() *Stack {
    return &Stack{list.New()}
}

func (s *Stack) Push(value interface{}) {
    s.l.PushFront(value)
}

func (s *Stack) Pop() (interface{}, os.Error) {
    if s.l.Front() == nil {
        return nil, EmptyStack
    }
    return s.l.Remove(s.l.Front()), nil
}

func (s *Stack) Peek() (interface{}, os.Error) {
    elem := s.l.Front()
    if elem == nil {
        return nil, EmptyStack
    }
    return elem.Value, nil
}

func (s *Stack) Clear() {
    s.l = nil
    s.l = list.New()
}

func (s *Stack) Size() int {
    return s.l.Len()
}

func (s *Stack) IsEmpty() bool {
    return s.Size() == 0
}
