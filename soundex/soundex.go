package soundex

import "strings"

func mapRune(rune int) string {
    m := make(map[int]string)
    m[65] = "0"
    m[66] = "1"
    m[67] = "2"
    m[68] = "3"
    m[69] = "0"
    m[70] = "1"
    m[71] = "2"
    m[72] = "0"
    m[73] = "0"
    m[74] = "2"
    m[75] = "2"
    m[76] = "4"
    m[77] = "5"
    m[78] = "5"
    m[79] = "0"
    m[80] = "1"
    m[81] = "2"
    m[82] = "6"
    m[83] = "2"
    m[84] = "3"
    m[85] = "0"
    m[86] = "1"
    m[87] = "0"
    m[88] = "2"
    m[89] = "0"
    m[90] = "2"
    
    if m[rune] != "" {
        return m[rune]
    }
    return "0"
}

func Encode(word string) string{
    if word == "" {
        return "0000"
    }
//    fmt.Println(strings.ToUpper(word))
    word_runes := []int(strings.ToUpper(word))
    result_runes := []int("0000")

    result_runes[0] = word_runes[0]

    stringIndex := 1
    resultIndex := 1

    for stringIndex < len(word_runes) && resultIndex < len(result_runes) {
        char := mapRune(word_runes[stringIndex])
        //fmt.Printf("%s maps to %s\n", string(word_runes[stringIndex]), char)

        if char != "0" && []int(char)[0] != result_runes[resultIndex-1] {
            result_runes[resultIndex] = []int(char)[0]
            resultIndex ++
        }
        stringIndex ++

    }
    return string(result_runes)
}

