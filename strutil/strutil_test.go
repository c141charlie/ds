package strutil

import "testing"
import "os"



func SearchFunction() func(string, string) (int, os.Error) {
    return Search
}


func TestFindOneLetterMatch(t *testing.T) {
    str := "a"
    pattern := "a"

    f := SearchFunction()

    pos, err := f(pattern, str)

    if pos != 0  || err != nil {
        t.Errorf("pos should be 0, but is %d and err should be nil\n", pos)
    }

}

func TestFindTwoLetterMatch(t *testing.T) {
    str := "cd"
    pattern := "cd"
    
    f := SearchFunction()

    pos, err := f(pattern, str)

    if pos != 0  || err != nil {
        t.Errorf("pos should be 0 and err should be nil\n")
    }

}

func TestFindThreeLetterMatch(t *testing.T) {
    str := "cde"
    pattern := "cde"
    f := SearchFunction()
    pos, err := f(pattern, str)

    if pos != 0  || err != nil {
        t.Errorf("pos should be 0 and err should be nil\n")
    }

}

func TestPatternGreaterThanText(t *testing.T) {
    str := "String"
    pattern := "StringP"
    f := SearchFunction()
    pos, err := f(pattern, str)

    if pos != -1  || err == nil {
        t.Errorf("pos should be -1 and err should not be nil\n")
    }
}

func TestPatternDoesNotExistInText(t *testing.T) {
    str := "String"
    pattern := "Foo"
    f := SearchFunction()
    pos, err := f(pattern, str)

    if pos != -1  || err == nil {
        t.Errorf("pos should be -1 and err should not be nil\n")
    }
}

