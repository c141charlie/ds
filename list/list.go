package list

//import "fmt"

type List struct {
    head_and_tail *Element 
    size int
}


func NewList() *List {
    elem := &Element{nil, nil, nil}
    elem.setNext(elem)
    elem.setPrev(elem)
    return &List{elem, 0}
}


func (l *List) Insert(pos int, val interface{}) {

    if pos < 0 || pos > l.size {
        return 
    }

    elem := newElement(val)

    elem.attachBefore(l.getElement(pos))

    l.size ++

}

func (l *List) getElement(pos int) *Element {
    cur := l.head_and_tail.getNext()

    for i := pos; i > 0; i-- {
        cur = cur.getNext()
    }

    return cur
}

func (l *List) Add(val interface{}) {
    l.Insert(l.size, val)
}

func (l *List) Get(pos int) interface{} {
    if pos < 0 || pos >= l.size {
        return nil
    }
    return l.getElement(pos).GetValue()
}

func (e *Element) GetValue() interface{} {
    return e.value
}

func (l *List) Set(pos int, val interface{}) {
    if pos < 0 || pos >= l.size {
        return
    }
    elem := l.getElement(pos)

    elem.setValue(val)
}


func (l *List) IndexOf(val interface{}) int {
    pos := 0

    for elem := l.head_and_tail.getNext();
        elem != l.head_and_tail;
        elem = elem.getNext() {

        if val == elem.value {
            return pos
        }

        pos ++
    }

    return -1
}

func (l *List) Contains(val interface{}) bool {
    return -1 != l.IndexOf(val)
}

func (l *List) Delete(pos int) interface{} {
    if pos < 0 || pos >= l.size {
        return nil
    }

    elem := l.getElement(pos)
    elem.detach()
    l.size --
    return elem.GetValue()
}

func (l *List) DeleteByVal(val interface{}) bool {
    for elem := l.head_and_tail.getNext();
        elem != l.head_and_tail;
        elem = elem.getNext() {

        if val == elem.GetValue() {
            elem.detach()
            l.size --
            return true
        }

    }
    return false
}

func (l *List) Size() int {
    return l.size
}

func (l *List) IsEmpty() bool {
    return l.size == 0
}

func (l *List) Clear() {


}

type Element struct { 
    value interface{}
    prev *Element
    next *Element
}

func newElement(val interface{}) *Element{
    return &Element{val, nil, nil}
}

func (e *Element) setValue(val interface{}) {
    e.value = val
}

func (e *Element) getValue() interface{} {
    return e.value
}

func (e *Element) getPrevious() *Element {
    return e.prev
}

func (e *Element) getNext() *Element {
    return e.next
}

func (e *Element) setPrev(elem *Element) {
    e.prev = elem
}

func (e *Element) setNext(elem *Element) {
    e.next = elem
}

func (e *Element) attachBefore(next *Element) {
    prev := next.getPrevious()
    e.prev = prev
    e.next = next
    next.setPrev(e)
    prev.setNext(e)

}

func (e *Element) detach() {
    e.prev.setNext(e.next)
    e.next.setPrev(e.prev)
}


type Iterator struct {
    list *List
    cur *Element
}

func NewIterator(list *List) *Iterator { 
    return &Iterator{list, list.head_and_tail}    
}

func (i *Iterator) First() *Element {
    i.cur = i.list.getElement(0)
    return i.cur
}

func (i *Iterator) Next() *Element {
    i.cur = i.cur.getNext()
    return i.cur
}

func (i *Iterator) Current() *Element {
    return i.cur
}

func (i *Iterator) IsDone() bool {
    if i.cur == i.list.head_and_tail {
        return true
    }
    return false
}

