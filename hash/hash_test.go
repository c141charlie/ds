package hash

import "testing"

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

func TestHash(t *testing.T) {
    
    h := NewHashtable()

    if h.Size() != 0 {
        t.Errorf("h.Size() should be 0.\n")
    }

    if h.numSlots() != 17 {
        t.Errorf("h.numSlots() should be 17.\n")
    }

    s := h.getBuckets(&Key{"hello"})

    if s == nil {
        t.Errorf("s should not be nil\n")
    }
}

func TestPrimeNumberGenerator(t *testing.T) {
    p := 3
    np := nextPrime(p)
    if np != 5 {
        t.Errorf("np should be 5, but was %d\n", np)
    }

    p = 5
    np = nextPrime(p)
    if np != 7 {
        t.Errorf("np should be 7, but was %d\n", np)
    }

    p = 7
    np = nextPrime(p)
    if np != 11 {
        t.Errorf("np should be 11, but was %d\n", np)
    }
}

func TestKey(t *testing.T) {
    k := &Key{"hello"}
    if k.Hash() != 99162322 {
        t.Errorf("k.Hash() should be 99162322.\n")
    }

}

func getVars() []*KeyValue {
    s := make([]*KeyValue, 6)
    s[0] = &KeyValue{"akey", "aval"}
    s[1] = &KeyValue{"bkey", "bval"}
    s[2] = &KeyValue{"ckey", "cval"}
    s[3] = &KeyValue{"dkey", "dval"}
    s[4] = &KeyValue{"ekey", "eval"}
    s[5] = &KeyValue{"fkey", "fval"}
    return s
}


func TestMapContainsExisting(t *testing.T) {
    
    s := getVars()

    m := NewMap()

    m.Set(s[2].Key, s[2].Value)
    m.Set(s[0].Key, s[0].Value)
    m.Set(s[1].Key, s[1].Value)
    m.Set(s[3].Key, s[3].Value)
    

    if m.Contains(s[0].Key) == false {
        t.Errorf("m.Contains(s[0].Key) should be true.\n")
    }

    if m.Contains(s[1].Key) == false {
        t.Errorf("m.Contains(s[1].Key) should be true.\n")
    }

    if m.Contains(s[2].Key) == false {
        t.Errorf("m.Contains(s[2].Key) should be true.\n")
    }

    if m.Contains(s[3].Key) == false {
        t.Errorf("m.Contains(s[4].Key) should be true.\n")
    }
    

}

func TestMapContainsNonExisting(t *testing.T) {
    s := getVars()
    m := NewMap()
    if m.Contains(s[4].Key) == true {
        t.Errorf("m.Contains(s[4].Key) should be false.\n")
    }
    
    if m.Contains(s[5].Key) == true {
        t.Errorf("m.Contains(s[5].Key) should be false.\n")
    }
}


func TestMapGetExisting(t *testing.T) {
    
    s := getVars()

    m := NewMap()

    m.Set(s[2].Key, s[2].Value)
    m.Set(s[0].Key, s[0].Value)
    m.Set(s[1].Key, s[1].Value)
    m.Set(s[3].Key, s[3].Value)
    
    if m.Get(s[0].Key) != s[0].Value {
        t.Errorf("m.Get(s[0].Key) should equal s[0].Value\n")
    }
    if m.Get(s[1].Key) != s[1].Value {
        t.Errorf("m.Get(s[1].Key) should equal s[1].Value\n")
    }

    if m.Get(s[2].Key) != s[2].Value {
        t.Errorf("m.Get(s[2].Key) should equal s[2].Value\n")
    }

    if m.Get(s[3].Key) != s[3].Value {
        t.Errorf("m.Get(s[3].Key) should equal s[3].Value\n")
    }
}

func TestMapGetNonExisting(t *testing.T) {
    
    s := getVars()

    m := NewMap()

    m.Set(s[2].Key, s[2].Value)
    m.Set(s[0].Key, s[0].Value)
    m.Set(s[1].Key, s[1].Value)
    m.Set(s[3].Key, s[3].Value)
    
    if m.Get(s[4].Key) == s[4].Value {
        t.Errorf("m.Get(s[4].Key) should not equal s[4].Value\n")
    }

    if m.Get(s[5].Key) == s[5].Value {
        t.Errorf("m.Get(s[5].Key) should not equal s[5].Value\n")
    }

    if m.Get(s[4].Key) != nil {
        t.Errorf("m.Get(s[4].Key should be nil.\n")
    }

    if m.Get(s[5].Key) != nil {
        t.Errorf("m.Get(s[5].Key should be nil.\n")
    }

}

func TestSetExistingKey(t *testing.T) {
    s := getVars()

    m := NewMap()

    m.Set(s[2].Key, s[2].Value)
    m.Set(s[0].Key, s[0].Value)
    
    if m.Get(s[2].Key) != s[2].Value {
        t.Errorf("m.Get(s[2].Key) should not equal s[2].Value\n")
    }
    
    m.Set(s[2].Key, "value 2")

    if m.Get(s[2].Key) != "value 2" {
        t.Errorf("m.Get(s[2].Key) should equal \"value 2\" but instead equals %s\n", m.Get(s[2].Key))
    }
}

