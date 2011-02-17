package binary

import "testing"

func TestSearch(t *testing.T) {

    str := []string{ "B", "C", "D", "E", "F", "G"}

    pos := Search(str, "G")

    if pos != 5 {
        t.Errorf("index was %d, but should be -1", pos)
    }

}
