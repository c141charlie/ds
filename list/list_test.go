package list 

import "fmt"
import "testing"

const A = "A"
const B = "B"
const C = "C"

func TestInsertIntoEmptyList(t *testing.T) {

    l := NewList()

    if l.Size() != 0 {
        t.Errorf("l.Size() should be 0.\n")
    }

    if l.IsEmpty() != true {
        t.Errorf("l.IsEmpty() should be true.\n")
    }

    l.Insert(0, "Hello World")

    if l.Size() != 1 {
        t.Errorf("l.Size() should be 1.\n")
    }

    if l.IsEmpty() != false {
        t.Errorf("l.IsEmpty() should be false.\n")
    }

    if l.Get(0) != "Hello World" {
        t.Errorf("l.Get(0) should be \"Hello World.\"\n")
    }
}

func TestInsertBetweenElements(t *testing.T) {
    l := NewList()
    l.Insert(0, A)
    l.Insert(1, B)
    l.Insert(1, C)

    if l.Size() != 3 {
        t.Errorf("l.Size() should be 3\n")
    }

    if l.Get(0) != A {
        t.Errorf("l.Get(0) should be \"A\"\n")
    }

    if l.Get(1) != C {
        t.Errorf("l.Get(1) should be \"C\"\n")
        
    }

    if l.Get(2) != B {
        t.Errorf("l.Get(2) should be \"B\"\n")
    }
}

func TestInsertBeforeFirstElement(t *testing.T) {
    l := NewList()
    l.Insert(0, A)
    l.Insert(0, B)
    
    if l.Size() != 2 {
        t.Errorf("l.Size() should be \"2\"\n")
    }

    if l.Get(0) != B {
        t.Errorf("l.Get(0) should be \"B\"\n")
    }
    if l.Get(1) != A {
        t.Errorf("l.Get(1) should be \"A\"\n")
    }
}

func TestInsertAfterLastElement(t *testing.T) {
    l := NewList()
    l.Insert(0, A)
    l.Insert(1, B)
    
    if l.Size() != 2 {
        t.Errorf("l.Size() should be \"2\"\n")
    }

    if l.Get(0) != A {
        t.Errorf("l.Get(0) should be \"A\"\n")
    }
    if l.Get(1) != B {
        t.Errorf("l.Get(1) should be \"B\"\n")
    }
}

func TestInsertOutOfBounds(t *testing.T) {
    l := NewList()
    
    defer func() {
        r := recover()
        
        if r != nil {
            fmt.Println("Recovered in TestInsertOutOfBounds()")
        }
    } ()

    l.Insert(-1, A)

}

func TestDeleteOnlyElement(t *testing.T) {
    l := NewList()

    l.Insert(0, A)
    
    if l.Size() != 1 {
        t.Errorf("l.Size() should be \"1\"\n")
    }

    l.Delete(0)

    if l.Size() != 0 {
        t.Errorf("l.Size() should be \"0\"\n")
    }
}

func TestDeleteFirstElement(t *testing.T) {
    l := NewList()

    l.Insert(0, A)
    l.Insert(1, B)
    l.Insert(2, C)

    if l.Size() != 3 {
        t.Errorf("l.Size() should be \"3\"\n")
    }

    l.Delete(0)

    if l.Size() != 2 {
        t.Errorf("l.Size() should be \"2\"\n")
    }

    if l.Get(0) != B {
        t.Errorf("l.Get(0) should be \"B\"\n")
    }
    
    if l.Get(1) != C {
        t.Errorf("l.Get(1) should be \"C\"\n")
    }
}


func TestDeleteLastElement(t *testing.T) {
    l := NewList()

    l.Insert(0, A)
    l.Insert(1, B)
    l.Insert(2, C)

    if l.Size() != 3 {
        t.Errorf("l.Size() should be \"3\"\n")
    }

    l.Delete(2)

    if l.Size() != 2 {
        t.Errorf("l.Size() should be \"2\"\n")
    }

    if l.Get(0) != A {
        t.Errorf("l.Get(0) should be \"A\"\n")
    }
    
    if l.Get(1) != B {
        t.Errorf("l.Get(1) should be \"B\"\n")
    }
}


func TestDeleteMiddleElement(t *testing.T) {
    l := NewList()

    l.Insert(0, A)
    l.Insert(1, B)
    l.Insert(2, C)

    if l.Size() != 3 {
        t.Errorf("l.Size() should be \"3\"\n")
    }

    l.Delete(1)

    if l.Size() != 2 {
        t.Errorf("l.Size() should be \"2\"\n")
    }

    if l.Get(0) != A {
        t.Errorf("l.Get(0) should be \"A\"\n")
    }
    
    if l.Get(1) != C {
        t.Errorf("l.Get(1) should be \"C\"\n")
    }
}

func TestDeleteOutOfBounds(t *testing.T) {
    l := NewList()

    l.Insert(0, A)
    l.Insert(1, B)
    l.Insert(2, C)

    if l.Size() != 3 {
        t.Errorf("l.Size() should be \"3\"\n")
    }

    elem := l.Delete(-1)

    if l.Size() != 3{
        t.Errorf("l.Size() should be \"3\"\n")
    }

    if l.Get(0) != A {
        t.Errorf("l.Get(0) should be \"A\"\n")
    }
    
    if l.Get(1) != B {
        t.Errorf("l.Get(1) should be \"B\"\n")
    }
    
    if l.Get(2) != C {
        t.Errorf("l.Get(1) should be \"C\"\n")
    }

    if elem != nil {
        t.Errorf("l.Delete(-1) should be nil")
    }
}

func TestEmptyIteration(t *testing.T) {

    l := NewList()
    i := NewIterator(l)

    if i.IsDone() != true {
        t.Errorf("i.IsDone() should be \"true\"\n.")
    }
}

func TestIteration(t *testing.T) {
    l := NewList()
    i := NewIterator(l)
    l.Insert(0, A)
    l.Insert(1, B)
    l.Insert(2, C)
    
    cur := i.First()

    if cur != l.getElement(0) {
        t.Errorf("l.First() should be element 0.")
    }

    if i.IsDone() == true {
        t.Errorf("i.IsDone() should be false.\n")
    }

    cur = i.Next()

    if cur != l.getElement(1) {
        t.Errorf("l.First() should be element 1.")
    }

    if i.IsDone() == true {
        t.Errorf("i.IsDone() should be false.\n")
    }

    cur = i.Next()

    if cur != l.getElement(2) {
        t.Errorf("l.First() should be element 2.")
    }

    if i.IsDone() == true {
        t.Errorf("i.IsDone() should be false.\n")
    }

    cur = i.Next()

    if i.IsDone() != true {
        t.Errorf("i.IsDone() should be true.\n")
    }
}