func TestDeleteExisting(t *testing.T) {
    s := getVars()

    m := NewMap()

    m.Set(s[2].Key, s[2].Value)
    m.Set(s[0].Key, s[0].Value)
    
    if m.Get(s[2].Key) != s[2].Value {
        t.Errorf("m.Get(s[2].Key) should not equal s[2].Value\n")
    }

    m.Delete(s[2].Key)

    if m.Get(s[2].Key) != nil {
        t.Errorf("m.Get(s[2].Key should be nil\n")
    }

    if m.Size() != 1 {
        t.Errorf("m.Size() should equal 1\n")
    }
    
}

func TestDeleteNonExisting(t *testing.T) {
    s := getVars()

    m := NewMap()

    m.Set(s[2].Key, s[2].Value)
    m.Set(s[0].Key, s[0].Value)

    m.Delete(s[3].Key) 
    
    if m.Get(s[2].Key) != s[2].Value {
        t.Errorf("m.Get(s[2].Key) should not equal s[2].Value\n")
    }
        
    if m.Get(s[0].Key) != s[0].Value {
        t.Errorf("m.Get(s[0].Key) should not equal s[0].Value\n")
    }

    if m.Size() != 2 {
        t.Errorf("m.Size() should equal 2\n")
    }
}

func TestClear(t *testing.T) {
    s := getVars()

    m := NewMap()

    m.Set(s[2].Key, s[2].Value)
    m.Set(s[0].Key, s[0].Value)
    
    m.Clear()

    if m.Size() != 0 {
        t.Errorf("m.Size() should be 0\n")
    }

    if m.Get(s[0].Key) != nil {
        t.Errorf("m.Get(s[0].Key) should be nil\n")
    }

    if m.Get(s[2].Key) != nil {
        t.Errorf("m.Get(s[2].Key) should be nil\n")
    }
}

func TestIterator(t *testing.T) {
    
    s := getVars()

    m := NewMap()

    m.Set(s[2].Key, s[2].Value)
    m.Set(s[0].Key, s[0].Value)
    m.Set(s[1].Key, s[1].Value)

    

    if m.First().Key != s[2].Key {
        t.Errorf("i.First().Key should equal s[2].Key\n")
    }

    if m.First().Value != s[2].Value {
        t.Errorf("i.First().Value should equal s[2].Value\n")
    }

    m.Next()

    if m.Current().Key != s[0].Key {
        t.Errorf("m.Current().Key should equal s[0].Key\n")
    }

    if m.Current().Value != s[0].Value {
        t.Errorf("m.Current().Value should equal s[0].Value\n")
    }

    m.Next()

    if m.Current().Key != s[1].Key {
        t.Errorf("m.Current().Key should equal s[1].Key\n")
    }

    if m.Current().Value != s[1].Value {
        t.Errorf("m.Current().Value should equal s[1].Value\n")
    }
    
    m.Next()

    if m.IsDone() != true {
        t.Errorf("m.IsDone() should be true\n")
    }   

}

func TestInsertHashtable(t *testing.T) {
    h := NewHashtable()

    k1 := &Key{"Lale"}
    h.Insert(k1, "Ismen")
    
    k2 := &Key{"Charles"}
    h.Insert(k2, "Thompson")

    if h.Size() != 2 {
        t.Errorf("h.Size() should be 2\n.")
    }

    if h.Get(k1) != "Ismen" {
        t.Errorf("h.Get(k1) should be \"Ismen\"\n")
    }
    
    if h.Get(k2) != "Thompson" {
        t.Errorf("h.Get(k1) should be \"Ismen\"\n")
    }


}

func TestDeleteHashtable(t *testing.T) {
    h := NewHashtable()

    k1 := &Key{"Lale"}
    h.Insert(k1, "Ismen")
    
    k2 := &Key{"Charles"}
    h.Insert(k2, "Thompson")

    if h.Size() != 2 {
        t.Errorf("h.Size() should be 2\n")
    }

    h.Delete(k1)

    if h.Size() != 1 {
        t.Errorf("h.Size() should be 1\n")
    }

    if h.Get(k1) != nil{
        t.Errorf("h.Get(k1) should be nil\n")
    }
    
    if h.Get(k2) != "Thompson" {
        t.Errorf("h.Get(k1) should be \"Ismen\"\n")
    }

}

func TestDeleteItemThatDoesNotExistFromHashtable(t *testing.T) {
    h := NewHashtable()

    k1 := &Key{"Lale"}
    h.Insert(k1, "Ismen")
    
    k2 := &Key{"Charles"}
    h.Insert(k2, "Thompson")

    if h.Size() != 2 {
        t.Errorf("h.Size() should be 2\n")
    }

    h.Delete(&Key{"Dude"})

    if h.Size() != 2 {
        t.Errorf("h.Size() should be 2\n")
    }

    if h.Get(k1) != "Ismen"{
        t.Errorf("h.Get(k1) should be \"Ismen\"\n")
    }
    
    if h.Get(k2) != "Thompson" {
        t.Errorf("h.Get(k1) should be \"Ismen\"\n")
    }

}

func TestHashtableContains(t *testing.T) {
    h := NewHashtable()

    k1 := &Key{"Lale"}
    h.Insert(k1, "Ismen")
    
    k2 := &Key{"Charles"}
    h.Insert(k2, "Thompson")

    if h.Size() != 2 {
        t.Errorf("h.Size() should be 2\n")
    }
    
    if !h.Contains(k1) {
        t.Errorf("h.Contains(k1) should be true\n")
    }
}






