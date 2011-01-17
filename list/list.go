package list

import "fmt"

type List struct {
	len int
	first *Element
	last *Element
}

type Element struct {
	Value interface{}
	prev *Element
	next *Element
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

func New() *List{
	return &List{0, nil, nil}
}

func (l *List) Length() int{
	return l.len
}

func (l *List) First() *Element {
	return l.First()
}

func (l *List) Last() *Element {
	return l.Last()
}

func (l *List) IsEmpty() bool {
	if l.len > 0 {
		return false
	}
	return true
}

func (l *List) Add(value interface{}) {
	l.Insert(l.len, value)
}

func (l *List) Set(position int, value interface{}) {
	e := l.Get(position)
	e.Value = value
	
}

func(l *List) Delete(position int) {
	e:= l.Get(position)
	l.delete(e)
}

func (l *List) DeleteByVal(value interface{}) {
	i := 0
	for e:= l.first; e != nil; e = e.next {
		if e.Value == value {
			l.delete(e)
			break
		}
		i++
	}
}

func (l *List) delete(e *Element) {
	if e != nil {
		if l.first == e && l.last == e {
			l.first, l.last = nil, nil
			l.len --
			return
		}
		
		if e == l.first {
			l.first = e.next
			e.next.prev = nil
			l.len--
			return
		}
		
		if e == l.last {
			l.last = e.prev
			e.prev.next = nil
			l.len--
			return
		}
		e.prev.next = e.next
		e.next.prev = e.prev
		l.len--
	}
}

func (l *List) GetByVal(value interface{}) *Element {
	for e:= l.first; e != nil; e = e.next {
		if e.Value == value {
			return e
		}
	}
	return nil
}

func (l *List) Insert(position int, value interface{}) {
	if position < 0 || position > l.len {
		panic (fmt.Sprintf("Insert(position: %d, value: %v) position is out-of-bounds.", position, value))
	}
	
	e := &Element{value, nil, nil}
	
	if l.last == nil {
		l.first, l.last = e, e
		l.len = 1
		return
	}
	
	if position == 0 {
		l.insertBefore(e, l.first)
		return
	}
	
	if position == l.len {
		l.insertAfter(l.last, e)
		return
	}
	
	mark := l.getElement(position)
	l.insertBefore(e, mark)
}

func (l *List) Get(position int) *Element {
	return l.getElement(position)
}


func (l *List) insertAfter(mark *Element, e *Element) {
	if mark.next == nil {
		l.last = e
	} else {
		mark.next.prev = e
	}
	e.next = mark.next
	mark.next = e
	e.prev = mark
	l.len ++
}

func (l *List) insertBefore(e *Element, mark *Element) {
	if mark.prev == nil {
		l.first = e
	} else {
		mark.prev.next = e
	}
	e.prev = mark.prev
	mark.prev = e
	e.next = mark
	l.len ++
}


func(l *List) getElement(position int) *Element {
	if position < 0 || position >= l.len {
		panic (fmt.Sprintf("getElement(%d) is out-of-bounds.", position))
	}
	
	current := l.first
	
	for i:=0; i <=position; i++ {
		if i == position {
			return current
		}
		current = current.next
	}
	return nil
}


