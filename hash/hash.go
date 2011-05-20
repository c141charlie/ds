package hash

import "list"

const default_capacity = 17
const load_factor = 0.75

type Map struct {
    buckets *list.List
}

func NewMap() *Map {
    return &Map{list.NewList()}
}

func (m *Map) Set(key interface{}, value interface{}) {
    kv := m.getKeyValueFor(key)
    
    if kv != nil {
        kv.Key = key
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

func (m *Map) Get(key interface{}) *KeyValue {
    i := list.NewIterator(m.buckets)

    for i.First(); !i.IsDone(); i.Next() {
        c := i.Current().GetValue().(*KeyValue)

        if c.Equals(key) {
            return c
        }
    }
    return nil
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
    size int
    slots []*list.List //slice of buckets
}

func NewHashtable() *Hashtable {
    return &Hashtable{0, make([]*list.List, default_capacity)}
}


func (h *Hashtable) Insert(key Hashable, value interface{}) bool{
    //get correct linked list
    //b := h.getBuckets(key)

    //see if key is in the list
        //if so update it
        //else insert it
    
    
    //increase the size by 1

    //resize the list if necessary
    return false
}

func (h *Hashtable) getBuckets(key Hashable) *list.List {
    bucket_list := h.slots[key.Hash() % len(h.slots)]

    if bucket_list == nil {
        bucket_list = list.NewList()
        h.slots[key.Hash() %len(h.slots)] = bucket_list
    }

    return bucket_list
}


func (h *Hashtable) Size() int {
    return h.size
}

func (h *Hashtable) numSlots() int {
    return len(h.slots)
}



type Key struct {
    key string
}

func (k *Key) Hash() int {
    a := []int(k.key)
    result := 0

    for i:= 0; i < len(a); i++ {
        result = result * 31 + a[i] 
    }
    return result
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



