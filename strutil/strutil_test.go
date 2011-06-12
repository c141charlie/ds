package strutil

import "testing"

func TestFindOneLetterMatch(t *testing.T) {
    str := "c"
    pattern := "c"

    pos, err := SearchUsingBruteForce(pattern, str)

    if pos != 0  || err != nil {
        t.Errorf("pos should be 0 and err should be nil\n")
    }

}

func TestFindTwoLetterMatch(t *testing.T) {
    str := "cd"
    pattern := "cd"

    pos, err := SearchUsingBruteForce(pattern, str)

    if pos != 0  || err != nil {
        t.Errorf("pos should be 0 and err should be nil\n")
    }

}

func TestFindThreeLetterMatch(t *testing.T) {
    str := "cde"
    pattern := "cde"

    pos, err := SearchUsingBruteForce(pattern, str)

    if pos != 0  || err != nil {
        t.Errorf("pos should be 0 and err should be nil\n")
    }

}

func TestPatternGreaterThanText(t *testing.T) {
    str := "String"
    pattern := "StringP"

    pos, err := SearchUsingBruteForce(pattern, str)

    if pos != -1  || err == nil {
        t.Errorf("pos should be -1 and err should not be nil\n")
    }
}

func TestPatternDoesNotExistInText(t *testing.T) {
    str := "String"
    pattern := "Foo"

    pos, err := SearchUsingBruteForce(pattern, str)

    if pos != -1  || err == nil {
        t.Errorf("pos should be -1 and err should not be nil\n")
    }
}
