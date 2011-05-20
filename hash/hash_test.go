package hash

import "testing"
import "fmt"

type Dude struct {
    name string
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
    
    a := m.Get(s[0].Key) 
    fmt.Println(a == a)

/*
    if m.Contains(s[1].Key) == false {
        t.Errorf("m.Contains(s[1].Key) should be true.\n")
    }

    if m.Contains(s[2].Key) == false {
        t.Errorf("m.Contains(s[2].Key) should be true.\n")
    }

    if m.Contains(s[3].Key) == false {
        t.Errorf("m.Contains(s[4].Key) should be true.\n")
    }
*/   

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
