package stack

import "testing"

const VALUE_A = "A"
const VALUE_B = "B"
const VALUE_C = "C"

func TestPushAndPop(t *testing.T) {
    s := New()
    s.Push(VALUE_A)
    s.Push(VALUE_B)
    s.Push(VALUE_C)

    if s.Size() != 3 {
        t.Errorf("s.Size() == 3")
    }

    if s.IsEmpty() != false {
        t.Errorf("s.IsEmpty() should be false\n")
    }

    //pop and test
    val, _ := s.Pop()
    if val != VALUE_C {
        t.Errorf("s.Pop() should be \"C\"\n")
    }
    if s.IsEmpty() != false {
        t.Errorf("s.IsEmpty() should be false\n")
    }

    //pop and test
    val, _ = s.Pop()
    if val != VALUE_B {
        t.Errorf("s.Pop() should be \"B\"\n")
    }   
    if s.IsEmpty() != false {
        t.Errorf("s.IsEmpty() should be false\n")
    }

    //pop and test
    val, _ = s.Pop()
    if val != VALUE_A {
        t.Errorf("s.Pop() should be \"A\"\n")
    }
    if s.IsEmpty() != true {
        t.Errorf("s.IsEmpty() shold be true\n")
    }
}

func TestCantPopFromEmptyStack(t *testing.T) {
    s := New()

    if s.Size() != 0 {
        t.Errorf("s.Size() should be 0\n")
    }

    _ , err := s.Pop()
    if err == nil {
        t.Errorf("s.Pop() should have thrown an error\n")
    }
}

func TestPeek(t *testing.T) {
    s := New()
    s.Push(VALUE_C)
    s.Push(VALUE_A)

    if s.Size() != 2 {
        t.Errorf("s.Size() should be 2\n")
    }
    val, _ := s.Peek()
    if val != VALUE_A {
        t.Errorf("s.Peek() should be \"A\", but is: %s\n", val)
    }

    if s.Size() != 2 {
        t.Errorf("s.Size() should be 2\n")
    }
}

func TestCantPeekIntoAnEmptyStack(t *testing.T) {
    s := New()

    if s.Size() != 0 {
        t.Errorf("s.Size() should be 0\n")
    }

    _, err := s.Peek()

    if err == nil {
        t.Errorf("err should not be nil.\n")
    }
}

func TestClear(t *testing.T) {
    s := New()
    s.Push(VALUE_A)
    s.Push(VALUE_B)
    s.Push(VALUE_C)
    
    if s.IsEmpty() != false {
        t.Errorf("s.IsEmpty() should be false\n")
    }

    if s.Size() != 3 {
        t.Errorf("s.Size() should be 3\n")
    }

    s.Clear()

    if s.IsEmpty() == false {
        t.Errorf("s.IsEmpty() should be true\n")
    }

    if s.Size() != 0 {
        t.Errorf("s.Size() should be 0\n")
    }

    _ , err := s.Pop()

    if err == nil {
        t.Errorf("err should not be nil\n")
    }
}
