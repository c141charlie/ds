package strutil

import "testing"
import "fmt"

func TestFind(t *testing.T) {
    str := "String Matching"
    pattern := "Str"

    pos, err := SearchUsingBruteForce(pattern, str)

    fmt.Printf("pos = %d\n", pos)

    if pos != 1 || err == nil {
        t.Errorf("pos should be 0 and err should be nil\n")
    }

}
