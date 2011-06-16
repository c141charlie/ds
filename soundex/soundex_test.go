package soundex

import "testing"
import "fmt"

func TestFirstLetterIsAlwaysUsed(t *testing.T) {

    for rune := []int("A")[0]; rune <= []int("Z")[0]; rune ++ {
        result := Encode(string(rune))
       
        if result == "" {
            t.Errorf("result should not be \"\"")
        }

        if len([]int(result)) != 4 {
            t.Errorf("result should be 4 runes")
        }

        if []int(result)[0] != rune {
            t.Errorf("result not correct. %d is %d\n", []int(result)[0], rune)
        }
        
    }

}

func TestVowelsAreIgnored(t *testing.T) {
    
    if assertAllEquals("0", "AEIOUHWY") != true {
        t.Errorf("AEIOUHWY should all have the value of 0")
    }
    
}

func TestLettesrRepresentedByOne(t *testing.T) {
    if assertAllEquals("1", "BFPV") != true {
        t.Errorf("BFPV should all have the value of 1")
    }
}

func TestLettesrRepresentedByTwo(t *testing.T) {
    if assertAllEquals("2", "CGJKQSXZ") != true {
        t.Errorf("CGJKQSXZ should all have the value of 2")
    }
}

func TestLettersRepresentedByThree(t *testing.T) {
    if assertAllEquals("3", "DT") != true {
        t.Errorf("DT should equal 3")
    }
}

func TestLettersRepresentedByFour(t *testing.T) {
    if assertAllEquals("4", "L") != true {
        t.Errorf("L should equal 4")
    }
}

func TestLettersRepresentedByFive(t *testing.T) {
    if assertAllEquals("5", "MN") != true {
        t.Errorf("L should equal 5")
    }
}

func TestLettersRepresentedBySix(t *testing.T) {
    if assertAllEquals("6", "R") != true {
        t.Errorf("R should equal 6")
    }
}

func assertAllEquals(char string, list string) bool {
    runes := []int(list)

    for _, rune := range(runes) {
        if char != string(mapRune(rune)) {
            fmt.Printf("failing because %d maps to %s\n", rune, mapRune(rune))
            return false
        }
    }

    return true
}

func TestDuplicateCodesAreDropped(t *testing.T) {
    assertEncodesEquals(t, "B100", "BFPV") 
    assertEncodesEquals(t, "C200", "CGJKQSXZ") 
    assertEncodesEquals(t, "D300", "DDT") 
    assertEncodesEquals(t, "L400", "LLL") 
    assertEncodesEquals(t, "M500", "MNMN") 
    assertEncodesEquals(t, "R600", "RRR") 
}

func TestSomeRealStrings(t *testing.T) {
    assertEncodesEquals(t, "S530", "Smith")
    assertEncodesEquals(t, "S530", "Smythe")
    assertEncodesEquals(t, "M235", "McDonald")
    assertEncodesEquals(t, "M235", "MacDonald")
    assertEncodesEquals(t, "H620", "Harris")
    assertEncodesEquals(t, "H620", "Harrys")
}

func assertEncodesEquals(t *testing.T, output string, input string) bool {
    if Encode(input) != output {
        t.Errorf("input doesn't match output. %s maps to %s", input, Encode(input))
    }
    return true
}
