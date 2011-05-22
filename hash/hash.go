package hash

import "list"

const default_capacity = 3
const load_factor = 75

type Map struct {
    buckets *list.List
    iterator *list.Iterator
}

func NewMap() *Map {
    l := list.NewList()
    i := list.NewIterator(l)
    return &Map{l, i}
}

func (m *Map) First() *KeyValue {
    return m.iterator.First().GetValue().(*KeyValue)
}

func (m *Map) Next() {
    m.iterator.Next()
}

func (m *Map) Current() *KeyValue {
    return m.iterator.Current().GetValue().(*KeyValue)
}

func (m *Map) IsDone() bool {
    return m.iterator.IsDone()
}

func (m *Map) Set(key interface{}, value interface{}) {
    kv := m.getKeyValueFor(key)
    
    if kv != nil {
        kv.Value = value
        return
    }
    
    key_val_pair := NewKeyValue(key, value)
    
    m.buckets.Add(key_val_pair)
    

}

func (m *Map) Contains(key interface{}) bool {
    i := list.NewIterator(m.buckets)

    for i.First(); !i.IsDone(); i.Next() {
        c := i.Current().GetValue().(*KeyValue)

        if c.Equals(key) {
            return true
        }
    }
    return false
}

func (m *Map) Get(key interface{}) interface{} {
    i := list.NewIterator(m.buckets)

    for i.First(); !i.IsDone(); i.Next() {
        c := i.Current().GetValue().(*KeyValue)

        if c.Equals(key) {
            return c.Value
        }
    }
    return nil
}

func (m *Map) Delete(key interface{}) bool {
   
    i := list.NewIterator(m.buckets)

    for i.First(); !i.IsDone(); i.Next() {
        c := i.Current().GetValue().(*KeyValue)

        if c.Equals(key) {
            return m.buckets.DeleteByVal(c)
        }
    }
    return false
}

func (m *Map) getKeyValueFor(key interface{}) *KeyValue{

    i := list.NewIterator(m.buckets)
    
    for i.First(); !i.IsDone(); i.Next() {
        c := i.Current().GetValue().(*KeyValue)
        
        if c.Equals(key) {
             return i.Current().GetValue().(*KeyValue)
        }
    }

    return nil
}

func (m *Map) Size() int {
    return m.buckets.Size()
}

func (m *Map) Clear() {
    m.buckets.Clear()
}

type KeyValue struct {
    Key interface{}
    Value interface{}
}

func NewKeyValue(key interface{}, value interface{}) *KeyValue{
    return  &KeyValue{key, value}
}

func (kv *KeyValue) Equals(arg interface{}) bool {
    if kv.Key == arg {
        return true
    }
    return false
}


type Hash func() int

type Hashable interface {
    Hash() int
}

type Hashtable struct {
    maps []*Map
    size int
}

func NewHashtable() *Hashtable {
    return &Hashtable{make([]*Map, default_capacity), 0}
}


func (h *Hashtable) Insert(key Hashable, value interface{}) {
    key_vals := h.getBuckets(key)
    key_vals.Set(key, value)
    h.size ++
    h.maintainLoad()
}

func (h *Hashtable) Get(key Hashable) interface{} {
    key_vals := h.getBuckets(key)
    return key_vals.Get(key)
}

func (h *Hashtable) Delete(key Hashable) {
    key_vals := h.getBuckets(key)
    
    if key_vals.Delete(key) {
        h.size --        
    }
}

func (h *Hashtable) Contains(key Hashable) bool {

    key_vals := h.getBuckets(key)

    return key_vals.Contains(key)

}

func (h *Hashtable) getBuckets(key Hashable) *Map{
    i := abs(key.Hash() % len(h.maps))
    key_vals := h.maps[i]

    if key_vals == nil {
        key_vals = NewMap()
        h.maps[i] = key_vals
    }

    return key_vals
}


func (h *Hashtable) Size() int {
    return h.size
}

func (h *Hashtable) numSlots() int {
    return len(h.maps)
}

func (h *Hashtable) maintainLoad() {
    if h.loadFactorExceeded() {
        h.resize()
    }
}

func (h *Hashtable) loadFactorExceeded() bool {
    return h.size > len(h.maps) * load_factor/100
}

func (h *Hashtable) resize() {
    maps_copy := h.maps
    np := nextPrime(2 * len(h.maps))
    replacement := make([]*Map, np)
    h.maps = replacement

    for i := 0; i < len(maps_copy); i++ {
        if maps_copy[i] != nil {
            for maps_copy[i].First(); !maps_copy[i].IsDone(); maps_copy[i].Next() {
                key_value := maps_copy[i].Current()
                h.Insert(key_value.Key.(Hashable), key_value.Value)
            }

        }

    }

}



func abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}

func nextPrime(i int) int {
    if i < 3{
        return 3
    }
    next_prime := i+1
    remainder := false
    for n:=2; n < next_prime; n++{
        if next_prime % n == 0 {
            remainder = true
            break
        } 
    }
    if remainder == false {
        return next_prime
    }
    return nextPrime(i+1)
}


