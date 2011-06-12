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

