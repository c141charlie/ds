package strutil

import "os" 

func SearchUsingBruteForce(pattern string, text string) (int, os.Error) {
    if text == "" || pattern == "" {
        return -1, os.NewError("text cannot be empty")
    }

    runes_pattern := []int(pattern)
    runes_text := []int(text)

    cur := 0

    for cur <= len(runes_text) - len(runes_pattern) {
        pos := 0    
        for pos < len(runes_pattern) && runes_pattern[pos] == runes_text[cur+pos] {
            pos ++
        }
        if pos == len(runes_pattern) {
            return cur, nil
        }
        cur ++
    }
    
    return -1, os.NewError("pattern not found")
}


func Search(pattern string, text string) (int, os.Error) {
    if pattern == "" || text == "" {
        return -1, os.NewError("pattern and/or text cannot be blank")
    }

    runes_pattern := []int(pattern)
    runes_text := []int(text)

    s := 0

    rune := -1

    last := computeLastOccurance(pattern)

    for s <= len(runes_text) - len(runes_pattern) {
    
        i := len(runes_pattern) - 1

        for i >= 0 && runes_pattern[i] == runes_text[s+i] {
            rune = runes_text[s+i]
            i --

        }
        
        if i < 0 {
            return s, nil
        }
        
        s += max(i - last[rune], 1)
        rune = -1
    }

    return -1, os.NewError("Not found yet")
}

func computeLastOccurance(pattern string) map[int]int {
    
    runes_pattern := []int(pattern)
    
    last := make(map[int]int)
    
    for i := 0; i < len(runes_pattern); i ++ {
        last[runes_pattern[i]] = i
    }
    return last
}


func max(i, j int) int {
    if i > j { return i }
    return j
}
