package binary

import "testing"

func TestSearch(t *testing.T) {

    str := []string{ "B", "C", "D", "F", "G"}

    pos := Search(str, "E")

    if pos != 3 {
        t.Errorf("index was %d, but should be 3", pos)
    }

}

func TestRsearch(t *testing.T) {
    
    str := []string{ "B", "C", "D", "F", "G"}

    pos := Rsearch(str, "E")

    if pos != 3 {
        t.Errorf("index was %d, but should be -1", pos)
    }
}
